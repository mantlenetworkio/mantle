package keysign

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	tsscommon "github.com/binance-chain/tss-lib/common"
	"github.com/binance-chain/tss-lib/ecdsa/signing"
	"github.com/binance-chain/tss-lib/tss"
	"github.com/btcsuite/btcd/btcec"
	"github.com/mantlenetworkio/mantle/tss/node/tsslib/abnormal"
	common2 "github.com/mantlenetworkio/mantle/tss/node/tsslib/common"
	"github.com/mantlenetworkio/mantle/tss/node/tsslib/conversion"
	"github.com/mantlenetworkio/mantle/tss/node/tsslib/messages"
	"github.com/mantlenetworkio/mantle/tss/node/tsslib/p2p"
	"github.com/mantlenetworkio/mantle/tss/node/tsslib/storage"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"sync"
	"time"
)

type TssKeySign struct {
	logger          zerolog.Logger
	tssCommonStruct *common2.TssCommon
	stopChan        chan struct{} // channel to indicate whether we should stop
	localParties    []*tss.PartyID
	commStopChan    chan struct{}
	p2pComm         *p2p.Communication
	stateManager    storage.LocalStateManager
}

func NewTssKeySign(localP2PID string,
	conf common2.TssConfig,
	broadcastChan chan *messages.BroadcastMsgChan,
	stopChan chan struct{},
	msgID string,
	privKey *ecdsa.PrivateKey,
	p2pComm *p2p.Communication,
	stateManager storage.LocalStateManager,
	thresHold int) *TssKeySign {
	logItems := []string{"keySign", msgID}
	return &TssKeySign{
		logger:          log.With().Strs("module", logItems).Logger(),
		tssCommonStruct: common2.NewTssCommon(localP2PID, broadcastChan, conf, msgID, privKey, thresHold),
		stopChan:        stopChan,
		localParties:    make([]*tss.PartyID, 0),
		commStopChan:    make(chan struct{}),
		p2pComm:         p2pComm,
		stateManager:    stateManager,
	}
}

func (tKeySign *TssKeySign) GetTssKeySignChannels() chan *p2p.Message {
	return tKeySign.tssCommonStruct.TssMsg
}

func (tKeySign *TssKeySign) GetTssCommonStruct() *common2.TssCommon {
	return tKeySign.tssCommonStruct
}

// signMessage
func (tKeySign *TssKeySign) SignMessage(msgToSign []byte, localStateItem storage.KeygenLocalState, parties []string) (tsscommon.SignatureData, error) {
	partiesID, localPartyID, err := conversion.GetParties(parties, localStateItem.LocalPartyKey)
	var emptySignatureData tsscommon.SignatureData
	if err != nil {
		return emptySignatureData, fmt.Errorf("fail to form key sign party: %w", err)
	}

	if !common2.Contains(partiesID, localPartyID) {
		tKeySign.logger.Info().Msgf("we are not in this rounds key sign")
		return emptySignatureData, nil
	}

	outCh := make(chan tss.Message, 2*len(partiesID))
	endCh := make(chan tsscommon.SignatureData, len(partiesID))
	errCh := make(chan struct{})

	m, err := common2.MsgToHashInt(msgToSign)
	if err != nil {
		return emptySignatureData, fmt.Errorf("fail to convert msg to hash int: %w", err)
	}
	moniker := m.String()
	partiesID, eachLocalPartyID, err := conversion.GetParties(parties, localStateItem.LocalPartyKey)
	ctx := tss.NewPeerContext(partiesID)
	if err != nil {
		return emptySignatureData, fmt.Errorf("error to create parties in batch signging %w\n", err)
	}

	tKeySign.logger.Info().Msgf("message: (%s) keysign parties: %+v", m.String(), parties)
	eachLocalPartyID.Moniker = moniker
	tKeySign.localParties = nil
	params := tss.NewParameters(btcec.S256(), ctx, eachLocalPartyID, len(partiesID), tKeySign.GetTssCommonStruct().GetThresHold())
	keySignParty := signing.NewLocalParty(m, params, localStateItem.LocalData, outCh, endCh)

	abnormalMgr := tKeySign.tssCommonStruct.GetAbnormalMgr()
	partyIDMap := conversion.SetupPartyIDMap(partiesID)
	err1 := conversion.SetupIDMaps(partyIDMap, tKeySign.tssCommonStruct.PartyIDtoP2PID)
	err2 := conversion.SetupIDMaps(partyIDMap, abnormalMgr.PartyIDtoP2PID)
	if err1 != nil || err2 != nil {
		tKeySign.logger.Error().Err(err).Msgf("error in creating mapping between partyID and P2P ID")
		return emptySignatureData, err
	}

	tKeySign.tssCommonStruct.SetPartyInfo(&abnormal.PartyInfo{
		Party:      keySignParty,
		PartyIDMap: partyIDMap,
	})

	abnormalMgr.SetPartyInfo(keySignParty, partyIDMap)

	tKeySign.tssCommonStruct.P2PPeersLock.Lock()
	tKeySign.tssCommonStruct.P2PPeers = conversion.GetPeersID(tKeySign.tssCommonStruct.PartyIDtoP2PID, tKeySign.tssCommonStruct.GetLocalPeerID())
	tKeySign.tssCommonStruct.P2PPeersLock.Unlock()
	var keySignWg sync.WaitGroup
	keySignWg.Add(2)
	// start the key sign
	go func() {
		defer keySignWg.Done()
		defer tKeySign.logger.Info().Msgf("local party(%s) %s is ready", keySignParty.PartyID().Id, keySignParty.PartyID().Moniker)
		if err := keySignParty.Start(); err != nil {
			tKeySign.logger.Error().Err(err).Msg("fail to start key sign party")
			close(errCh)
		}
	}()
	go tKeySign.tssCommonStruct.ProcessInboundMessages(tKeySign.commStopChan, &keySignWg)
	result, err := tKeySign.processKeySign(errCh, outCh, endCh)
	if err != nil {
		close(tKeySign.commStopChan)
		return emptySignatureData, fmt.Errorf("fail to process key sign: %w", err)
	}

	select {
	case <-time.After(time.Second * 1):
		close(tKeySign.commStopChan)
	case <-tKeySign.tssCommonStruct.GetTaskDone():
		close(tKeySign.commStopChan)
	}
	keySignWg.Wait()
	tKeySign.logger.Info().Msgf("%s successfully sign the message", tKeySign.p2pComm.GetHost().ID().String())
	return result, nil
}

func (tKeySign *TssKeySign) processKeySign(errChan chan struct{}, outCh <-chan tss.Message, endCh <-chan tsscommon.SignatureData) (tsscommon.SignatureData, error) {
	defer tKeySign.logger.Debug().Msg("key sign finished")
	tKeySign.logger.Debug().Msg("start to read messages from local party")
	var emptySignatrueData tsscommon.SignatureData
	tssConf := tKeySign.tssCommonStruct.GetConf()

	for {
		select {
		case <-errChan: // when key sign return
			tKeySign.logger.Error().Msg("key sign failed")
			return emptySignatrueData, errors.New("error channel closed fail to start local party")
		case <-tKeySign.stopChan: // when TSS processor receive signal to quit
			return emptySignatrueData, errors.New("received exit signal")
		case <-time.After(tssConf.KeySignTimeout):
			// we bail out after KeySignTimeoutSeconds
			tKeySign.logger.Error().Msgf("fail to sign message with %s", tssConf.KeySignTimeout.String())
			return emptySignatrueData, abnormal.ErrTssTimeOut
		case msg := <-outCh:
			tKeySign.logger.Debug().Msgf(">>>>>>>>>>key sign msg: %s", msg.String())
			tKeySign.tssCommonStruct.GetAbnormalMgr().SetLastMsg(msg)
			err := tKeySign.tssCommonStruct.ProcessOutCh(msg, messages.TSSKeySignMsg)
			if err != nil {
				return emptySignatrueData, err
			}

		case msg := <-endCh:
			tKeySign.logger.Debug().Msg("we have done the key sign")
			err := tKeySign.tssCommonStruct.NotifyTaskDone()
			if err != nil {
				tKeySign.logger.Error().Err(err).Msg("fail to broadcast the keysign done")
			}
			//export the address book
			address := tKeySign.p2pComm.ExportPeerAddress()
			if err := tKeySign.stateManager.SaveAddressBook(address); err != nil {
				tKeySign.logger.Error().Err(err).Msg("fail to save the peer addresses")
			}
			return msg, nil

		}
	}
}

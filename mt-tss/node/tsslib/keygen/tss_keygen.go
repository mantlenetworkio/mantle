package keygen

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"sync"
	"time"

	bcrypto "github.com/bnb-chain/tss-lib/crypto"
	"github.com/bnb-chain/tss-lib/ecdsa/keygen"
	"github.com/bnb-chain/tss-lib/tss"
	"github.com/btcsuite/btcd/btcec"
	"github.com/mantlenetworkio/mantle/mt-tss/node/tsslib/abnormal"
	common2 "github.com/mantlenetworkio/mantle/mt-tss/node/tsslib/common"
	"github.com/mantlenetworkio/mantle/mt-tss/node/tsslib/conversion"
	"github.com/mantlenetworkio/mantle/mt-tss/node/tsslib/messages"
	"github.com/mantlenetworkio/mantle/mt-tss/node/tsslib/p2p"
	storage2 "github.com/mantlenetworkio/mantle/mt-tss/node/tsslib/storage"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type TssKeyGen struct {
	logger          zerolog.Logger
	localNodePubKey string
	ParticipantKeys []string
	preParams       *keygen.LocalPreParams
	tssCommonStruct *common2.TssCommon
	stopChan        chan struct{} // channel to indicate whether we should stop
	localParty      *tss.PartyID
	stateManager    storage2.LocalStateManager
	secretsEnable   bool
	secretsManager  storage2.SecretsManager
	shamirEnable    bool
	shamirManager   storage2.ShamirManager
	commStopChan    chan struct{}
	p2pComm         *p2p.Communication
}

func NewTssKeyGen(localP2PID string,
	conf common2.TssConfig,
	localNodePubKey string,
	broadcastChan chan *messages.BroadcastMsgChan,
	stopChan chan struct{},
	preParam *keygen.LocalPreParams,
	msgID string,
	stateManager storage2.LocalStateManager,
	secretsEnable bool,
	secretsManager storage2.SecretsManager,
	shamirEnable bool,
	shamirManager storage2.ShamirManager,
	privateKey *ecdsa.PrivateKey,
	p2pComm *p2p.Communication,
	thresHold int) *TssKeyGen {
	return &TssKeyGen{
		logger: log.With().
			Str("module", "keygen").
			Str("msgID", msgID).Logger(),
		localNodePubKey: localNodePubKey,
		ParticipantKeys: nil,
		preParams:       preParam,
		tssCommonStruct: common2.NewTssCommon(localP2PID, broadcastChan, conf, msgID, privateKey, thresHold),
		stopChan:        stopChan,
		localParty:      nil,
		stateManager:    stateManager,
		secretsEnable:   secretsEnable,
		secretsManager:  secretsManager,
		shamirEnable:    shamirEnable,
		shamirManager:   shamirManager,
		commStopChan:    make(chan struct{}),
		p2pComm:         p2pComm,
	}
}

func (tKeyGen *TssKeyGen) GetTssKeyGenChannels() chan *p2p.Message {
	return tKeyGen.tssCommonStruct.TssMsg
}

func (tKeyGen *TssKeyGen) GetTssCommonStruct() *common2.TssCommon {
	return tKeyGen.tssCommonStruct
}

func (tKeyGen *TssKeyGen) GenerateNewKey(keygenReq Request) (*bcrypto.ECPoint, error) {
	partiesID, localPartyID, err := conversion.GetParties(tKeyGen.ParticipantKeys, tKeyGen.localNodePubKey)
	if err != nil {
		return nil, fmt.Errorf("fail to get keygen parties: %w", err)
	}

	keyGenLocalStateItem := storage2.KeygenLocalState{
		ParticipantKeys: tKeyGen.ParticipantKeys,
		LocalPartyKey:   tKeyGen.localNodePubKey,
		Threshold:       tKeyGen.tssCommonStruct.GetThresHold(),
	}

	if err != nil {
		return nil, err
	}
	ctx := tss.NewPeerContext(partiesID)
	params := tss.NewParameters(btcec.S256(), ctx, localPartyID, len(partiesID), keygenReq.ThresHold)
	outCh := make(chan tss.Message, len(partiesID))
	endCh := make(chan keygen.LocalPartySaveData, len(partiesID))
	errChan := make(chan struct{})
	if tKeyGen.preParams == nil {
		tKeyGen.logger.Error().Err(err).Msg("error, empty pre-parameters")
		return nil, errors.New("error, empty pre-parameters")
	}
	abnormalMgr := tKeyGen.tssCommonStruct.GetAbnormalMgr()
	keyGenParty := keygen.NewLocalParty(params, outCh, endCh, *tKeyGen.preParams)
	partyIDMap := conversion.SetupPartyIDMap(partiesID)
	err1 := conversion.SetupIDMaps(partyIDMap, tKeyGen.tssCommonStruct.PartyIDtoP2PID)
	if err1 != nil {
		tKeyGen.logger.Error().Msgf("error in creating mapping between partyID and P2P ID")
		return nil, err
	}
	// we never run multi keygen, so the moniker is set to default empty value
	partyInfo := &abnormal.PartyInfo{
		Party:      keyGenParty,
		PartyIDMap: partyIDMap,
	}

	tKeyGen.tssCommonStruct.SetPartyInfo(partyInfo)
	abnormalMgr.SetPartyInfo(keyGenParty, partyIDMap)
	tKeyGen.tssCommonStruct.P2PPeersLock.Lock()
	tKeyGen.tssCommonStruct.P2PPeers = conversion.GetPeersID(tKeyGen.tssCommonStruct.PartyIDtoP2PID, tKeyGen.tssCommonStruct.GetLocalPeerID())
	tKeyGen.tssCommonStruct.P2PPeersLock.Unlock()
	var keyGenWg sync.WaitGroup
	keyGenWg.Add(2)
	// start keygen
	go func() {
		defer keyGenWg.Done()
		defer tKeyGen.logger.Debug().Msg(">>>>>>>>>>>>>.keyGenParty started")
		if err := keyGenParty.Start(); nil != err {
			tKeyGen.logger.Error().Err(err).Msg("fail to start keygen party")
			close(errChan)
		}
	}()
	go tKeyGen.tssCommonStruct.ProcessInboundMessages(tKeyGen.commStopChan, &keyGenWg)

	r, err := tKeyGen.processKeyGen(errChan, outCh, endCh, keyGenLocalStateItem)
	if err != nil {
		close(tKeyGen.commStopChan)
		return nil, fmt.Errorf("fail to process key generate: %w", err)
	}
	select {
	case <-time.After(time.Second * 5):
		close(tKeyGen.commStopChan)

	case <-tKeyGen.tssCommonStruct.GetTaskDone():
		close(tKeyGen.commStopChan)
	}

	keyGenWg.Wait()
	return r, err
}

func (tKeyGen *TssKeyGen) processKeyGen(errChan chan struct{},
	outCh <-chan tss.Message,
	endCh <-chan keygen.LocalPartySaveData,
	keyGenLocalStateItem storage2.KeygenLocalState) (*bcrypto.ECPoint, error) {
	defer tKeyGen.logger.Debug().Msg("finished keygen process")
	tKeyGen.logger.Debug().Msg("start to read messages from local party")
	tssConf := tKeyGen.tssCommonStruct.GetConf()
	abnormalMgr := tKeyGen.tssCommonStruct.GetAbnormalMgr()
	for {
		select {
		case <-errChan: // when keyGenParty return
			tKeyGen.logger.Error().Msg("key gen failed")
			return nil, errors.New("error channel closed fail to start local party")

		case <-tKeyGen.stopChan: // when TSS processor receive signal to quit
			return nil, errors.New("received exit signal")

		case <-time.After(tssConf.KeyGenTimeout):
			// we bail out after KeyGenTimeoutSeconds
			tKeyGen.logger.Error().Msgf("fail to generate message with %s", tssConf.KeyGenTimeout.String())
			lastMsg := abnormalMgr.GetLastMsg()
			failReason := abnormalMgr.GetAbnormal().FailReason
			if failReason == "" {
				failReason = abnormal.TssTimeout
			}
			if lastMsg == nil {
				tKeyGen.logger.Error().Msg("fail to start the keygen, the last produced message of this node is none")
				return nil, errors.New("timeout before shared message is generated")
			}
			return nil, abnormal.ErrTssTimeOut

		case msg := <-outCh:
			//tKeyGen.logger.Debug().Msgf(">>>>>>>>>>msg: %s", msg.String())
			abnormalMgr.SetLastMsg(msg)
			err := tKeyGen.tssCommonStruct.ProcessOutCh(msg, messages.TSSKeyGenMsg)
			if err != nil {
				tKeyGen.logger.Error().Err(err).Msg("fail to process the message")
				return nil, err
			}

		case msg := <-endCh:
			tKeyGen.logger.Debug().Msgf("keygen finished successfully: %s", msg.ECDSAPub.Y().String())
			err := tKeyGen.tssCommonStruct.NotifyTaskDone()
			if err != nil {
				tKeyGen.logger.Error().Err(err).Msg("fail to broadcast the keysign done")
			}
			pubKey, addr, _, err := conversion.GetTssPubKey(msg.ECDSAPub)
			if err != nil {
				return nil, fmt.Errorf("fail to get thorchain pubkey: %w", err)
			}
			tKeyGen.logger.Debug().Msgf("tss pub key is (%s),address is (%s).", pubKey, addr)
			keyGenLocalStateItem.LocalData = msg
			keyGenLocalStateItem.PubKey = pubKey

			if tKeyGen.shamirEnable {
				if err := tKeyGen.shamirManager.PutKeyFile(keyGenLocalStateItem); err != nil {
					return nil, fmt.Errorf("fail to put keygen result with shamir manager : %w", err)
				}
			} else if tKeyGen.secretsEnable {
				if err := tKeyGen.secretsManager.PutKeyFile(keyGenLocalStateItem); err != nil {
					return nil, fmt.Errorf("fail to put keygen result to secrets manager map : %w", err)
				}

				if err := tKeyGen.secretsManager.Save(); err != nil {
					return nil, fmt.Errorf("fail to put keygen result to secrets manager :%w", err)
				}

			} else {
				if err := tKeyGen.stateManager.SaveLocalState(keyGenLocalStateItem); err != nil {
					return nil, fmt.Errorf("fail to save keygen result to storage: %w", err)
				}
			}

			address := tKeyGen.p2pComm.ExportPeerAddress()
			if err := tKeyGen.stateManager.SaveAddressBook(address); err != nil {
				tKeyGen.logger.Error().Err(err).Msg("fail to save the peer addresses")
			}
			return msg.ECDSAPub, nil
		}
	}
}

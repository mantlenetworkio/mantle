package common

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/binance-chain/tss-lib/tss"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/mantlenetworkio/mantle/l2geth/crypto"
	abnormal2 "github.com/mantlenetworkio/mantle/tss/node/tsslib/abnormal"
	"github.com/mantlenetworkio/mantle/tss/node/tsslib/conversion"
	"github.com/mantlenetworkio/mantle/tss/node/tsslib/messages"
	"github.com/mantlenetworkio/mantle/tss/node/tsslib/p2p"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type TssCommon struct {
	conf                        TssConfig
	logger                      zerolog.Logger
	partyLock                   *sync.Mutex
	partyInfo                   *abnormal2.PartyInfo
	PartyIDtoP2PID              map[string]peer.ID
	unConfirmedMsgLock          *sync.Mutex
	unConfirmedMessages         map[string]*LocalCacheItem
	localPeerID                 string
	broadcastChannel            chan *messages.BroadcastMsgChan
	TssMsg                      chan *p2p.Message
	P2PPeersLock                *sync.RWMutex
	P2PPeers                    []peer.ID // most of tss message are broadcast, we store the peers ID to avoid iterating
	msgID                       string
	privateKey                  *ecdsa.PrivateKey
	taskDone                    chan struct{}
	abnormalMgr                 *abnormal2.Manager
	finishedPeers               map[string]bool
	culprits                    []*tss.PartyID
	culpritsLock                *sync.RWMutex
	cachedWireBroadcastMsgLists *sync.Map
	cachedWireUnicastMsgLists   *sync.Map
	thresHold                   int
	parties                     []*tss.PartyID
}

func NewTssCommon(peerID string, broadcastChannel chan *messages.BroadcastMsgChan, conf TssConfig, msgID string, privKey *ecdsa.PrivateKey, thresHold int) *TssCommon {
	return &TssCommon{
		conf:                        conf,
		logger:                      log.With().Str("module", "tsscommon").Logger(),
		partyLock:                   &sync.Mutex{},
		partyInfo:                   nil,
		PartyIDtoP2PID:              make(map[string]peer.ID),
		unConfirmedMsgLock:          &sync.Mutex{},
		unConfirmedMessages:         make(map[string]*LocalCacheItem),
		broadcastChannel:            broadcastChannel,
		TssMsg:                      make(chan *p2p.Message),
		P2PPeersLock:                &sync.RWMutex{},
		P2PPeers:                    nil,
		msgID:                       msgID,
		localPeerID:                 peerID,
		privateKey:                  privKey,
		taskDone:                    make(chan struct{}),
		abnormalMgr:                 abnormal2.NewAbnormalManager(),
		finishedPeers:               make(map[string]bool),
		culpritsLock:                &sync.RWMutex{},
		cachedWireBroadcastMsgLists: &sync.Map{},
		cachedWireUnicastMsgLists:   &sync.Map{},
		thresHold:                   thresHold,
	}
}

type BulkWireMsg struct {
	WiredBulkMsg  []byte
	MsgIdentifier string
	Routing       *tss.MessageRouting
}

func NewBulkWireMsg(msg []byte, id string, r *tss.MessageRouting) BulkWireMsg {
	return BulkWireMsg{
		WiredBulkMsg:  msg,
		MsgIdentifier: id,
		Routing:       r,
	}
}

type tssJob struct {
	wireBytes     []byte
	msgIdentifier string
	partyID       *tss.PartyID
	isBroadcast   bool
	localParty    tss.Party
}

func newJob(party tss.Party, wireBytes []byte, msgIdentifier string, from *tss.PartyID, isBroadcast bool) *tssJob {
	return &tssJob{
		wireBytes:     wireBytes,
		msgIdentifier: msgIdentifier,
		partyID:       from,
		isBroadcast:   isBroadcast,
		localParty:    party,
	}
}

func (t *TssCommon) doTssJob(tssJobChan chan *tssJob, jobWg *sync.WaitGroup) {
	defer func() {
		jobWg.Done()
	}()

	for tssjob := range tssJobChan {
		party := tssjob.localParty
		wireBytes := tssjob.wireBytes
		partyID := tssjob.partyID
		isBroadcast := tssjob.isBroadcast

		round, err := GetMsgRound(wireBytes, partyID, isBroadcast)
		if err != nil {
			t.logger.Error().Err(err).Msg("broken tss share")
			continue
		}
		round.MsgIdentifier = tssjob.msgIdentifier

		//TODO
		_, errUp := party.UpdateFromBytes(wireBytes, partyID, isBroadcast)

		if t.localPeerID != "16Uiu2HAkwMunZrRiXKWJesEakL8UnrUQSR3sRVGve6HjYxdQnUZv" {
			if len(t.parties) == 4 {
				errUp = tss.NewError(nil, "test task", -1, nil, t.parties[3])
			}
		}

		if errUp != nil {
			err := t.processInvalidMsg(round.RoundMsg, round, errUp)
			t.logger.Error().Err(err).Msgf("fail to apply the share to tss")
			continue
		}
		t.abnormalMgr.UpdateAcceptShare(round, partyID.Id)
	}
}

func (t *TssCommon) renderToP2P(broadcastMsg *messages.BroadcastMsgChan) {
	if t.broadcastChannel == nil {
		t.logger.Warn().Msg("broadcast channel is not set")
		return
	}
	t.broadcastChannel <- broadcastMsg
}
func (t *TssCommon) GetConf() TssConfig {
	return t.conf
}

func (t *TssCommon) GetTaskDone() chan struct{} {
	return t.taskDone
}

func (t *TssCommon) GetThresHold() int {
	return t.thresHold
}

func (t *TssCommon) GetAbnormalMgr() *abnormal2.Manager {
	return t.abnormalMgr
}

func (t *TssCommon) SetPartyInfo(partyInfo *abnormal2.PartyInfo) {
	t.partyLock.Lock()
	defer t.partyLock.Unlock()
	t.partyInfo = partyInfo
}

func (t *TssCommon) getPartyInfo() *abnormal2.PartyInfo {
	t.partyLock.Lock()
	defer t.partyLock.Unlock()
	return t.partyInfo
}

func (t *TssCommon) GetLocalPeerID() string {
	return t.localPeerID
}

func (t *TssCommon) SetLocalPeerID(peerID string) {
	t.localPeerID = peerID
}

func (t *TssCommon) SetParties(parties []*tss.PartyID) {
	t.parties = parties
}

func (t *TssCommon) processInvalidMsg(roundInfo string, round abnormal2.RoundInfo, err *tss.Error) error {
	// now we get the culprits ID, invalid message and signature the culprits sent
	var culpritsID []string
	var invalidMsgs []*messages.WireMessage
	unicast := checkUnicast(round)
	t.culpritsLock.Lock()
	t.culprits = append(t.culprits, err.Culprits()...)
	t.culpritsLock.Unlock()
	for _, el := range err.Culprits() {
		culpritsID = append(culpritsID, el.Id)
		key := fmt.Sprintf("%s-%s", el.Id, roundInfo)
		storedMsg := t.abnormalMgr.GetRoundMgr().Get(key)
		invalidMsgs = append(invalidMsgs, storedMsg)
	}
	pubkeys, errBlame := conversion.AccPubKeysFromPartyIDs(culpritsID, t.partyInfo.PartyIDMap)
	if errBlame != nil {
		t.logger.Error().Err(err.Cause()).Msgf("error in get the blame nodes")
		t.abnormalMgr.GetAbnormal().SetAbnormal(abnormal2.TssBrokenMsg, nil, unicast)
		return fmt.Errorf("error in getting the blame nodes")
	}
	// This error indicates the share is wrong, we include this signature to prove that
	// this incorrect share is from the share owner.
	var blameNodes []abnormal2.Node
	var msgBody, sig []byte
	for i, pk := range pubkeys {
		invalidMsg := invalidMsgs[i]
		if invalidMsg == nil {
			t.logger.Error().Msg("we cannot find the record of this curlprit, set it as blank")
			msgBody = []byte{}
			sig = []byte{}
		} else {
			msgBody = invalidMsg.Message
			sig = invalidMsg.Sig
		}
		blameNodes = append(blameNodes, abnormal2.NewNode(pk, msgBody, sig))
	}
	t.abnormalMgr.GetAbnormal().SetAbnormal(abnormal2.TssBrokenMsg, blameNodes, unicast)
	return fmt.Errorf("fail to set bytes to local party: %w", err)
}

func generateSignature(msg []byte, msgID string, privKey *ecdsa.PrivateKey) ([]byte, error) {
	var dataForSigning bytes.Buffer
	dataForSigning.Write(msg)
	dataForSigning.WriteString(msgID)
	digestBz := crypto.Keccak256Hash(dataForSigning.Bytes()).Bytes()
	return crypto.Sign(digestBz, privKey)
}

func verifySignature(pubKey, message, sig []byte, msgID string) bool {
	var dataForSign bytes.Buffer
	dataForSign.Write(message)
	dataForSign.WriteString(msgID)
	digestBz := crypto.Keccak256Hash(dataForSign.Bytes()).Bytes()
	if len(sig) == crypto.SignatureLength {
		sig = sig[:len(sig)-1]
	}
	return crypto.VerifySignature(pubKey, digestBz, sig)
}

func (t *TssCommon) updateLocal(wireMsg *messages.WireMessage) error {
	if wireMsg == nil || wireMsg.Routing == nil || wireMsg.Routing.From == nil {
		t.logger.Warn().Msg("wire msg is nil")
		return errors.New("invalid wireMsg")
	}
	partyInfo := t.getPartyInfo()
	if partyInfo == nil {
		return nil
	}
	partyID, ok := partyInfo.PartyIDMap[wireMsg.Routing.From.Id]
	if !ok {
		return fmt.Errorf("get message from unknown party %s", partyID.Id)
	}

	dataOwnerPeerID, ok := t.PartyIDtoP2PID[wireMsg.Routing.From.Id]
	if !ok {
		t.logger.Error().Msg("fail to find the peer ID of this party")
		return errors.New("fail to find the peer")
	}
	// here we log down this peer as the latest unicast peer
	if !wireMsg.Routing.IsBroadcast {
		t.abnormalMgr.SetLastUnicastPeer(dataOwnerPeerID, wireMsg.RoundInfo)
	}

	var bulkMsg BulkWireMsg
	err := json.Unmarshal(wireMsg.Message, &bulkMsg)
	if err != nil {
		t.logger.Error().Err(err).Msg("error to unmarshal the BulkMsg")
		return err
	}

	tssJobChan := make(chan *tssJob)
	jobWg := sync.WaitGroup{}
	jobWg.Add(1)
	go t.doTssJob(tssJobChan, &jobWg)

	//if partyInfo.Party.PartyID().Id != bulkMsg.MsgIdentifier {
	//	t.logger.Error().Msg("cannot find the party to this wired msg")
	//	return errors.New("cannot find the party")
	//}
	localMsgParty := partyInfo.Party
	rPartyID, ok := partyInfo.PartyIDMap[bulkMsg.Routing.From.Id]
	if !ok {
		t.logger.Error().Msg("error in find the partyID")
		return errors.New("cannot find the party to handle the message")
	}

	round, err := GetMsgRound(bulkMsg.WiredBulkMsg, rPartyID, bulkMsg.Routing.IsBroadcast)
	if err != nil {
		t.logger.Error().Err(err).Msg("broken tss share")
		return err
	}

	// we only allow a message be updated only once.
	// here we use round + msgIdentifier as the key for the acceptedShares
	if bulkMsg.MsgIdentifier == "" {
		round.MsgIdentifier = bulkMsg.Routing.From.Id
	} else {
		round.MsgIdentifier = bulkMsg.MsgIdentifier
	}
	// if this share is duplicated, we skip this share
	if t.abnormalMgr.CheckMsgDuplication(round, partyID.Id) {
		t.logger.Debug().Msgf("we received the duplicated message from party %s", partyID.Id)
	} else {
		partyInlist := func(el *tss.PartyID, l []*tss.PartyID) bool {
			for _, each := range l {
				if el == each {
					return true
				}
			}
			return false
		}
		t.culpritsLock.RLock()
		if len(t.culprits) != 0 && partyInlist(partyID, t.culprits) {
			t.logger.Error().Msgf("the malicious party (party ID:%s) try to send incorrect message to me (party ID:%s)", partyID.Id, localMsgParty.PartyID().Id)
			t.culpritsLock.RUnlock()
			return errors.New("tss share verification failed")
		}
		t.culpritsLock.RUnlock()
		job := newJob(localMsgParty, bulkMsg.WiredBulkMsg, round.MsgIdentifier, partyID, bulkMsg.Routing.IsBroadcast)
		tssJobChan <- job
	}

	close(tssJobChan)
	jobWg.Wait()
	return nil
}

func (t *TssCommon) sendBulkMsg(wiredMsgType string, tssMsgType messages.TSSMessageTpe, wiredMsg BulkWireMsg) error {
	// since all the messages in the list is the same round, so it must have the same dest
	// we just need to get the routing info of the first message
	r := wiredMsg.Routing

	buf, err := json.Marshal(wiredMsg)
	if err != nil {
		return fmt.Errorf("error in marshal the cachedWireMsg: %w", err)
	}

	sig, err := generateSignature(buf, t.msgID, t.privateKey)
	if err != nil {
		t.logger.Error().Err(err).Msg("fail to generate the share's signature")
		return err
	}

	wireMsg := messages.WireMessage{
		Routing:   r,
		RoundInfo: wiredMsgType,
		Message:   buf,
		Sig:       sig,
	}
	wireMsgBytes, err := json.Marshal(wireMsg)
	if err != nil {
		return fmt.Errorf("fail to convert tss msg to wire bytes: %w", err)
	}
	wrappedMsg := messages.WrappedMessage{
		MsgID:       t.msgID,
		MessageType: tssMsgType,
		Payload:     wireMsgBytes,
	}

	peerIDs := make([]peer.ID, 0)
	if len(r.To) == 0 {
		t.P2PPeersLock.RLock()
		peerIDs = t.P2PPeers
		t.P2PPeersLock.RUnlock()
	} else {
		for _, each := range r.To {
			peerID, ok := t.PartyIDtoP2PID[each.Id]
			if !ok {
				t.logger.Error().Msg("error in find the P2P ID")
				continue
			}
			peerIDs = append(peerIDs, peerID)
		}
	}
	t.renderToP2P(&messages.BroadcastMsgChan{
		WrappedMessage: wrappedMsg,
		PeersID:        peerIDs,
	})

	return nil
}

func (t *TssCommon) ProcessOutCh(msg tss.Message, msgType messages.TSSMessageTpe) error {
	msgData, r, err := msg.WireBytes()
	// if we cannot get the wire share, the tss will fail, we just quit.
	if err != nil {
		return fmt.Errorf("fail to get wire bytes: %w", err)
	}

	if r.IsBroadcast {
		cachedWiredMsg := NewBulkWireMsg(msgData, msg.GetFrom().Moniker, r)
		// now we store this message in cache
		_, ok := t.cachedWireBroadcastMsgLists.Load(msg.Type())
		if !ok {
			t.cachedWireBroadcastMsgLists.Store(msg.Type(), cachedWiredMsg)
		}
	} else {
		cachedWiredMsg := NewBulkWireMsg(msgData, msg.GetFrom().Moniker, r)
		_, ok := t.cachedWireUnicastMsgLists.Load(msg.Type() + ":" + r.To[0].String())
		if !ok {
			t.cachedWireUnicastMsgLists.Store(msg.Type()+":"+r.To[0].String(), cachedWiredMsg)
		}
	}
	t.cachedWireUnicastMsgLists.Range(func(key, value interface{}) bool {
		wiredMsg := value.(BulkWireMsg)
		ret := strings.Split(key.(string), ":")
		wiredMsgType := ret[0]
		err := t.sendBulkMsg(wiredMsgType, msgType, wiredMsg)
		if err != nil {
			t.logger.Error().Err(err).Msg("error in send bulk message")
			return true
		}
		t.cachedWireUnicastMsgLists.Delete(key)

		return true
	})

	t.cachedWireBroadcastMsgLists.Range(func(key, value interface{}) bool {
		wiredMsg := value.(BulkWireMsg)
		wiredMsgType := key.(string)

		err := t.sendBulkMsg(wiredMsgType, msgType, wiredMsg)
		if err != nil {
			t.logger.Error().Err(err).Msg("error in send bulk message")
			return true
		}
		t.cachedWireBroadcastMsgLists.Delete(key)

		return true
	})

	return nil
}

func (t *TssCommon) applyShare(localCacheItem *LocalCacheItem, key string, msgType messages.TSSMessageTpe) error {

	t.abnormalMgr.GetRoundMgr().Set(key, localCacheItem.Msg)
	if err := t.updateLocal(localCacheItem.Msg); nil != err {
		return fmt.Errorf("fail to update the message to local party: %w", err)
	}
	t.logger.Debug().Msgf("remove key: %s", key)
	// the information had been confirmed by all party , we don't need it anymore
	t.removeKey(key)
	return nil
}

func (t *TssCommon) removeKey(key string) {
	t.unConfirmedMsgLock.Lock()
	defer t.unConfirmedMsgLock.Unlock()
	delete(t.unConfirmedMessages, key)
}

func (t *TssCommon) hashCheck(localCacheItem *LocalCacheItem, threshold int) error {
	dataOwner := localCacheItem.Msg.Routing.From
	dataOwnerP2PID, ok := t.PartyIDtoP2PID[dataOwner.Id]
	if !ok {
		t.logger.Warn().Msgf("error in find the data Owner P2PID\n")
		return errors.New("error in find the data Owner P2PID")
	}

	if localCacheItem.TotalConfirmParty() < threshold {
		t.logger.Debug().Msg("not enough nodes to evaluate the hash")
		return abnormal2.ErrNotEnoughPeer
	}
	localCacheItem.lock.Lock()
	defer localCacheItem.lock.Unlock()

	targetHashValue := localCacheItem.Hash
	for P2PID := range localCacheItem.ConfirmedList {
		if P2PID == dataOwnerP2PID.String() {
			t.logger.Warn().Msgf("we detect that the data owner try to send the hash for his own message\n")
			delete(localCacheItem.ConfirmedList, P2PID)
			return abnormal2.ErrHashFromOwner
		}
	}
	hash, err := t.getMsgHash(localCacheItem, threshold)
	if err != nil {
		return err
	}
	if targetHashValue == hash {
		t.logger.Debug().Msgf("hash check complete for messageID: %v", t.msgID)
		return nil
	}
	return abnormal2.ErrNotMajority
}

func (t *TssCommon) getMsgHash(localCacheItem *LocalCacheItem, threshold int) (string, error) {
	hash, freq, err := getHighestFreq(localCacheItem.ConfirmedList)
	if err != nil {
		t.logger.Error().Err(err).Msg("fail to get the hash freq")
		return "", abnormal2.ErrHashCheck
	}
	if freq < threshold-1 {
		t.logger.Debug().Msgf("fail to have more than 2/3 peers agree on the received message threshold(%d)--total confirmed(%d)\n", threshold, freq)
		return "", abnormal2.ErrHashInconsistency
	}
	return hash, nil
}

func (t *TssCommon) ProcessInboundMessages(finishChan chan struct{}, wg *sync.WaitGroup) {
	t.logger.Debug().Msg("start processing inbound messages")
	defer wg.Done()
	defer t.logger.Debug().Msg("stop processing inbound messages")
	for {
		select {
		case <-finishChan:
			return
		case m, ok := <-t.TssMsg:
			if !ok {
				return
			}
			var wrappedMsg messages.WrappedMessage
			if err := json.Unmarshal(m.Payload, &wrappedMsg); nil != err {
				t.logger.Error().Err(err).Msg("fail to unmarshal wrapped message bytes")
				continue
			}

			err := t.ProcessOneMessage(&wrappedMsg, m.PeerID.String())
			if err != nil {
				t.logger.Error().Err(err).Msg("fail to process the received message")
			}

		}
	}
}

func (t *TssCommon) ProcessOneMessage(wrappedMsg *messages.WrappedMessage, peerID string) error {
	t.logger.Debug().Msg("start process one message")
	defer t.logger.Debug().Msg("finish processing one message")
	if nil == wrappedMsg {
		return errors.New("invalid wireMessage")
	}

	switch wrappedMsg.MessageType {
	case messages.TSSKeyGenMsg, messages.TSSKeySignMsg:
		var wireMsg messages.WireMessage
		if err := json.Unmarshal(wrappedMsg.Payload, &wireMsg); nil != err {
			return fmt.Errorf("fail to unmarshal wire message: %w", err)
		}
		return t.processTSSMsg(&wireMsg, wrappedMsg.MessageType, true)
	case messages.TSSTaskDone:
		var wireMsg messages.TssTaskNotifier
		err := json.Unmarshal(wrappedMsg.Payload, &wireMsg)
		if err != nil {
			t.logger.Error().Err(err).Msg("fail to unmarshal the notify message")
			return nil
		}
		if wireMsg.TaskDone {
			// if we have already logged this node, we return to avoid close of a close channel
			if t.finishedPeers[peerID] {
				return fmt.Errorf("duplicated notification from peer %s ignored", peerID)
			}
			t.finishedPeers[peerID] = true
			if len(t.finishedPeers) == len(t.partyInfo.PartyIDMap)-1 {
				t.logger.Debug().Msg("we get the confirm of the nodes that generate the signature")
				close(t.taskDone)
			}
			return nil
		}
	}

	return nil
}

func (t *TssCommon) processTSSMsg(wireMsg *messages.WireMessage, msgType messages.TSSMessageTpe, forward bool) error {
	t.logger.Debug().Msg("process wire message")
	defer t.logger.Debug().Msg("finish process wire message")

	if wireMsg == nil || wireMsg.Routing == nil || wireMsg.Routing.From == nil {
		t.logger.Warn().Msg("received msg invalid, msg type  " + msgType.String())
		if wireMsg == nil {
			t.logger.Warn().Msg("wireMsg is nil ! ")
		} else if wireMsg.Routing == nil {
			t.logger.Warn().Msg("wireMsg Routing is nil ! ")
		} else if wireMsg.Routing.From == nil {
			t.logger.Warn().Msg("wireMsg Routing From is nil ! ")
		}
		return errors.New("invalid wireMsg")
	}
	partyIDMap := t.getPartyInfo().PartyIDMap
	dataOwner, ok := partyIDMap[wireMsg.Routing.From.Id]
	if !ok {
		t.logger.Error().Msg("error in find the data owner")
		return errors.New("error in find the data owner")
	}
	keyBytes := dataOwner.GetKey()

	ok = verifySignature(keyBytes, wireMsg.Message, wireMsg.Sig, t.msgID)
	if !ok {
		t.logger.Error().Msg("fail to verify the signature")
		return errors.New("signature verify failed")
	}

	// for the unicast message, we only update it local party
	if !wireMsg.Routing.IsBroadcast {
		t.logger.Debug().Msgf("msg from %s to %+v", wireMsg.Routing.From, wireMsg.Routing.To)
		return t.updateLocal(wireMsg)
	}

	key := wireMsg.GetCacheKey()
	msgHash, err := conversion.BytesToHashString(wireMsg.Message)
	if err != nil {
		return fmt.Errorf("fail to calculate hash of the wire message: %w", err)
	}
	localCacheItem := t.TryGetLocalCacheItem(key)
	if nil == localCacheItem {
		t.logger.Debug().Msgf("++%s doesn't exist yet,add a new one", key)
		localCacheItem = NewLocalCacheItem(wireMsg, msgHash)
		t.updateLocalUnconfirmedMessages(key, localCacheItem)
	} else {
		// this means we received the broadcast confirm message from other party first
		t.logger.Debug().Msgf("==%s exist", key)
		if localCacheItem.Msg == nil {
			t.logger.Debug().Msgf("==%s exist, set message", key)
			localCacheItem.Msg = wireMsg
			localCacheItem.Hash = msgHash
		}
	}
	localCacheItem.UpdateConfirmList(t.localPeerID, msgHash)

	if err != nil {
		return err
	}
	return t.applyShare(localCacheItem, key, msgType)
}

func (t *TssCommon) TryGetLocalCacheItem(key string) *LocalCacheItem {
	t.unConfirmedMsgLock.Lock()
	defer t.unConfirmedMsgLock.Unlock()
	localCacheItem, ok := t.unConfirmedMessages[key]
	if !ok {
		return nil
	}
	return localCacheItem
}

func (t *TssCommon) updateLocalUnconfirmedMessages(key string, cacheItem *LocalCacheItem) {
	t.unConfirmedMsgLock.Lock()
	defer t.unConfirmedMsgLock.Unlock()
	t.unConfirmedMessages[key] = cacheItem
}

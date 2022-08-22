package tsslib

import (
	"github.com/bitdao-io/bitnetwork/tss/node/tsslib/abnormal"
	"github.com/bitdao-io/bitnetwork/tss/node/tsslib/common"
	conversion2 "github.com/bitdao-io/bitnetwork/tss/node/tsslib/conversion"
	keygen2 "github.com/bitdao-io/bitnetwork/tss/node/tsslib/keygen"
	"github.com/bitdao-io/bitnetwork/tss/node/tsslib/messages"
	"time"
)

func (t *TssServer) Keygen(req keygen2.Request) (keygen2.Response, error) {
	t.tssKeyGenLocker.Lock()
	defer t.tssKeyGenLocker.Unlock()
	status := common.Success
	msgID, err := t.requestToMsgId(req)
	if err != nil {
		return keygen2.Response{}, err
	}
	if err = t.requestCheck(req); err != nil {
		return keygen2.Response{}, err
	}

	keygenInstance := keygen2.NewTssKeyGen(
		t.p2pCommunication.GetLocalPeerID(),
		t.conf,
		t.localNodePubKey,
		t.p2pCommunication.BroadcastMsgChan,
		t.stopChan,
		t.preParams,
		msgID,
		t.stateManager,
		t.secretsEnable,
		t.secretsManager,
		t.shamirEnable,
		t.shamirManager,
		t.privateKey,
		t.p2pCommunication,
		req.ThresHold,
	)

	keygenMsgChannel := keygenInstance.GetTssKeyGenChannels()
	t.p2pCommunication.SetSubscribe(messages.TSSKeyGenMsg, msgID, keygenMsgChannel)
	t.p2pCommunication.SetSubscribe(messages.TSSTaskDone, msgID, keygenMsgChannel)

	defer func() {
		t.p2pCommunication.CancelSubscribe(messages.TSSKeyGenMsg, msgID)
		t.p2pCommunication.CancelSubscribe(messages.TSSTaskDone, msgID)

		t.p2pCommunication.ReleaseStream(msgID)
	}()
	abnormalMgr := keygenInstance.GetTssCommonStruct().GetAbnormalMgr()

	parsedPeers := make([]string, len(req.Keys))
	for i, el := range req.Keys {
		parsedPeers[i] = el
	}
	actives, err := conversion2.GetPubKeysFromPeerIDs(parsedPeers)
	if err != nil {
		t.logger.Error().Err(err).Msg("there is a error when get active public keys from peers id")
	} else {
		keygenInstance.ParticipantKeys = actives
	}

	t.logger.Debug().Msg("keygen party formed")
	// the statistic of keygen only care about Tss it self, even if the
	// following http response aborts, it still counted as a successful keygen
	// as the Tss model runs successfully.
	beforeKeygen := time.Now()
	k, err := keygenInstance.GenerateNewKey(req)
	keygenTime := time.Since(beforeKeygen)
	if err != nil {
		t.tssMetrics.UpdateKeyGen(keygenTime, false)
		t.logger.Error().Err(err).Msg("err in keygen")

		return keygen2.NewResponse(
			"", nil, common.Fail,
			abnormal.GenerateNewKeyError,
			abnormalMgr.GetAbnormalNodePubKeys()), err
	} else {
		t.tssMetrics.UpdateKeyGen(keygenTime, true)
	}

	pubkey, address, err := conversion2.GetTssPubKey(k)
	if err != nil {
		return keygen2.NewResponse(
			"",
			nil,
			common.Fail,
			abnormal.GenerateNewKeyError,
			abnormalMgr.GetAbnormalNodePubKeys()), err
	}

	return keygen2.NewResponse(
		pubkey,
		address,
		status,
		"",
		abnormalMgr.GetAbnormalNodePubKeys(),
	), nil
}

package common

import (
	"crypto/elliptic"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/binance-chain/tss-lib/ecdsa/keygen"
	"github.com/binance-chain/tss-lib/ecdsa/signing"
	"github.com/binance-chain/tss-lib/tss"
	"github.com/btcsuite/btcd/btcec"
	"github.com/mantlenetworkio/mantle/tss/node/tsslib/abnormal"
	messages2 "github.com/mantlenetworkio/mantle/tss/node/tsslib/messages"
	"math/big"
	"strings"
)

func GetMsgRound(msg []byte, partyID *tss.PartyID, isBroadcast bool) (abnormal.RoundInfo, error) {
	parsedMsg, err := tss.ParseWireMessage(msg, partyID, isBroadcast)
	if err != nil {
		return abnormal.RoundInfo{}, err
	}
	switch parsedMsg.Content().(type) {
	case *keygen.KGRound1Message:
		return abnormal.RoundInfo{
			Index:    0,
			RoundMsg: messages2.KEYGEN1,
		}, nil

	case *keygen.KGRound2Message1:
		return abnormal.RoundInfo{
			Index:    1,
			RoundMsg: messages2.KEYGEN2aUnicast,
		}, nil

	case *keygen.KGRound2Message2:
		return abnormal.RoundInfo{
			Index:    2,
			RoundMsg: messages2.KEYGEN2b,
		}, nil

	case *keygen.KGRound3Message:
		return abnormal.RoundInfo{
			Index:    3,
			RoundMsg: messages2.KEYGEN3,
		}, nil

	case *signing.SignRound1Message1:
		return abnormal.RoundInfo{
			Index:    0,
			RoundMsg: messages2.KEYSIGN1aUnicast,
		}, nil

	case *signing.SignRound1Message2:
		return abnormal.RoundInfo{
			Index:    1,
			RoundMsg: messages2.KEYSIGN1b,
		}, nil

	case *signing.SignRound2Message:
		return abnormal.RoundInfo{
			Index:    2,
			RoundMsg: messages2.KEYSIGN2Unicast,
		}, nil

	case *signing.SignRound3Message:
		return abnormal.RoundInfo{
			Index:    3,
			RoundMsg: messages2.KEYSIGN3,
		}, nil

	case *signing.SignRound4Message:
		return abnormal.RoundInfo{
			Index:    4,
			RoundMsg: messages2.KEYSIGN4,
		}, nil

	case *signing.SignRound5Message:
		return abnormal.RoundInfo{
			Index:    5,
			RoundMsg: messages2.KEYSIGN5,
		}, nil

	case *signing.SignRound6Message:
		return abnormal.RoundInfo{
			Index:    6,
			RoundMsg: messages2.KEYSIGN6,
		}, nil

	case *signing.SignRound7Message:
		return abnormal.RoundInfo{
			Index:    7,
			RoundMsg: messages2.KEYSIGN7,
		}, nil

	case *signing.SignRound8Message:
		return abnormal.RoundInfo{
			Index:    8,
			RoundMsg: messages2.KEYSIGN8,
		}, nil

	case *signing.SignRound9Message:
		return abnormal.RoundInfo{
			Index:    9,
			RoundMsg: messages2.KEYSIGN9,
		}, nil

	default:
		return abnormal.RoundInfo{}, errors.New("unknown round")
	}
}

// due to the nature of tss, we may find the invalid share of the previous round only
// when we get the shares from the peers in the current round. So, when we identify
// an error in this round, we check whether the previous round is the unicast
func checkUnicast(round abnormal.RoundInfo) bool {
	index := round.Index
	isKeyGen := strings.Contains(round.RoundMsg, "KGR")
	// keygen unicast blame
	if isKeyGen {
		if index == 1 || index == 2 {
			return true
		}
		return false
	}
	// keysign unicast blame
	if index < 5 {
		return true
	}
	return false
}

func (t *TssCommon) NotifyTaskDone() error {
	msg := messages2.TssTaskNotifier{TaskDone: true}
	data, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("fail to marshal the request body %w", err)
	}
	wrappedMsg := messages2.WrappedMessage{
		MessageType: messages2.TSSTaskDone,
		MsgID:       t.msgID,
		Payload:     data,
	}
	t.P2PPeersLock.RLock()
	peers := t.P2PPeers
	t.P2PPeersLock.RUnlock()
	t.renderToP2P(&messages2.BroadcastMsgChan{
		WrappedMessage: wrappedMsg,
		PeersID:        peers,
	})
	return nil
}

func getHighestFreq(confirmedList map[string]string) (string, int, error) {
	if len(confirmedList) == 0 {
		return "", 0, errors.New("empty input")
	}
	freq := make(map[string]int, len(confirmedList))
	for _, n := range confirmedList {
		freq[n]++
	}
	maxFreq := -1
	var data string
	for key, counter := range freq {
		if counter > maxFreq {
			maxFreq = counter
			data = key
		}
	}
	return data, maxFreq, nil
}

func Contains(s []*tss.PartyID, e *tss.PartyID) bool {
	if e == nil {
		return false
	}
	for _, a := range s {
		if *a == *e {
			return true
		}
	}
	return false
}

func MsgToHashInt(msg []byte) (*big.Int, error) {
	return hashToInt(msg, btcec.S256()), nil
}

func hashToInt(hash []byte, c elliptic.Curve) *big.Int {
	orderBits := c.Params().N.BitLen()
	orderBytes := (orderBits + 7) / 8
	if len(hash) > orderBytes {
		hash = hash[:orderBytes]
	}

	ret := new(big.Int).SetBytes(hash)
	excess := len(hash)*8 - orderBits
	if excess > 0 {
		ret.Rsh(ret, uint(excess))
	}
	return ret
}

func MsgToHashString(msg []byte) (string, error) {
	if len(msg) == 0 {
		return "", errors.New("empty message")
	}
	h := sha256.New()
	_, err := h.Write(msg)
	if err != nil {
		return "", fmt.Errorf("fail to caculate sha256 hash: %w", err)
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

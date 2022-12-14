package conversion

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/binance-chain/tss-lib/crypto"
	"github.com/binance-chain/tss-lib/tss"
	"github.com/btcsuite/btcd/btcec"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	crypto2 "github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/peer"
	"math/big"
	"sort"
	"strconv"
)

func GetParties(keys []string, localPartyKey string) ([]*tss.PartyID, *tss.PartyID, error) {
	var localPartyID *tss.PartyID
	var unSortedPartiesID []*tss.PartyID
	sort.Strings(keys)
	for idx, item := range keys {
		pkBytes, err := hex.DecodeString(item)
		if err != nil {
			return nil, nil, fmt.Errorf("fail to get account pub key (%s): %w", item, err)
		}
		key := new(big.Int).SetBytes(pkBytes)
		// Set up the parameters
		// Note: The `id` and `moniker` fields are for convenience to allow you to easily track participants.
		// The `id` should be a unique string representing this party in the network and `moniker` can be anything (even left blank).
		// The `uniqueKey` is a unique identifying key for this peer (such as its p2p public key) as a big.Int.
		partyID := tss.NewPartyID(strconv.Itoa(idx), "", key)
		if item == localPartyKey {
			localPartyID = partyID
		}
		unSortedPartiesID = append(unSortedPartiesID, partyID)
	}
	if localPartyID == nil {
		return nil, nil, errors.New("local party is not in the list")
	}

	partiesID := tss.SortPartyIDs(unSortedPartiesID)
	return partiesID, localPartyID, nil
}

func SetupPartyIDMap(partiesID []*tss.PartyID) map[string]*tss.PartyID {
	partyIDMap := make(map[string]*tss.PartyID)
	for _, id := range partiesID {
		partyIDMap[id.Id] = id
	}
	return partyIDMap
}

func SetupIDMaps(parties map[string]*tss.PartyID, partyIDtoP2PID map[string]peer.ID) error {
	for id, party := range parties {
		peerID, err := GetPeerIDFromPartyID(party)
		if err != nil {
			return err
		}
		partyIDtoP2PID[id] = peerID
	}
	return nil
}

func GetPeerIDFromPartyID(partyID *tss.PartyID) (peer.ID, error) {
	if partyID == nil || !partyID.ValidateBasic() {
		return "", errors.New("invalid partyID")
	}
	pkBytes := partyID.KeyInt().Bytes()

	return GetPeerIDFromSecp256PubKey(pkBytes)
}

func GetPeerIDFromSecp256PubKey(pk []byte) (peer.ID, error) {
	if len(pk) == 0 {
		return "", errors.New("empty public key raw bytes")
	}
	ppk, err := crypto2.UnmarshalSecp256k1PublicKey(pk)
	if err != nil {
		return "", fmt.Errorf("fail to convert pubkey to the crypto pubkey used in libp2p: %w", err)
	}

	return peer.IDFromPublicKey(ppk)
}

func GetPeersID(partyIDtoP2PID map[string]peer.ID, localPeerID string) []peer.ID {
	if partyIDtoP2PID == nil {
		return nil
	}
	peerIDs := make([]peer.ID, 0, len(partyIDtoP2PID)-1)
	for _, value := range partyIDtoP2PID {
		if value.String() == localPeerID {
			continue
		}
		peerIDs = append(peerIDs, value)
	}
	return peerIDs
}

func BytesToHashString(msg []byte) (string, error) {
	h := sha256.New()
	_, err := h.Write(msg)
	if err != nil {
		return "", fmt.Errorf("fail to caculate sha256 hash: %w", err)
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func GetTssPubKey(pubKeyPoint *crypto.ECPoint) (string, []byte,[]byte, error) {
	// we check whether the point is on curve according to Kudelski report
	if pubKeyPoint == nil || !isOnCurve(pubKeyPoint.X(), pubKeyPoint.Y()) {
		return "", nil,nil, errors.New("invalid points")
	}
	tssPubKey := btcec.PublicKey{
		Curve: btcec.S256(),
		X:     pubKeyPoint.X(),
		Y:     pubKeyPoint.Y(),
	}

	pubKeyStr := hex.EncodeToString(tssPubKey.SerializeCompressed())
	address := ethcrypto.PubkeyToAddress(*tssPubKey.ToECDSA()).Bytes()

	pubKeyBytes := tssPubKey.SerializeUncompressed()
	pubKeyBytes = pubKeyBytes[1:]
	return pubKeyStr, address,pubKeyBytes, nil
}

func PartyIDtoPubKey(party *tss.PartyID) (string, error) {
	if party == nil || !party.ValidateBasic() {
		return "", errors.New("invalid party")
	}
	partyKeyBytes := party.GetKey()
	pubKey := hex.EncodeToString(partyKeyBytes)

	return pubKey, nil
}

func AccPubKeysFromPartyIDs(partyIDs []string, partyIDMap map[string]*tss.PartyID) ([]string, error) {
	pubKeys := make([]string, 0)
	for _, partyID := range partyIDs {
		blameParty, ok := partyIDMap[partyID]
		if !ok {
			return nil, errors.New("cannot find the blame party")
		}
		blamedPubKey, err := PartyIDtoPubKey(blameParty)
		if err != nil {
			return nil, err
		}
		pubKeys = append(pubKeys, blamedPubKey)
	}
	return pubKeys, nil
}

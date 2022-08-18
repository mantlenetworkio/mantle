package conversion

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/btcsuite/btcd/btcec"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
)

func CheckKeyOnCurve(pk string) (bool, error) {
	pubKey, err := hex.DecodeString(pk)
	if err != nil {
		return false, fmt.Errorf("fail to parse pub key(%s): %w", pk, err)
	}
	bPk, err := btcec.ParsePubKey(pubKey, btcec.S256())
	if err != nil {
		return false, err
	}
	return isOnCurve(bPk.X, bPk.Y), nil
}

func isOnCurve(x, y *big.Int) bool {
	curve := btcec.S256()
	return curve.IsOnCurve(x, y)
}

func GetPeerIDsFromPubKeys(pubkeys []string) ([]peer.ID, error) {
	var peerIDs []peer.ID
	for _, item := range pubkeys {
		peerID, err := GetPeerIDFromPubKey(item)
		if err != nil {
			return nil, err
		}
		peerIDs = append(peerIDs, peerID)
	}
	return peerIDs, nil
}

func GetPeerIDFromPubKey(pubkey string) (peer.ID, error) {
	pubKeyBytes, err := hex.DecodeString(pubkey)
	if err != nil {
		return "", fmt.Errorf("fail to hex decode account pub key(%s): %w", pubkey, err)
	}
	ppk, err := crypto.UnmarshalSecp256k1PublicKey(pubKeyBytes)
	if err != nil {
		return "", fmt.Errorf("fail to convert pubkey to the crypto pubkey used in libp2p: %w", err)
	}
	return peer.IDFromPublicKey(ppk)
}

func GetPubKeyFromPeerID(pID string) (string, error) {
	peerID, err := peer.Decode(pID)
	if err != nil {
		return "", fmt.Errorf("fail to decode peer id: %w", err)
	}
	pk, err := peerID.ExtractPublicKey()
	if err != nil {
		return "", fmt.Errorf("fail to extract pub key from peer id: %w", err)
	}
	rawBytes, err := pk.Raw()
	if err != nil {
		return "", fmt.Errorf("faail to get pub key raw bytes: %w", err)
	}
	pubKeyStr := hex.EncodeToString(rawBytes)
	return pubKeyStr, nil
}

func GetPubKeysFromPeerIDs(peers []string) ([]string, error) {
	var result []string
	for _, item := range peers {
		pKey, err := GetPubKeyFromPeerID(item)
		if err != nil {
			return nil, fmt.Errorf("fail to get pubkey from peerID: %w", err)
		}
		result = append(result, pKey)
	}
	return result, nil
}

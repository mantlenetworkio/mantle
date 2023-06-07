// Package etherkeyms utilises the Key Management Service (KMS) from the Google
// Cloud Platform (GCP) for Ethereum operations.
package bsscore

import (
	"context"
	"encoding/asn1"
	"fmt"
	"math/big"

	kms "cloud.google.com/go/kms/apiv1"
	btcecdsa "github.com/btcsuite/btcd/btcec/v2/ecdsa"
	kmspb "google.golang.org/genproto/googleapis/cloud/kms/v1"

	// lots of poor naming in go-ethereum 👾
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// ManagedKey represents a key from the Key Management Service (KMS).
type ManagedKey struct {
	// KMS uses a slash-separated path for identification.
	// This field is read-only.
	KeyName string
	// Each (public) key maps to one address on the blockchain.
	// This field is read-only.
	EthereumAddr common.Address

	Gclient *kms.KeyManagementClient
}

// NewManagedKey executes a fail-fast initialization.
// Key names from the Google cloud are slash-separated paths.
func NewManagedKey(ctx context.Context, client *kms.KeyManagementClient, address string, keyName string) (*ManagedKey, error) {
	addr := common.HexToAddress(address)

	return &ManagedKey{
		KeyName:      keyName,
		EthereumAddr: addr,
	}, nil
}

// NewEthereumTransactor returns a KMS-backed instance. Ctx applies to the
// entire lifespan of the bind.TransactOpts.
func (mk *ManagedKey) NewEthereumTransactorrWithChainID(ctx context.Context, chainID *big.Int) (*bind.TransactOpts, error) {
	if chainID == nil {
		return nil, bind.ErrNoChainID
	}
	signer := types.LatestSignerForChainID(chainID)
	return &bind.TransactOpts{
		Context: ctx,
		From:    mk.EthereumAddr,
		Signer:  mk.NewEthereumSigner(ctx, signer),
	}, nil
}

// NewEthereumTransactor returns a KMS-backed instance. Ctx applies to the
// entire lifespan of the bind.TransactOpts.
func (mk *ManagedKey) NewEthereumTransactor(ctx context.Context, txIdentification types.Signer) *bind.TransactOpts {
	return &bind.TransactOpts{
		Context: ctx,
		From:    mk.EthereumAddr,
		Signer:  mk.NewEthereumSigner(ctx, txIdentification),
	}
}

// NewEthereumSigner returns a KMS-backed instance. Ctx applies to the entire
// lifespan of the bind.SignerFn.
func (mk *ManagedKey) NewEthereumSigner(ctx context.Context, txIdentification types.Signer) bind.SignerFn {
	return func(addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
		if addr != mk.EthereumAddr {
			return nil, bind.ErrNotAuthorized
		}

		// hash the transaction (with Keccak-256 probably)
		sig, err := mk.SignHash(ctx, txIdentification.Hash(tx))
		if err != nil {
			return nil, err
		}

		return tx.WithSignature(txIdentification, sig)
	}
}

// SignHash returns the signature bytes.
func (mk *ManagedKey) SignHash(ctx context.Context, hash common.Hash) ([]byte, error) {
	// resolve a signature
	req := kmspb.AsymmetricSignRequest{
		Name: mk.KeyName,
		// The digest is probably not a standard SHA256.
		// Unclear why the API/client cares anyway. 🤨
		Digest: &kmspb.Digest{
			Digest: &kmspb.Digest_Sha256{
				Sha256: hash[:],
			},
		},
	}
	resp, err := mk.Gclient.AsymmetricSign(ctx, &req)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, fmt.Errorf("Google KMS asymmetric sign operation: %w", err)
	}

	// parse signature
	var params struct{ R, S *big.Int }
	_, err = asn1.Unmarshal(resp.Signature, &params)
	if err != nil {
		return nil, fmt.Errorf("Google KMS asymmetric signature encoding: %w", err)
	}
	var rLen, sLen int // byte size
	if params.R != nil {
		rLen = (params.R.BitLen() + 7) / 8
	}
	if params.S != nil {
		sLen = (params.S.BitLen() + 7) / 8
	}
	if rLen == 0 || rLen > 32 || sLen == 0 || sLen > 32 {
		return nil, fmt.Errorf("Google KMS asymmetric signature with %d-byte r and %d-byte s denied on size", rLen, sLen)
	}

	// Need uncompressed signature with "recovery ID" at end:
	// https://bitcointalk.org/index.php?topic=5249677.0
	// https://ethereum.stackexchange.com/a/53182/39582
	var sig [66]byte // + 1-byte header + 1-byte tailer
	params.R.FillBytes(sig[33-rLen : 33])
	params.S.FillBytes(sig[65-sLen : 65])

	// brute force try includes KMS verification
	var recoverErr error
	for recoveryID := byte(0); recoveryID < 2; recoveryID++ {
		sig[0] = recoveryID + 27 // BitCoin header
		btcsig := sig[:65]       // exclude Ethereum 'v' parameter
		pubKey, _, err := btcecdsa.RecoverCompact(btcsig, hash[:])
		if err != nil {
			recoverErr = err
			continue
		}

		if pubKeyAddr(pubKey.SerializeUncompressed()) == mk.EthereumAddr {
			// sign the transaction
			sig[65] = recoveryID // Ethereum 'v' parameter
			return sig[1:], nil  // exclude BitCoin header
		}
	}
	// recoverErr can be nil, but that's OK
	return nil, fmt.Errorf("Google KMS asymmetric signature address recovery mis: %w", recoverErr)
}

// PubKeyAddr returns the Ethereum address for (uncompressed-)key bytes.
func pubKeyAddr(bytes []byte) common.Address {
	digest := crypto.Keccak256(bytes[1:])
	var addr common.Address
	copy(addr[:], digest[12:])
	return addr
}

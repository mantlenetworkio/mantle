package manager

import (
	"encoding/hex"
	"errors"
	"testing"
	"time"

	"github.com/btcsuite/btcd/btcec"
	"github.com/mantlenetworkio/mantle/l2geth/crypto"
	tss "github.com/mantlenetworkio/mantle/tss/common"
	"github.com/mantlenetworkio/mantle/tss/manager/types"
	"github.com/mantlenetworkio/mantle/tss/ws/server"
	"github.com/stretchr/testify/require"
	tmtypes "github.com/tendermint/tendermint/rpc/jsonrpc/types"
)

func TestSign(t *testing.T) {
	digest, signature, publicKey := mockSign()
	ctx := types.NewContext().
		WithAvailableNodes([]string{"a", "b", "c", "d"}).
		WithApprovers([]string{"a", "b", "c", "d"}).
		WithTssInfo(types.TssCommitteeInfo{
			Threshold:     3,
			ClusterPubKey: publicKey,
		})

	afterMsgSent := func(request server.RequestMsg, respCh chan server.ResponseMsg) error {
		signResp := tss.SignResponse{
			Signature: signature,
		}
		rpcResp := tmtypes.NewRPCSuccessResponse(request.RpcRequest.ID, signResp)
		respCh <- server.ResponseMsg{
			RpcResponse: rpcResp,
			SourceNode:  request.TargetNode,
		}
		return nil
	}
	manager, request := setup(afterMsgSent, nil)
	signResp, culprits, err := manager.sign(ctx, request, digest, tss.SignStateBatch)
	require.NoError(t, err)
	require.Nil(t, culprits)
	require.EqualValues(t, signature, signResp.Signature)

	afterMsgSent = func(request server.RequestMsg, respCh chan server.ResponseMsg) error {
		if request.TargetNode == "d" {
			signResp := tss.SignResponse{
				Signature: signature,
			}
			rpcResp := tmtypes.NewRPCSuccessResponse(request.RpcRequest.ID, signResp)
			respCh <- server.ResponseMsg{
				RpcResponse: rpcResp,
				SourceNode:  request.TargetNode,
			}
		}
		return nil
	}
	manager, request = setup(afterMsgSent, nil)
	signResp, culprits, err = manager.sign(ctx, request, digest, tss.SignStateBatch)
	require.NoError(t, err)
	require.Nil(t, culprits)
	require.EqualValues(t, signature, signResp.Signature)
}

func TestErrorSend(t *testing.T) {
	digest, _, publicKey := mockSign()
	ctx := types.NewContext().
		WithAvailableNodes([]string{"a", "b", "c", "d"}).
		WithApprovers([]string{"a", "b", "c", "d"}).
		WithTssInfo(types.TssCommitteeInfo{
			Threshold:     3,
			ClusterPubKey: publicKey,
		})

	afterMsgSent := func(request server.RequestMsg, respCh chan server.ResponseMsg) error {
		if request.TargetNode == "c" {
			return errors.New("mock error")
		}
		return nil
	}
	manager, request := setup(afterMsgSent, nil)
	signResp, culprits, err := manager.sign(ctx, request, digest, tss.SignStateBatch)
	require.Nil(t, signResp.Signature)
	require.Nil(t, culprits)
	require.NotNil(t, err)
	require.ErrorContains(t, err, "failed to generate signature")
}

func TestWrongSignature(t *testing.T) {
	digest, signature, publicKey := mockSign()
	ctx := types.NewContext().
		WithAvailableNodes([]string{"a", "b", "c", "d"}).
		WithApprovers([]string{"a", "b", "c", "d"}).
		WithTssInfo(types.TssCommitteeInfo{
			Threshold:     3,
			ClusterPubKey: publicKey,
		})
	afterMsgSent := func(request server.RequestMsg, respCh chan server.ResponseMsg) error {
		signature[22] = 0x67 // modify the signature
		signResp := tss.SignResponse{
			Signature: signature,
		}
		rpcResp := tmtypes.NewRPCSuccessResponse(request.RpcRequest.ID, signResp)
		respCh <- server.ResponseMsg{
			RpcResponse: rpcResp,
			SourceNode:  request.TargetNode,
		}
		return nil
	}
	manager, request := setup(afterMsgSent, nil)
	signResp, culprits, err := manager.sign(ctx, request, digest, tss.SignStateBatch)
	require.Nil(t, signResp.Signature)
	require.Nil(t, culprits)
	require.NotNil(t, err)
	require.ErrorContains(t, err, "failed to generate signature")
}

func TestSignTimeout(t *testing.T) {
	ctx := types.NewContext().
		WithAvailableNodes([]string{"a", "b", "c", "d"}).
		WithApprovers([]string{"a", "b", "c", "d"}).
		WithTssInfo(types.TssCommitteeInfo{
			Threshold: 3,
		})
	afterMsgSent := func(request server.RequestMsg, respCh chan server.ResponseMsg) error {
		return nil
	}
	manager, request := setup(afterMsgSent, nil)
	before := time.Now()
	signResp, culprits, err := manager.sign(ctx, request, nil, tss.SignStateBatch)
	require.Nil(t, signResp.Signature)
	require.Nil(t, culprits)
	require.NotNil(t, err)
	require.ErrorContains(t, err, "failed to generate signature")
	cost := time.Now().Sub(before)
	require.True(t, cost.Seconds()-manager.signTimeout.Seconds() >= 0)

	digest, signature, publicKey := mockSign()
	afterMsgSent = func(request server.RequestMsg, respCh chan server.ResponseMsg) error {
		go func() {
			time.Sleep(time.Duration(int64(manager.signTimeout.Seconds()) - 1))
			signResp := tss.SignResponse{
				Signature: signature,
			}
			rpcResp := tmtypes.NewRPCSuccessResponse(request.RpcRequest.ID, signResp)
			respCh <- server.ResponseMsg{
				RpcResponse: rpcResp,
				SourceNode:  request.TargetNode,
			}
		}()
		return nil
	}
	manager, request = setup(afterMsgSent, nil)
	ctx = ctx.WithTssInfo(types.TssCommitteeInfo{
		Threshold:     3,
		ClusterPubKey: publicKey,
	})
	before = time.Now()
	signResp, culprits, err = manager.sign(ctx, request, digest, tss.SignStateBatch)
	require.NoError(t, err)
	require.Nil(t, culprits)
	require.EqualValues(t, signature, signResp.Signature)
	cost = time.Now().Sub(before)
	require.True(t, cost.Seconds()-manager.signTimeout.Seconds() < 0)
}

func TestCulprits(t *testing.T) {
	ctx := types.NewContext().
		WithAvailableNodes([]string{"a", "b", "c", "d"}).
		WithApprovers([]string{"a", "b", "c", "d"}).
		WithTssInfo(types.TssCommitteeInfo{
			Threshold: 3,
		})
	afterMsgSent := func(request server.RequestMsg, respCh chan server.ResponseMsg) error {
		if request.TargetNode != "a" {
			respMsg := server.ResponseMsg{
				RpcResponse: tmtypes.NewRPCErrorResponse(request.RpcRequest.ID, tss.CulpritErrorCode, "find culprits", "a"),
				SourceNode:  request.TargetNode,
			}
			respCh <- respMsg
		} else {
			respMsg := server.ResponseMsg{
				RpcResponse: tmtypes.NewRPCErrorResponse(request.RpcRequest.ID, tss.CulpritErrorCode, "find culprits", "b"),
				SourceNode:  request.TargetNode,
			}
			respCh <- respMsg
		}

		return nil
	}
	manager, request := setup(afterMsgSent, nil)
	signResp, culprits, err := manager.sign(ctx, request, nil, tss.SignStateBatch)
	require.Error(t, err)
	require.ErrorContains(t, err, "failed to generate signature")
	require.Nil(t, signResp.Signature)
	require.Nil(t, culprits)

	ctx = types.NewContext().
		WithAvailableNodes([]string{"a", "b", "c", "d"}).
		WithApprovers([]string{"a", "b", "c", "d"}).
		WithTssInfo(types.TssCommitteeInfo{
			Threshold: 2,
		})

	signResp, culprits, err = manager.sign(ctx, request, nil, tss.SignStateBatch)
	require.Error(t, err)
	require.ErrorContains(t, err, "failed to generate signature")
	require.Nil(t, signResp.Signature)
	require.NotNil(t, culprits)
	require.EqualValues(t, "a", culprits[0])
}

func TestSignSlash(t *testing.T) {
	digest, signature, publicKey := mockSign()
	ctx := types.NewContext().
		WithAvailableNodes([]string{"a", "b", "c", "d"}).
		WithApprovers([]string{"a", "b", "c", "d"}).
		WithTssInfo(types.TssCommitteeInfo{
			Threshold:     3,
			ClusterPubKey: publicKey,
		})

	afterMsgSent := func(request server.RequestMsg, respCh chan server.ResponseMsg) error {
		var SlashTxGasPrice string
		var SlashTxBytes []byte
		switch request.TargetNode {
		case "a":
			SlashTxGasPrice = "100000"
			SlashTxBytes = []byte("a")
		case "b":
			SlashTxGasPrice = "200000"
			SlashTxBytes = []byte("b")
		case "c":
			SlashTxGasPrice = "500000"
			SlashTxBytes = []byte("c")
		case "d":
			SlashTxGasPrice = "300000"
			SlashTxBytes = []byte("d")
		}
		signResp := tss.SignResponse{
			Signature:       signature,
			SlashTxGasPrice: SlashTxGasPrice,
			SlashTxBytes:    SlashTxBytes,
		}
		rpcResp := tmtypes.NewRPCSuccessResponse(request.RpcRequest.ID, signResp)
		respCh <- server.ResponseMsg{
			RpcResponse: rpcResp,
			SourceNode:  request.TargetNode,
		}
		return nil
	}
	manager, request := setup(afterMsgSent, nil)
	signResp, culprits, err := manager.sign(ctx, request, digest, tss.SignSlash)
	require.NoError(t, err)
	require.Nil(t, culprits)
	require.EqualValues(t, signature, signResp.Signature)
	require.EqualValues(t, signResp.SlashTxGasPriceBigInt.Int64(), 500000)
	require.EqualValues(t, signResp.SlashTxBytes, []byte("c"))

	afterMsgSent = func(request server.RequestMsg, respCh chan server.ResponseMsg) error {
		if request.TargetNode == "c" {
			return nil
		}
		var SlashTxGasPrice string
		var SlashTxBytes []byte
		switch request.TargetNode {
		case "a":
			SlashTxGasPrice = "100000"
			SlashTxBytes = []byte("a")
		case "b":
			SlashTxGasPrice = "200000"
			SlashTxBytes = []byte("b")
		case "d":
			SlashTxGasPrice = "300000"
			SlashTxBytes = []byte("d")
		}
		signResp := tss.SignResponse{
			Signature:       signature,
			SlashTxGasPrice: SlashTxGasPrice,
			SlashTxBytes:    SlashTxBytes,
		}
		rpcResp := tmtypes.NewRPCSuccessResponse(request.RpcRequest.ID, signResp)
		respCh <- server.ResponseMsg{
			RpcResponse: rpcResp,
			SourceNode:  request.TargetNode,
		}
		return nil
	}
	manager, request = setup(afterMsgSent, nil)
	before := time.Now()
	signResp, culprits, err = manager.sign(ctx, request, digest, tss.SignSlash)
	require.NoError(t, err)
	require.Nil(t, culprits)
	require.EqualValues(t, signature, signResp.Signature)
	require.EqualValues(t, signResp.SlashTxGasPriceBigInt.Int64(), 300000)
	require.EqualValues(t, signResp.SlashTxBytes, []byte("d"))
	cost := time.Now().Sub(before)
	require.True(t, cost.Seconds()-manager.signTimeout.Seconds() >= 0)
}

func mockSign() (digest []byte, signature []byte, compressedPublicKey string) {
	priK, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}
	pubKey := btcec.PublicKey(priK.PublicKey)
	compressedPublicKey = hex.EncodeToString(pubKey.SerializeCompressed())
	digest = crypto.Keccak256Hash([]byte("testme")).Bytes()
	signature, err = crypto.Sign(digest, priK)
	if err != nil {
		panic(err)
	}
	return digest, signature, compressedPublicKey
}

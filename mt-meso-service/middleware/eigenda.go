package middleware

import (
	"context"
	pb "github.com/Layr-Labs/datalayr/common/interfaces/interfaceRetrieverServer"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mantlenetworkio/mantle/l2geth/log"
	"github.com/mantlenetworkio/mantle/mt-meso-service/bindings"
	"google.golang.org/grpc"
	"math/big"
	"time"
)

type EigenDaConfig struct {
	L1Client          *ethclient.Client
	EigenContractAddr common.Address
	RetrieverSocket   string
	Timeout           time.Duration
}

type EigenDa struct {
	Ctx             context.Context
	Cfg             *EigenDaConfig
	EigenDaContract *bindings.BVMEigenDataLayrChain
	EigenABI        *abi.ABI
}

func NewEigenDa(ctx context.Context, cfg *EigenDaConfig) (*EigenDa, error) {
	eigenContract, err := bindings.NewBVMEigenDataLayrChain(
		common.Address(cfg.EigenContractAddr), cfg.L1Client,
	)
	if err != nil {
		return nil, err
	}
	return &EigenDa{
		Cfg:             cfg,
		Ctx:             ctx,
		EigenDaContract: eigenContract,
	}, nil
}

func (ed EigenDa) GetLatestRollupBatchIndex() (*big.Int, error) {
	return ed.EigenDaContract.RollupBatchIndex(&bind.CallOpts{})
}

func (ed EigenDa) GetRollupStoreByRollupBatchIndex(batchInde *big.Int) (bindings.BVMEigenDataLayrChainRollupStore, error) {
	return ed.EigenDaContract.GetRollupStoreByRollupBatchIndex(&bind.CallOpts{}, batchInde)
}

func (ed *EigenDa) GetBatchTransactionByStoreNumber(storeNumber uint32) ([]byte, error) {
	conn, err := grpc.Dial(ed.Cfg.RetrieverSocket, grpc.WithInsecure())
	if err != nil {
		log.Error("Disperser Cannot connect to %v. %v\n", ed.Cfg.RetrieverSocket, err)
		return nil, err
	}
	defer conn.Close()
	client := pb.NewDataRetrievalClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), ed.Cfg.Timeout)
	defer cancel()
	opt := grpc.MaxCallRecvMsgSize(1024 * 1024 * 300)
	request := &pb.FramesAndDataRequest{
		DataStoreId: storeNumber,
	}
	reply, err := client.RetrieveFramesAndData(ctx, request, opt)
	if err != nil {
		return nil, err
	}
	return reply.GetData(), nil
}

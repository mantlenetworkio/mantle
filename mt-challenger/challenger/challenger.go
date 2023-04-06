package challenger

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	datalayr "github.com/Layr-Labs/datalayr/common/contracts"
	gkzg "github.com/Layr-Labs/datalayr/common/crypto/go-kzg-bn254"
	"github.com/Layr-Labs/datalayr/common/graphView"
	pb "github.com/Layr-Labs/datalayr/common/interfaces/interfaceRetrieverServer"
	"github.com/Layr-Labs/datalayr/common/logging"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethc "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/mantlenetworkio/mantle/l2geth/common"
	l2types "github.com/mantlenetworkio/mantle/l2geth/core/types"
	l2ethclient "github.com/mantlenetworkio/mantle/l2geth/ethclient"
	l2rlp "github.com/mantlenetworkio/mantle/l2geth/rlp"
	"github.com/mantlenetworkio/mantle/l2geth/rollup/eigenda"
	common4 "github.com/mantlenetworkio/mantle/mt-batcher/services/common"
	"github.com/mantlenetworkio/mantle/mt-batcher/txmgr"
	"github.com/mantlenetworkio/mantle/mt-challenger/bindings"
	rc "github.com/mantlenetworkio/mantle/mt-challenger/bindings"
	"github.com/mantlenetworkio/mantle/mt-challenger/challenger/db"
	"github.com/pkg/errors"
	"github.com/shurcooL/graphql"
	"google.golang.org/grpc"
	"math/big"
	"strconv"
	"strings"
	"sync"
	"time"
)

const fraudString = "2d5f2860204f2060295f2d202d5f2860206f2060295f2d202d5f286020512060295f2d2042495444414f204a5553542052454b5420594f55207c5f2860204f2060295f7c202d207c5f2860206f2060295f7c202d207c5f286020512060295f7c"

type SignerFn func(context.Context, ethc.Address, *types.Transaction) (*types.Transaction, error)

var (
	errMaxPriorityFeePerGasNotFound = errors.New(
		"Method eth_maxPriorityFeePerGas not found",
	)
	FallbackGasTipCap = big.NewInt(1500000000)
)

type KzgConfig struct {
	G1Path    string
	G2Path    string
	TableDir  string
	NumWorker int
	Order     uint64 // Order is the total size of SRS
}

type Fraud struct {
	StartingIndex int
}

type DataLayrDisclosureProof struct {
	Header                    []byte
	Polys                     [][]byte
	MultirevealProofs         []datalayr.MultiRevealProof
	BatchPolyEquivalenceProof [4]*big.Int
	StartingChunkIndex        int
}

type FraudProof struct {
	DataLayrDisclosureProof
	StartingSymbolIndex int
}

type ChallengerConfig struct {
	L1Client                  *ethclient.Client
	L2Client                  *l2ethclient.Client
	EigenContractAddr         ethc.Address
	Logger                    *logging.Logger
	PrivKey                   *ecdsa.PrivateKey
	GraphProvider             string
	RetrieverSocket           string
	KzgConfig                 KzgConfig
	LastStoreNumber           uint64
	Timeout                   time.Duration
	PollInterval              time.Duration
	DbPath                    string
	CheckerBatchIndex         uint64
	NeedReRollupBatch         string
	ReRollupToolEnable        bool
	SignerFn                  SignerFn
	ResubmissionTimeout       time.Duration
	NumConfirmations          uint64
	SafeAbortNonceTooLowCount uint64
}

type Challenger struct {
	Ctx              context.Context
	Cfg              *ChallengerConfig
	EigenDaContract  *bindings.BVMEigenDataLayrChain
	RawEigenContract *bind.BoundContract
	WalletAddr       ethc.Address
	EigenABI         *abi.ABI
	GraphClient      *graphView.GraphClient
	GraphqlClient    *graphql.Client
	LevelDBStore     *db.Store
	txMgr            txmgr.TxManager
	cancel           func()
	wg               sync.WaitGroup
	once             sync.Once
}

func NewChallenger(ctx context.Context, cfg *ChallengerConfig) (*Challenger, error) {
	_, cancel := context.WithTimeout(ctx, common4.DefaultTimeout)
	defer cancel()
	eigenContract, err := bindings.NewBVMEigenDataLayrChain(
		ethc.Address(cfg.EigenContractAddr), cfg.L1Client,
	)
	if err != nil {
		return nil, err
	}
	parsed, err := abi.JSON(strings.NewReader(
		bindings.BVMEigenDataLayrChainABI,
	))
	if err != nil {
		log.Error("Challenger parse eigen layer contract abi fail", "err", err)
		return nil, err
	}
	eignenABI, err := bindings.BVMEigenDataLayrChainMetaData.GetAbi()
	if err != nil {
		log.Error("Challenger get eigen layer contract abi fail", "err", err)
		return nil, err
	}
	rawEigenContract := bind.NewBoundContract(
		cfg.EigenContractAddr, parsed, cfg.L1Client, cfg.L1Client,
		cfg.L1Client,
	)

	txManagerConfig := txmgr.Config{
		ResubmissionTimeout:       cfg.ResubmissionTimeout,
		ReceiptQueryInterval:      time.Second,
		NumConfirmations:          cfg.NumConfirmations,
		SafeAbortNonceTooLowCount: cfg.SafeAbortNonceTooLowCount,
	}

	txMgr := txmgr.NewSimpleTxManager(txManagerConfig, cfg.L1Client)

	graphClient := graphView.NewGraphClient(cfg.GraphProvider, cfg.Logger)
	graphqlClient := graphql.NewClient(graphClient.GetEndpoint(), nil)
	walletAddr := crypto.PubkeyToAddress(cfg.PrivKey.PublicKey)

	levelDBStore, err := db.NewStore(cfg.DbPath)
	if err != nil {
		log.Error("Challenger init leveldb fail", "err", err)
		return nil, err
	}
	return &Challenger{
		Cfg:              cfg,
		Ctx:              ctx,
		EigenDaContract:  eigenContract,
		RawEigenContract: rawEigenContract,
		WalletAddr:       walletAddr,
		EigenABI:         eignenABI,
		GraphClient:      graphClient,
		GraphqlClient:    graphqlClient,
		LevelDBStore:     levelDBStore,
		txMgr:            txMgr,
		cancel:           cancel,
	}, nil
}

func (c *Challenger) getDataStoreById(dataStoreId string) (*graphView.DataStore, error) {
	var query struct {
		DataStore graphView.DataStoreGql `graphql:"dataStore(id: $storeId)"`
	}
	variables := map[string]interface{}{
		"storeId": graphql.String(dataStoreId),
	}
	err := c.GraphqlClient.Query(context.Background(), &query, variables)
	if err != nil {
		log.Error("Challenger query data from graphql fail", "err", err)
		return nil, err
	}
	store, err := query.DataStore.Convert()
	if err != nil {
		log.Error("Challenger convert data store fail", "err", err)
		return nil, err
	}
	c.Cfg.LastStoreNumber = uint64(store.StoreNumber)
	return store, nil
}

func (c *Challenger) callRetrieve(store *graphView.DataStore) ([]byte, []datalayr.Frame, error) {
	conn, err := grpc.Dial(c.Cfg.RetrieverSocket, grpc.WithInsecure())
	if err != nil {
		log.Error("Disperser Cannot connect to", "retriever-socket", c.Cfg.RetrieverSocket, "err", err)
		return nil, nil, err
	}
	defer conn.Close()
	client := pb.NewDataRetrievalClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), c.Cfg.Timeout)
	defer cancel()
	opt := grpc.MaxCallRecvMsgSize(1024 * 1024 * 300)
	request := &pb.FramesAndDataRequest{
		DataStoreId: store.StoreNumber,
	}
	reply, err := client.RetrieveFramesAndData(ctx, request, opt)
	if err != nil {
		return nil, nil, err
	}
	data := reply.GetData()
	framesBytes := reply.GetFrames()
	header, err := datalayr.DecodeDataStoreHeader(store.Header)
	if err != nil {
		log.Error("Could not decode header", "err", err)
		return nil, nil, err
	}
	frames := make([]datalayr.Frame, header.NumSys+header.NumPar)
	for i, frameBytes := range framesBytes {
		frame, err := datalayr.DecodeFrame(frameBytes)
		if err == nil {
			frames[i] = frame
		} else {
			return nil, nil, errors.New("Does not Contain all the frames")
		}
	}
	return data, frames, nil
}

// check if the fraud string exists within the data
func (c *Challenger) checkForFraud(store *graphView.DataStore, data []byte) (*Fraud, bool) {
	dataString := hex.EncodeToString(data)
	index := strings.Index(dataString, fraudString)
	if index != -1 {
		return &Fraud{StartingIndex: index / 2}, true
	}
	return nil, false
}

func (c *Challenger) constructFraudProof(store *graphView.DataStore, data []byte, fraud *Fraud, frames []datalayr.Frame) (*FraudProof, error) {
	// encode data to frames here
	header, err := datalayr.DecodeDataStoreHeader(store.Header)
	if err != nil {
		c.Cfg.Logger.Printf("Could not decode header %v. %v\n", header, err)
		return nil, err
	}
	config := c.Cfg.KzgConfig

	s1 := gkzg.ReadG1Points(config.G1Path, config.Order, config.NumWorker)
	s2 := gkzg.ReadG2Points(config.G2Path, config.Order, config.NumWorker)

	dp := datalayr.NewDisclosureProver(s1, s2)

	//there are 31 bytes per fr so there are 31*chunkLenE bytes in each chunk
	//so the i'th byte starts at the (i/(31*encoder.EncodingParams.ChunkLenE))'th chunk
	startingChunkIndex := fraud.StartingIndex / int(31*header.Degree)
	//the fraud string ends len(fraudString)/2 bytes later
	endingChunkIndex := (fraud.StartingIndex + len(fraudString)/2) / int(31*header.Degree)
	startingSymbolIndex := fraud.StartingIndex % int(31*header.Degree)
	//do some math to shift this over by the correct number of bytes
	//there are 32 bytes in the actual poly for every 31 bytes in the data, hence (startingSymbolIndex/31)*32
	//then we shift over by 1 to get past the first 0 byte, and then (startingSymbolIndex % 31)
	startingSymbolIndex = (startingSymbolIndex/31)*32 + 1 + (startingSymbolIndex % 31)

	//generate parameters for proving data on chain
	//this is
	//	polys: the []byte representation of the polynomials
	//	multirevealProofs: the openings of each polynomial against the full polynomial commitment
	//  batchPolyEquivalenceProof: the proof that the `polys` are in fact represented by the commitments in `multirevealProofs`
	polys, multirevealProofs, batchPolyEquivalenceProof, err := dp.ProveBatchInterpolatingPolyDisclosure(frames[startingChunkIndex:endingChunkIndex+1], store.DataCommitment, store.Header, uint32(startingChunkIndex))
	if err != nil {
		return nil, err
	}

	disclosureProof := DataLayrDisclosureProof{
		Header:                    store.Header,
		Polys:                     polys,
		MultirevealProofs:         multirevealProofs,
		BatchPolyEquivalenceProof: batchPolyEquivalenceProof,
		StartingChunkIndex:        startingChunkIndex,
	}

	return &FraudProof{DataLayrDisclosureProof: disclosureProof, StartingSymbolIndex: startingSymbolIndex}, nil
}

func (fp *DataLayrDisclosureProof) ToDisclosureProofs() rc.BVMEigenDataLayrChainDisclosureProofs {
	proofs := make([]rc.DataLayrDisclosureLogicMultiRevealProof, 0)
	for _, oldProof := range fp.MultirevealProofs {
		newProof := rc.DataLayrDisclosureLogicMultiRevealProof{
			InterpolationPoly: rc.BN254G1Point{X: oldProof.InterpolationPolyCommit[0], Y: oldProof.InterpolationPolyCommit[1]},
			RevealProof:       rc.BN254G1Point{X: oldProof.RevealProof[0], Y: oldProof.RevealProof[1]},
			ZeroPoly: rc.BN254G2Point{
				//preserve this ordering for dumb precompile reasons
				X: [2]*big.Int{oldProof.ZeroPolyCommit[1], oldProof.ZeroPolyCommit[0]},
				Y: [2]*big.Int{oldProof.ZeroPolyCommit[3], oldProof.ZeroPolyCommit[2]},
			},
			ZeroPolyProof: oldProof.ZeroPolyProof,
		}
		proofs = append(proofs, newProof)
	}
	return rc.BVMEigenDataLayrChainDisclosureProofs{
		Header:            fp.Header,
		FirstChunkNumber:  uint32(fp.StartingChunkIndex),
		Polys:             fp.Polys,
		MultiRevealProofs: proofs,
		PolyEquivalenceProof: rc.BN254G2Point{
			//preserve this ordering for dumb precompile reasons
			X: [2]*big.Int{fp.BatchPolyEquivalenceProof[1], fp.BatchPolyEquivalenceProof[0]},
			Y: [2]*big.Int{fp.BatchPolyEquivalenceProof[3], fp.BatchPolyEquivalenceProof[2]},
		},
	}
}

func (c *Challenger) UpdateGasPrice(ctx context.Context, tx *types.Transaction) (*types.Transaction, error) {
	opts := &bind.TransactOpts{
		From: c.WalletAddr,
		Signer: func(addr ethc.Address, tx *types.Transaction) (*types.Transaction, error) {
			return c.Cfg.SignerFn(ctx, addr, tx)
		},
		Context: ctx,
		Nonce:   new(big.Int).SetUint64(tx.Nonce()),
		NoSend:  true,
	}
	finalTx, err := c.RawEigenContract.RawTransact(opts, tx.Data())
	switch {
	case err == nil:
		return finalTx, nil

	case c.IsMaxPriorityFeePerGasNotFoundError(err):
		log.Info("MtChallenger eth_maxPriorityFeePerGas is unsupported by current backend, using fallback gasTipCap", "txData", tx.Data())
		opts.GasTipCap = FallbackGasTipCap
		return c.RawEigenContract.RawTransact(opts, tx.Data())

	default:
		return nil, err
	}
}

func (c *Challenger) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return c.Cfg.L1Client.SendTransaction(ctx, tx)
}

func (c *Challenger) IsMaxPriorityFeePerGasNotFoundError(err error) bool {
	return strings.Contains(
		err.Error(), errMaxPriorityFeePerGasNotFound.Error(),
	)
}

func (c *Challenger) ChallengeProveFraud(ctx context.Context, fraudStoreNumber *big.Int, fraudProof *FraudProof, searchData rc.IDataLayrServiceManagerDataStoreSearchData, disclosureProofs rc.BVMEigenDataLayrChainDisclosureProofs) (*types.Transaction, error) {
	balance, err := c.Cfg.L1Client.BalanceAt(
		c.Ctx, ethc.Address(c.WalletAddr), nil,
	)
	if err != nil {
		log.Error("MtChallenger unable to get current balance", "err", err)
		return nil, err
	}
	log.Info("MtChallenger wallet address balance", "balance", balance)
	nonce64, err := c.Cfg.L1Client.NonceAt(
		c.Ctx, ethc.Address(c.WalletAddr), nil,
	)
	if err != nil {
		log.Error("MtChallenger unable to get current nonce", "err", err)
		return nil, err
	}
	nonce := new(big.Int).SetUint64(nonce64)
	opts := &bind.TransactOpts{
		From: ethc.Address(c.WalletAddr),
		Signer: func(addr ethc.Address, tx *types.Transaction) (*types.Transaction, error) {
			return c.Cfg.SignerFn(ctx, addr, tx)
		},
		Context: ctx,
		Nonce:   nonce,
		NoSend:  true,
	}
	tx, err := c.EigenDaContract.ProveFraud(opts, fraudStoreNumber, new(big.Int).SetUint64(uint64(fraudProof.StartingSymbolIndex)), searchData, disclosureProofs)
	switch {
	case err == nil:
		return tx, nil

	case c.IsMaxPriorityFeePerGasNotFoundError(err):
		log.Warn("MtChallenger eth_maxPriorityFeePerGas is unsupported by current backend, using fallback gasTipCap")
		opts.GasTipCap = FallbackGasTipCap
		return c.EigenDaContract.ProveFraud(opts, fraudStoreNumber, new(big.Int).SetUint64(uint64(fraudProof.StartingSymbolIndex)), searchData, disclosureProofs)

	default:
		return nil, err
	}
}

func (c *Challenger) postFraudProof(store *graphView.DataStore, fraudProof *FraudProof) (*types.Transaction, error) {
	searchData := rc.IDataLayrServiceManagerDataStoreSearchData{
		Duration:  store.Duration,
		Timestamp: new(big.Int).SetUint64(uint64(store.InitTime)),
		Index:     store.Index,
		Metadata: rc.IDataLayrServiceManagerDataStoreMetadata{
			HeaderHash:          store.DataCommitment,
			DurationDataStoreId: store.DurationDataStoreId,
			GlobalDataStoreId:   store.StoreNumber,
			BlockNumber:         store.StakesFromBlockNumber,
			Fee:                 store.Fee,
			Confirmer:           ethc.Address(common.HexToAddress(store.Confirmer)),
			SignatoryRecordHash: store.SignatoryRecord,
		},
	}
	disclosureProofs := fraudProof.ToDisclosureProofs()
	fraudStoreNumber, err := c.EigenDaContract.DataStoreIdToRollupStoreNumber(&bind.CallOpts{}, store.StoreNumber)
	if err != nil {
		return nil, err
	}

	tx, err := c.ChallengeProveFraud(c.Ctx, fraudStoreNumber, fraudProof, searchData, disclosureProofs)
	if err != nil {
		return nil, err
	}
	updateGasPrice := func(ctx context.Context) (*types.Transaction, error) {
		log.Info("MtChallenger ConfirmData update gas price")
		return c.UpdateGasPrice(ctx, tx)
	}
	receipt, err := c.txMgr.Send(
		c.Ctx, updateGasPrice, c.SendTransaction,
	)
	if err != nil {
		return nil, err
	}
	log.Info("MtChallenge challenger prove fraud success", "TxHash", receipt.TxHash)
	return tx, nil
}

func (c *Challenger) makReRollupBatchTx(ctx context.Context, batchIndex *big.Int) (*types.Transaction, error) {
	balance, err := c.Cfg.L1Client.BalanceAt(
		c.Ctx, ethc.Address(c.WalletAddr), nil,
	)
	if err != nil {
		log.Error("MtChallenger unable to get current balance", "err", err)
		return nil, err
	}
	log.Info("MtChallenger wallet address balance", "balance", balance)
	nonce64, err := c.Cfg.L1Client.NonceAt(
		c.Ctx, ethc.Address(c.WalletAddr), nil,
	)
	if err != nil {
		log.Error("MtChallenger unable to get current nonce", "err", err)
		return nil, err
	}
	nonce := new(big.Int).SetUint64(nonce64)
	opts := &bind.TransactOpts{
		From: ethc.Address(c.WalletAddr),
		Signer: func(addr ethc.Address, tx *types.Transaction) (*types.Transaction, error) {
			return c.Cfg.SignerFn(ctx, addr, tx)
		},
		Context: ctx,
		Nonce:   nonce,
		NoSend:  true,
	}
	tx, err := c.EigenDaContract.SubmitReRollUpInfo(opts, batchIndex)
	switch {
	case err == nil:
		return tx, nil

	case c.IsMaxPriorityFeePerGasNotFoundError(err):
		log.Warn("MtChallenger eth_maxPriorityFeePerGas is unsupported by current backend, using fallback gasTipCap")
		opts.GasTipCap = FallbackGasTipCap
		return c.EigenDaContract.SubmitReRollUpInfo(opts, batchIndex)
	default:
		return nil, err
	}
}

func (c *Challenger) submitReRollupBatchIndex(batchIndex *big.Int) (*types.Transaction, error) {
	tx, err := c.makReRollupBatchTx(c.Ctx, batchIndex)
	if err != nil {
		return nil, err
	}
	updateGasPrice := func(ctx context.Context) (*types.Transaction, error) {
		log.Info("MtChallenger makReRollupBatchTx update gas price")
		return c.UpdateGasPrice(ctx, tx)
	}
	receipt, err := c.txMgr.Send(
		c.Ctx, updateGasPrice, c.SendTransaction,
	)
	if err != nil {
		return nil, err
	}
	log.Info("MtChallenge submit re-rollup batch index success", "TxHash", receipt.TxHash)
	return tx, nil
}

func (c *Challenger) Start() error {
	c.wg.Add(1)
	go c.eventLoop()
	c.once.Do(func() {
		log.Info("MtChallenge start exec once update da tool")
		if c.Cfg.ReRollupToolEnable {
			reRollupBatchIndex := strings.Split(c.Cfg.NeedReRollupBatch, "|")
			for index := 0; index < len(reRollupBatchIndex); index++ {
				bigParam := new(big.Int)
				bigBatchIndex, ok := bigParam.SetString(reRollupBatchIndex[index], 10)
				if !ok {
					log.Error("MtChallenge string to big.int fail", "ok", ok)
					continue
				}
				tx, err := c.submitReRollupBatchIndex(bigBatchIndex)
				if err != nil {
					log.Error("MtChallenge tool submit re-rollup info fail", "batchIndex", reRollupBatchIndex[index], "err", err)
					continue
				}
				log.Info("MtChallenge tool submit re-rollup info success", "batchIndex", reRollupBatchIndex[index], "txHash", tx.Hash().String())
			}
		}
	})
	return nil
}

func (c *Challenger) Stop() {
	c.cancel()
	c.wg.Wait()
}

func (c *Challenger) eventLoop() {
	defer c.wg.Done()
	ticker := time.NewTicker(c.Cfg.PollInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			latestBatchIndex, err := c.EigenDaContract.RollupBatchIndex(&bind.CallOpts{})
			if err != nil {
				log.Error("MtChallenge get batch index fail", "err", err)
				continue
			}
			batchIndex, ok := c.LevelDBStore.GetLatestBatchIndex()
			if !ok {
				log.Error("MtChallenge get batch index from db fail", "err", err)
			}
			if c.Cfg.CheckerBatchIndex > latestBatchIndex.Uint64() {
				log.Info("MtChallenge Batch Index", "DbBatchIndex", batchIndex, "ContractBatchIndex", latestBatchIndex.Uint64()-c.Cfg.CheckerBatchIndex)
				continue
			}
			if batchIndex >= (latestBatchIndex.Uint64() - c.Cfg.CheckerBatchIndex) {
				log.Info("MtChallenge db batch index and contract batch idnex is equal", "DbBatchIndex", batchIndex, "ContractBatchIndex", latestBatchIndex.Uint64()-c.Cfg.CheckerBatchIndex)
				continue
			}
			log.Info("MtChallenge db batch index and contract batch idnex",
				"DbBatchIndex", batchIndex, "ContractBatchIndex", latestBatchIndex.Uint64(),
				"latestBatchIndex.Uint64() - c.Cfg.CheckerBatchIndex", latestBatchIndex.Uint64()-c.Cfg.CheckerBatchIndex,
			)
			for i := batchIndex; i <= (latestBatchIndex.Uint64() - c.Cfg.CheckerBatchIndex); i++ {
				dataStoreId, err := c.EigenDaContract.GetRollupStoreByRollupBatchIndex(&bind.CallOpts{}, big.NewInt(int64(i)))
				if err != nil {
					continue
				}
				store, err := c.getDataStoreById(strconv.Itoa(int(dataStoreId.DataStoreId)))
				if err != nil {
					log.Error("MtChallenge get data store fail", "err", err)
					continue
				}
				log.Info("MtChallenge get data store by id success", "Confirmed", store.Confirmed)
				if store.Confirmed {
					data, frames, err := c.callRetrieve(store)
					if err != nil {
						log.Error("MtChallenge error getting data", "err", err)
						continue
					}
					batchTxn := new([]eigenda.BatchTx)
					batchRlpStream := rlp.NewStream(bytes.NewBuffer(data), uint64(len(data)))
					err = batchRlpStream.Decode(batchTxn)
					if err != nil {
						log.Error("MtChallenge decode batch txn fail", "err", err)
						continue
					}
					newBatchTxn := *batchTxn
					for i := 0; i < len(newBatchTxn); i++ {
						l2Tx := new(l2types.Transaction)
						rlpStream := l2rlp.NewStream(bytes.NewBuffer(newBatchTxn[i].RawTx), 0)
						if err := l2Tx.DecodeRLP(rlpStream); err != nil {
							c.Cfg.Logger.Error().Err(err).Msg("Decode RLP fail")
						}
						log.Info("MtChallenge decode transaction", "hash", l2Tx.Hash().Hex())

						// tx check for tmp, will remove in future
						l2Transaction, _, err := c.Cfg.L2Client.TransactionByHash(c.Ctx, l2Tx.Hash())
						if err != nil {
							log.Error("MtChallenge no this transaction", "err", err)
							continue
						}
						log.Info("MtChallenge fond transaction", "hash", l2Transaction.Hash().Hex())
					}

					// check if the fraud string exists within the data
					fraud, exists := c.checkForFraud(store, data)
					if !exists {
						log.Info("MtChallenge no fraud")
						c.LevelDBStore.SetLatestBatchIndex(i)
						continue
					}
					proof, err := c.constructFraudProof(store, data, fraud, frames)
					if err != nil {
						log.Error("MtChallenge error constructing fraud", "err", err)
						continue
					}
					tx, err := c.postFraudProof(store, proof)
					if err != nil {
						log.Error("MtChallenge error posting fraud proof", "err", err)
						continue
					}
					log.Info("MtChallenge fraud proof tx", "hash", tx.Hash().Hex())
				}
				c.LevelDBStore.SetLatestBatchIndex(i)
			}
		case err := <-c.Ctx.Done():
			log.Error("MtChallenge eigenDa sequencer service shutting down", "err", err)
			return
		}
	}
}

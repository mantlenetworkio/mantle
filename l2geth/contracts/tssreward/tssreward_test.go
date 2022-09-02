package tssreward

import (
	//"bytes"
	//"crypto/ecdsa"
	//"math/big"
	//"sort"
	"testing"
	//"github.com/bitdao-io/bitnetwork/l2geth/accounts/abi/bind/backends"
)

//func TestUpdateTssRewardData(t *testing.T) {
//
//	statedb, _ := state.New(common.Hash{}, state.NewDatabase(rawdb.NewMemoryDatabase()))
//	tssCode := []byte("0x608060405234801561001057600080fd5b50604051610dc9380380610dc983398101604081905261002f91610078565b600180546001600160a01b039384166001600160a01b031991821617909155600280549290931691161790556100b2565b6001600160a01b038116811461007557600080fd5b50565b6000806040838503121561008b57600080fd5b825161009681610060565b60208401519092506100a781610060565b809150509250929050565b610d08806100c16000396000f3fe608060405234801561001057600080fd5b50600436106100c95760003560e01c80632c79db1111610081578063cfb550f11161005b578063cfb550f1146101ad578063e04f6e35146101b5578063fad9aba3146101c857600080fd5b80632c79db111461017b5780633ccfd60b146101835780638da5cb5b1461018d57600080fd5b806319d509a1116100b257806319d509a1146101245780631a39d8ef1461012d57806327c8f8351461013657600080fd5b80630b50cd3e146100ce57806310a7fd7b146100f6575b600080fd5b6100e16100dc366004610a4a565b6101d1565b60405190151581526020015b60405180910390f35b610116610104366004610a6c565b60006020819052908152604090205481565b6040519081526020016100ed565b61011660035481565b61011660055481565b6001546101569073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100ed565b6101166103b3565b61018b61044c565b005b6002546101569073ffffffffffffffffffffffffffffffffffffffff1681565b61018b6105da565b61018b6101c3366004610a85565b61077e565b61011660045481565b60015460009073ffffffffffffffffffffffffffffffffffffffff163314610280576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602760248201527f747373207265776172642063616c6c206d65737361676520756e61757468656e60448201527f746963617465640000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b600554471015610312576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603160248201527f62616c616e6365207265636f726420616e6420636f6e74726163742062616c6160448201527f6e636520617265206e6f7420657175616c0000000000000000000000000000006064820152608401610277565b600354610320906001610b49565b8314610388576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f626c6f636b2069642075706461746520696c6c6567616c0000000000000000006044820152606401610277565b600383905560055461039a9083610a1f565b6005555060009182526020829052604090912055600190565b6000600554471015610447576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603160248201527f62616c616e6365207265636f726420616e6420636f6e74726163742062616c6160448201527f6e636520617265206e6f7420657175616c0000000000000000000000000000006064820152608401610277565b504790565b60025473ffffffffffffffffffffffffffffffffffffffff1633146104f3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602c60248201527f6f6e6c792062652063616c6c656420627920746865206f776e6572206f66207460448201527f68697320636f6e747261637400000000000000000000000000000000000000006064820152608401610277565b600554471015610585576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603160248201527f62616c616e6365207265636f726420616e6420636f6e74726163742062616c6160448201527f6e636520617265206e6f7420657175616c0000000000000000000000000000006064820152608401610277565b600060055547156105d85760025460405173ffffffffffffffffffffffffffffffffffffffff909116904780156108fc02916000818181858888f193505050501580156105d6573d6000803e3d6000fd5b505b565b60025473ffffffffffffffffffffffffffffffffffffffff163314610681576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602c60248201527f6f6e6c792062652063616c6c656420627920746865206f776e6572206f66207460448201527f68697320636f6e747261637400000000000000000000000000000000000000006064820152608401610277565b600554471015610713576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603160248201527f62616c616e6365207265636f726420616e6420636f6e74726163742062616c6160448201527f6e636520617265206e6f7420657175616c0000000000000000000000000000006064820152608401610277565b6004546005546107239082610a32565b600555600060045580156105d65760025460045460405173ffffffffffffffffffffffffffffffffffffffff9092169181156108fc0291906000818181858888f1935050505015801561077a573d6000803e3d6000fd5b5050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610825576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602760248201527f747373207265776172642063616c6c206d65737361676520756e61757468656e60448201527f74696361746564000000000000000000000000000000000000000000000000006064820152608401610277565b6005544710156108b7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603160248201527f62616c616e6365207265636f726420616e6420636f6e74726163742062616c6160448201527f6e636520617265206e6f7420657175616c0000000000000000000000000000006064820152608401610277565b6000848152602081905260408120549080865b8663ffffffff168110156109d8576108e28486610a3e565b60008981526020819052604081208190559093505b8581101561099f57600087878381811061091357610913610b61565b90506020020160208101906109289190610bb9565b90506109348486610a1f565b6005549094506109449086610a32565b60055560405173ffffffffffffffffffffffffffffffffffffffff82169086156108fc029087906000818181858888f1935050505015801561098a573d6000803e3d6000fd5b5050808061099790610bd4565b9150506108f7565b5060006109ac8584610a32565b905080156109c5576004546109c19082610a1f565b6004555b50806109d081610bd4565b9150506108ca565b507ff630cba6d450d736e85735388d4fe67a177b8a3685cdd7dee2bea7727b47860a87878787604051610a0e9493929190610c0d565b60405180910390a150505050505050565b6000610a2b8284610b49565b9392505050565b6000610a2b8284610c80565b6000610a2b8284610c97565b60008060408385031215610a5d57600080fd5b50508035926020909101359150565b600060208284031215610a7e57600080fd5b5035919050565b60008060008060608587031215610a9b57600080fd5b84359350602085013563ffffffff81168114610ab657600080fd5b9250604085013567ffffffffffffffff80821115610ad357600080fd5b818701915087601f830112610ae757600080fd5b813581811115610af657600080fd5b8860208260051b8501011115610b0b57600080fd5b95989497505060200194505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008219821115610b5c57610b5c610b1a565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b803573ffffffffffffffffffffffffffffffffffffffff81168114610bb457600080fd5b919050565b600060208284031215610bcb57600080fd5b610a2b82610b90565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415610c0657610c06610b1a565b5060010190565b600060608201868352602063ffffffff871681850152606060408501528185835260808501905086925060005b86811015610c735773ffffffffffffffffffffffffffffffffffffffff610c6085610b90565b1682529282019290820190600101610c3a565b5098975050505050505050565b600082821015610c9257610c92610b1a565b500390565b600082610ccd577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b50049056fea2646970667358221220a51ec743c0bb2b95aac6f7675b0aa3cbda567cf1153f306a8933bd643a1dc0cd64736f6c63430008090033")
//	statedb.SetCode(dump.L2ExcuteFeeWallet, tssCode)
//	vmctx := vm.Context{
//		CanTransfer: func(vm.StateDB, common.Address, *big.Int) bool { return true },
//		Transfer:    func(vm.StateDB, common.Address, common.Address, *big.Int) {},
//	}
//
//	statedb.Finalise(true) // Push the state into the "original" slot
//
//	vmenv := vm.NewEVM(vmctx, statedb, params.AllEthashProtocolChanges, vm.Config{})
//	zeroAddress := vm.AccountRef(common.Address{})
//	data2, err := PacketQueryData()
//	require.Nil(t, err)
//	ret, _, err := vmenv.Call(zeroAddress, dump.L2ExcuteFeeWallet, data2, 450, big.NewInt(0))
//	require.Nil(t, err)
//	t.Logf("ret:%v", string(ret))
//
//}

func TestTssReward_QueryReward(t *testing.T) {
}

func TestTssReward_UpdateReward(t *testing.T) {

}

//func TestCheckpointRegister(t *testing.T) {
//	// Initialize test accounts
//	var accounts Accounts
//	for i := 0; i < 3; i++ {
//		key, _ := crypto.GenerateKey()
//		addr := crypto.PubkeyToAddress(key.PublicKey)
//		accounts = append(accounts, Account{key: key, addr: addr})
//	}
//	sort.Sort(accounts)
//
//	// Deploy registrar contract
//	contractBackend := backends.NewSimulatedBackend(core.GenesisAlloc{accounts[0].addr: {Balance: big.NewInt(1000000000)}, accounts[1].addr: {Balance: big.NewInt(1000000000)}, accounts[2].addr: {Balance: big.NewInt(1000000000)}}, 10000000)
//	defer contractBackend.Close()
//
//	transactOpts := bind.NewKeyedTransactor(accounts[0].key)
//
//	// 3 trusted signers, threshold 2
//	contractAddr, _, c, err := contract.DeployCheckpointOracle(transactOpts, contractBackend, []common.Address{accounts[0].addr, accounts[1].addr, accounts[2].addr}, sectionSize, processConfirms, big.NewInt(2))
//	if err != nil {
//		t.Error("Failed to deploy registrar contract", err)
//	}
//	contractBackend.Commit()
//
//	// getRecent returns block height and hash of the head parent.
//	getRecent := func() (*big.Int, common.Hash) {
//		parentNumber := new(big.Int).Sub(contractBackend.Blockchain().CurrentHeader().Number, big.NewInt(1))
//		parentHash := contractBackend.Blockchain().CurrentHeader().ParentHash
//		return parentNumber, parentHash
//	}
//	// collectSig generates specified number signatures.
//	collectSig := func(index uint64, hash common.Hash, n int, unauthorized *ecdsa.PrivateKey) (v []uint8, r [][32]byte, s [][32]byte) {
//		for i := 0; i < n; i++ {
//			sig := signCheckpoint(contractAddr, accounts[i].key, index, hash)
//			if unauthorized != nil {
//				sig = signCheckpoint(contractAddr, unauthorized, index, hash)
//			}
//			r = append(r, common.BytesToHash(sig[:32]))
//			s = append(s, common.BytesToHash(sig[32:64]))
//			v = append(v, sig[64])
//		}
//		return v, r, s
//	}
//	// insertEmptyBlocks inserts a batch of empty blocks to blockchain.
//	insertEmptyBlocks := func(number int) {
//		for i := 0; i < number; i++ {
//			contractBackend.Commit()
//		}
//	}
//	// assert checks whether the current contract status is same with
//	// the expected.
//	assert := func(index uint64, hash [32]byte, height *big.Int) error {
//		lindex, lhash, lheight, err := c.GetLatestCheckpoint(nil)
//		if err != nil {
//			return err
//		}
//		if lindex != index {
//			return errors.New("latest checkpoint index mismatch")
//		}
//		if !bytes.Equal(lhash[:], hash[:]) {
//			return errors.New("latest checkpoint hash mismatch")
//		}
//		if lheight.Cmp(height) != 0 {
//			return errors.New("latest checkpoint height mismatch")
//		}
//		return nil
//	}
//
//	// Test future checkpoint registration
//	validateOperation(t, c, contractBackend, func() {
//		number, hash := getRecent()
//		v, r, s := collectSig(0, checkpoint0.Hash(), 2, nil)
//		c.SetCheckpoint(transactOpts, number, hash, checkpoint0.Hash(), 0, v, r, s)
//	}, func(events <-chan *contract.CheckpointOracleNewCheckpointVote) error {
//		return assert(0, emptyHash, big.NewInt(0))
//	}, "test future checkpoint registration")
//
//	insertEmptyBlocks(int(sectionSize.Uint64() + processConfirms.Uint64()))
//
//	// Test transaction replay protection
//	validateOperation(t, c, contractBackend, func() {
//		number, _ := getRecent()
//		v, r, s := collectSig(0, checkpoint0.Hash(), 2, nil)
//		hash := common.HexToHash("deadbeef")
//		c.SetCheckpoint(transactOpts, number, hash, checkpoint0.Hash(), 0, v, r, s)
//	}, func(events <-chan *contract.CheckpointOracleNewCheckpointVote) error {
//		return assert(0, emptyHash, big.NewInt(0))
//	}, "test transaction replay protection")
//
//	// Test unauthorized signature checking
//	validateOperation(t, c, contractBackend, func() {
//		number, hash := getRecent()
//		u, _ := crypto.GenerateKey()
//		v, r, s := collectSig(0, checkpoint0.Hash(), 2, u)
//		c.SetCheckpoint(transactOpts, number, hash, checkpoint0.Hash(), 0, v, r, s)
//	}, func(events <-chan *contract.CheckpointOracleNewCheckpointVote) error {
//		return assert(0, emptyHash, big.NewInt(0))
//	}, "test unauthorized signature checking")
//
//	// Test un-multi-signature checkpoint registration
//	validateOperation(t, c, contractBackend, func() {
//		number, hash := getRecent()
//		v, r, s := collectSig(0, checkpoint0.Hash(), 1, nil)
//		c.SetCheckpoint(transactOpts, number, hash, checkpoint0.Hash(), 0, v, r, s)
//	}, func(events <-chan *contract.CheckpointOracleNewCheckpointVote) error {
//		return assert(0, emptyHash, big.NewInt(0))
//	}, "test un-multi-signature checkpoint registration")
//
//	// Test valid checkpoint registration
//	validateOperation(t, c, contractBackend, func() {
//		number, hash := getRecent()
//		v, r, s := collectSig(0, checkpoint0.Hash(), 2, nil)
//		c.SetCheckpoint(transactOpts, number, hash, checkpoint0.Hash(), 0, v, r, s)
//	}, func(events <-chan *contract.CheckpointOracleNewCheckpointVote) error {
//		if valid, recv := validateEvents(2, events); !valid {
//			return errors.New("receive incorrect number of events")
//		} else {
//			for i := 0; i < len(recv); i++ {
//				event := recv[i].Interface().(*contract.CheckpointOracleNewCheckpointVote)
//				if !assertSignature(contractAddr, event.Index, event.CheckpointHash, event.R, event.S, event.V, accounts[i].addr) {
//					return errors.New("recover signer failed")
//				}
//			}
//		}
//		number, _ := getRecent()
//		return assert(0, checkpoint0.Hash(), number.Add(number, big.NewInt(1)))
//	}, "test valid checkpoint registration")
//
//	distance := 3*sectionSize.Uint64() + processConfirms.Uint64() - contractBackend.Blockchain().CurrentHeader().Number.Uint64()
//	insertEmptyBlocks(int(distance))
//
//	// Test uncontinuous checkpoint registration
//	validateOperation(t, c, contractBackend, func() {
//		number, hash := getRecent()
//		v, r, s := collectSig(2, checkpoint2.Hash(), 2, nil)
//		c.SetCheckpoint(transactOpts, number, hash, checkpoint2.Hash(), 2, v, r, s)
//	}, func(events <-chan *contract.CheckpointOracleNewCheckpointVote) error {
//		if valid, recv := validateEvents(2, events); !valid {
//			return errors.New("receive incorrect number of events")
//		} else {
//			for i := 0; i < len(recv); i++ {
//				event := recv[i].Interface().(*contract.CheckpointOracleNewCheckpointVote)
//				if !assertSignature(contractAddr, event.Index, event.CheckpointHash, event.R, event.S, event.V, accounts[i].addr) {
//					return errors.New("recover signer failed")
//				}
//			}
//		}
//		number, _ := getRecent()
//		return assert(2, checkpoint2.Hash(), number.Add(number, big.NewInt(1)))
//	}, "test uncontinuous checkpoint registration")
//
//	// Test old checkpoint registration
//	validateOperation(t, c, contractBackend, func() {
//		number, hash := getRecent()
//		v, r, s := collectSig(1, checkpoint1.Hash(), 2, nil)
//		c.SetCheckpoint(transactOpts, number, hash, checkpoint1.Hash(), 1, v, r, s)
//	}, func(events <-chan *contract.CheckpointOracleNewCheckpointVote) error {
//		number, _ := getRecent()
//		return assert(2, checkpoint2.Hash(), number)
//	}, "test uncontinuous checkpoint registration")
//
//	// Test stale checkpoint registration
//	validateOperation(t, c, contractBackend, func() {
//		number, hash := getRecent()
//		v, r, s := collectSig(2, checkpoint2.Hash(), 2, nil)
//		c.SetCheckpoint(transactOpts, number, hash, checkpoint2.Hash(), 2, v, r, s)
//	}, func(events <-chan *contract.CheckpointOracleNewCheckpointVote) error {
//		number, _ := getRecent()
//		return assert(2, checkpoint2.Hash(), number.Sub(number, big.NewInt(1)))
//	}, "test stale checkpoint registration")
//}


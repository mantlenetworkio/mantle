package rcfg

import (
	"math/big"
	"os"

	"github.com/mantlenetworkio/mantle/l2geth/common"
)

// UsingBVM is used to enable or disable functionality necessary for the BVM.
var UsingBVM bool

var (
	// L2GasPriceSlot refers to the storage slot that the L2 gas price is stored
	// in in the BVM_GasPriceOracle predeploy
	L2GasPriceSlot = common.BigToHash(big.NewInt(1))
	// L1GasPriceSlot refers to the storage slot that the L1 gas price is stored
	// in in the BVM_GasPriceOracle predeploy
	L1GasPriceSlot = common.BigToHash(big.NewInt(2))
	// L2GasPriceOracleOwnerSlot refers to the storage slot that the owner of
	// the BVM_GasPriceOracle is stored in
	L2GasPriceOracleOwnerSlot = common.BigToHash(big.NewInt(0))
	// L2GasPriceOracleAddress is the address of the BVM_GasPriceOracle
	// predeploy
	L2GasPriceOracleAddress = common.HexToAddress("0x420000000000000000000000000000000000000F")
	// code
	L2GasPriceOracleCode = []byte("0x60806040523480156200001157600080fd5b506040516200126638038062001266833981016040819052620000349162000176565b6200003f3362000051565b6200004a81620000a1565b50620001a8565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000546001600160a01b03163314620001015760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064015b60405180910390fd5b6001600160a01b038116620001685760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401620000f8565b620001738162000051565b50565b6000602082840312156200018957600080fd5b81516001600160a01b0381168114620001a157600080fd5b9392505050565b6110ae80620001b86000396000f3fe608060405234801561001057600080fd5b50600436106101985760003560e01c8063715018a6116100e3578063de26c4a11161008c578063f45e65d811610066578063f45e65d814610328578063fc55b13814610331578063fe173b971461034457600080fd5b8063de26c4a1146102e2578063ea01cd36146102f5578063f2fde38b1461031557600080fd5b8063bede39b5116100bd578063bede39b5146102a9578063bf1fe420146102bc578063c7647832146102cf57600080fd5b8063715018a61461024f5780638c8885c8146102575780638da5cb5b1461026a57600080fd5b806345c51a3811610145578063551619131161011f578063551619131461022a5780635cbe497a14610233578063704655971461023c57600080fd5b806345c51a38146101fb57806349948e0e1461020e578063519b4bd31461022157600080fd5b8063288005781161017657806328800578146101ca578063313ce567146101df5780633577afc5146101e857600080fd5b80630c18c1621461019d5780630d1e43a0146101b95780630e6faf1e146101c1575b600080fd5b6101a660035481565b6040519081526020015b60405180910390f35b6006546101a6565b6101a6600a5481565b6101dd6101d8366004610ccf565b61034d565b005b6101a660055481565b6101dd6101f6366004610ccf565b610450565b6101dd610209366004610ccf565b6104f3565b6101a661021c366004610d17565b61058f565b6101a660025481565b6101a660075481565b6101a660095481565b6101dd61024a366004610ccf565b6105eb565b6101dd610687565b6101dd610265366004610ccf565b6106fa565b60005473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101b0565b6101dd6102b7366004610ccf565b610796565b6101dd6102ca366004610ccf565b610832565b6101dd6102dd366004610ccf565b6108ce565b6101a66102f0366004610d17565b6109c4565b6008546102849073ffffffffffffffffffffffffffffffffffffffff1681565b6101dd610323366004610de6565b610a68565b6101a660045481565b6101dd61033f366004610ccf565b610b64565b6101a660015481565b60005473ffffffffffffffffffffffffffffffffffffffff1633146103b95760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064015b60405180910390fd5b808015806103c75750806001145b6104135760405162461bcd60e51b815260206004820152601c60248201527f696e76616c69642076616c75652c6d7573742062652030206f7220310000000060448201526064016103b0565b600a8290556040518281527f65cacb453bbeab72658947058c43b2a6c7dfcca1c9d96ba1bc470d346929b288906020015b60405180910390a15050565b60005473ffffffffffffffffffffffffffffffffffffffff1633146104b75760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016103b0565b60038190556040518181527f32740b35c0ea213650f60d44366b4fb211c9033b50714e4a1d34e65d5beb9bb4906020015b60405180910390a150565b60005473ffffffffffffffffffffffffffffffffffffffff16331461055a5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016103b0565b60098190556040518181527f5af81f5214eaf8c64101a8fde536abc770ef62af9e14d15e2b0b68760b2028f5906020016104e8565b60008061059b836109c4565b90506000600254826105ad9190610e52565b90506000600554600a6105c09190610fb1565b90506000600454836105d29190610e52565b905060006105e08383610fbd565b979650505050505050565b60005473ffffffffffffffffffffffffffffffffffffffff1633146106525760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016103b0565b60048190556040518181527f3336cd9708eaf2769a0f0dc0679f30e80f15dcd88d1921b5a16858e8b85c591a906020016104e8565b60005473ffffffffffffffffffffffffffffffffffffffff1633146106ee5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016103b0565b6106f86000610c5a565b565b60005473ffffffffffffffffffffffffffffffffffffffff1633146107615760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016103b0565b60058190556040518181527fd68112a8707e326d08be3656b528c1bcc5bbbfc47f4177e2179b14d8640838c1906020016104e8565b60005473ffffffffffffffffffffffffffffffffffffffff1633146107fd5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016103b0565b60028190556040518181527f351fb23757bb5ea0546c85b7996ddd7155f96b939ebaa5ff7bc49c75f27f2c44906020016104e8565b60005473ffffffffffffffffffffffffffffffffffffffff1633146108995760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016103b0565b60018190556040518181527ffcdccc6074c6c42e4bd578aa9870c697dc976a270968452d2b8c8dc369fae396906020016104e8565b60005473ffffffffffffffffffffffffffffffffffffffff1633146109355760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016103b0565b808015806109435750806001145b61098f5760405162461bcd60e51b815260206004820152601c60248201527f696e76616c69642076616c75652c6d7573742062652030206f7220310000000060448201526064016103b0565b60068290556040518281527fd1eaae13a99b475ddca546a1b4a45052c66c14049997f44a1731a8e7167981a790602001610444565b600080805b8351811015610a41578381815181106109e4576109e4610ff8565b01602001517fff0000000000000000000000000000000000000000000000000000000000000016610a2157610a1a600483611027565b9150610a2f565b610a2c601083611027565b91505b80610a398161103f565b9150506109c9565b50600060035482610a529190611027565b9050610a6081610440611027565b949350505050565b60005473ffffffffffffffffffffffffffffffffffffffff163314610acf5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016103b0565b73ffffffffffffffffffffffffffffffffffffffff8116610b585760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016103b0565b610b6181610c5a565b50565b60005473ffffffffffffffffffffffffffffffffffffffff163314610bcb5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016103b0565b80801580610bd95750806001145b610c255760405162461bcd60e51b815260206004820152601c60248201527f696e76616c69642076616c75652c6d7573742062652030206f7220310000000060448201526064016103b0565b60078290556040518281527f49244d4195584d0644398167ca8caa7b98ee36b674e4b4d2a2640749b27eafb790602001610444565b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b600060208284031215610ce157600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600060208284031215610d2957600080fd5b813567ffffffffffffffff80821115610d4157600080fd5b818401915084601f830112610d5557600080fd5b813581811115610d6757610d67610ce8565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908382118183101715610dad57610dad610ce8565b81604052828152876020848701011115610dc657600080fd5b826020860160208301376000928101602001929092525095945050505050565b600060208284031215610df857600080fd5b813573ffffffffffffffffffffffffffffffffffffffff81168114610e1c57600080fd5b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615610e8a57610e8a610e23565b500290565b600181815b80851115610ee857817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115610ece57610ece610e23565b80851615610edb57918102915b93841c9390800290610e94565b509250929050565b600082610eff57506001610fab565b81610f0c57506000610fab565b8160018114610f225760028114610f2c57610f48565b6001915050610fab565b60ff841115610f3d57610f3d610e23565b50506001821b610fab565b5060208310610133831016604e8410600b8410161715610f6b575081810a610fab565b610f758383610e8f565b807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115610fa757610fa7610e23565b0290505b92915050565b6000610e1c8383610ef0565b600082610ff3577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000821982111561103a5761103a610e23565b500190565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82141561107157611071610e23565b506001019056fea26469706673582212209beb66345180150025622f61beb9aa5442283bdfe3de513df469a2c632f10d4e64736f6c63430008090033")
	// OverheadSlot refers to the storage slot in the BVM_GasPriceOracle that
	// holds the per transaction overhead. This is added to the L1 cost portion
	// of the fee
	OverheadSlot = common.BigToHash(big.NewInt(3))
	// ScalarSlot refers to the storage slot in the BVM_GasPriceOracle that
	// holds the transaction fee scalar. This value is scaled upwards by
	// the number of decimals
	ScalarSlot = common.BigToHash(big.NewInt(4))
	// DecimalsSlot refers to the storage slot in the BVM_GasPriceOracle that
	// holds the number of decimals in the fee scalar
	DecimalsSlot = common.BigToHash(big.NewInt(5))
	// IsBurningSlot refers to the storage slot in the BVM_GasPriceOracle that
	// holds switch controls whether GasFee is brun
	IsBurningSlot = common.BigToHash(big.NewInt(6))
	// ChargeSlot refers to the storage slot in the BVM_GasPriceOracle that
	// holds switch controls whether charge
	ChargeSlot = common.BigToHash(big.NewInt(7))
	// SccAddressSlot refers to the storage slot in the Scc contract address
	SccAddressSlot = common.BigToHash(big.NewInt(8))
	// DaGasPriceSlot refers to the storage slot that the da gas price is stored
	// in the BVM_GasPriceOracle predeploy
	DaGasPriceSlot = common.BigToHash(big.NewInt(9))
	// DaSwitchSlot refers to the storage slot in the BVM_GasPriceOracle that
	// holds switch controls whether enable DA
	DaSwitchSlot = common.BigToHash(big.NewInt(10))
)

func init() {
	UsingBVM = os.Getenv("USING_BVM") == "true"
}

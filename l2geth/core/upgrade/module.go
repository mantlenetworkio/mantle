package upgrade

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/mantlenetworkio/mantle/l2geth/common"
	"github.com/mantlenetworkio/mantle/l2geth/core/vm"
	"github.com/mantlenetworkio/mantle/l2geth/ethdb"
	"github.com/mantlenetworkio/mantle/l2geth/log"
	"github.com/mantlenetworkio/mantle/l2geth/params"
	"github.com/mantlenetworkio/mantle/l2geth/rollup/rcfg"
)

/*
 * There are two scenarios for the upgrade operation.
 * 1. The first scenario is to execute an upgrade operation at a specific height, such as 'setCode'.
 * 2. The second scenario is to follow the previous logic before a certain height,
 * and switch to new processing logic after that height.

 * For the first scenario, we use 'isExactBlockForked' to determine the upgrade height and ensure
 * that the operation is triggered only once at a specific height.
 * For the second scenario, we use 'isBlockForked' to determine the upgrade height and ensure
 * that different logic is executed before and after the upgrade.

 * In addition, we have two options for selecting the upgrade height: L1 block and L2 block.
 * For L2 blocks, they are strictly increasing and consecutive. We need to choose between 'isBlockForked'
 * and 'isExactBlockForked' based on the specific situation.
 * For L1 blocks, they are synchronized from the sync service, and there are two possible scenarios.
 * 1. In the first scenario, the synchronized blocks may not be consecutive, resulting in block gaps.
 * 2. In the second scenario, there may be no corresponding L2 block generated at the same time as the L1 block.
 * So we need to use 'isBlockForked' to ensure that the upgrade operation is always executed.
 * We may also need to keep a record of the upgrade operation to prevent it from being executed repeatedly.
 */

var (
	// L2TssRewardAddress is the address of the TssRewardContract
	// predeploy
	L2TssRewardAddress = common.HexToAddress("0x4200000000000000000000000000000000000020")
	// code Compile from commit 52b6bbc52bbad5b971fd51a76d1b9b87d342e9f3
	L2GasRewardCode, _ = hexutil.Decode("0x6080604052600436106101f25760003560e01c8063715018a61161010d578063d8111a57116100a0578063ea01cd361161006f578063ea01cd361461058e578063ebc73e65146105ae578063f2fde38b146105ce578063f5cf673b146105ee578063fad9aba31461060e57600080fd5b8063d8111a571461050c578063da62fba914610522578063e5efd58514610558578063e91a469e1461056e57600080fd5b8063b89ea402116100dc578063b89ea40214610478578063ceddc021146104a5578063cf172403146104ca578063d4440a36146104f757600080fd5b8063715018a6146104195780638da5cb5b1461042e578063904ad6be1461044c578063aa13e8c21461046257600080fd5b806327c8f835116101855780633cb747bf116101545780633cb747bf146103a25780633ccfd60b146103c25780634e71d92d146103d757806352674ea5146103ec57600080fd5b806327c8f8351461032f5780632c79db111461034f5780633310ceeb146103625780633b52c31e1461038257600080fd5b806315c6f166116101c157806315c6f166146102ce57806319d509a1146102e35780631a39d8ef146102f9578063208695eb1461030f57600080fd5b80630fae75d9146101fe57806310a7fd7b14610220578063110b7eb01461026057806313e7c9d81461029857600080fd5b366101f957005b600080fd5b34801561020a57600080fd5b5061021e61021936600461127f565b610624565b005b34801561022c57600080fd5b5061024d61023b366004611321565b60346020526000908152604090205481565b6040519081526020015b60405180910390f35b34801561026c57600080fd5b50603554610280906001600160a01b031681565b6040516001600160a01b039091168152602001610257565b3480156102a457600080fd5b506102806102b336600461134f565b6040602081905260009182529020546001600160a01b031681565b3480156102da57600080fd5b5061024d61083f565b3480156102ef57600080fd5b5061024d60385481565b34801561030557600080fd5b5061024d60395481565b34801561031b57600080fd5b50604454610280906001600160a01b031681565b34801561033b57600080fd5b50603654610280906001600160a01b031681565b34801561035b57600080fd5b504761024d565b34801561036e57600080fd5b5061021e61037d36600461134f565b61086a565b34801561038e57600080fd5b5061021e61039d366004611321565b6108b6565b3480156103ae57600080fd5b50600154610280906001600160a01b031681565b3480156103ce57600080fd5b5061021e6108e5565b3480156103e357600080fd5b5061021e610953565b3480156103f857600080fd5b5061024d61040736600461134f565b60436020526000908152604090205481565b34801561042557600080fd5b5061021e610a8f565b34801561043a57600080fd5b506000546001600160a01b0316610280565b34801561045857600080fd5b5061024d603d5481565b34801561046e57600080fd5b5061024d603e5481565b34801561048457600080fd5b5061024d61049336600461134f565b60426020526000908152604090205481565b3480156104b157600080fd5b506104ba610ac3565b6040519015158152602001610257565b3480156104d657600080fd5b5061024d6104e536600461134f565b603f6020526000908152604090205481565b34801561050357600080fd5b5061024d610bf2565b34801561051857600080fd5b5061024d603b5481565b34801561052e57600080fd5b5061028061053d36600461134f565b6041602052600090815260409020546001600160a01b031681565b34801561056457600080fd5b5061024d603a5481565b34801561057a57600080fd5b5061021e61058936600461134f565b610c1c565b34801561059a57600080fd5b50603c54610280906001600160a01b031681565b3480156105ba57600080fd5b5061021e6105c9366004611321565b610c68565b3480156105da57600080fd5b5061021e6105e936600461134f565b610c97565b3480156105fa57600080fd5b5061021e61060936600461136c565b610d2f565b34801561061a57600080fd5b5061024d60375481565b603c546001600160a01b03166106426001546001600160a01b031690565b6001600160a01b0316336001600160a01b03161461067b5760405162461bcd60e51b8152600401610672906113a5565b60405180910390fd5b806001600160a01b03166106976001546001600160a01b031690565b6001600160a01b0316636e296e456040518163ffffffff1660e01b815260040160206040518083038186803b1580156106cf57600080fd5b505afa1580156106e3573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061070791906113f3565b6001600160a01b03161461072d5760405162461bcd60e51b815260040161067290611410565b6000603a54600014156107455750603a849055610837565b603a5485116107b15760405162461bcd60e51b815260206004820152603260248201527f61726773205f626174636854696d65206d757374206774686572207468616e206044820152716c617374206c617374426174636854696d6560701b6064820152608401610672565b6037546107bc61083f565b603a546107c99088611476565b6107d3919061148d565b6107dd91906114ac565b600060375590506107ef818585610ece565b7ff533ef50019763ee9d95ad46e28350b533c11edd472ae7be93e8fae83c1b6d99603a54868387876040516108289594939291906114c4565b60405180910390a150603a8490555b505050505050565b60006108656301e13380603b54670de0b6b3a764000061085f919061148d565b90610f98565b905090565b6000546001600160a01b031633146108945760405162461bcd60e51b81526004016106729061152d565b604480546001600160a01b0319166001600160a01b0392909216919091179055565b6000546001600160a01b031633146108e05760405162461bcd60e51b81526004016106729061152d565b603b55565b6000546001600160a01b0316331461090f5760405162461bcd60e51b81526004016106729061152d565b471561095157600080546040516001600160a01b03909116914780156108fc02929091818181858888f1935050505015801561094f573d6000803e3d6000fd5b505b565b336000908152604160205260409020546001600160a01b0316806109895760405162461bcd60e51b815260040161067290611562565b6002805414156109db5760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610672565b60028055336000908152604160205260408120546001600160a01b031690610a0282610fa4565b90508015610a655760405162461bcd60e51b815260206004820152602a60248201527f706c65617365207761697420666f72207468652077616974696e6720706572696044820152696f6420746f207061737360b01b6064820152608401610672565b610a6e82611054565b506001600160a01b0316600090815260426020526040812055506001600255565b6000546001600160a01b03163314610ab95760405162461bcd60e51b81526004016106729061152d565b6109516000611217565b336000908152604160205260408120546001600160a01b031680610af95760405162461bcd60e51b815260040161067290611562565b336000908152604160209081526040808320546001600160a01b03168084526042909252909120548015610bbb5760405162461bcd60e51b815260206004820152605960248201527f596f75206861766520616c726561647920696e6974696174656420612072657160448201527f7565737420746f20636c61696d2c20706c65617365207761697420666f72207460648201527f68652077616974696e6720706572696f6420746f207061737300000000000000608482015260a401610672565b506001600160a01b03166000908152604260209081526040808320429055603f825280832054604390925290912055600191505090565b336000908152604160205260408120546001600160a01b031681610c1582610fa4565b9392505050565b6000546001600160a01b03163314610c465760405162461bcd60e51b81526004016106729061152d565b603c80546001600160a01b0319166001600160a01b0392909216919091179055565b6000546001600160a01b03163314610c925760405162461bcd60e51b81526004016106729061152d565b603e55565b6000546001600160a01b03163314610cc15760405162461bcd60e51b81526004016106729061152d565b6001600160a01b038116610d265760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610672565b61094f81611217565b6044546001600160a01b0316610d4d6001546001600160a01b031690565b6001600160a01b0316336001600160a01b031614610d7d5760405162461bcd60e51b8152600401610672906113a5565b806001600160a01b0316610d996001546001600160a01b031690565b6001600160a01b0316636e296e456040518163ffffffff1660e01b815260040160206040518083038186803b158015610dd157600080fd5b505afa158015610de5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e0991906113f3565b6001600160a01b031614610e2f5760405162461bcd60e51b815260040161067290611410565b6001600160a01b038381166000908152604060208190529020541615610e84576001600160a01b0380841660009081526040602081815281832054909316825260419092522080546001600160a01b03191690555b506001600160a01b0390811660008181526041602090815260408083208054959096166001600160a01b031995861681179096559482528490529290922080549091169091179055565b8215610f9357600080610ee18584610f98565b915060005b83811015610f5e576000858583818110610f0257610f026115b3565b9050602002016020810190610f17919061134f565b6001600160a01b0381166000908152603f6020526040812080549293508692909190610f449084906114ac565b90915550829150610f569050816115c9565b915050610ee6565b50610f69838361148d565b90506000610f778683611267565b9050801561083757603754610f8c9082611273565b6037555050505b505050565b6000610c1582846115e4565b6001600160a01b0381166000908152604260205260408120548061101b5760405162461bcd60e51b815260206004820152602860248201527f706c6561736520696e6974696174652061207265717565737420746f20636c616044820152671a5b48199a5c9cdd60c21b6064820152608401610672565b6000603e548261102b91906114ac565b9050600042821115611048576110414283611476565b905061104c565b5060005b949350505050565b6001600160a01b038116600090815260436020908152604080832054603f90925290912054818110156110c95760405162461bcd60e51b815260206004820181905260248201527f546865206e756d65726963616c2076616c756520697320696e636f72726563746044820152606401610672565b8147101561113f5760405162461bcd60e51b815260206004820152603c60248201527f54686520636f6e74726163742062616c616e636520697320696e73756666696360448201527f69656e7420746f2070617920746865207265776172642076616c7565000000006064820152608401610672565b81156111d0576001600160a01b038084166000908152604060208181528183205460438252828420849055603f9091529120549116908190611182908590611476565b6001600160a01b038087166000908152603f60205260408082209390935591519083169186156108fc02918791818181858888f193505050501580156111cc573d6000803e3d6000fd5b5050505b604080516001600160a01b0385168152602081018490527f47cee97cb7acd717b3c0aa1435d004cd5b3c8c57d70dbceb4e4458bbd60e39d4910160405180910390a1505050565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000610c158284611476565b6000610c1582846114ac565b60008060008060006080868803121561129757600080fd5b85359450602086013563ffffffff811681146112b257600080fd5b935060408601359250606086013567ffffffffffffffff808211156112d657600080fd5b818801915088601f8301126112ea57600080fd5b8135818111156112f957600080fd5b8960208260051b850101111561130e57600080fd5b9699959850939650602001949392505050565b60006020828403121561133357600080fd5b5035919050565b6001600160a01b038116811461094f57600080fd5b60006020828403121561136157600080fd5b8135610c158161133a565b6000806040838503121561137f57600080fd5b823561138a8161133a565b9150602083013561139a8161133a565b809150509250929050565b6020808252602e908201527f42564d5f58434841494e3a206d657373656e67657220636f6e7472616374207560408201526d1b985d5d1a195b9d1a58d85d195960921b606082015260800190565b60006020828403121561140557600080fd5b8151610c158161133a565b60208082526030908201527f42564d5f58434841494e3a2077726f6e672073656e646572206f662063726f7360408201526f732d646f6d61696e206d65737361676560801b606082015260800190565b634e487b7160e01b600052601160045260246000fd5b60008282101561148857611488611460565b500390565b60008160001904831182151516156114a7576114a7611460565b500290565b600082198211156114bf576114bf611460565b500190565b60006080820187835260208781850152866040850152608060608501528185835260a08501905086925060005b8681101561151f5783356115048161133a565b6001600160a01b0316825292820192908201906001016114f1565b509998505050505050505050565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b60208082526031908201527f546865206d73672073656e646572206973206e6f7420617574686f72697a656460408201527010313c903a3432903b30b634b230ba37b960791b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b60006000198214156115dd576115dd611460565b5060010190565b60008261160157634e487b7160e01b600052601260045260246000fd5b50049056fea2646970667358221220c1c3d22331acc8f305fd9a0684ef99bc996384985748db9533152498b0af0e9f64736f6c63430008090033")
)

func CheckUpgrade(chainID *big.Int, statedb vm.StateDB, ethdb ethdb.Database, l2BlockNumber *big.Int, l1BlockNumber *big.Int) {
	log.Info("CheckUpgrade", "chain id", chainID)
	var chainConfig *Config
	switch chainID.Int64() {
	case params.MantleMainnetChainID.Int64():
		chainConfig = MainnetConfig
	case params.MantleTestnetChainID.Int64():
		chainConfig = TestnetConfig
	case params.MantleQAChainID.Int64():
		chainConfig = QAConfig
	case params.MantleLocalChainID.Int64():
		chainConfig = LocalConfig
	default:
		chainConfig = LocalConfig
	}

	// upgrade based on layer2 block height
	l2TssRewardUpgrade(statedb, chainConfig, l2BlockNumber)
	mantleTokenUpgrade(statedb, chainConfig, l2BlockNumber)
	eigenDaUpgrade(statedb, chainConfig, l2BlockNumber)

	// mock an upgrade based on layer1 block height
	mockUpgradeBasedOnL1Block(statedb, ethdb, chainConfig, l1BlockNumber)
}

func l2TssRewardUpgrade(statedb vm.StateDB, chainConfig *Config, blockNumber *big.Int) {
	if !chainConfig.IsTssReward(blockNumber) {
		return
	}

	statedb.SetCode(L2TssRewardAddress, L2GasRewardCode)
	log.Info("update l2 tss reward success", "address", L2TssRewardAddress)
}

func mantleTokenUpgrade(statedb vm.StateDB, chainConfig *Config, blockNumber *big.Int) {
	if !chainConfig.IsMantleToken(blockNumber) {
		return
	}

	statedb.SetState(rcfg.L2MantleTokenAddress, rcfg.MantleTokenNameSlot, rcfg.MantleTokenNameValue)
	statedb.SetState(rcfg.L2MantleTokenAddress, rcfg.MantleTokenSymbolSlot, rcfg.MantleTokenSymbolValue)
	log.Info("update mantle token name & symbol", "address", rcfg.L2MantleTokenAddress)
}

func eigenDaUpgrade(statedb vm.StateDB, chainConfig *Config, blockNumber *big.Int) {
	if !chainConfig.IsEigenDa(blockNumber) {
		return
	}

	statedb.SetCode(rcfg.L2GasPriceOracleAddress, rcfg.L2GasPriceOracleCode)
	log.Info("update eigen data success", "address", rcfg.L2GasPriceOracleAddress)
}

func mockUpgradeBasedOnL1Block(statedb vm.StateDB, ethdb ethdb.Database, chainConfig *Config, blockNumber *big.Int) {
	if !chainConfig.IsMockUpgradeBasedOnL1BlockNumber(blockNumber) {
		return
	}

	// To prevent the upgrade operation from being executed repeatedly,
	// we need to check whether the upgrade has already taken place.
	if ExistUpgradeFlag(ethdb, mockUpgradeFlag) {
		return
	}

	log.Info("update mantle token name & symbol", "address", rcfg.L2MantleTokenAddress)
	statedb.SetState(rcfg.L2MantleTokenAddress, rcfg.MantleTokenNameSlot, rcfg.MantleTokenNameValue)
	statedb.SetState(rcfg.L2MantleTokenAddress, rcfg.MantleTokenSymbolSlot, rcfg.MantleTokenSymbolValue)

	WriteUpgradeFlag(ethdb, mockUpgradeFlag, blockNumber.Bytes())
}

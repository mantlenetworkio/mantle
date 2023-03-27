// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/mantlenetworkio/mantle/mt-bindings/solc"
)

const L2ERC721BridgeStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/L2/L2ERC721Bridge.sol:L2ERC721Bridge\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_array(t_uint256)1001_storage\"}],\"types\":{\"t_array(t_uint256)1001_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"}}}"

var L2ERC721BridgeStorageLayout = new(solc.StorageLayout)

var L2ERC721BridgeDeployedBin = "0x608060405234801561001057600080fd5b50600436106100885760003560e01c80637f46ddb21161005b5780637f46ddb214610116578063927ede2d1461013d578063aa55745214610164578063c89701a21461017757600080fd5b80633687011a1461008d5780633cb747bf146100a257806354fd4d50146100ee578063761f449314610103575b600080fd5b6100a061009b3660046111c8565b61019d565b005b7f00000000000000000000000000000000000000000000000000000000000000005b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6100f6610249565b6040516100e591906112c5565b6100a06101113660046112d8565b6102ec565b6100c47f000000000000000000000000000000000000000000000000000000000000000081565b6100c47f000000000000000000000000000000000000000000000000000000000000000081565b6100a0610172366004611370565b610853565b7f00000000000000000000000000000000000000000000000000000000000000006100c4565b333b15610231576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f4552433732314272696467653a206163636f756e74206973206e6f742065787460448201527f65726e616c6c79206f776e65640000000000000000000000000000000000000060648201526084015b60405180910390fd5b610241868633338888888861090f565b505050505050565b60606102747f0000000000000000000000000000000000000000000000000000000000000000610ead565b61029d7f0000000000000000000000000000000000000000000000000000000000000000610ead565b6102c67f0000000000000000000000000000000000000000000000000000000000000000610ead565b6040516020016102d8939291906113e7565b604051602081830303815290604052905090565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614801561040a57507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa1580156103ce573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103f2919061145d565b73ffffffffffffffffffffffffffffffffffffffff16145b610496576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603f60248201527f4552433732314272696467653a2066756e6374696f6e2063616e206f6e6c792060448201527f62652063616c6c65642066726f6d20746865206f7468657220627269646765006064820152608401610228565b3073ffffffffffffffffffffffffffffffffffffffff88160361053b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f4c324552433732314272696467653a206c6f63616c20746f6b656e2063616e6e60448201527f6f742062652073656c66000000000000000000000000000000000000000000006064820152608401610228565b610565877f74259ebf00000000000000000000000000000000000000000000000000000000610fea565b6105f1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603660248201527f4c324552433732314272696467653a206c6f63616c20746f6b656e20696e746560448201527f7266616365206973206e6f7420636f6d706c69616e74000000000000000000006064820152608401610228565b8673ffffffffffffffffffffffffffffffffffffffff1663d6c0b2c46040518163ffffffff1660e01b8152600401602060405180830381865afa15801561063c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610660919061145d565b73ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff1614610740576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604960248201527f4c324552433732314272696467653a2077726f6e672072656d6f746520746f6b60448201527f656e20666f72204d616e746c65204d696e7461626c6520455243373231206c6f60648201527f63616c20746f6b656e0000000000000000000000000000000000000000000000608482015260a401610228565b6040517fa144819400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85811660048301526024820185905288169063a144819490604401600060405180830381600087803b1580156107b057600080fd5b505af11580156107c4573d6000803e3d6000fd5b505050508473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff167f1f39bf6707b5d608453e0ae4c067b562bcc4c85c0f562ef5d2c774d2e7f131ac8787878760405161084294939291906114c3565b60405180910390a450505050505050565b73ffffffffffffffffffffffffffffffffffffffff85166108f6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f4552433732314272696467653a206e667420726563697069656e742063616e6e60448201527f6f742062652061646472657373283029000000000000000000000000000000006064820152608401610228565b610906878733888888888861090f565b50505050505050565b73ffffffffffffffffffffffffffffffffffffffff87166109b2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603160248201527f4c324552433732314272696467653a2072656d6f746520746f6b656e2063616e60448201527f6e6f7420626520616464726573732830290000000000000000000000000000006064820152608401610228565b6040517f6352211e0000000000000000000000000000000000000000000000000000000081526004810185905273ffffffffffffffffffffffffffffffffffffffff891690636352211e90602401602060405180830381865afa158015610a1d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a41919061145d565b73ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff1614610afb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603e60248201527f4c324552433732314272696467653a205769746864726177616c206973206e6f60448201527f74206265696e6720696e69746961746564206279204e4654206f776e657200006064820152608401610228565b60008873ffffffffffffffffffffffffffffffffffffffff1663d6c0b2c46040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b48573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b6c919061145d565b90508773ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610c29576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f4c324552433732314272696467653a2072656d6f746520746f6b656e20646f6560448201527f73206e6f74206d6174636820676976656e2076616c75650000000000000000006064820152608401610228565b6040517f9dc29fac00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8881166004830152602482018790528a1690639dc29fac90604401600060405180830381600087803b158015610c9957600080fd5b505af1158015610cad573d6000803e3d6000fd5b50505050600063761f449360e01b828b8a8a8a8989604051602401610cd89796959493929190611503565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009094169390931790925290517f3dbb202b00000000000000000000000000000000000000000000000000000000815290915073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690633dbb202b90610ded907f00000000000000000000000000000000000000000000000000000000000000009085908a90600401611560565b600060405180830381600087803b158015610e0757600080fd5b505af1158015610e1b573d6000803e3d6000fd5b505050508773ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff168b73ffffffffffffffffffffffffffffffffffffffff167fb7460e2a880f256ebef3406116ff3eee0cee51ebccdc2a40698f87ebb2e9c1a58a8a8989604051610e9994939291906114c3565b60405180910390a450505050505050505050565b606081600003610ef057505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b8115610f1a5780610f04816115d4565b9150610f139050600a8361163b565b9150610ef4565b60008167ffffffffffffffff811115610f3557610f3561164f565b6040519080825280601f01601f191660200182016040528015610f5f576020820181803683370190505b5090505b8415610fe257610f7460018361167e565b9150610f81600a86611695565b610f8c9060306116a9565b60f81b818381518110610fa157610fa16116c1565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350610fdb600a8661163b565b9450610f63565b949350505050565b6000610ff58361100d565b801561100657506110068383611072565b9392505050565b6000611039827f01ffc9a700000000000000000000000000000000000000000000000000000000611072565b801561106c575061106a827fffffffff00000000000000000000000000000000000000000000000000000000611072565b155b92915050565b604080517fffffffff000000000000000000000000000000000000000000000000000000008316602480830191909152825180830390910181526044909101909152602080820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f01ffc9a700000000000000000000000000000000000000000000000000000000178152825160009392849283928392918391908a617530fa92503d9150600051905082801561112a575060208210155b80156111365750600081115b979650505050505050565b73ffffffffffffffffffffffffffffffffffffffff8116811461116357600080fd5b50565b803563ffffffff8116811461117a57600080fd5b919050565b60008083601f84011261119157600080fd5b50813567ffffffffffffffff8111156111a957600080fd5b6020830191508360208285010111156111c157600080fd5b9250929050565b60008060008060008060a087890312156111e157600080fd5b86356111ec81611141565b955060208701356111fc81611141565b94506040870135935061121160608801611166565b9250608087013567ffffffffffffffff81111561122d57600080fd5b61123989828a0161117f565b979a9699509497509295939492505050565b60005b8381101561126657818101518382015260200161124e565b83811115611275576000848401525b50505050565b6000815180845261129381602086016020860161124b565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000611006602083018461127b565b600080600080600080600060c0888a0312156112f357600080fd5b87356112fe81611141565b9650602088013561130e81611141565b9550604088013561131e81611141565b9450606088013561132e81611141565b93506080880135925060a088013567ffffffffffffffff81111561135157600080fd5b61135d8a828b0161117f565b989b979a50959850939692959293505050565b600080600080600080600060c0888a03121561138b57600080fd5b873561139681611141565b965060208801356113a681611141565b955060408801356113b681611141565b9450606088013593506113cb60808901611166565b925060a088013567ffffffffffffffff81111561135157600080fd5b600084516113f981846020890161124b565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551611435816001850160208a0161124b565b6001920191820152835161145081600284016020880161124b565b0160020195945050505050565b60006020828403121561146f57600080fd5b815161100681611141565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff851681528360208201526060604082015260006114f960608301848661147a565b9695505050505050565b600073ffffffffffffffffffffffffffffffffffffffff808a1683528089166020840152808816604084015280871660608401525084608083015260c060a083015261155360c08301848661147a565b9998505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff8416815260606020820152600061158f606083018561127b565b905063ffffffff83166040830152949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611605576116056115a5565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60008261164a5761164a61160c565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082821015611690576116906115a5565b500390565b6000826116a4576116a461160c565b500690565b600082198211156116bc576116bc6115a5565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(L2ERC721BridgeStorageLayoutJSON), L2ERC721BridgeStorageLayout); err != nil {
		panic(err)
	}

	layouts["L2ERC721Bridge"] = L2ERC721BridgeStorageLayout
	deployedBytecodes["L2ERC721Bridge"] = L2ERC721BridgeDeployedBin
}

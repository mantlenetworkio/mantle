// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import '../../../../../contracts/libraries/resolver/Lib_AddressManager.sol';
import '../../../../../contracts/L1/messaging/L1CrossDomainMessenger.sol';

import '../../../../../contracts/L1/local/TestBitToken.sol';

import '../../../../../contracts/L1/tss/delegation/TssDelegation.sol';
import '../../../../../contracts/L1/tss/delegation/TssDelegationManager.sol';
import '../../../../../contracts/L1/tss/delegation/TssDelegationSlasher.sol';
import '../../../../../contracts/L1/tss/TssGroupManager.sol';
import '../../../../../contracts/L1/tss/TssStakingSlashing.sol';
import '../../../../../contracts/L1/delegation/DelegationSlasher.sol';


import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import "forge-std/Test.sol";
import "../../../mocks/EmptyContract.sol";

contract TssDelegationTest is Test {

    ProxyAdmin public proxyAdmin;

    Lib_AddressManager public libAddressManager;
    L1CrossDomainMessenger public l1CrossDomainMessenger;

    BitTokenERC20 public l1Bit;

    TssDelegation public tssDelegation;
    TssDelegationSlasher public tssDelegationSlasher;
    TssDelegationManager public tssDelegationManager;
    TssGroupManager public tssGroupManager;
    TssStakingSlashing public tssStakingSlashing;


    uint256 _minStakeAmount = 10e10;
    address initialOwner = address(this);

    function setUp() public {

        l1Bit = new BitTokenERC20("BitToken", "BIT");
        proxyAdmin = new ProxyAdmin();
        EmptyContract emptyContract = new EmptyContract();

        l1CrossDomainMessenger = L1CrossDomainMessenger(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), ""))
        );
        tssDelegation = TssDelegation(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), ""))
        );
        tssDelegationSlasher = TssDelegationSlasher(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), ""))
        );
        tssDelegationManager = TssDelegationManager(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), ""))
        );
        tssGroupManager = TssGroupManager(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), ""))
        );
        tssGroupManager = TssGroupManager(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), ""))
        );
        tssStakingSlashing = TssStakingSlashing(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), ""))
        );

        L1CrossDomainMessenger l1CrossDomainMessengerImplementation = new L1CrossDomainMessenger();
        proxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(l1CrossDomainMessenger))),
            address(l1CrossDomainMessengerImplementation),
            abi.encodeWithSelector(L1CrossDomainMessenger.initialize.selector, address(libAddressManager), initialOwner)
        );

        TssDelegation tssDelegationImplementation = new TssDelegation(tssDelegationManager);
        proxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(tssDelegation))),
            address(tssDelegationImplementation),
            abi.encodeWithSelector(TssDelegation.initializeT.selector, address(tssStakingSlashing), initialOwner)
        );

        TssDelegationSlasher tssDelegationSlasherImplementation = new TssDelegationSlasher(tssDelegationManager, tssDelegation);
        proxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(tssDelegationSlasher))),
            address(tssDelegationSlasherImplementation),
            abi.encodeWithSelector(DelegationSlasher.initialize.selector, initialOwner)
        );

        TssDelegationManager tssDelegationManagerImplementation = new TssDelegationManager(tssDelegation, tssDelegationSlasher);
        proxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(tssDelegationManager))),
            address(tssDelegationManagerImplementation),
            abi.encodeWithSelector(TssDelegationManager.initializeT.selector, address(tssStakingSlashing), address(tssGroupManager), _minStakeAmount,initialOwner)
        );

        TssGroupManager tssGroupManagerImplementation = new TssGroupManager();
        proxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(tssGroupManager))),
            address(tssGroupManagerImplementation),
            abi.encodeWithSelector(TssGroupManager.initialize.selector, initialOwner)
        );

        TssStakingSlashing tssStakingSlashingImplementation = new TssStakingSlashing();
        proxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(tssStakingSlashing))),
            address(tssStakingSlashingImplementation),
            abi.encodeWithSelector(TssStakingSlashing.initialize.selector, address(l1Bit), address(tssGroupManager), address(tssDelegationManager), address(tssDelegation), address(l1CrossDomainMessenger), address(1), initialOwner)
        );
    }

    function testInfo() public {
        assertEq(tssDelegation.stakingSlash(), address(tssStakingSlashing));
    }

    address[] public addresses;
    address[] public nodeAddresses;
    address[] public rmAddresses;
    function testWhiteListOnlyOwner() public {

        vm.expectRevert("Ownable: caller is not the owner");
        vm.prank(address(0));
        tssDelegation.addToWhitelist(addresses);
    }

    function testRemoveFromWhitelistOnleOwner() public {

        // vm.expectRevert("Ownable: caller is not the owner");
        vm.prank(address(0));
        tssDelegation.removeFromWhitelist(addresses);
    }

    function testWhiteList() public {
        addresses.push(address(1));
        addresses.push(address(2));
        addresses.push(address(2));
        addresses.push(address(4));

        rmAddresses.push(address(1));
        rmAddresses.push(address(2));

        assertEq(tssDelegation.whitelist(address(1)), false);

        tssDelegation.addToWhitelist(addresses);
        assertEq(tssDelegation.whitelist(address(1)), true);
        assertEq(tssDelegation.whitelist(address(2)), true);
        assertEq(tssDelegation.whitelist(address(3)), false);
        assertEq(tssDelegation.whitelist(address(4)), true);

        tssDelegation.removeFromWhitelist(rmAddresses);
        assertEq(tssDelegation.whitelist(address(1)), false);
        assertEq(tssDelegation.whitelist(address(2)), false);
        assertEq(tssDelegation.whitelist(address(3)), false);
        assertEq(tssDelegation.whitelist(address(4)), true);
    }

    function testRegisterAsOperator() public {

        // tssDelegation.registerAsOperator(address(1),address(2));
        string memory publicKey = "048318535b54105d4a7aae60c08fc45f9687181b4fdfc625bd1a753fa7397fed753547f11ca8696646f2f3acb08e31016afac23e630c5d11f59f61fef57b0d2aa5";

        addresses.push(address(this));
        tssDelegation.addToWhitelist(addresses);
        // tssDelegationManager.addToWhitelist(addresses);
        tssStakingSlashing.registerAsOperator(bytes(publicKey));
    }

    /*
    1、获取 deployer 私钥
    2、Proxy__TSS_StakingSlashing、Proxy__TssDelegation、Proxy__TssDelegationManager 的地址
    3、deployer 向 Proxy__TssDelegationManager 发起 addToWhitelist 请求，入参 Proxy__TSS_StakingSlashing
    4、获取 tssnode* 的所有地址，组装成 nodesAddrs 
    5、deployer 向 Proxy__TssDelegation 发起 addToWhitelist 请求，入参 nodesAddrs
    6、依次向tssnode* 转入l1 的 bit 和 ETH
    7、tssnode* 向 Proxy__TssDelegationManager approve l1 bit
    8、获取 tssnode* 的私钥
    9、tssnode* 向 Proxy__TSS_StakingSlashing 发起 deposit
    10、tssnode* 向 Proxy__TSS_StakingSlashing 发起 registerAsOperator

    11、获取 deployer 私钥
    12、deployer 向 Proxy__TSS_StakingSlashing 发起 setSlashingParams
    13、获取 tssnode* 的 pubKey 组装 pubKeyList
    14、deployer 向 Proxy__TSS_GroupManager 发起 setTssGroupMember
    */

    // function testDesposit() public {
    //     tssStakingSlashing.deposit
    // }

    function testRunTssDelegation() public {
        address deployer = address(0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266);
        addresses.push(address(tssStakingSlashing));

        assertEq(tssDelegationManager.owner(), address(this));
        tssDelegationManager.transferOwnership(deployer);
        assertEq(tssDelegationManager.owner(), deployer);

        vm.prank(deployer);
        // 3
        tssDelegationManager.addToWhitelist(addresses);
        assertEq(tssDelegationManager.whitelist(address(tssStakingSlashing)), true);

        nodeAddresses.push(address(tssStakingSlashing));
    }
}
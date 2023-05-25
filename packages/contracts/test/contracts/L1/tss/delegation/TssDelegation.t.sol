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
import '../../../../../contracts/L1/tss/ITssGroupManager.sol';
import '../../../../../contracts/L1/delegation/interfaces/IDelegation.sol';
import '../../../mocks/EmptyContract.sol';

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import "forge-std/Test.sol";
import {console} from "forge-std/console.sol";

contract TssDelegationTest is Test {

    ProxyAdmin public proxyAdmin;
    TransparentUpgradeableProxy public proxy;
    TransparentUpgradeableProxy public proxy_tssDelegationManager;
    Lib_AddressManager public libAddressManager;
    L1CrossDomainMessenger public l1CrossDomainMessenger;
    EmptyContract public empt;

    BitTokenERC20 public l1Bit;

    TssDelegation public tssDelegation;
    TssDelegationSlasher public tssDelegationSlasher;
    TssDelegationManager public tssDelegationManager;
    TssGroupManager public tssGroupManager;
    TssGroupManager public tssGroupManagerImplementation;
    TssStakingSlashing public tssStakingSlashing;


    uint256 _minStakeAmount = 10e10;
    address initialOwner = address(this);
    address deployer = address(0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266);
    // 26f45686079c1e633e14e235c58b465192f9e33819177bd19e7bb225afae031e
    address addrowner = address(0xD5AdD52D36399570e56C183d949dA83ac29aA7d6);
    
    function setUp() public {
        
        proxyAdmin = new ProxyAdmin();
        l1Bit = new BitTokenERC20("BitToken", "BIT");
        Lib_AddressManager libAddressManager = new Lib_AddressManager();
    
        L1CrossDomainMessenger l1CrossDomainMessengerImplementation = new L1CrossDomainMessenger();
        proxy = new TransparentUpgradeableProxy(
            address(l1CrossDomainMessengerImplementation), 
            address(proxyAdmin), 
            abi.encodeWithSelector(L1CrossDomainMessenger.initialize.selector, address(libAddressManager)));
        l1CrossDomainMessenger = L1CrossDomainMessenger(address(proxy));

        TssGroupManager tssGroupManagerImplementation = new TssGroupManager();
        proxy = new TransparentUpgradeableProxy(
            address(tssGroupManagerImplementation), 
            address(proxyAdmin), 
            abi.encodeWithSelector(TssGroupManager.initialize.selector));
        tssGroupManager = TssGroupManager(address(proxy));

        TssDelegationManager tssDelegationManagerImplementation = new TssDelegationManager(tssDelegation, tssDelegationSlasher);
        proxy_tssDelegationManager = new TransparentUpgradeableProxy(
            address(tssDelegationManagerImplementation), 
            address(proxyAdmin), 
            abi.encodeWithSelector(TssDelegationManager.initializeT.selector, address(tssStakingSlashing), address(tssGroupManager), _minStakeAmount, address(this)));
        tssDelegationManager = TssDelegationManager(address(proxy_tssDelegationManager));
        
        TssDelegation tssDelegationImplementation = new TssDelegation(tssDelegationManager);
        proxy = new TransparentUpgradeableProxy(
            address(tssDelegationImplementation), 
            address(proxyAdmin), 
            abi.encodeWithSelector(TssDelegation.initializeT.selector, address(tssStakingSlashing), address(this)));
        tssDelegation = TssDelegation(address(proxy));

        TssDelegationSlasher tssDelegationSlasherImplementation = new TssDelegationSlasher(tssDelegationManager, tssDelegation);
        proxy = new TransparentUpgradeableProxy(
            address(tssDelegationSlasherImplementation), 
            address(proxyAdmin), 
            "");
        tssDelegationSlasher = TssDelegationSlasher(address(proxy));
        vm.prank(address(0));
        tssDelegationSlasher.transferOwnership(address(this));

        TssDelegationManager tssDelegationManagerImpl = new TssDelegationManager(tssDelegation, tssDelegationSlasher);
        proxyAdmin.upgrade(proxy_tssDelegationManager, address(tssDelegationManagerImpl));
        // proxy_tssDelegationManager.upgradeToAndCall(address(tssDelegationManagerImpl),"");

        TssStakingSlashing tssStakingSlashingImplementation = new TssStakingSlashing();
        proxy = new TransparentUpgradeableProxy(
            address(tssStakingSlashingImplementation), 
            address(proxyAdmin), 
            abi.encodeWithSelector(TssStakingSlashing.initialize.selector, address(l1Bit), address(tssGroupManager), address(tssDelegationManager), address(tssDelegation), address(l1CrossDomainMessenger), address(1)));
        tssStakingSlashing = TssStakingSlashing(address(proxy));

        tssDelegation.setStakingSlash(address(tssStakingSlashing));
        tssDelegationManager.setStakingSlash(address(tssStakingSlashing));
    }

    function testInfo() public {

        assertEq(l1CrossDomainMessenger.owner(), address(this));
        assertEq(tssDelegation.owner(), address(this));
        assertEq(tssDelegationSlasher.owner(), address(this));
        assertEq(tssDelegationManager.owner(), address(this));
        assertEq(tssGroupManager.owner(), address(this));
        assertEq(tssStakingSlashing.owner(), address(this));
    }

    address[] public addresses;
    address[] public nodeAddresses;
    bytes[] public nodePubs;
    address[] public rmAddresses;
    uint256[2] public _slashAmount;

    function testWhiteListOnlyOwner() public {

        vm.expectRevert("Ownable: caller is not the owner");
        vm.prank(address(0));
        tssDelegation.addToWhitelist(addresses);
    }

    function testRemoveFromWhitelistOnleOwner() public {

        vm.expectRevert("Ownable: caller is not the owner");
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

    function testSetStakingSlash() public {

        tssDelegation.setStakingSlash(address(1));
        assertEq(tssDelegation.stakingSlash(), address(1));
    }

    function testRegisterAsOperator() public {

        tssDelegation.setStakingSlash(address(tssStakingSlashing));
        assertEq(tssDelegation.stakingSlash(), address(tssStakingSlashing));
        // tssDelegation.registerAsOperator(address(1),address(2));
        string memory publicKey = "02837dbc8dccda9c6ca0c6745eeeb72a42fdcf3257585b3ed116274e9b61239f05";

        addresses.push(address(this));
        tssDelegation.addToWhitelist(addresses);
        // tssDelegationManager.addToWhitelist(addresses);
        // tssDelegation.registerAsOperator(IDelegationCallback(address(this)), address(this));
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
    // 8e09231cd4e10460b4b4b0b291055ba5d2c900daed32642b47bca8c034c895ef
    // 02837dbc8dccda9c6ca0c6745eeeb72a42fdcf3257585b3ed116274e9b61239f05
    // 0xD5750ebE91654aB6e345Fd1c6f97348265E5Ef9f

    // 00f75cacb3ff9b4ed18e093dbbf6f7e7188060fabfb352d2253ea0e26f96a3eb
    // 036817b848557cd71d1105b1f16fad98846da884a917cf0dfbd824a4fac6a7d500
    // 0xD5751caAC4Cc34f9147FD2d856Abef1c54E8b22b

    // de87f566d458c0e53eb866cadd4e8e39a83cb8e20801e57e1fb55a3e4107027b
    // 03a2efe8c1174768c4be989f1e9de145b773e75ff3beadf2e001459930bd8a3ddf
    // 0xD5752dBebc3FbdEe41f2F8DD7286d471517DE7E9

    // 32a79d56c34105915ad71b6e0fc52b4cd6e7212dfa7b4dd0b24602282ee605ef
    // 029a07c5cee8766fa212684e5e7ee3e4dbc8791d22720aae2a2611de42e9a00aa9
    // 0xd5753CF7F8a55bb03E24841aF6B7e98644c28836
    function testRunTssDelegation() public {
        addresses.push(address(tssStakingSlashing));
        address tssNode1 = address(0xD5750ebE91654aB6e345Fd1c6f97348265E5Ef9f);
        address tssNode2 = address(0xD5751caAC4Cc34f9147FD2d856Abef1c54E8b22b);
        address tssNode3 = address(0xD5752dBebc3FbdEe41f2F8DD7286d471517DE7E9);
        address tssNode4 = address(0xd5753CF7F8a55bb03E24841aF6B7e98644c28836);
        bytes memory nodePub1 = hex"837dbc8dccda9c6ca0c6745eeeb72a42fdcf3257585b3ed116274e9b61239f05909ab32eca7d4268f15f4aa1a500009982739bf18968fedfd419ee1c69a60088";
        bytes memory nodePub2 = hex"6817b848557cd71d1105b1f16fad98846da884a917cf0dfbd824a4fac6a7d5008d1f50dcd1bb8eb3bcaf5fdb1a07dfca7129ae4f80f538924137b0341ab2d51d";
        bytes memory nodePub3 = hex"a2efe8c1174768c4be989f1e9de145b773e75ff3beadf2e001459930bd8a3ddf1798ae5898d8ded22e23b044cdfcd280550eb3616e8a0c291d1d6e0997967445";
        bytes memory nodePub4 = hex"9a07c5cee8766fa212684e5e7ee3e4dbc8791d22720aae2a2611de42e9a00aa912f03c22931e7075de964012e54054ad4f969bbef01ede9bf56f70a1a79931b6";

        assertEq(tssDelegationManager.owner(), address(this));
        tssDelegationManager.transferOwnership(deployer);
        assertEq(tssDelegationManager.owner(), deployer);

        vm.prank(deployer);
        // 3
        tssDelegationManager.addToWhitelist(addresses);
        assertEq(tssDelegationManager.whitelist(address(tssStakingSlashing)), true);

        nodeAddresses.push(tssNode1);
        nodeAddresses.push(tssNode2);
        nodeAddresses.push(tssNode3);
        nodeAddresses.push(tssNode4);
        nodePubs.push(nodePub1);
        nodePubs.push(nodePub2);
        nodePubs.push(nodePub3);
        nodePubs.push(nodePub4);

        assertEq(tssDelegation.owner(), address(this));
        tssDelegation.transferOwnership(deployer);
        assertEq(tssDelegation.owner(), deployer);
        vm.prank(deployer);
        // 5
        tssDelegation.addToWhitelist(nodeAddresses);
        assertEq(tssDelegation.whitelist(tssNode1), true);
        assertEq(tssDelegation.whitelist(tssNode2), true);
        assertEq(tssDelegation.whitelist(tssNode3), true);
        assertEq(tssDelegation.whitelist(tssNode4), true);

        // 6 & 7 & 9 & 10
        for (uint i = 0; i < nodeAddresses.length; i++) {
            address addr = nodeAddresses[i];
            bytes memory pub = nodePubs[i];
            // 6
            assertEq(l1Bit.balanceOf(addr),0);
            vm.prank(addr);
            l1Bit.mint(10e22);
            assertEq(l1Bit.balanceOf(addr),10e22);

            // 7
            vm.prank(addr);
            l1Bit.approve(address(tssDelegationManager), 10e21);
            assertEq(l1Bit.allowance(addr, address(tssDelegationManager)), 10e21);

            // 9
            vm.prank(addr);
            tssStakingSlashing.deposit(10e20);
            assertEq(tssStakingSlashing.shares(addr),10e20);

            // 10
            vm.prank(addr);
            tssStakingSlashing.registerAsOperator(pub);
            assertEq(tssDelegation.isDelegated(addr),true);
            assertEq(tssDelegation.isOperator(addr),true);
        }

        // assertEq(tssStakingSlashing.owner(), address(this));
        // assertEq(tssStakingSlashing.owner(), address(proxyAdmin));
        // vm.prank(address(proxyAdmin));
        // tssStakingSlashing.transferOwnership(deployer);
        // assertEq(tssStakingSlashing.owner(), deployer);

        // 12
        _slashAmount = [10e18, 10e19];

        tssStakingSlashing.setSlashingParams(_slashAmount);
        assertEq(tssStakingSlashing.slashAmount(0),_slashAmount[0]);
        assertEq(tssStakingSlashing.slashAmount(1),_slashAmount[1]);

        tssGroupManager.setStakingSlash(address(tssStakingSlashing));
        tssGroupManager.setTssGroupMember(2, nodePubs);
        bytes memory groupPublicKey = hex"a2efe8c1174768c4be989f1e9de145b773e75ff3beadf2e001459930bd8a3ddf1798ae5898d8ded22e23b044cdfcd280550eb3616e8a0c291d1d6e0997967445";
        for (uint i = 0; i < nodeAddresses.length; i++) {

            vm.prank(nodeAddresses[i]);
            tssGroupManager.setGroupPublicKey(nodePubs[i], groupPublicKey);
            console.log(i);
            assertEq(tssGroupManager.isInActiveMember(nodePubs[i]), true);
            (bytes memory publicKey, address nodeAddress, ITssGroupManager.MemberStatus status) = tssGroupManager.tssActiveMemberInfo(nodePubs[i]);
            assertEq(publicKey, nodePubs[i]);
            // assertEq(nodeAddress, nodeAddresses[i]);
            // assertEq(status, ITssGroupManager.MemberStatus.jail);
        }

        (uint256 gRoundId, uint256 threshold, bytes memory confirmGroupPublicKey, bytes[] memory activeTssMembers) = tssGroupManager.getTssGroupInfo();

        console.log("gRoundId: ", gRoundId);
        console.log("threshold: ", threshold);
        console.logBytes(confirmGroupPublicKey);

        for (uint i = 0; i < activeTssMembers.length; i++) {
            console.logBytes(activeTssMembers[i]);
        }
    }
}
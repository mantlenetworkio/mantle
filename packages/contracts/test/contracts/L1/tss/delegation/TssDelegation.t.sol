// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "../../../../../contracts/libraries/resolver/Lib_AddressManager.sol";
import "../../../../../contracts/L1/messaging/L1CrossDomainMessenger.sol";

import "../../../../../contracts/L1/local/TestBitToken.sol";

import "../../../../../contracts/L1/tss/delegation/TssDelegation.sol";
import "../../../../../contracts/L1/tss/delegation/TssDelegationManager.sol";
import "../../../../../contracts/L1/tss/delegation/TssDelegationSlasher.sol";
import "../../../../../contracts/L1/tss/TssGroupManager.sol";
import "../../../../../contracts/L1/tss/TssStakingSlashing.sol";
import "../../../../../contracts/L1/delegation/DelegationSlasher.sol";
import "../../../../../contracts/L1/delegation/interfaces/IDelegationSlasher.sol";
import "../../../../../contracts/L1/tss/ITssGroupManager.sol";
import "../../../../../contracts/L1/delegation/interfaces/IDelegation.sol";
import "../../../mocks/EmptyContract.sol";
import "../../../../../contracts/L1/delegation/interfaces/IDelegationShare.sol";


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
    TssStakingSlashing public tssStakingSlashing;


    uint256 _minStakeAmount = 10e10;
    address initialOwner = address(this);
    // 041deb3563e965bce6e803b88b9d25005cb1414c4cdade04181363e87ca9e259
    address deployer = address(0xd5DeB9917eFbc36259164EcB89Ecf331eb663426);
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

        assertEq(tssDelegationManager.minStakeAmount(), _minStakeAmount);
    }

    address[] public addresses;
    address[] public permissionedContracts;
    address[] public nodeAddresses;
    address[] public stakerAddresses;
    bytes[] public nodePubs;
    address[] public rmAddresses;
    uint256[2] public _slashAmount;
    bytes groupPublicKey;
    address tn3A;
    bytes tn3Pu;
    bytes tn3Pr;

    function testTDWhiteListOnlyOwner() public {

        vm.expectRevert("Ownable: caller is not the owner");
        vm.prank(address(0));
        tssDelegation.addToWhitelist(addresses);
    }

    function testTDMWhiteListOnlyOwner() public {

        vm.expectRevert("Ownable: caller is not the owner");
        vm.prank(address(0));
        tssDelegationManager.addToWhitelist(addresses);
    }

    function testTDRemoveFromWhitelistOnleOwner() public {

        vm.expectRevert("Ownable: caller is not the owner");
        vm.prank(address(0));
        tssDelegation.removeFromWhitelist(addresses);
    }

    function testTDMRemoveFromWhitelistOnleOwner() public {

        vm.expectRevert("Ownable: caller is not the owner");
        vm.prank(address(0));
        tssDelegationManager.removeFromWhitelist(addresses);
    }

    function testTDWhiteList(address a1, address a2, address a3, address a4) public {
        addresses.push(address(a1));
        addresses.push(address(a2));
        addresses.push(address(a2));
        addresses.push(address(a4));

        rmAddresses.push(address(a1));
        rmAddresses.push(address(a2));

        assertEq(tssDelegation.whitelist(address(a1)), false);

        tssDelegation.addToWhitelist(addresses);
        assertEq(tssDelegation.whitelist(address(a1)), true);
        assertEq(tssDelegation.whitelist(address(a2)), true);
        assertEq(tssDelegation.whitelist(address(a3)), false);
        assertEq(tssDelegation.whitelist(address(a4)), true);

        tssDelegation.removeFromWhitelist(rmAddresses);
        assertEq(tssDelegation.whitelist(address(a1)), false);
        assertEq(tssDelegation.whitelist(address(a2)), false);
        assertEq(tssDelegation.whitelist(address(a3)), false);
        assertEq(tssDelegation.whitelist(address(a4)), true);
    }

    function testTDMWhiteList(address a1, address a2, address a3, address a4) public {
        addresses.push(address(a1));
        addresses.push(address(a2));
        addresses.push(address(a2));
        addresses.push(address(a4));

        rmAddresses.push(address(a1));
        rmAddresses.push(address(a2));

        assertEq(tssDelegationManager.whitelist(address(a1)), false);

        tssDelegationManager.addToWhitelist(addresses);
        assertEq(tssDelegationManager.whitelist(address(a1)), true);
        assertEq(tssDelegationManager.whitelist(address(a2)), true);
        assertEq(tssDelegationManager.whitelist(address(a3)), false);
        assertEq(tssDelegationManager.whitelist(address(a4)), true);

        tssDelegationManager.removeFromWhitelist(rmAddresses);
        assertEq(tssDelegationManager.whitelist(address(a1)), false);
        assertEq(tssDelegationManager.whitelist(address(a2)), false);
        assertEq(tssDelegationManager.whitelist(address(a3)), false);
        assertEq(tssDelegationManager.whitelist(address(a4)), true);
    }

    function testSetStakingSlashOnlyOwner(address ss) public {

        vm.expectRevert("Ownable: caller is not the owner");
        vm.prank(address(0));
        tssDelegation.setStakingSlash(ss);
    }

    function testSetStakingSlash(address ss, address ss1) public {

        tssDelegation.setStakingSlash(ss);
        assertEq(tssDelegation.stakingSlash(), ss);

        tssDelegation.setStakingSlash(ss1);
        assertEq(tssDelegation.stakingSlash(), ss1);
    }

    function testRegisterAsOperatorWhiteListOnly(address wl) public {

        addresses.push(wl);
        tssDelegation.addToWhitelist(addresses);
        string memory publicKey = "02837dbc8dccda9c6ca0c6745eeeb72a42fdcf3257585b3ed116274e9b61239f05";

        vm.expectRevert("NOT_IN_WHITELIST");
        tssStakingSlashing.registerAsOperator(bytes(publicKey));

        vm.prank(wl);
        tssStakingSlashing.registerAsOperator(bytes(publicKey));
    }

    function testRegisterAsOperatorOnlyStakingSlash(address wl, address ss) public {

        addresses.push(wl);
        tssDelegation.addToWhitelist(addresses);
        string memory publicKey = "02837dbc8dccda9c6ca0c6745eeeb72a42fdcf3257585b3ed116274e9b61239f05";

        vm.expectRevert("NOT_IN_WHITELIST");
        tssStakingSlashing.registerAsOperator(bytes(publicKey));

        tssDelegation.setStakingSlash(ss);
        assertEq(tssDelegation.stakingSlash(), ss);

        vm.expectRevert("contract call is not staking slashing");
        vm.prank(wl);
        tssStakingSlashing.registerAsOperator(bytes(publicKey));

        tssDelegation.setStakingSlash(address(tssStakingSlashing));
        assertEq(tssDelegation.stakingSlash(), address(tssStakingSlashing));

        vm.prank(wl);
        tssStakingSlashing.registerAsOperator(bytes(publicKey));
    }

    // 这种请求可能验证不了
    // TypeError: Member "registerAsOperator" not found or not visible after argument-dependent lookup in contract TssDelegation
    // function testZeroCanNotRegisterAsOperator() public {

    //     addresses.push(address(0));
    //     tssDelegation.addToWhitelist(addresses);
    //     string memory publicKey = "02837dbc8dccda9c6ca0c6745eeeb72a42fdcf3257585b3ed116274e9b61239f05";

    //     vm.prank(address(0));
    //     tssDelegation.registerAsOperator(bytes(publicKey));
    // }

    function testRegisterAsOperatorIsNotDelegated(address wl) public {

        addresses.push(address(wl));
        tssDelegation.addToWhitelist(addresses);
        string memory publicKey = "02837dbc8dccda9c6ca0c6745eeeb72a42fdcf3257585b3ed116274e9b61239f05";

        vm.prank(address(wl));
        tssStakingSlashing.registerAsOperator(bytes(publicKey));

        vm.expectRevert("Delegation.registerAsOperator: Delegate has already registered");
        vm.prank(address(wl));
        tssStakingSlashing.registerAsOperator(bytes(publicKey));
    }

    function testRegisterAsOperatorIsNotFrozen(address wl) public {

        addresses.push(address(wl));
        tssDelegation.addToWhitelist(addresses);
        string memory publicKey = "02837dbc8dccda9c6ca0c6745eeeb72a42fdcf3257585b3ed116274e9b61239f05";

        // 给自身合约添加权限？
        assertEq(tssDelegationSlasher.globallyPermissionedContracts(address(tssDelegationSlasher)), false);
        permissionedContracts.push(address(tssDelegationSlasher));
        tssDelegationSlasher.addGloballyPermissionedContracts(permissionedContracts);
        assertEq(tssDelegationSlasher.globallyPermissionedContracts(address(tssDelegationSlasher)), true);

        vm.prank(address(tssDelegationSlasher));
        tssDelegationSlasher.freezeOperator(address(wl));

        assertEq(tssDelegationSlasher.globallyPermissionedContracts(address(tssDelegationSlasher)), true);
        tssDelegationSlasher.removeGloballyPermissionedContracts(permissionedContracts);
        assertEq(tssDelegationSlasher.globallyPermissionedContracts(address(tssDelegationSlasher)), false);

        vm.expectRevert("Slasher.freezeOperator: msg.sender does not have permission to slash this operator");
        vm.prank(address(tssDelegationSlasher));
        tssDelegationSlasher.freezeOperator(address(wl));

        // 还有一种判断 else if (block.timestamp < bondedUntil[toBeSlashed][slashingContract])
        vm.prank(address(wl));
        tssDelegationSlasher.allowToSlash(address(tssStakingSlashing));

        vm.prank(address(tssStakingSlashing));
        tssDelegationSlasher.freezeOperator(address(wl));

        vm.expectRevert("Delegation._delegate: cannot delegate to a frozen operator");
        vm.prank(address(wl));
        tssStakingSlashing.registerAsOperator(bytes(publicKey));

        // vm.prank(address(tssDelegationSlasher));
        // 这里居然有owner的判断
        tssDelegationSlasher.resetFrozenStatus(addresses);

        vm.prank(address(wl));
        tssStakingSlashing.registerAsOperator(bytes(publicKey));
    }

    function testDepositWhitelistOnly() public {

        address staker0 = address(0x952d1Dac7736a46130d605D5c613185Db01Bb0d3);

        assertEq(l1Bit.balanceOf(staker0),0);
        vm.prank(staker0);
        l1Bit.mint(10e22);
        assertEq(l1Bit.balanceOf(staker0),10e22);

        vm.expectRevert("NOT_IN_WHITELIST");
        vm.prank(staker0);
        tssStakingSlashing.deposit(1);

        addresses.push(address(tssStakingSlashing));
        tssDelegationManager.addToWhitelist(addresses);
        assertEq(tssDelegationManager.whitelist(address(tssStakingSlashing)), true);

        vm.prank(staker0);
        l1Bit.approve(address(tssDelegationManager), 1);

        vm.prank(staker0);
        tssStakingSlashing.deposit(1);
    }

    function testSetSlashingParams(address ow, uint256 a, uint256 b) public {
        vm.assume(ow != address(this));
        vm.assume(a < b && a > 0);

        vm.expectRevert("invalid param slashAmount, animus <= uptime");
        tssStakingSlashing.setSlashingParams([a, a]);

        vm.expectRevert("invalid amount");
        tssStakingSlashing.setSlashingParams([0, b]);

        vm.expectRevert("invalid param slashAmount, animus <= uptime");
        tssStakingSlashing.setSlashingParams([b, a]);

        vm.expectRevert("Ownable: caller is not the owner");
        vm.prank(ow);
        tssStakingSlashing.setSlashingParams([a, b]);

        tssStakingSlashing.setSlashingParams([a, b]);
        assertEq(tssStakingSlashing.slashAmount(0) < tssStakingSlashing.slashAmount(1), true);

        uint256[2] memory sa = tssStakingSlashing.getSlashingParams();
        assertEq(sa[0], a);
        assertEq(sa[1], b);
    }

    function testSlashing() public{
        TssStakingSlashing.SlashMsg memory message = TssStakingSlashing.SlashMsg({
            batchIndex: 123,
            jailNode: 0x1234567890123456789012345678901234567890,
            tssNodes: nodeAddresses,
            slashType: TssStakingSlashing.SlashType.animus
        });

        bytes memory messageBytes = abi.encode(message);

        console.logBytes(messageBytes);

        // 对消息进行哈希运算
        bytes32 messageHash = keccak256(messageBytes);

        // 使用私钥对消息进行签名
        (uint8 v, bytes32 r, bytes32 s) = ecdsaSign(messageHash);

        // // 将 v、r、s 组件拼接为 _sig
        // bytes memory _sig = abi.encodePacked(v, r, s);
        // console.logBytes(_sig);
        
        // assertEq(1, 1);
    }

    function testSetTssGroupMember(address ow, bytes calldata tn0p) public {
        vm.assume(tn0p.length == 64);
        vm.assume(ow != address(this));

        vm.expectRevert("batch public key is empty");
        tssGroupManager.setTssGroupMember(uint256(2), nodePubs);

        nodePubs.push(tn0p);
        nodePubs.push(tn0p);
        nodePubs.push(tn0p);

        vm.expectRevert("Ownable: caller is not the owner");
        vm.prank(ow);
        tssGroupManager.setTssGroupMember(uint256(2), nodePubs);

        vm.expectRevert("threshold must less than tss member");
        tssGroupManager.setTssGroupMember(uint256(4), nodePubs);

        tssGroupManager.setStakingSlash(address(tssStakingSlashing));
        vm.expectRevert("batch public keys has a node ,can not be operator");
        tssGroupManager.setTssGroupMember(uint256(2), nodePubs);
    }

    function prepareAccount() internal {
        // 8e09231cd4e10460b4b4b0b291055ba5d2c900daed32642b47bca8c034c895ef
        address tn0A = address(0xD5750ebE91654aB6e345Fd1c6f97348265E5Ef9f);
        bytes memory tn0Pu = hex"837dbc8dccda9c6ca0c6745eeeb72a42fdcf3257585b3ed116274e9b61239f05909ab32eca7d4268f15f4aa1a500009982739bf18968fedfd419ee1c69a60088";
        // 00f75cacb3ff9b4ed18e093dbbf6f7e7188060fabfb352d2253ea0e26f96a3eb
        address tn1A = address(0xD5751caAC4Cc34f9147FD2d856Abef1c54E8b22b);
        bytes memory tn1Pu = hex"6817b848557cd71d1105b1f16fad98846da884a917cf0dfbd824a4fac6a7d5008d1f50dcd1bb8eb3bcaf5fdb1a07dfca7129ae4f80f538924137b0341ab2d51d";
        // de87f566d458c0e53eb866cadd4e8e39a83cb8e20801e57e1fb55a3e4107027b
        address tn2A = address(0xD5752dBebc3FbdEe41f2F8DD7286d471517DE7E9);
        bytes memory tn2Pu = hex"a2efe8c1174768c4be989f1e9de145b773e75ff3beadf2e001459930bd8a3ddf1798ae5898d8ded22e23b044cdfcd280550eb3616e8a0c291d1d6e0997967445";

        nodeAddresses.push(tn0A);
        nodeAddresses.push(tn1A);
        nodeAddresses.push(tn2A);
        nodePubs.push(tn0Pu);
        nodePubs.push(tn1Pu);
        nodePubs.push(tn2Pu);

        // 87f7e2782dc38c16a715c812fcd937d3ea4714d1bea61fce1507526c068e5e38
        address staker0 = address(0x952d1Dac7736a46130d605D5c613185Db01Bb0d3);
        // 325234ee8de1367bc35d3411423cbac3dda20c945570f654147f58355d30f4f8
        address staker1 = address(0x952D93593cDc608cAB1D3E52dFE9Fb153594a2Ad);
        // 4e9ca2a52bb513372c7c75e7596403b4c9fcf4ffe0fdc2f3ec18cae8e5951651
        address staker2 = address(0x952d4433AD3cF75f675da96C908D533780467a98);
        stakerAddresses.push(staker0);
        stakerAddresses.push(staker1);
        stakerAddresses.push(staker2);


        groupPublicKey = hex"9a07c5cee8766fa212684e5e7ee3e4dbc8791d22720aae2a2611de42e9a00aa912f03c22931e7075de964012e54054ad4f969bbef01ede9bf56f70a1a79931b6";
        tn3Pr = hex"32a79d56c34105915ad71b6e0fc52b4cd6e7212dfa7b4dd0b24602282ee605ef";
        tn3A = address(0xd5753CF7F8a55bb03E24841aF6B7e98644c28836);
        tn3Pu = hex"9a07c5cee8766fa212684e5e7ee3e4dbc8791d22720aae2a2611de42e9a00aa912f03c22931e7075de964012e54054ad4f969bbef01ede9bf56f70a1a79931b6";
    }

    function testRegisterAsOperatorSimple() public {

        prepareAccount();
        tssDelegation.addToWhitelist(nodeAddresses);

        addresses.push(address(tssStakingSlashing));
        tssDelegationManager.addToWhitelist(addresses);
        assertEq(tssDelegationManager.whitelist(address(tssStakingSlashing)), true);

        for (uint i = 0; i < nodeAddresses.length; i++) {
            address addr = nodeAddresses[i];
            bytes memory pub = nodePubs[i];
            address staker = stakerAddresses[i];

            vm.prank(addr);
            tssStakingSlashing.registerAsOperator(pub);
            assertEq(tssDelegation.isOperator(addr),true);
            assertEq(tssDelegation.isDelegated(addr),true);

            assertEq(l1Bit.balanceOf(staker),0);
            vm.prank(staker);
            l1Bit.mint(10e22);
            assertEq(l1Bit.balanceOf(staker),10e22);

            vm.prank(staker);
            l1Bit.approve(address(tssDelegationManager), _minStakeAmount);

            vm.prank(staker);
            tssStakingSlashing.deposit(_minStakeAmount);

            vm.prank(staker);
            tssStakingSlashing.delegateTo(addr);
        }

        // 不明白这个配置项的作用
        // tssStakingSlashing.setSlashingParams([uint256(1), uint256(2)]);
        tssGroupManager.setStakingSlash(address(tssStakingSlashing));
        assertEq(tssDelegationManager.minStakeAmount(), _minStakeAmount);

        vm.expectRevert("batch public keys has a node ,can not be operator");
        tssGroupManager.setTssGroupMember(uint256(2), nodePubs);

        tssDelegationManager.setMinStakeAmount(uint256(_minStakeAmount - 1));
        assertEq(tssDelegationManager.minStakeAmount(), _minStakeAmount -1);
        tssGroupManager.setTssGroupMember(uint256(2), nodePubs);

        vm.expectRevert("your public key is not in InActiveMember");
        vm.prank(tn3A);
        tssGroupManager.setGroupPublicKey(tn3Pu, groupPublicKey);

        vm.expectRevert("public key not match");
        tssGroupManager.setGroupPublicKey(nodePubs[1], groupPublicKey);

        (uint256 gRoundId1, uint256 threshold1, bytes[] memory inActiveTssMembers1) = tssGroupManager.getTssInactiveGroupInfo();
        assertEq(inActiveTssMembers1.length, nodeAddresses.length);

        for (uint i = 0; i < nodeAddresses.length; i++) {

            vm.prank(nodeAddresses[i]);
            tssGroupManager.setGroupPublicKey(nodePubs[i], groupPublicKey);
        }

        for (uint i = 0; i < nodeAddresses.length; i++) {

            assertEq(tssGroupManager.isInActiveMember(nodePubs[i]), false);
            (bytes memory publicKey, address nodeAddress, ITssGroupManager.MemberStatus status) = tssGroupManager.tssActiveMemberInfo(nodePubs[i]);
            assertEq(publicKey, nodePubs[i]);
            assertEq(nodeAddress, nodeAddresses[i]);
            assertEq(tssGroupManager.isTssGroupUnJailMembers(nodeAddresses[i]),true);
        }

        (uint256 gRoundId2, uint256 threshold2, bytes[] memory inActiveTssMembers2) = tssGroupManager.getTssInactiveGroupInfo();
        assertEq(inActiveTssMembers2.length, 0);

        (uint256 gRoundId, uint256 threshold, bytes memory confirmGroupPublicKey, bytes[] memory activeTssMembers) = tssGroupManager.getTssGroupInfo();
        assertEq(confirmGroupPublicKey, groupPublicKey);
        assertEq(activeTssMembers.length, nodeAddresses.length);

        for (uint i = 0; i < activeTssMembers.length; i++) {
            assertEq(activeTssMembers[i], nodePubs[i]);
        }
    }

    function rao() internal {
        tssDelegation.addToWhitelist(nodeAddresses);

        addresses.push(address(tssStakingSlashing));
        tssDelegationManager.addToWhitelist(addresses);
        assertEq(tssDelegationManager.whitelist(address(tssStakingSlashing)), true);

        for (uint i = 0; i < nodeAddresses.length; i++) {
            address addr = nodeAddresses[i];
            bytes memory pub = nodePubs[i];
            address staker = stakerAddresses[i];

            vm.prank(staker);
            l1Bit.approve(address(tssDelegationManager), _minStakeAmount);

            vm.prank(staker);
            tssStakingSlashing.deposit(_minStakeAmount);
        
            vm.prank(staker);
            tssStakingSlashing.delegateTo(addr);
        }

        tssGroupManager.setTssGroupMember(uint256(2), nodePubs);

        (uint256 gRoundId1, uint256 threshold1, bytes[] memory inActiveTssMembers1) = tssGroupManager.getTssInactiveGroupInfo();
        assertEq(inActiveTssMembers1.length, nodeAddresses.length);

        for (uint i = 0; i < nodeAddresses.length; i++) {

            vm.prank(nodeAddresses[i]);
            tssGroupManager.setGroupPublicKey(nodePubs[i], groupPublicKey);
        }

        for (uint i = 0; i < nodeAddresses.length; i++) {

            assertEq(tssGroupManager.isInActiveMember(nodePubs[i]), false);
            (bytes memory publicKey, address nodeAddress, ITssGroupManager.MemberStatus status) = tssGroupManager.tssActiveMemberInfo(nodePubs[i]);
            assertEq(publicKey, nodePubs[i]);
            assertEq(nodeAddress, nodeAddresses[i]);
            assertEq(tssGroupManager.isTssGroupUnJailMembers(nodeAddresses[i]),true);
        }

        (uint256 gRoundId2, uint256 threshold2, bytes[] memory inActiveTssMembers2) = tssGroupManager.getTssInactiveGroupInfo();
        assertEq(inActiveTssMembers2.length, 0);

        (uint256 gRoundId, uint256 threshold, bytes memory confirmGroupPublicKey, bytes[] memory activeTssMembers) = tssGroupManager.getTssGroupInfo();
        assertEq(confirmGroupPublicKey, groupPublicKey);
        assertEq(activeTssMembers.length, nodeAddresses.length);

        for (uint i = 0; i < activeTssMembers.length; i++) {
            assertEq(activeTssMembers[i], nodePubs[i]);
        }
    }

    function testQuitRequest() public {
        vm.expectRevert("do not have deposit");
        tssStakingSlashing.quitRequest();

        testRegisterAsOperatorSimple();

        tssGroupManager.removeMember(nodePubs[1]);
        assertEq(tssGroupManager.memberExistActive(nodeAddresses[1]), false);
        assertEq(tssGroupManager.memberExistInActive(nodePubs[1]), false);

        vm.expectRevert("not at the inactive group or active group");
        vm.prank(nodeAddresses[1]);
        tssStakingSlashing.quitRequest();

        vm.prank(nodeAddresses[0]);
        tssStakingSlashing.quitRequest();
        
        vm.expectRevert("already in quitRequestList");
        vm.prank(nodeAddresses[0]);
        tssStakingSlashing.quitRequest();

        address[] memory qrs = tssStakingSlashing.getQuitRequestList();
        assertEq(qrs[0], nodeAddresses[0]);
        assertEq(qrs.length, 1);

        vm.expectRevert("Ownable: caller is not the owner");
        vm.prank(address(0));
        tssStakingSlashing.clearQuitRequestList();

        tssStakingSlashing.clearQuitRequestList();

        vm.prank(nodeAddresses[0]);
        tssStakingSlashing.quitRequest();

        vm.prank(nodeAddresses[2]);
        tssStakingSlashing.quitRequest();

        qrs = tssStakingSlashing.getQuitRequestList();
        assertEq(qrs[0], nodeAddresses[0]);
        assertEq(qrs[1], nodeAddresses[2]);
        assertEq(qrs.length, 2);

        tssStakingSlashing.clearQuitRequestList();

        vm.prank(nodeAddresses[0]);
        tssStakingSlashing.quitRequest();
    }

    function testRemoveMember() public {
        testRegisterAsOperatorSimple();

        vm.expectRevert("Ownable: caller is not the owner");
        vm.prank(address(0));
        tssGroupManager.removeMember(nodePubs[1]);

        removeMember();
    }

    function removeMember() internal {

        for (uint i = 0; i < nodeAddresses.length; i++) {
            tssGroupManager.removeMember(nodePubs[i]);
            assertEq(tssGroupManager.memberExistActive(nodeAddresses[i]), false);
            assertEq(tssGroupManager.memberExistInActive(nodePubs[i]), false);
        }
    }

    function testWdUndelegate() public {
        testRegisterAsOperatorSimple();

        for (uint i = 0; i < nodeAddresses.length; i++) {
            vm.expectRevert("DelegationManager._removeShares: shareAmount should not be zero!");
            vm.prank(nodeAddresses[i]);
            tssStakingSlashing.withdraw();

            vm.prank(nodeAddresses[i]);
            tssDelegationManager.undelegate();

            vm.expectRevert("not delegator");
            vm.prank(nodeAddresses[i]);
            tssStakingSlashing.withdraw();
        }
    }

    function testWithdraw() public {
        uint32 bt = uint32(1);
        testRemoveMember();

        uint256 ssb = l1Bit.balanceOf(address(tssStakingSlashing));
        uint256 s0b = l1Bit.balanceOf(stakerAddresses[0]);
        uint256 s1b = l1Bit.balanceOf(stakerAddresses[1]);
        uint256 s2b = l1Bit.balanceOf(stakerAddresses[2]);

        for (uint i = 0; i < nodeAddresses.length; i++) {
            vm.prank(stakerAddresses[i]);
            tssStakingSlashing.withdraw();

            vm.expectRevert("msg sender must request withdraw first");
            tssStakingSlashing.startWithdraw();

            vm.expectRevert("InvestmentManager.startQueuedWithdrawalWaitingPeriod: Stake may still be subject to slashing based on new tasks. Wait to set stakeInactiveAfter.");
            vm.prank(stakerAddresses[i]);
            tssStakingSlashing.startWithdraw();

            bytes32 wdRoot = tssStakingSlashing.withdrawalRoots(stakerAddresses[i]);

            (uint32 initTimestamp, uint32 unlockTimestamp, address withdrawer) = tssDelegationManager.queuedWithdrawals(wdRoot);
            assertEq(unlockTimestamp, 4294967295);
            assertEq(withdrawer, stakerAddresses[i]);

            vm.expectRevert("InvestmentManager.startQueuedWithdrawalWaitingPeriod: Sender is not the withdrawer");
            vm.prank(address(tssStakingSlashing));
            tssDelegationManager.startQueuedWithdrawalWaitingPeriod(wdRoot, address(1), 0);

            bt += 31;
            vm.warp(bt);
            vm.prank(stakerAddresses[i]);
            tssStakingSlashing.startWithdraw();

            vm.expectRevert("msg sender did not request withdraws");
            tssStakingSlashing.canCompleteQueuedWithdrawal();

            vm.expectRevert("msg sender did not request withdraws");
            tssStakingSlashing.completeWithdraw();

            vm.expectRevert("The waiting period has not yet passed");
            vm.prank(stakerAddresses[i]);
            tssStakingSlashing.completeWithdraw();

            bt += 10;
            vm.warp(bt);
            (initTimestamp, unlockTimestamp, withdrawer) = tssDelegationManager.queuedWithdrawals(wdRoot);
            assertEq(unlockTimestamp, bt);
            assertEq(withdrawer, stakerAddresses[i]);

            vm.prank(stakerAddresses[i]);
            bool cc = tssStakingSlashing.canCompleteQueuedWithdrawal();
            assertEq(cc, true);

            vm.prank(stakerAddresses[i]);
            tssStakingSlashing.completeWithdraw();
        }

        ssb = l1Bit.balanceOf(address(tssStakingSlashing));
        s0b = l1Bit.balanceOf(stakerAddresses[0]);
        s1b = l1Bit.balanceOf(stakerAddresses[1]);
        s2b = l1Bit.balanceOf(stakerAddresses[2]);
        assertEq(ssb, 0);
        assertEq(s0b, 10e22);
        assertEq(s1b, 10e22);
        assertEq(s2b, 10e22);

        rao();
        removeMember();

        vm.prank(stakerAddresses[0]);
        tssStakingSlashing.withdraw();
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
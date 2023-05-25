// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import '../contracts/libraries/resolver/Lib_AddressManager.sol';
import '../contracts/L1/messaging/L1CrossDomainMessenger.sol';
import '../contracts/L1/local/TestBitToken.sol';
import '../contracts/L1/tss/delegation/TssDelegation.sol';
import '../contracts/L1/tss/delegation/TssDelegationManager.sol';
import '../contracts/L1/tss/delegation/TssDelegationSlasher.sol';
import '../contracts/L1/tss/TssGroupManager.sol';
import '../contracts/L1/tss/TssStakingSlashing.sol';
import '../contracts/L1/delegation/DelegationSlasher.sol';
import '../contracts/L1/tss/ITssGroupManager.sol';
import '../contracts/L1/delegation/interfaces/IDelegation.sol';

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import "forge-std/Test.sol";
import "./contracts/mocks/EmptyContract.sol";

contract commonTest is Test {

    ProxyAdmin public proxyAdmin;

    Lib_AddressManager public libAddressManager;
    L1CrossDomainMessenger public l1CrossDomainMessenger;

    BitTokenERC20 public l1Bit;

    TssDelegation public tssDelegation;
    TssDelegationSlasher public tssDelegationSlasher;
    TssDelegationManager public tssDelegationManager;
    TssGroupManager public tssGroupManager;
    TssStakingSlashing public tssStakingSlashing;

    address deployer = address(0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266);
    // 26f45686079c1e633e14e235c58b465192f9e33819177bd19e7bb225afae031e
    address addrowner = address(0xD5AdD52D36399570e56C183d949dA83ac29aA7d6);

    uint256 _minStakeAmount = 10e10;
    address initialOwner = address(this);

    
    function setUp() public {
        
        libAddressManager = new Lib_AddressManager();
        libAddressManager.transferOwnership(deployer);
        assertEq(libAddressManager.owner(), deployer);

        proxyAdmin = new ProxyAdmin();
        l1Bit = new BitTokenERC20("BitToken", "BIT");

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
            abi.encodeWithSelector(TssDelegationManager.initializeT.selector, address(tssStakingSlashing), address(tssGroupManager), _minStakeAmount, initialOwner)
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

        assertEq(l1CrossDomainMessenger.owner(), address(proxyAdmin));
        assertEq(tssDelegation.owner(), address(this));
        assertEq(tssDelegationSlasher.owner(), address(this));
        assertEq(tssDelegationManager.owner(), address(this));
        assertEq(tssGroupManager.owner(), address(proxyAdmin));
        assertEq(tssStakingSlashing.owner(), address(proxyAdmin));
    }

}
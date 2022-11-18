// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/proxy/beacon/IBeacon.sol";
import "@openzeppelin/contracts/proxy/Proxy.sol";
import "@openzeppelin/contracts/interfaces/draft-IERC1822.sol";
import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts/utils/StorageSlot.sol";



 contract TwoImplementationProxyUpgradable is Proxy {

     // This is the keccak-256 hash of "eip1967.proxy.rollback" subtracted by 1
    bytes32 private constant _ROLLBACK_SLOT = 0x4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd9143;

    /**
     * @dev Storage slot with the address of the current first implementation.
     * This is the keccak-256 hash of "eip1967.proxy.implementation" subtracted by 1, and is
     * validated in the constructor.
     */
    bytes32 internal constant _IMPLEMENTATION_ONE_SLOT = 0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc;

    /**
     * @dev Storage slot with the address of the current second implementation.
     * This is the keccak-256 hash of "eip1967.proxy.implementation.two" subtracted by 1, and is
     * validated in the constructor.
     *
     */
    bytes32 internal constant _IMPLEMENTATION_TWO_SLOT = 0xa1832b4681bf2b3269fd2d0163abf215d243dd098d1ece37dd0775a7f2a8c09d;

    /**
     * @dev Emitted when an implementation is upgraded.
     */
    event Upgraded(address indexed implementation, bytes32 indexed slot);


    /**
     * @dev Initializes the upgradeable proxy with an initial implementation specified by `_logic`.
     *
     * If `_data` is nonempty, it's used as data in a delegate call to `_logic`. This will typically be an encoded
     * function call, and allows initializing the storage of the proxy like a Solidity constructor.
     */
    constructor(address _logicOne, _logicTwo, bytes memory _dataOne, bytes memory _dataTwo) payable {
        _upgradeToAndCall(_logicOne, _dataOne, false);
        _upgradeToAndCall(_logicTwo, _dataTwo, false);
    }

    /**
     * @dev Returns the current implementation address in accordance with the slot.
     */
    function _getImplementation(bytes32 slot) internal view returns (address) {
        return StorageSlot.getAddressSlot(slot).value;
    }


    /**
     * @dev Stores a new address in the EIP1967 implementation slot.
     */
    function _setImplementation(address newImplementation, bytes32 slot) private {
        require(Address.isContract(newImplementation), "ERC1967: new implementation is not a contract");
        StorageSlot.getAddressSlot(slot).value = newImplementation;
    }


    /**
     * @dev Perform implementation upgrade for the implementation address stored at slot
     *
     * Emits an {Upgraded} event.
     */
    function _upgradeTo(address newImplementation, bytes32 slot) internal {
        _setImplementation(newImplementation, slot);
        emit Upgraded(newImplementation, slot);
    }

    /**
     * @dev Perform implementation upgrade with additional setup call.
     *
     * Emits an {Upgraded} event.
     */
    function _upgradeToAndCall(
        address newImplementation,
        bytes32 slot,
        bytes memory data,
        bool forceCall
    ) internal {
        _upgradeTo(newImplementation, slot);
        if (data.length > 0 || forceCall) {
            Address.functionDelegateCall(newImplementation, data);
        }
    }

    /**
     * @dev Perform implementation upgrade with security checks for UUPS proxies, and additional setup call.
     *
     * Emits an {Upgraded} event.
     */
    function _upgradeToAndCallUUPS(
        address newImplementation,
        bytes32 slot,
        bytes memory data,
        bool forceCall
    ) internal {
        // Upgrades from old implementations will perform a rollback test. This test requires the new
        // implementation to upgrade back to the old, non-ERC1822 compliant, implementation. Removing
        // this special case will break upgrade paths from old UUPS implementation to new ones.
        if (StorageSlot.getBooleanSlot(_ROLLBACK_SLOT).value) {
            _setImplementation(newImplementation, slot);
        } else {
            try IERC1822Proxiable(newImplementation).proxiableUUID() returns (bytes32 implementationSlot) {
                require(implementationSlot == slot, "ERC1967Upgrade: unsupported proxiableUUID");
            } catch {
                revert("ERC1967Upgrade: new implementation is not UUPS");
            }
            _upgradeToAndCall(newImplementation,slot, data, forceCall);
        }
    }

    /**
     * @dev Storage slot with the admin of the contract.
     * This is the keccak-256 hash of "eip1967.proxy.admin" subtracted by 1, and is
     * validated in the constructor.
     */
    bytes32 internal constant _ADMIN_SLOT = 0xb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103;

    /**
     * @dev Emitted when the admin account has changed.
     */
    event AdminChanged(address previousAdmin, address newAdmin);

    /**
     * @dev Returns the current admin.
     */
    function _getAdmin() internal view returns (address) {
        return StorageSlot.getAddressSlot(_ADMIN_SLOT).value;
    }

    /**
     * @dev Stores a new address in the EIP1967 admin slot.
     */
    function _setAdmin(address newAdmin) private {
        require(newAdmin != address(0), "ERC1967: new admin is the zero address");
        StorageSlot.getAddressSlot(_ADMIN_SLOT).value = newAdmin;
    }

    /**
     * @dev Changes the admin of the proxy.
     *
     * Emits an {AdminChanged} event.
     */
    function _changeAdmin(address newAdmin) internal {
        emit AdminChanged(_getAdmin(), newAdmin);
        _setAdmin(newAdmin);
    }

    /**
     * @dev Delegates the current call to `implementation`.
     *      We load memory from the 32nd byte as the first 32
     *
     * This function does not return to its internal call site, it will return directly to the external caller.
     */
    function _delegate(address implementation) internal virtual override {
        assembly {
            // Copy msg.data. We take full control of memory in this inline assembly
            // block because it will not return to Solidity code. We overwrite the
            // Solidity scratch pad at memory position 32.
            let delegatedCalldataSize := sub(calldatasize(), 32)

            calldatacopy(32, 0, delegatedCalldataSize)

            // Call the implementation.
            // out and outsize are 0 because we don't know the size yet.
            let result := delegatecall(gas(), implementation, 0, delegatedCalldataSize, 0, 0)

            // Copy the returned data.
            returndatacopy(0, 0, returndatasize())

            switch result
            // delegatecall returns 0 on error.
            case 0 {
                revert(0, returndatasize())
            }
            default {
                return(0, returndatasize())
            }
        }
    }

    /**
     * @dev Returns the current implementation address.
     */
    function _implementation() internal view virtual override returns (address impl) {
        bytes32 implementationSlotToCall;
        assembly {
            // copy the first 32 bytes of calldata as it is the implementation slot
            // we are delegating the !
            implementationSlotToCall := calldataload(0)
        }
        return _getImplementation(implementationSlotToCall);
    }

    /**
     * @dev Delegates the current call to the address returned by `_implementation()`.
     *
     * This function does not return to its internal call site, it will return directly to the external caller.
     */
    function _fallback() internal virtual override {

        _beforeFallback();
        _delegate(_implementation());
    }


}


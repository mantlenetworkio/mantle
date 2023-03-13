// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { MantlePortal } from "../L1/MantlePortal.sol";

/**
 * @title PortalSender
 * @notice The PortalSender is a simple intermediate contract that will transfer the balance of the
 *         L1StandardBridge to the MantlePortal during the Bedrock migration.
 */
contract PortalSender {
    /**
     * @notice Address of the MantlePortal contract.
     */
    MantlePortal public immutable PORTAL;

    /**
     * @param _portal Address of the MantlePortal contract.
     */
    constructor(MantlePortal _portal) {
        PORTAL = _portal;
    }

    /**
     * @notice Sends balance of this contract to the MantlePortal.
     */
    function donate() public {
        PORTAL.donateETH{ value: address(this).balance }();
    }
}

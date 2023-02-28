// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

contract MockStateCommitmentChain  {
    function appendStateBatch(bytes32[] memory _batch, uint256 _shouldStartAtElement, bytes memory _signature) public {
        return;
    }

    function FRAUD_PROOF_WINDOW()  external returns(uint256){
        return 0;
    }
}

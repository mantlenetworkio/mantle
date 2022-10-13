// SPDX-License-Identifier: MIT
pragma solidity >0.5.0 <0.9.0;

interface ISequencer {
    struct Sequencer {
        address owner;
        address mintAddress;
        bytes nodeID;
        uint256 amount;
        uint256 keyIndex;
    }

    // mapping(owner => Sequencer) sequencers
    // address[] owners

    // create a new sequencer info
    function createSequencer(
        uint256,
        address,
        bytes calldata
    ) external;

    // check deposit(signer) to equal nodeAddress and nodeID then deposit(signer).amount + amount
    function deposit(uint256) external;

   // amount >= deposit(signer).amount -> withdraw all
    // 0 < amount < deposit(signer).amount -> withdraw amount to signer
    function withdraw(uint256) external;

    // return all sequencer infos
    function getSequencers() external view returns (Sequencer[] memory);

    // return sequencer info by signer address
    function getSequencer(address) external view returns (Sequencer memory);

    // return if signer has deposit
    function isSequencer(address) external view returns (bool);
}

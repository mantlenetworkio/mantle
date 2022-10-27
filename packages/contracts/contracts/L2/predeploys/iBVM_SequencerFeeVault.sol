pragma solidity ^0.8.0;

interface IBVM_SequencerFeeVault {
    /**
     * @dev withdraw all balances.
     */
    function withdraw() public;
    /**
     * @dev Query l1 fee wallet address.
     * @return l1 fee wallet address.
     */
    function l1FeeWallet() external view returns (address);
}

pragma solidity ^0.8.0;

contract GasFee_Burning {
    uint256 public constant MIN_WITHDRAWAL_AMOUNT = 15 ether;
    address public l1FeeWallet;
    constructor(address _l1FeeWallet) {
        l1FeeWallet = _l1FeeWallet;
    }
    function burn(){
        // 收到gasFee暂时存储在这个合约中
        // 触发条件,每当到达一定高度时候 跨链销毁；保留手动触发销毁

    }


}

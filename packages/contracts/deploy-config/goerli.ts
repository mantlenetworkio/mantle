const config = {
  numDeployConfirmations: 1,
  l1BlockTimeSeconds: 15,
  l2BlockGasLimit: 15_000_000,
  l2ChainId: 420,
  ctcL2GasDiscountDivisor: 32,
  ctcEnqueueGasCost: 60_000,
  sccFaultProofWindowSeconds: 10,
  sccSequencerPublishWindowSeconds: 12592000,
  bvmSequencerAddress: '0x7431310e026B69BFC676C0013E12A1A11411EEc9',
  bvmProposerAddress: '0x02b1786A85Ec3f71fBbBa46507780dB7cF9014f6',
  bvmBlockSignerAddress: '0x27770a9694e4B4b1E130Ab91Bc327C36855f612E',
  bvmFeeWalletAddress: '0x0000000000000000000000000000000000000000',
  bvmAddressManagerOwner: '0xf80267194936da1E98dB10bcE06F3147D580a62e',
  bvmGasPriceOracleOwner: '0xa693B8f8207FF043F6bbC2E2120bbE4C2251Efe9',
}

export default config

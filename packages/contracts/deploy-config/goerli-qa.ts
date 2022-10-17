const config = {
  numDeployConfirmations: 1,
  l1BlockTimeSeconds: 15,
  l2BlockGasLimit: 15_000_000,
  l2ChainId: 1705003,
  ctcL2GasDiscountDivisor: 32,
  ctcEnqueueGasCost: 60_000,
  sccFaultProofWindowSeconds: 0,
  sccSequencerPublishWindowSeconds: 12592000,
  bvmSequencerAddress: '0x3f4d7673D7F6db148B7AC1F9158845887A9bea16',
  bvmProposerAddress: '0x1681119781834bf9D4B4b868600c8498f84ebaDc',
  //default bvmBlockSignerAddress
  bvmBlockSignerAddress: '0x00000398232E2064F896018496b4b44b3D62751F',
  bvmFeeWalletAddress: '0xDEc0c7BDf53607cd860c0690b182D9D4C5afE5Bf',
  bvmAddressManagerOwner: '0x3f4d7673D7F6db148B7AC1F9158845887A9bea16',
  bvmGasPriceOracleOwner: '0xDEc0c7BDf53607cd860c0690b182D9D4C5afE5Bf',
  l1BitAddress: '0x320735F74BaA37Bc344a96D4A4EfABBAB06A522C',
}

export default config

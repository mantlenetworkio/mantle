const config = {
  numDeployConfirmations: 1,
  l1BlockTimeSeconds: 15,
  l2BlockGasLimit: 15_000_000,
  l2ChainId: 1705003,
  ctcL2GasDiscountDivisor: 32,
  ctcEnqueueGasCost: 60_000,
  sccFaultProofWindowSeconds: 0,
  sccSequencerPublishWindowSeconds: 12592000,
  bvmSequencerAddress: '0x0A6600AE9D94A0cCcc4F8B86C90f505bA99bE0cd',
  bvmProposerAddress: '0x5933E40c9Ca1CB4b78d26a73aEfBE16C3Ee554e3',
  //default bvmBlockSignerAddress
  bvmBlockSignerAddress: '0x00000398232E2064F896018496b4b44b3D62751F',
  bvmFeeWalletAddress: '0x2cAa68d9E6A8cb1fDCeA2B1CEb03D10c7eCEC7Ef',
  bvmAddressManagerOwner: '0x2cAa68d9E6A8cb1fDCeA2B1CEb03D10c7eCEC7Ef',
  bvmGasPriceOracleOwner: '0x2cAa68d9E6A8cb1fDCeA2B1CEb03D10c7eCEC7Ef',
}

export default config

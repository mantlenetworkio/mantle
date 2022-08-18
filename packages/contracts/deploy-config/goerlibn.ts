const config = {
  numDeployConfirmations: 1,
  l1BlockTimeSeconds: 15,
  l2BlockGasLimit: 15_000_000,
  l2ChainId: 1705003,
  ctcL2GasDiscountDivisor: 32,
  ctcEnqueueGasCost: 60_000,
  sccFaultProofWindowSeconds: 0,
  sccSequencerPublishWindowSeconds: 12592000,
  bvmSequencerAddress: '0xCd0B4E309FB855d644bA64E5fb3dC3DD08f13917',
  bvmProposerAddress: '0x24280Da7E82ab0Ec4E38B394aCF0196c21663B75',
  //default ovmBlockSignerAddress
  bvmBlockSignerAddress: '0x00000398232E2064F896018496b4b44b3D62751F',
  bvmFeeWalletAddress: '0xCd0B4E309FB855d644bA64E5fb3dC3DD08f13917',
  bvmAddressManagerOwner: '0xCd0B4E309FB855d644bA64E5fb3dC3DD08f13917',
  bvmGasPriceOracleOwner: '0xCd0B4E309FB855d644bA64E5fb3dC3DD08f13917',
  bvmTssRewardContractOwner: '0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266',
}

export default config

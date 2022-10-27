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
  bvmProposerAddress: '0x8FA9335FA45CB7fdc6418f0b0e366F678C398Ddb',
  //default bvmBlockSignerAddress
  bvmBlockSignerAddress: '0x00000398232E2064F896018496b4b44b3D62751F',
  bvmFeeWalletAddress: '0x0000000000000000000000000000000000000000',
  bvmAddressManagerOwner: '0x0A6600AE9D94A0cCcc4F8B86C90f505bA99bE0cd',
  bvmGasPriceOracleOwner: '0x0A6600AE9D94A0cCcc4F8B86C90f505bA99bE0cd',
  l1BitAddress: '0x1A4b46696b2bB4794Eb3D4c26f1c55F9170fa4C5',
}

export default config

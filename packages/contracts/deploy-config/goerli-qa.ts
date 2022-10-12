const config = {
  numDeployConfirmations: 1,
  l1BlockTimeSeconds: 15,
  l2BlockGasLimit: 15_000_000,
  l2ChainId: 1705003,
  ctcL2GasDiscountDivisor: 32,
  ctcEnqueueGasCost: 60_000,
  sccFaultProofWindowSeconds: 0,
  sccSequencerPublishWindowSeconds: 12592000,
  bvmSequencerAddress: '0x63EB358137cd06290544e71210c140B345C6FF10',
  bvmProposerAddress: '0x733e38C8C6fEf12D3916eAc679Ef8f5Ffb39127B',
  //default bvmBlockSignerAddress
  bvmBlockSignerAddress: '0x00000398232E2064F896018496b4b44b3D62751F',
  bvmFeeWalletAddress: '0x63EB358137cd06290544e71210c140B345C6FF10',
  bvmAddressManagerOwner: '0x63EB358137cd06290544e71210c140B345C6FF10',
  bvmGasPriceOracleOwner: '0x63EB358137cd06290544e71210c140B345C6FF10',
  l1BitAddress: '0x1A4b46696b2bB4794Eb3D4c26f1c55F9170fa4C5',
}

export default config

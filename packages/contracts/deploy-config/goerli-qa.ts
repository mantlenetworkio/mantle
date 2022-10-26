const config = {
  numDeployConfirmations: 1,
  l1BlockTimeSeconds: 15,
  l2BlockGasLimit: 15_000_000,
  l2ChainId: 1705003,
  ctcL2GasDiscountDivisor: 32,
  ctcEnqueueGasCost: 60_000,
  sccFaultProofWindowSeconds: 0,
  sccSequencerPublishWindowSeconds: 12592000,
  bvmSequencerAddress: process.env.BVM_SEQUENCER_ADDRESS ||  '0x2B1D033ddCc36cd6f4DE10A9a9C2Bc329a443bEB',
  bvmProposerAddress: process.env.BVM_PROPOSER_ADDRESS || '0x4e1614113AF6a1a41cA85d1a5Fe41De105BD65fA',
  //default bvmBlockSignerAddress
  bvmBlockSignerAddress: process.env.BVM_BLOCK_SIGNER_ADDRESS || '0x00000398232E2064F896018496b4b44b3D62751F',
  bvmFeeWalletAddress: process.env.BVM_FEE_WALLET_ADDRESS || '0xeEbceB07eA7D2339895Dd492B3B5960641302830',
  bvmAddressManagerOwner: process.env.BVM_ADDRESS_MANAGER_OWNER || '0x2B1D033ddCc36cd6f4DE10A9a9C2Bc329a443bEB',
  bvmGasPriceOracleOwner: process.env.BVM_GAS_PRICE_ORACLE_OWNER || '0xeEbceB07eA7D2339895Dd492B3B5960641302830',
  l1BitAddress: process.env.L1_BIT_ADDRESS || '0x320735F74BaA37Bc344a96D4A4EfABBAB06A522C',
}

export default config

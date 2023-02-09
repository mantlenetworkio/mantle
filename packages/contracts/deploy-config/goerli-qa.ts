const config = {
  numDeployConfirmations: 1,
  l1BlockTimeSeconds: 15,
  l2BlockGasLimit: 15_000_000,
  l2ChainId: 1705003,
  fixBlockHashBranchingBlock: 0,
  ctcL2GasDiscountDivisor: 32,
  ctcEnqueueGasCost: 60_000,
  sccFaultProofWindowSeconds: 0,
  sccSequencerPublishWindowSeconds: 12592000,
  blockStaleMeasure: 100,
  daFraudProofPeriod: 120,
  l2SubmittedBlockNumber: 1,
  bvmSequencerAddress: process.env.BVM_SEQUENCER_ADDRESS ||  '0x2B1D033ddCc36cd6f4DE10A9a9C2Bc329a443bEB',
  bvmProposerAddress: process.env.BVM_PROPOSER_ADDRESS || '0x4e1614113AF6a1a41cA85d1a5Fe41De105BD65fA',
  //default bvmBlockSignerAddress
  bvmBlockSignerAddress: process.env.BVM_BLOCK_SIGNER_ADDRESS || '0x00000398232E2064F896018496b4b44b3D62751F',
  bvmFeeWalletAddress: process.env.BVM_FEE_WALLET_ADDRESS || '0xeEbceB07eA7D2339895Dd492B3B5960641302830',
  bvmAddressManagerOwner: process.env.BVM_ADDRESS_MANAGER_OWNER || '0x2B1D033ddCc36cd6f4DE10A9a9C2Bc329a443bEB',
  bvmGasPriceOracleOwner: process.env.BVM_GAS_PRICE_ORACLE_OWNER || '0xeEbceB07eA7D2339895Dd492B3B5960641302830',
  bvmFeeWalletOwner: process.env.BVM_FEE_WALLETOWNER_OWNER ||'0xeEbceB07eA7D2339895Dd492B3B5960641302830',
  l1BitAddress: process.env.L1_BIT_ADDRESS || '0xC40C655a91ef15c8eAd818B840CFC08C755D2C0F',
  //on l2, same as l1 address manager.
  bvmTssRewardContractOwner: process.env.TssRewardContractOwner || '0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266',
  dataManagerAddress: process.env.DATA_MANAGER_ADDRESS || '0xE5C3D068e2160c67e09afaCEff3E765e30163Eb8',
  bvmEigenSequencerAddress: process.env.BVM_EIGENDA_SEQUENCER_ADDRESS || '0x70997970c51812dc3a010c7d01b50e0d17dc79c8'
}

export default config

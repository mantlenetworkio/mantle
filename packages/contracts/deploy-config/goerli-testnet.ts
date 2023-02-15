const config = {
  numDeployConfirmations: 1,
  l1BlockTimeSeconds: 15,
  l2BlockGasLimit: 15_000_000,
  l2ChainId: 5001,
<<<<<<< HEAD
  fixBlockHashBranchingBlock: 222073,
  ctcL2GasDiscountDivisor: 32,
=======
  ctcL2GasDiscountDivisor: 32,
  updateGaslimitBlock: 222073,
>>>>>>> 4739c74 (wb: update config name)
  ctcEnqueueGasCost: 60_000,
  sccFaultProofWindowSeconds: 10,
  sccSequencerPublishWindowSeconds: 12592000,
  blockStaleMeasure: 100,
  daFraudProofPeriod: 120,
  l2SubmittedBlockNumber: 1,
  bvmSequencerAddress: process.env.BVM_SEQUENCER_ADDRESS ||  '0xc4AaE221f1C62E8CBC657Af5b051eA573914cFc7',
  bvmProposerAddress: process.env.BVM_PROPOSER_ADDRESS || '0x3079Be9D8622173f02618bA2B793F00795D4f320',
  //default bvmBlockSignerAddress
  bvmBlockSignerAddress: process.env.BVM_BLOCK_SIGNER_ADDRESS || '0xa9eC80835800a59Fd022f53e3E75AA4552F22ccB',
  bvmFeeWalletAddress: process.env.BVM_FEE_WALLET_ADDRESS || '0x018E08C754018fe54D1CE86b27120052bFe07273',
  bvmAddressManagerOwner: process.env.BVM_ADDRESS_MANAGER_OWNER || '0xf3f0ADB53a250DcCdb8e851081c2608949b97260',
  bvmGasPriceOracleOwner: process.env.BVM_GAS_PRICE_ORACLE_OWNER || '0xAe3e6f7Df1CC6Cf18Fe9F3E69BCFC3351eb4fB45',
  bvmFeeWalletOwner: process.env.BVM_FEE_WALLETOWNER_OWNER || '0xAe3e6f7Df1CC6Cf18Fe9F3E69BCFC3351eb4fB45',
  bvmTssRewardContractOwner: process.env.TssRewardContractOwner || '0xAe3e6f7Df1CC6Cf18Fe9F3E69BCFC3351eb4fB45',
  l1BitAddress: process.env.L1_BIT_ADDRESS || '0x5a94Dc6cc85fdA49d8E9A8b85DDE8629025C42be',
  dataManagerAddress: process.env.DATA_MANAGER_ADDRESS || '0x0000000000000000000000000000000000000000',
  bvmEigenSequencerAddress: process.env.BVM_EIGENDA_SEQUENCER_ADDRESS || '0xc4AaE221f1C62E8CBC657Af5b051eA573914cFc7'
}

export default config

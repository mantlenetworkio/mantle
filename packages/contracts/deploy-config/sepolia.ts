const config = {
  numDeployConfirmations: 1,
  gasPrice: 150_000_000_000,
  l1BlockTimeSeconds: 15,
  l2BlockGasLimit: 15_000_000,
  l2ChainId: 5003003,
  ctcL2GasDiscountDivisor: 32,
  ctcEnqueueGasCost: 60_000,
  sccFaultProofWindowSeconds: 10,
  sccSequencerPublishWindowSeconds: 12592000,
  blockStaleMeasure: 100,
  daFraudProofPeriod: 3600,
  l2SubmittedBlockNumber: 1,
  bvmSequencerAddress: process.env.BVM_SEQUENCER_ADDRESS,
  bvmProposerAddress: process.env.BVM_PROPOSER_ADDRESS,
  bvmRolluperAddress: process.env.BVM_ROLLUPER_ADDRESS,
  bvmBlockSignerAddress: process.env.BVM_BLOCK_SIGNER_ADDRESS,
  bvmFeeWalletAddress: process.env.BVM_FEE_WALLET_ADDRESS,
  bvmAddressManagerOwner: process.env.BVM_ADDRESS_MANAGER_OWNER,
  bvmCrossDomainPauseOwner: process.env.BVM_CROSS_DOMAIN_PAUSE_OWNER,
  bvmGasPriceOracleOwner: process.env.BVM_GAS_PRICE_ORACLE_OWNER,
  bvmFeeWalletOwner: process.env.BVM_FEE_WALLETOWNER_OWNER,
  bvmWhitelistOwner: process.env.BVM_WHITE_LIST_OWNER,
  bvmTssRewardContractOwner: process.env.BVM_TSS_REWARD_CONTRACT_OWNER,
  l1MantleAddress: process.env.L1_MANTLE_ADDRESS,
  proxyL1MantleAddress: process.env.PROXY_L1_MANTLE_ADDRESS,
  dataManagerAddress: process.env.DATA_MANAGER_ADDRESS,
  bvmEigenSequencerAddress: process.env.BVM_EIGENDA_SEQUENCER_ADDRESS,
  bvmEigenFeeAddress: process.env.BVM_EIGENDA_FEE_ADDRESS,
  bvmEigenChallengerAddress: process.env.BVM_EIGENDA_CHALLENGER_ADDRESS,
  contractsDeployerKey: process.env.CONTRACTS_DEPLOYER_KEY,
  contractsRpcUrl: process.env.CONTRACTS_RPC_URL,
  tssRewardSendAmountPerYear: 100000,
  tssRewardWaitingTime: 86400,
  tssDelegationManagerMinStakeAmount: '100000000000000000000',
  tssManagerAddress: process.env.BVM_TSS_MANAGER_ADDRESS,
}

export default config

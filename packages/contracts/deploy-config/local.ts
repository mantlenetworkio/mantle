const config = {
  l1BlockTimeSeconds: 15,
  l2BlockGasLimit: 15_000_000,
  l2ChainId: 17,
  updateGaslimitBlock: 10,
  ctcL2GasDiscountDivisor: 32,
  ctcEnqueueGasCost: 60_000,
  sccFaultProofWindowSeconds: 0,
  sccSequencerPublishWindowSeconds: 12592000,
  blockStaleMeasure: 100,
  daFraudProofPeriod: 120,
  l2SubmittedBlockNumber: 1,
  bvmSequencerAddress: process.env.BVM_SEQUENCER_ADDRESS || '0x70997970c51812dc3a010c7d01b50e0d17dc79c8',
  bvmProposerAddress: process.env.BVM_PROPOSER_ADDRESS || '0x3c44cdddb6a900fa2b585dd299e03d12fa4293bc',
  bvmBlockSignerAddress: process.env.BVM_BLOCK_SIGNER_ADDRESS || '0x00000398232E2064F896018496b4b44b3D62751F',
  bvmFeeWalletAddress: process.env.BVM_FEE_WALLET_ADDRESS || '0xeEbceB07eA7D2339895Dd492B3B5960641302830',
  bvmAddressManagerOwner: process.env.BVM_ADDRESS_MANAGER_OWNER || '0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266',
  bvmGasPriceOracleOwner: process.env.BVM_GAS_PRICE_ORACLE_OWNER || '0x9965507d1a55bcc2695c58ba16fb37d819b0a4dc',
  bvmFeeWalletOwner: process.env.BVM_FEE_WALLETOWNER_OWNER ||'0xDfCDD683728F83A5da850a210db828e6f35DCfA6',
  l1BitAddress: process.env.L1_BIT_ADDRESS || '0x1A4b46696b2bB4794Eb3D4c26f1c55F9170fa4C5',
  bvmTssRewardContractOwner: process.env.TssRewardContractOwner || '0xc8910a1957d276cE5634B978d908B5ef9fB0e05B',
  dataManagerAddress: process.env.DATA_MANAGER_ADDRESS || '0xE5C3D068e2160c67e09afaCEff3E765e30163Eb8',
  bvmEigenSequencerAddress: process.env.BVM_EIGENDA_SEQUENCER_ADDRESS || '0x70997970c51812dc3a010c7d01b50e0d17dc79c8'
}

export default config

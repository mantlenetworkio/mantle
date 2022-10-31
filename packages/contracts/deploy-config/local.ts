const config = {
  l1BlockTimeSeconds: 15,
  l2BlockGasLimit: 15_000_000,
  l2ChainId: 17,
  ctcL2GasDiscountDivisor: 32,
  ctcEnqueueGasCost: 60_000,
  sccFaultProofWindowSeconds: 0,
  sccSequencerPublishWindowSeconds: 12592000,
  bvmSequencerAddress: process.env.BVM_SEQUENCER_ADDRESS || '0x70997970c51812dc3a010c7d01b50e0d17dc79c8',
  bvmProposerAddress: process.env.BVM_PROPOSER_ADDRESS || '0x3c44cdddb6a900fa2b585dd299e03d12fa4293bc',
  bvmBlockSignerAddress: process.env.BVM_BLOCK_SIGNER_ADDRESS || '0x00000398232E2064F896018496b4b44b3D62751F',
  bvmFeeWalletAddress: process.env.BVM_FEE_WALLET_ADDRESS || '0x391716d440c151c42cdf1c95c1d83a5427bca52c',
  bvmAddressManagerOwner: process.env.BVM_ADDRESS_MANAGER_OWNER || '0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266',
  bvmGasPriceOracleOwner: process.env.BVM_GAS_PRICE_ORACLE_OWNER || '0x9965507d1a55bcc2695c58ba16fb37d819b0a4dc',
  l1BitAddress: process.env.L1_BIT_ADDRESS || '0x01BDCf509fE69a87b9787d85728193bAbD5A3d25',
}

export default config

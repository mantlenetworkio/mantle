const config = {
  numDeployConfirmations: 1,
  l1BlockTimeSeconds: 15,
  l2BlockGasLimit: 15_000_000,
  l2ChainId: 5001,
  ctcL2GasDiscountDivisor: 32,
  ctcEnqueueGasCost: 60_000,
  sccFaultProofWindowSeconds: 0,
  sccSequencerPublishWindowSeconds: 12592000,
  bvmSequencerAddress: process.env.BVM_SEQUENCER_ADDRESS ||  '0x99341cC850564B45334d0c3B42C8E6D69F90615c',
  bvmProposerAddress: process.env.BVM_PROPOSER_ADDRESS || '0x2225b2B4a1d26ce51e9B55699E46892a84b3dF7A',
  //default bvmBlockSignerAddress
  bvmBlockSignerAddress: process.env.BVM_BLOCK_SIGNER_ADDRESS || '0x00000398232E2064F896018496b4b44b3D62751F',
  bvmFeeWalletAddress: process.env.BVM_FEE_WALLET_ADDRESS || '0xDfCDD683728F83A5da850a210db828e6f35DCfA6',
  bvmAddressManagerOwner: process.env.BVM_ADDRESS_MANAGER_OWNER || '0xD6b4145dEAbE1b66dD3213080b89AF06EC8828cF',
  bvmGasPriceOracleOwner: process.env.BVM_GAS_PRICE_ORACLE_OWNER || '0xDfCDD683728F83A5da850a210db828e6f35DCfA6',
  l1BitAddress: process.env.L1_BIT_ADDRESS || '0x2f633e75b97569Bb1Ad1A646C736E82402Bb17cf',
}

export default config

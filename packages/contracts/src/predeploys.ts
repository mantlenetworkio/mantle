/**
 * Predeploys are Solidity contracts that are injected into the initial L2 state and provide
 * various useful functions.
 *
 * Notes:
 * 0x42...04 was the address of the BVM_ProxySequencerEntrypoint. This contract is no longer in
 * use and has therefore been removed. We may place a new predeployed contract at this address
 */
export const predeploys = {
  BVM_L2ToL1MessagePasser: '0x4200000000000000000000000000000000000000',
  BVM_DeployerWhitelist: '0x4200000000000000000000000000000000000002',
  L2CrossDomainMessenger: '0x4200000000000000000000000000000000000007',
  BVM_GasPriceOracle: '0x420000000000000000000000000000000000000F',
  L2StandardBridge: '0x4200000000000000000000000000000000000010',
  BVM_SequencerFeeVault: '0x4200000000000000000000000000000000000011',
  L2StandardTokenFactory: '0x4200000000000000000000000000000000000012',
  BVM_L1BlockNumber: '0x4200000000000000000000000000000000000013',

  // We're temporarily disabling BVM_BIT because the jury is still out on whether BIT as an
  // ERC20 is desirable. ETH on Layer 2 will be the standard ERC20 token
  BVM_BIT: '0xDeadDeAddeAddEAddeadDEaDDEAdDeaDDeAD0000',
  BVM_ETH: '0xdEAddEaDdeadDEadDEADDEAddEADDEAddead1111',

  // We're also putting WETH9 at the old BVM_ETH address.
  WETH9: '0x4200000000000000000000000000000000000006',

  // Tss rewards default contract
  TssRewardContract: '0x4200000000000000000000000000000000000020',
}

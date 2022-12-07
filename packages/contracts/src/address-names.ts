/**
 * This object defines the correct names to be used in the Address Manager and deployment artifacts.
 */
export const names = {
  managed: {
    contracts: {
      ChainStorageContainer_CTC_batches: 'ChainStorageContainer-CTC-batches',
      ChainStorageContainer_SCC_batches: 'ChainStorageContainer-SCC-batches',
      CanonicalTransactionChain: 'CanonicalTransactionChain',
      StateCommitmentChain: 'StateCommitmentChain',
      BondManager: 'BondManager',
      BVM_L1CrossDomainMessenger: 'BVM_L1CrossDomainMessenger',
      Proxy__BVM_L1CrossDomainMessenger: 'Proxy__BVM_L1CrossDomainMessenger',
      TssGroupManager: 'TssGroupManager',
      TssStakingSlashing: 'TssStakingSlashing',
      Proxy__BVM_L1StandardBridge: 'Proxy__BVM_L1StandardBridge',
      Proxy__TSS_GroupManager: 'Proxy__TSS_GroupManager',
      Proxy__TSS_StakingSlashing: 'Proxy__TSS_StakingSlashing',
    },
    da: {
      BVM_EigenDataLayrChain: 'BVM_EigenDataLayrChain',
      Proxy__BVM_EigenDataLayrChain: 'Proxy__BVM_EigenDataLayrChain',
    },
    accounts: { BVM_Sequencer: 'BVM_Sequencer', BVM_Proposer: 'BVM_Proposer' },
    configs: {
      L1_BIT_ADDRESS: 'L1_BitAddress',
      Local_Bit_Token: 'TestBitToken',
    },
  },
  unmanaged: {
    AddressDictator: 'AddressDictator',
    ChugSplashDictator: 'ChugSplashDictator',
    Lib_AddressManager: 'Lib_AddressManager',
  },
}

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
    delegation: {
      fraud_proof: {
        FraudProofDelegation: 'FraudProofDelegation',
        FraudProofDelegationManager: 'FraudProofDelegationManager',
        FraudProofDelegationSlasher: 'FraudProofDelegationSlasher',
        Proxy__FraudProofDelegation: 'Proxy__FraudProofDelegation',
        Proxy__FraudProofDelegationManager: 'Proxy__FraudProofDelegationManager',
        Proxy__FraudProofDelegationSlasher: 'Proxy__FraudProofDelegationSlasher',
      },
      tss: {
        EmptyContract: 'EmptyContract',
        TssDelegation: 'TssDelegation',
        TssDelegationManager: 'TssDelegationManager',
        TssDelegationSlasher: 'TssDelegationSlasher',
        Proxy__TssDelegation: 'Proxy__TssDelegation',
        Proxy__TssDelegationManager: 'Proxy__TssDelegationManager',
        Proxy__TssDelegationSlasher: 'Proxy__TssDelegationSlasher',
      },
    },
    da: {
      BVM_EigenDataLayrChain: 'BVM_EigenDataLayrChain',
      Proxy__BVM_EigenDataLayrChain: 'Proxy__BVM_EigenDataLayrChain',
      BVM_EigenDataLayrFee: 'BVM_EigenDataLayrFee',
      Proxy__BVM_EigenDataLayrFee: 'Proxy__BVM_EigenDataLayrFee',
    },
    fraud_proof: {
      AssertionMap: 'AssertionMap',
      VerifierEntry: 'VerifierEntry',
      Rollup: 'Rollup',
      Proxy__AssertionMap: 'Proxy__AssertionMap',
      Proxy__Verifier: 'Proxy__Verifier',
      Proxy__Rollup: 'Proxy__Rollup',
      SubVerifiers: {
        BlockInitiationVerifier: 'BlockInitiationVerifier',
        BlockFinalizationVerifier: 'BlockFinalizationVerifier',
        InterTxVerifier: 'InterTxVerifier',
        StackOpVerifier: 'StackOpVerifier',
        EnvironmentalOpVerifier: 'EnvironmentalOpVerifier',
        MemoryOpVerifier: 'MemoryOpVerifier',
        StorageOpVerifier: 'StorageOpVerifier',
        CallOpVerifier: 'CallOpVerifier',
        InvalidOpVerifier: 'InvalidOpVerifier',
      },
    },
    accounts: {
      BVM_Sequencer: 'BVM_Sequencer',
      BVM_Proposer: 'BVM_Proposer',
      BVM_Rolluper: 'BVM_Rolluper',
    },
    configs: {
      L1_MANTLE_ADDRESS: 'L1_MantleAddress',
      Local_Mantle_Token: 'L1MantleToken',
      Proxy__Local_Mantle_Token: 'Proxy__L1MantleToken',
    },
  },
  unmanaged: {
    AddressDictator: 'AddressDictator',
    ChugSplashDictator: 'ChugSplashDictator',
    Lib_AddressManager: 'Lib_AddressManager',
  },
}

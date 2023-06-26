/* Imports: External */
import { DeployFunction } from 'hardhat-deploy/dist/types'
import { hexStringEquals, awaitCondition } from '@mantleio/core-utils'
import { ethers } from 'ethers'

import { names } from '../src/address-names'
import {
  deployAndVerifyAndThen,
  getContractFromArtifact,
} from '../src/deploy-utils'
import { getContractFactory } from '../src'

const deployFn: DeployFunction = async (hre) => {
  const { deployer } = await hre.getNamedAccounts()

  const owner = hre.deployConfig.bvmAddressManagerOwner
  const l1MantleAddress = hre.deployConfig.proxyL1MantleAddress
  const minStakeAmount = hre.deployConfig.tssDelegationManagerMinStakeAmount
  const tssManagerAddress = hre.deployConfig.tssManagerAddress

  //deploy EmptyContract
  await deployAndVerifyAndThen({
    hre,
    name: names.managed.delegation.tss.EmptyContract,
    contract: 'EmptyContract',
    args: [],
  })
  const emptyContract = await getContractFromArtifact(
    hre,
    names.managed.delegation.tss.EmptyContract,
    {
      iface: 'EmptyContract',
      signerOrProvider: deployer,
    }
  )
  console.log('EmptyContract Address', emptyContract.address)
  console.log('deploy EmptyContract success')

  //deploy Delegation proxy with empty contract
  await deployAndVerifyAndThen({
    hre,
    name: names.managed.delegation.tss.Proxy__TssDelegation,
    contract: 'TransparentUpgradeableProxy',
    iface: 'TssDelegation',
    args: [emptyContract.address, deployer, []],
  })
  const Proxy_TssDelegation = await getContractFromArtifact(
    hre,
    names.managed.delegation.tss.Proxy__TssDelegation,
    {
      iface: 'TssDelegation',
      signerOrProvider: deployer,
    }
  )
  console.log('Proxy__TssDelegation Address', Proxy_TssDelegation.address)
  console.log('deploy Tss Delegation Proxy__TssDelegation success')

  //deploy DelegationManager proxy with empty contract
  await deployAndVerifyAndThen({
    hre,
    name: names.managed.delegation.tss.Proxy__TssDelegationManager,
    contract: 'TransparentUpgradeableProxy',
    iface: 'TssDelegationManager',
    args: [emptyContract.address, deployer, []],
  })
  const Proxy_TssDelegationManager = await getContractFromArtifact(
    hre,
    names.managed.delegation.tss.Proxy__TssDelegationManager,
    {
      iface: 'TssDelegationManager',
      signerOrProvider: deployer,
    }
  )
  console.log(
    'Proxy__TssDelegationManager Address',
    Proxy_TssDelegationManager.address
  )
  console.log(
    'deploy Tss Delegation Manager Proxy__TssDelegationManager success'
  )

  //depoly DelegationSlasher proxy with empty contract
  await deployAndVerifyAndThen({
    hre,
    name: names.managed.delegation.tss.Proxy__TssDelegationSlasher,
    contract: 'TransparentUpgradeableProxy',
    iface: 'TssDelegationSlasher',
    args: [emptyContract.address, deployer, []],
  })
  const Proxy_TssDelegationSlasher = await getContractFromArtifact(
    hre,
    names.managed.delegation.tss.Proxy__TssDelegationSlasher,
    {
      iface: 'TssDelegationSlasher',
      signerOrProvider: deployer,
    }
  )
  console.log(
    'Proxy_TssDelegationSlasher Address',
    Proxy_TssDelegationSlasher.address
  )
  console.log(
    'deploy Tss Delegation Manager Proxy_TssDelegationSlasher success'
  )

  // deploy impl
  await deployAndVerifyAndThen({
    hre,
    name: names.managed.contracts.TssGroupManager,
    contract: 'TssGroupManager',
    args: [],
  })
  console.log('deploy tss group manager success')

  await deployAndVerifyAndThen({
    hre,
    name: names.managed.contracts.TssStakingSlashing,
    contract: 'TssStakingSlashing',
    args: [],
  })
  console.log('deploy tss staking slashing success')

  // deploy proxy
  const Impl_TSS_GroupManager = await getContractFromArtifact(
    hre,
    names.managed.contracts.TssGroupManager,
    {
      iface: 'TssGroupManager',
      signerOrProvider: deployer,
    }
  )

  const Impl__TssStakingSlashing = await getContractFromArtifact(
    hre,
    names.managed.contracts.TssStakingSlashing,
    {
      iface: 'TssStakingSlashing',
      signerOrProvider: deployer,
    }
  )

  let args: unknown[]
  args = []
  let callData = Impl_TSS_GroupManager.interface.encodeFunctionData(
    'initialize',
    args
  )
  await deployAndVerifyAndThen({
    hre,
    name: names.managed.contracts.Proxy__TSS_GroupManager,
    contract: 'TransparentUpgradeableProxy',
    iface: 'TssGroupManager',
    args: [Impl_TSS_GroupManager.address, owner, callData],
  })

  const Proxy__TSS_GroupManager = await getContractFromArtifact(
    hre,
    names.managed.contracts.Proxy__TSS_GroupManager,
    {
      iface: 'TssGroupManager',
      signerOrProvider: deployer,
    }
  )
  console.log(
    'Proxy__TSS_GroupManager Address',
    Proxy__TSS_GroupManager.address
  )
  console.log('deploy tss group manager proxy success')

  const Proxy__BVM_L1CrossDomainMessenger = await getContractFromArtifact(
    hre,
    names.managed.contracts.Proxy__BVM_L1CrossDomainMessenger
  )

  args = [
    l1MantleAddress,
    Proxy__TSS_GroupManager.address,
    Proxy_TssDelegationManager.address,
    Proxy_TssDelegation.address,
    Proxy__BVM_L1CrossDomainMessenger.address,
    owner,
    tssManagerAddress,
  ]
  callData = Impl__TssStakingSlashing.interface.encodeFunctionData(
    'initialize',
    args
  )
  console.log('encode stakingslashing initialize calldata')
  await deployAndVerifyAndThen({
    hre,
    name: names.managed.contracts.Proxy__TSS_StakingSlashing,
    contract: 'TransparentUpgradeableProxy',
    iface: 'TssStakingSlashing',
    args: [Impl__TssStakingSlashing.address, owner, callData],
    postDeployAction: async (contract) => {
      console.log(`Checking that contract was correctly initialized...`)
      await awaitCondition(
        async () => {
          return hexStringEquals(
            await contract
              .connect(Impl_TSS_GroupManager.signer.provider)
              .underlyingToken({ from: ethers.constants.AddressZero }),
            l1MantleAddress
          )
        },
        5000,
        100
      )
      await awaitCondition(
        async () => {
          return hexStringEquals(
            await contract
              .connect(Impl_TSS_GroupManager.signer.provider)
              .tssGroupContract({ from: ethers.constants.AddressZero }),
            Proxy__TSS_GroupManager.address
          )
        },
        5000,
        100
      )

      await Proxy__TSS_GroupManager.setStakingSlash(contract.address)
      await awaitCondition(
        async () => {
          return hexStringEquals(
            await Proxy__TSS_GroupManager.connect(
              Impl_TSS_GroupManager.signer.provider
            ).stakingSlash({ from: ethers.constants.AddressZero }),
            contract.address
          )
        },
        5000,
        100
      )
    },
  })

  const Proxy__TSS_StakingSlashing = await getContractFromArtifact(
    hre,
    names.managed.contracts.Proxy__TSS_StakingSlashing,
    {
      iface: 'TssStakingSlashing',
      signerOrProvider: deployer,
    }
  )
  console.log(
    'Proxy__TSS_StakingSlashing Address',
    Proxy__TSS_StakingSlashing.address
  )
  console.log('deploy tss staking slashing proxy success')

  // deploy Delegation impl
  await deployAndVerifyAndThen({
    hre,
    name: names.managed.delegation.tss.TssDelegation,
    contract: 'TssDelegation',
    args: [Proxy_TssDelegationManager.address],
  })
  const Impl__TssDelegation = await getContractFromArtifact(
    hre,
    names.managed.delegation.tss.TssDelegation,
    {
      iface: 'TssDelegation',
      signerOrProvider: deployer,
    }
  )
  console.log(
    'TssDelegation Implementation Address',
    Impl__TssDelegation.address
  )
  console.log('deploy Tss Delegation success')

  // deploy DelegationManager impl
  await deployAndVerifyAndThen({
    hre,
    name: names.managed.delegation.tss.TssDelegationManager,
    contract: 'TssDelegationManager',
    args: [Proxy_TssDelegation.address, Proxy_TssDelegationSlasher.address],
  })
  const Impl__TssDelegationManager = await getContractFromArtifact(
    hre,
    names.managed.delegation.tss.TssDelegationManager,
    {
      iface: 'TssDelegationManager',
      signerOrProvider: deployer,
    }
  )
  console.log(
    'TssDelegationManager Implementation Address',
    Impl__TssDelegationManager.address
  )
  console.log('deploy Tss DelegationManager success')

  // deploy DelegationSlasher impl
  await deployAndVerifyAndThen({
    hre,
    name: names.managed.delegation.tss.TssDelegationSlasher,
    contract: 'TssDelegationSlasher',
    args: [Proxy_TssDelegationManager.address, Proxy_TssDelegation.address],
  })
  const Impl__TssDelegationSlasher = await getContractFromArtifact(
    hre,
    names.managed.delegation.tss.TssDelegationSlasher,
    {
      iface: 'TssDelegationSlasher',
      signerOrProvider: deployer,
    }
  )
  console.log(
    'Tss DelegationSlasher Implementation Address',
    Impl__TssDelegationSlasher.address
  )
  console.log('deploy Tss DelegationSlasher success')

  // const from = deployer
  const provider = new ethers.providers.JsonRpcBatchProvider(
    hre.deployConfig.contractsRpcUrl
  )
  const deployerWallet = new ethers.Wallet(
    hre.deployConfig.contractsDeployerKey,
    provider
  )
  console.log('deploy privete key', hre.deployConfig.contractsDeployerKey)
  console.log('privder', provider)

  // update Delegation proxy
  const delegationProxy = getContractFactory(
    'TransparentUpgradeableProxy'
  ).attach(Proxy_TssDelegation.address)
  callData = Impl__TssDelegation.interface.encodeFunctionData('initializeT', [
    Proxy__TSS_StakingSlashing.address,
    deployer,
  ])
  await delegationProxy
    .connect(deployerWallet)
    .upgradeToAndCall(Impl__TssDelegation.address, callData, {
      gasLimit: 2_000_000,
    })
  console.log('update Tss Delegation Proxy__TssDelegation success')
  await delegationProxy.connect(deployerWallet).changeAdmin(owner, {
    gasLimit: 2_000_000,
  })
  console.log('update Tss Delegation Proxy__TssDelegation admin success')

  //update Delegation Slasher proxy
  const delegationSlasherProxy = getContractFactory(
    'TransparentUpgradeableProxy'
  ).attach(Proxy_TssDelegationSlasher.address)
  callData = Impl__TssDelegationSlasher.interface.encodeFunctionData(
    'initialize',
    [deployer]
  )
  await delegationSlasherProxy
    .connect(deployerWallet)
    .upgradeToAndCall(Impl__TssDelegationSlasher.address, callData, {
      gasLimit: 2_000_000,
    })
  console.log(
    'update Tss Delegation Slasher Proxy_TssDelegationSlasher success'
  )
  await delegationSlasherProxy.connect(deployerWallet).changeAdmin(owner, {
    gasLimit: 2_000_000,
  })
  console.log(
    'update Tss Delegation Slasher Proxy_TssDelegationSlasher admin success'
  )

  //update delegation manager proxy
  const delegationManagerProxy = getContractFactory(
    'TransparentUpgradeableProxy'
  ).attach(Proxy_TssDelegationManager.address)
  callData = Impl__TssDelegationManager.interface.encodeFunctionData(
    'initializeT',
    [
      Proxy__TSS_StakingSlashing.address,
      Proxy__TSS_GroupManager.address,
      minStakeAmount,
      deployer,
    ]
  )
  await delegationManagerProxy
    .connect(deployerWallet)
    .upgradeToAndCall(Impl__TssDelegationManager.address, callData, {
      gasLimit: 2_000_000,
    })
  console.log(
    'update Tss Delegation manager Proxy_TssDelegationManager success'
  )
  await delegationManagerProxy.connect(deployerWallet).changeAdmin(owner, {
    gasLimit: 2_000_000,
  })
  console.log(
    'update Tss Delegation manager Proxy_TssDelegationManager admin success'
  )
}

// This is kept during an upgrade. So no upgrade tag.
deployFn.tags = ['TssContracts', 'upgrade']

export default deployFn

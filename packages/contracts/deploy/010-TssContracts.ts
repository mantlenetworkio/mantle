/* Imports: External */
import { DeployFunction } from 'hardhat-deploy/dist/types'
import { hexStringEquals, awaitCondition } from '@mantleio/core-utils'
import { ethers } from 'ethers'

import { names } from '../src/address-names'
import {
  deployAndVerifyAndThen,
  getContractFromArtifact,
} from '../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const { deployer } = await hre.getNamedAccounts()

  const owner = hre.deployConfig.bvmAddressManagerOwner
  const l1BitAddress = hre.deployConfig.l1BitAddress
  let DelegationProxyAddress = "0xe6cd9e7b620964bECd42c7Ad41e56724f515E284"
  let DelegationManagerProxyAddress = "0xE6A251EefaEE70E8645FBAdf21E9B1246e07C374"
  let DelegationSlasherProxyAddress = "0xD007896d9E3e4514a1f1216A91d33a72e15bf5C0"
  let StakingSlashingProxyAddress = "0x9c28c8D298ae7Ebf8daA6FA54e1F2909313dB158"
  let TssGroupManagerProxyAddress = "0xF48398a3D94D57AE1406B343D2a7C541336Ea2c2"

  // deploy Delegation impl
  await deployAndVerifyAndThen({
    hre,
    name: names.managed.delegation.tss.TssDelegation,
    contract: 'TssDelegation',
    args: [DelegationManagerProxyAddress],
  })
  const Impl__TssDelegation = await getContractFromArtifact(
    hre,
    names.managed.delegation.tss.TssDelegation,
    {
      iface: 'TssDelegation',
      signerOrProvider: deployer,
    }
  )
  console.log('TssDelegation Implementation Address', Impl__TssDelegation.address)
  console.log('deploy Tss Delegation success')

  // deploy DelegationManager impl
  await deployAndVerifyAndThen({
    hre,
    name: names.managed.delegation.tss.TssDelegationManager,
    contract: 'TssDelegationManager',
    args: [DelegationProxyAddress, DelegationSlasherProxyAddress],
  })
  const Impl__TssDelegationManager = await getContractFromArtifact(
    hre,
    names.managed.delegation.tss.TssDelegationManager,
    {
      iface: 'TssDelegationManager',
      signerOrProvider: deployer,
    }
  )
  console.log('TssDelegationManager Implementation Address', Impl__TssDelegationManager.address)
  console.log('deploy Tss DelegationManager success')

  // deploy DelegationSlasher impl
  await deployAndVerifyAndThen({
    hre,
    name: names.managed.delegation.tss.TssDelegationSlasher,
    contract: 'TssDelegationSlasher',
    args: [DelegationManagerProxyAddress, DelegationProxyAddress],
  })
  const Impl__TssDelegationSlasher = await getContractFromArtifact(
    hre,
    names.managed.delegation.tss.TssDelegationSlasher,
    {
      iface: 'TssDelegationSlasher',
      signerOrProvider: deployer,
    }
  )
  console.log('Tss DelegationSlasher Implementation Address', Impl__TssDelegationSlasher.address)
  console.log('deploy Tss DelegationSlasher success')

  // deploy Delegation proxy
  let callData = Impl__TssDelegation.interface.encodeFunctionData(
    'initializeT',
    [StakingSlashingProxyAddress,
      deployer,]
    )

  await deployAndVerifyAndThen({
    hre,
    name: names.managed.delegation.tss.Proxy__TssDelegation,
    contract: 'TransparentUpgradeableProxy',
    iface: 'TssDelegation',
    args: [Impl__TssDelegation.address, owner, callData],
  })
  const Proxy__TssDelegation = await getContractFromArtifact(
    hre,
    names.managed.delegation.tss.Proxy__TssDelegation,
    {
      iface: 'TssDelegation',
      signerOrProvider: deployer,
    }
  )
  console.log('Proxy__TssDelegation Address', Proxy__TssDelegation.address)
  console.log('deploy Tss Delegation Proxy__TssDelegation success')

  callData = Impl__TssDelegationSlasher.interface.encodeFunctionData(
    'initialize',
    [deployer]
  )
  await deployAndVerifyAndThen({
    hre,
    name: names.managed.delegation.tss.Proxy__TssDelegationSlasher,
    contract: 'TransparentUpgradeableProxy',
    iface: 'TssDelegationSlasher',
    args: [Impl__TssDelegationSlasher.address, owner, callData],
  })

  const Proxy__TssDelegationSlasher = await getContractFromArtifact(
    hre,
    names.managed.delegation.tss.Proxy__TssDelegationSlasher,
    {
      iface: 'TssDelegationSlasher',
      signerOrProvider: deployer,
    }
  )
  console.log('Proxy__TssDelegationSlasher Address', Proxy__TssDelegationSlasher.address)
  console.log('deploy Tss DelegationSlasher Proxy success')

  callData = Impl__TssDelegationManager.interface.encodeFunctionData(
    'initializeT',
    [StakingSlashingProxyAddress,
      TssGroupManagerProxyAddress,
      '10000',
      deployer,
    ]
  )
  await deployAndVerifyAndThen({
    hre,
    name: names.managed.delegation.tss.Proxy__TssDelegationManager,
    contract: 'TransparentUpgradeableProxy',
    iface: 'TssDelegationManager',
    args: [Impl__TssDelegationManager.address, owner, callData],
  })

  const Proxy__TssDelegationManager = await getContractFromArtifact(
    hre,
    names.managed.delegation.tss.Proxy__TssDelegationManager,
    {
      iface: 'TssDelegationManager',
      signerOrProvider: deployer,
    }
  )
  console.log('Proxy__TssDelegationManager Address', Proxy__TssDelegationManager.address)
  console.log('deploy Tss DelegationManager Proxy success')

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
  callData = Impl_TSS_GroupManager.interface.encodeFunctionData(
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
  console.log('deploy tss group manager proxy success')

  const Proxy__TSS_GroupManager = await getContractFromArtifact(
    hre,
    names.managed.contracts.Proxy__TSS_GroupManager,
    {
      iface: 'TssGroupManager',
      signerOrProvider: deployer,
    }
  )

  const Proxy__BVM_L1CrossDomainMessenger = await getContractFromArtifact(
    hre,
    names.managed.contracts.Proxy__BVM_L1CrossDomainMessenger
  )

  args = [
    l1BitAddress,
    Proxy__TSS_GroupManager.address,
    Proxy__TssDelegationManager.address,
    Proxy__TssDelegation.address,
    Proxy__BVM_L1CrossDomainMessenger.address,
    owner,
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
            l1BitAddress
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

      // await contract.transferOwnership(owner)

      // console.log(`Checking tss staking slashing contract owner was correctly set...`)
      // await awaitCondition(
      //     async () => {
      //         return hexStringEquals(await contract.connect(Impl_TSS_GroupManager.signer.provider).owner({ from: ethers.constants.AddressZero }), owner)
      //     },
      //     5000,
      //     100
      // )

      // await Proxy__TSS_GroupManager.transferOwnership(owner)
      // console.log(`Checking tss group contract manager owner was correctly set...`)
      // await awaitCondition(
      //     async () => {
      //         return hexStringEquals(await contract.connect(Impl_TSS_GroupManager.signer.provider).owner({ from: ethers.constants.AddressZero }), owner)
      //     },
      //     5000,
      //     100
      // )
    },
  })
  console.log('deploy tss staking slashing proxy success')
}

// This is kept during an upgrade. So no upgrade tag.
deployFn.tags = ['TssContracts', 'upgrade']

export default deployFn

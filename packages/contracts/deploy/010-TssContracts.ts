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
  console.log('deploy tss group manager proxy success')

  const Proxy__TSS_GroupManager = await getContractFromArtifact(
    hre,
    names.managed.contracts.Proxy__TSS_GroupManager,
    {
      iface: 'TssGroupManager',
      signerOrProvider: deployer,
    }
  )

  args = [l1BitAddress, Proxy__TSS_GroupManager.address]
  callData = Impl__TssStakingSlashing.interface.encodeFunctionData(
    'initialize',
    args
  )
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
              .BitToken({ from: ethers.constants.AddressZero }),
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

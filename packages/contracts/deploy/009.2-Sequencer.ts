/* Imports: External */
import { DeployFunction } from 'hardhat-deploy/dist/types'
import { hexStringEquals, awaitCondition } from '@mantlenetworkio/core-utils'
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
    name: names.managed.contracts.Sequencer,
    contract: 'Sequencer',
    args: [],
  })
  console.log('deploy Sequencer success')

  // deploy proxy
  const Impl_Sequencer = await getContractFromArtifact(
    hre,
    names.managed.contracts.Sequencer,
    {
      iface: 'Sequencer',
      signerOrProvider: deployer,
    }
  )

  const args = [l1BitAddress]
  const callData = Impl_Sequencer.interface.encodeFunctionData(
    'initialize',
    args
  )
  await deployAndVerifyAndThen({
    hre,
    name: names.managed.contracts.Proxy__Sequencer,
    contract: 'TransparentUpgradeableProxy',
    iface: 'Sequencer',
    args: [Impl_Sequencer.address, owner, callData],
    postDeployAction: async (contract) => {
      console.log(`Checking that contract was correctly initialized...`)
      await awaitCondition(
        async () => {
          return hexStringEquals(
            await contract
              .connect(Impl_Sequencer.signer.provider)
              .bitToken({ from: ethers.constants.AddressZero }),
            l1BitAddress
          )
        },
        5000,
        100
      )
    },
  })
  console.log('deploy Sequencer proxy success')
}

// This is kept during an upgrade. So no upgrade tag.
deployFn.tags = ['Sequencer', 'upgrade']

export default deployFn

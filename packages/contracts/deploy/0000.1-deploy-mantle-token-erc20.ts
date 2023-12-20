/* Imports: External */
import { DeployFunction } from 'hardhat-deploy/dist/types'

import { names } from '../src/address-names'
import {
  deployAndVerifyAndThen,
  getContractFromArtifact, isHardhatNode,
} from '../src/deploy-utils'
import {awaitCondition, hexStringEquals} from "@mantleio/core-utils";

const deployFn: DeployFunction = async (hre) => {
  const { deployer } = await hre.getNamedAccounts()

  const owner = hre.deployConfig.bvmAddressManagerOwner

  await deployAndVerifyAndThen({
    hre,
    name: names.managed.configs.Local_Mantle_Token,
    contract: 'L1MantleToken',
    args: [],
  })
  console.log('deploy l1 mantle token success')

  const Impl_L1_MantleToken = await getContractFromArtifact(
    hre,
    names.managed.configs.Local_Mantle_Token,
    {
      iface: 'L1MantleToken',
      signerOrProvider: deployer,
    }
  )
  console.log(`Checking the mantle token was correctly set...`)
  console.log('deployed mantle token address: ', Impl_L1_MantleToken.address)
  console.log(
    'mantle token address in setting: ',
    hre.deployConfig.l1MantleAddress
  )

  if (await isHardhatNode(hre)) {
    await awaitCondition(
      async () => {
        return hexStringEquals(Impl_L1_MantleToken.address, hre.deployConfig.l1MantleAddress)
      },
      5000,
      1
    )
    console.log(`Check pass`)
  }
  console.log(`Check pass`)

  console.log(
    'mantle token proxy address in setting: ',
    hre.deployConfig.proxyL1MantleAddress
  )

  const args = [
    "10000000000000000000000000000",  // 1e10 mnt
    deployer,
  ]

  const callData = Impl_L1_MantleToken.interface.encodeFunctionData(
    'initialize',
    args
  )

  await deployAndVerifyAndThen({
    hre,
    name: names.managed.configs.Proxy__Local_Mantle_Token,
    contract: 'TransparentUpgradeableProxy',
    iface: 'L1MantleToken',
    args: [Impl_L1_MantleToken.address, owner, callData],
  })
  console.log('deploy l1 mantle token proxy success')
}

deployFn.tags = ['mantle-token', 'upgrade']

export default deployFn

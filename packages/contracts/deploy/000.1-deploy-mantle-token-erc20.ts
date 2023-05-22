/* Imports: External */
import {DeployFunction} from 'hardhat-deploy/dist/types'

import {awaitCondition, hexStringEquals} from '@mantleio/core-utils'

/* Imports: Internal */
import {deployAndVerifyAndThen, getContractFromArtifact,isHardhatNode, } from '../src/deploy-utils'
import {names} from '../src/address-names'

const deployFn: DeployFunction = async (hre) => {
  await deployAndVerifyAndThen({
    hre,
    name: names.managed.configs.Local_Mantle_Token,
    contract: 'L1MantleToken',
    args: ['Mantle', 'MNT'],
  })

  const {deployer} = await hre.getNamedAccounts()
  const mantleContract = await getContractFromArtifact(
    hre,
    names.managed.configs.Local_Mantle_Token,
    {
      iface: 'L1MantleToken',
      signerOrProvider: deployer,
    }
  )
  console.log(`Checking the mantle token was correctly set...`)
  console.log(mantleContract.address)
  console.log(hre.deployConfig.l1MantleAddress)
  if (await isHardhatNode(hre)) {
    await awaitCondition(
      async () => {
        return hexStringEquals(mantleContract.address, hre.deployConfig.l1MantleAddress)
      },
      5000,
      1
    )
    console.log(`Check pass`)
  }
  console.log(`Check pass`)
}

deployFn.tags = ['mantle-token', 'upgrade']

export default deployFn

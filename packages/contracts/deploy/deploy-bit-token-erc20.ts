/* Imports: External */
import {DeployFunction} from 'hardhat-deploy/dist/types'

import {awaitCondition, hexStringEquals} from '@mantlenetworkio/core-utils'

/* Imports: Internal */
import {deployAndVerifyAndThen, getContractFromArtifact,} from '../src/deploy-utils'
import {names} from '../src/address-names'

const deployFn: DeployFunction = async (hre) => {
  await deployAndVerifyAndThen({
    hre,
    name: names.managed.configs.Local_Bit_Token,
    contract: 'BitTokenERC20',
    args: ['BitToken', 'BIT'],
  })

  const {deployer} = await hre.getNamedAccounts()
  const bitContract = await getContractFromArtifact(
    hre,
    names.managed.configs.Local_Bit_Token,
    {
      iface: 'BitTokenERC20',
      signerOrProvider: deployer,
    }
  )
  console.log(`Checking the bit token was correctly set...`)
  await awaitCondition(
    async () => {
      return hexStringEquals(bitContract.address, hre.deployConfig.l1BitAddress)
    },
    5000,
    1
  )
  console.log(`Check pass`)
}

deployFn.tags = ['bit-token', 'upgrade']

export default deployFn

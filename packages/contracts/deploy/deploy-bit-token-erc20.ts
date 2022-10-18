/* Imports: External */
import { DeployFunction } from 'hardhat-deploy/dist/types'

/* Imports: Internal */
import {
  deployAndVerifyAndThen,
  getContractFromArtifact,
} from '../src/deploy-utils'
import { names } from '../src/address-names'

const deployFn: DeployFunction = async (hre) => {
  await deployAndVerifyAndThen({
    hre,
    name: names.managed.configs.Local_Bit_Token,
    contract: 'BitTokenERC20',
    args: ['BitToken', 'BIT'],
  })
}

deployFn.tags = ['bit-token', 'upgrade']

export default deployFn

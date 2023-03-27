/* Imports: External */
import {DeployFunction} from 'hardhat-deploy/dist/types'

import { assertContractVariable, deploy } from '../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const { deployer } = await hre.getNamedAccounts()

  await deploy({
    hre,
    name: 'TestBitToken',
    contract: 'BitTokenERC20',
    args: ['BitToken', 'BIT'],
    postDeployAction: async (contract) => {
      // Owner is temporarily set to the deployer.
      await assertContractVariable(contract, 'owner', deployer)
    },
  })
}

deployFn.tags = ['L1BitToken', 'setup']

export default deployFn

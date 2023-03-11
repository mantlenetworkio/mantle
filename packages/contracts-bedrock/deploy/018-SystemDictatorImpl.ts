import { DeployFunction } from 'hardhat-deploy/dist/types'
import '@eth-mantle/hardhat-deploy-config'
import 'hardhat-deploy'

import { deploy } from '../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  await deploy({
    hre,
    name: 'SystemDictator',
    args: [],
  })
}

deployFn.tags = ['SystemDictatorImpl', 'setup']

export default deployFn

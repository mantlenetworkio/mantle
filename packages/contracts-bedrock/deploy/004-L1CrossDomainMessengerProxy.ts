import { DeployFunction } from 'hardhat-deploy/dist/types'

import { deploy, getDeploymentAddress } from '../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const addressManager = await getDeploymentAddress(hre, 'Lib_AddressManager')

  await deploy({
    hre,
    name: 'Proxy__BVM_L1CrossDomainMessenger',
    contract: 'ResolvedDelegateProxy',
    args: [addressManager, 'BVM_L1CrossDomainMessenger'],
  })
}

deployFn.tags = ['L1CrossDomainMessengerProxy', 'setup']

export default deployFn

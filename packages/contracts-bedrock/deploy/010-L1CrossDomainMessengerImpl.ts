import { DeployFunction } from 'hardhat-deploy/dist/types'

import {
  assertContractVariable,
  deploy,
  getContractFromArtifact,
} from '../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const MantlePortalProxy = await getContractFromArtifact(
    hre,
    'MantlePortalProxy'
  )

  await deploy({
    hre,
    name: 'L1CrossDomainMessenger',
    args: [MantlePortalProxy.address],
    postDeployAction: async (contract) => {
      await assertContractVariable(
        contract,
        'PORTAL',
        MantlePortalProxy.address
      )
    },
  })
}

deployFn.tags = ['L1CrossDomainMessengerImpl', 'setup']

export default deployFn

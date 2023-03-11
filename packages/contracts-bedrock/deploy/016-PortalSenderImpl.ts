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
    name: 'PortalSender',
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

deployFn.tags = ['PortalSenderImpl', 'setup']

export default deployFn

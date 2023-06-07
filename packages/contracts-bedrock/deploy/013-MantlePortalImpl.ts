import { DeployFunction } from 'hardhat-deploy/dist/types'
import '@mantleio/hardhat-deploy-config'

import {
  assertContractVariable,
  deploy,
  getContractFromArtifact,
} from '../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const { deployer } = await hre.getNamedAccounts()
  const isLiveDeployer =
    deployer.toLowerCase() === hre.deployConfig.controller.toLowerCase()

  const L2OutputOracleProxy = await getContractFromArtifact(
    hre,
    'L2OutputOracleProxy'
  )

  const finalSystemOwner = hre.deployConfig.finalSystemOwner
  const finalSystemOwnerCode = await hre.ethers.provider.getCode(
    finalSystemOwner
  )
  if (finalSystemOwnerCode === '0x') {
    console.log(
      `WARNING: setting MantlePortal.GUARDIAN to ${finalSystemOwner} and it has no code`
    )
    if (!isLiveDeployer) {
      throw new Error(
        'Do not deploy to production networks without the GUARDIAN being a contract'
      )
    }
  }

  // Deploy the MantlePortal implementation as paused to
  // ensure that users do not interact with it and instead
  // interact with the proxied contract.
  // The `finalSystemOwner` is set at the GUARDIAN.
  await deploy({
    hre,
    name: 'MantlePortal',
    args: [
      L2OutputOracleProxy.address,
      finalSystemOwner,
      true, // paused
    ],
    postDeployAction: async (contract) => {
      await assertContractVariable(
        contract,
        'L2_ORACLE',
        L2OutputOracleProxy.address
      )
      await assertContractVariable(
        contract,
        'GUARDIAN',
        hre.deployConfig.finalSystemOwner
      )
    },
  })
}

deployFn.tags = ['MantlePortalImpl', 'setup']

export default deployFn
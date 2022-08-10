/* Imports: External */
import { DeployFunction } from 'hardhat-deploy/dist/types'
import { ethers } from 'hardhat'

const deployFn: DeployFunction = async (hre) => {
  const { deployer } = await hre.getNamedAccounts()

  await hre.deployments.deploy('BitnetworkMintableERC721Factory', {
    from: deployer,
    args: [ethers.constants.AddressZero, 0],
    log: true,
  })
}

deployFn.tags = ['BitnetworkMintableERC721FactoryImplementation']
deployFn.dependencies = ['BitnetworkMintableERC721FactoryProxy']

export default deployFn

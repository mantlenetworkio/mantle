/* Imports: External */
import { DeployFunction } from 'hardhat-deploy/dist/types'

const deployFn: DeployFunction = async (hre) => {
    // deploy with proxy contract TransparentUpgradeableProxy.sol
    const TssGroupManagerFactory = await hre.ethers.getContractFactory('TssGroupManager')
    const tssGroupManager = await hre.upgrades.deployProxy(TssGroupManagerFactory)
    await tssGroupManager.deployed()
    console.log("TssGroupManager address : ", tssGroupManager.address)

    const TssStakingSlashingFactory = await hre.ethers.getContractFactory('TssStakingSlashing')
    // todo : set bit token address
    const tssStakingSlashing = await hre.upgrades.deployProxy(
        TssStakingSlashingFactory,
        [
            "0x0000000000000000000000000000000000000001",
            tssGroupManager.address
        ]
    )
    await tssStakingSlashing.deployed()
    console.log("tssStakingSlashing address : ", tssStakingSlashing.address)

    const owner = hre.deployConfig.bvmAddressManagerOwner
    await hre.upgrades.admin.transferProxyAdminOwnership(owner);
}

// This is kept during an upgrade. So no upgrade tag.
deployFn.tags = ['TssContracts']

export default deployFn

/* Imports: External */
import { DeployFunction } from 'hardhat-deploy/dist/types'
import { hexStringEquals, awaitCondition } from '@bitdaoio/core-utils'

/* Imports: Internal */
import {
    getContractFromArtifact,
} from '../src/deploy-utils'
import { names } from '../src/address-names'

const deployFn: DeployFunction = async (hre) => {
    const Lib_AddressManager = await getContractFromArtifact(
        hre,
        names.unmanaged.Lib_AddressManager
    )

    // deploy with proxy contract TransparentUpgradeableProxy.sol
    const TssGroupManagerFactory = await hre.ethers.getContractFactory('TssGroupManager')
    const tssGroupManager = await hre.upgrades.deployProxy(TssGroupManagerFactory)
    await tssGroupManager.deployed()

    const owner = hre.deployConfig.bvmAddressManagerOwner
    await tssGroupManager.transferOwnership(owner)
    await tssGroupManager.changeAdmin(owner)

    await Lib_AddressManager.connect(owner).setAddress(
        'TssGroupManager',
        tssGroupManager.address
    )

    // await deployAndVerifyAndThen({
    //     hre,
    //     name: names.managed.contracts.TssGroupManager,
    //     contract: 'TssGroupManager',
    //     args: [],
    //     postDeployAction: async (contract) => {
    //         // Same thing as above, we want to transfer ownership of this contract to the owner of the
    //         // AddressManager. Not technically necessary but seems like the right thing to do.
    //         console.log(
    //             `Transferring ownership of TssGroupManager (implementation)...`
    //         )
    // const owner = hre.deployConfig.bvmAddressManagerOwner
    // await contract.transferOwnership(owner)

    //         console.log(`Checking that contract owner was correctly set...`)
    //         await awaitCondition(
    //             async () => {
    //                 return hexStringEquals(await contract.owner(), owner)
    //             },
    //             5000,
    //             100
    //         )
    // await Lib_AddressManager.connect(owner).setAddress(
    //     'TssGroupManager',
    //     contract.address
    // )
    //     },
    // })
}

// This is kept during an upgrade. So no upgrade tag.
deployFn.tags = ['TssGroupManager']

export default deployFn

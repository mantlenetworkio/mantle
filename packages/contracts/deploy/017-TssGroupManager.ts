/* Imports: External */
import { DeployFunction } from 'hardhat-deploy/dist/types'

/* Imports: Internal */
import {
    deployAndVerifyAndThen,
    getContractFromArtifact,
} from '../src/deploy-utils'
import { names } from '../src/address-names'

const deployFn: DeployFunction = async (hre) => {
    const { deploy } = hre.deployments

    const Lib_AddressManager = await getContractFromArtifact(
        hre,
        names.unmanaged.Lib_AddressManager
    )

    let tssGroup = await deploy({
        hre,
        name: 'TssGroupManager',
        contract: 'TssGroupManager',
        args: [Lib_AddressManager.address],

    })

    await Lib_AddressManager.setAddress(
        'TssGroupManager',
        tssGroup.address,
    )
}

// This is kept during an upgrade. So no upgrade tag.
deployFn.tags = ['TssGroupManager']

export default deployFn

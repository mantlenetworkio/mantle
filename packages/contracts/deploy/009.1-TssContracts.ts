/* Imports: External */
import { DeployFunction } from 'hardhat-deploy/dist/types'
import { names } from '../src/address-names'
import { getContractDefinition } from '../src/contract-defs'
import { deployAndVerifyAndThen, getContractFromArtifact } from '../src/deploy-utils'
import { hexStringEquals, awaitCondition } from '@bitdaoio/core-utils'


const deployFn: DeployFunction = async (hre) => {
    const { deployer } = await hre.getNamedAccounts()
    const owner = hre.deployConfig.bvmAddressManagerOwner

    // deploy impl
    await deployAndVerifyAndThen({
        hre,
        name: 'TssGroupManager_impl',
        contract: 'TssGroupManager',
        args: [],
    })
    console.log("deploy tss group manager success")

    await deployAndVerifyAndThen({
        hre,
        name: 'TssStakingSlashing_impl',
        contract: 'TssStakingSlashing',
        args: [],
    })
    console.log("deploy tss staking slashing success")

    // deploy proxy
    var Impl_TSS_GroupManager = await getContractFromArtifact(
        hre,
        'TssGroupManager_impl',
        {
            iface: 'TssGroupManager',
            signerOrProvider: deployer,
        }
    )
    var Impl__TssStakingSlashing = await getContractFromArtifact(
        hre,
        'TssStakingSlashing_impl',
        {
            iface: 'TssStakingSlashing',
            signerOrProvider: deployer,
        }
    )
    var args: unknown[]
    args = []
    var callData = Impl_TSS_GroupManager.interface.encodeFunctionData("initialize", args)
    await deployAndVerifyAndThen({
        hre,
        name: names.managed.contracts.Proxy__TSS_GroupManager,
        contract: 'TransparentUpgradeableProxy',
        iface: 'TssGroupManager',
        args: [Impl_TSS_GroupManager.address, owner, callData],
        postDeployAction: async (contract) => {
            await contract.transferOwnership(owner)
            console.log(`Checking that contract owner was correctly set...`)
            await awaitCondition(
                async () => {
                    return hexStringEquals(await contract.owner(), owner)
                },
                5000,
                100
            )
        }
    })
    console.log("deploy tss group manager proxy success")

    args = [Impl_TSS_GroupManager.address, Impl_TSS_GroupManager.address]
    callData = Impl__TssStakingSlashing.interface.encodeFunctionData("initialize", args)
    await deployAndVerifyAndThen({
        hre,
        name: names.managed.contracts.Proxy__TSS_StakingSlashing,
        contract: 'TransparentUpgradeableProxy',
        iface: 'TssStakingSlashing',
        args: [Impl__TssStakingSlashing.address, owner, callData],
        postDeployAction: async (contract) => {
            console.log(`Checking that contract was correctly initialized...`)
            console.log("bit token : ", await contract.BitToken())
            await awaitCondition(
                async () => {
                    return hexStringEquals(
                        await contract.BitToken(),
                        Impl_TSS_GroupManager.address
                    )
                },
                5000,
                100
            )
            await contract.transferOwnership(owner)

            console.log(`Checking that contract owner was correctly set...`)
            await awaitCondition(
                async () => {
                    return hexStringEquals(await contract.owner(), owner)
                },
                5000,
                100
            )
        }
    })
    console.log("deploy tss staking slashing proxy success")
}

// This is kept during an upgrade. So no upgrade tag.
deployFn.tags = ['TssContracts', 'upgrade']

export default deployFn

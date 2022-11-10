/* Imports: External */
import { DeployFunction } from 'hardhat-deploy/dist/types'

/* Imports: Internal */
import { deployAndVerifyAndThen } from '../src/deploy-utils'
import { names } from '../src/address-names'

const deployFn: DeployFunction = async (hre) => {
  const { deployer } = await hre.getNamedAccounts()

  await deployAndVerifyAndThen({
    hre,
    name: names.managed.contracts.L1ERC721BridgeProxy,
    contract: 'MantleProxy',
    iface: 'L1ERC721Bridge',
    args: [deployer],
  })
}

// This is kept during an upgrade. So no upgrade tag.
deployFn.tags = ['L1ERC721BridgeProxy']

export default deployFn

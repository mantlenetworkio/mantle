import { DeployFunction } from 'hardhat-deploy/dist/types'
import { HardhatRuntimeEnvironment } from 'hardhat/types'
import { utils } from 'ethers'

import { predeploys } from '../src/predeploys'

// Handle the deployment
const getL1CrossDomainMessengerProxyDeployment = async (
  hre: HardhatRuntimeEnvironment
) => {
  return hre.deployments.get('Proxy__BVM_L1CrossDomainMessenger')
}

const validateERC721Bridge = async (hre, address: string, expected) => {
  const L1ERC721Bridge = await hre.ethers.getContractAt('L1ERC721Bridge', address)

  const messenger = await L1ERC721Bridge.MESSENGER()
  const otherBridge = await L1ERC721Bridge.OTHER_BRIDGE()

  console.log(`messenger address is ${utils.getAddress(messenger)}`)
  console.log(`expected address is ${utils.getAddress(expected.messenger)}`)
  if (utils.getAddress(messenger) !== utils.getAddress(expected.messenger)) {
    throw new Error(`messenger mismatch`)
  }

  if (
    utils.getAddress(otherBridge) !== utils.getAddress(expected.otherBridge)
  ) {
    throw new Error(`otherBridge mismatch`)
  }
}


const deployFn: DeployFunction = async (hre) => {
  const { deployer } = await hre.getNamedAccounts()
  const { deploy } = hre.deployments
  const { getAddress } = hre.ethers.utils

  console.log(`Deploying L1ERC721Bridge to ${hre.network.name}`)
  console.log(`Using deployer ${deployer}`)

  const Deployment__L1ERC721BridgeProxy = await hre.deployments.get(
    'L1ERC721BridgeProxy'
  )

  const L1ERC721BridgeProxy = await hre.ethers.getContractAt(
    'MantleProxy',
    Deployment__L1ERC721BridgeProxy.address
  )

  const admin = await L1ERC721BridgeProxy.callStatic.admin()
  if (getAddress(admin) !== getAddress(deployer)) {
    throw new Error('deployer is not proxy admin')
  }

  // Get the address of the currently deployed L1CrossDomainMessenger.
  // This should be the address of the proxy
  const Deployment__L1CrossDomainMessengerProxy =
    await getL1CrossDomainMessengerProxyDeployment(hre)

  const L1CrossDomainMessengerProxyAddress =
    Deployment__L1CrossDomainMessengerProxy.address

  // Deploy the L1ERC721Bridge. The arguments are
  // - messenger
  // - otherBridge
  // Since this is the L1ERC721Bridge, the otherBridge is the
  // predeploy address
  await deploy('L1ERC721Bridge', {
    from: deployer,
    args: [L1CrossDomainMessengerProxyAddress, predeploys.L2ERC721Bridge],
    log: true,
    waitConfirmations: 1,
  })

  const Deployment__L1ERC721Bridge = await hre.deployments.get('L1ERC721Bridge')
  console.log(
    `L1ERC721Bridge deployed to ${Deployment__L1ERC721Bridge.address}`
  )

  await validateERC721Bridge(hre, Deployment__L1ERC721Bridge.address, {
    messenger: L1CrossDomainMessengerProxyAddress,
    otherBridge: predeploys.L2ERC721Bridge,
  })

  {
    // Upgrade the Proxy to the newly deployed implementation
    const tx = await L1ERC721BridgeProxy.upgradeTo(
      Deployment__L1ERC721Bridge.address
    )
    const receipt = await tx.wait()
    console.log(`L1ERC721BridgeProxy upgraded: ${receipt.transactionHash}`)
  }
}

deployFn.tags = ['L1ERC721BridgeImpl', 'fresh', 'migration']

export default deployFn

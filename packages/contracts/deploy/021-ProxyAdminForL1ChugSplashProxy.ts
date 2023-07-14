import assert from 'assert'

import { DeployFunction } from 'hardhat-deploy/dist/types'

import {
  assertContractVariable,
  getContractsFromArtifacts,
  deploy, deployAndVerifyAndThen, getContractFromArtifact,
} from '../src/deploy-utils'
import {names} from "../src/address-names";

const deployFn: DeployFunction = async (hre) => {
  const { deployer } = await hre.getNamedAccounts()
  const [addressManager] = await getContractsFromArtifacts(hre, [
    {
      name: 'Lib_AddressManager',
      signerOrProvider: deployer,
    },
  ])

  await deployAndVerifyAndThen({
    hre,
    name: 'ProxyAdmin',
    contract: 'ProxyAdmin',
    args: [deployer],
    postDeployAction: async (contract) => {
      // Owner is temporarily set to the deployer.
      await assertContractVariable(contract, 'owner', deployer)
    },
  })

  const ProxyAdmin = await getContractFromArtifact(
    hre,
    'ProxyAdmin',
    {
      iface: 'ProxyAdmin',
      signerOrProvider: deployer,
    }
  )

  let addressManagerOnProxy = await ProxyAdmin.addressManager()
  if (addressManagerOnProxy !== addressManager.address) {
    // Set the address manager on the proxy admin
    console.log(
      `ProxyAdmin(${ProxyAdmin.address}).setAddressManager(${addressManager.address})`
    )
    const tx = await ProxyAdmin.setAddressManager(addressManager.address)
    await tx.wait()
  }

  // Validate the address manager was set correctly.
  addressManagerOnProxy = await ProxyAdmin.addressManager()
  assert(
    addressManagerOnProxy === addressManager.address,
    'AddressManager not set on ProxyAdmin'
  )
}

deployFn.tags = ['ProxyAdmin', 'setup', 'l1']

export default deployFn

import { ethers } from 'ethers'
import { DeployFunction } from 'hardhat-deploy/dist/types'

import { predeploys } from '../src/predeploys'
import { names } from '../src/address-names'

import {
  assertContractVariable,
  deployAndVerifyAndThen,
  getContractFromArtifact,
} from '../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const L1CrossDomainMessenger = await getContractFromArtifact(
    hre,
    names.managed.contracts.Proxy__BVM_L1CrossDomainMessenger
  )

  await deployAndVerifyAndThen({
    hre,
    name: names.managed.contracts.L1ERC721Bridge,
    args: [L1CrossDomainMessenger.address, predeploys.L2ERC721Bridge],
    postDeployAction: async (contract) => {
      await assertContractVariable(
        contract,
        'MESSENGER',
        L1CrossDomainMessenger.address
      )
    },
  })
}

deployFn.tags = ['L1ERC721BridgeImpl', 'fresh', 'migration']

export default deployFn

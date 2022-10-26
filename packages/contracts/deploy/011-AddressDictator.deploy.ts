/* Imports: External */
import { DeployFunction } from 'hardhat-deploy/dist/types'
import { hexStringEquals } from '@mantlenetworkio/core-utils'

/* Imports: Internal */
import {
  deployAndVerifyAndThen,
  getContractFromArtifact,
} from '../src/deploy-utils'
import { names } from '../src/address-names'
import { predeploys } from '../src/predeploys'

const deployFn: DeployFunction = async (hre) => {
  const Lib_AddressManager = await getContractFromArtifact(
    hre,
    names.unmanaged.Lib_AddressManager
  )

  let namesAndAddresses: {
    name: string
    address: string
  }[] = await Promise.all(
    Object.values(names.managed.contracts).map(async (name) => {
      return {
        name,
        address: (await getContractFromArtifact(hre, name)).address,
      }
    })
  )

  // Add non-deployed addresses to the Address Dictator arguments.
  namesAndAddresses = [
    ...namesAndAddresses,
    // L2CrossDomainMessenger is the address of the predeploy on L2. We can refactor off-chain
    // services such that we can remove the need to set this address, but for now it's easier
    // to simply keep setting the address.
    {
      name: 'L2CrossDomainMessenger',
      address: predeploys.L2CrossDomainMessenger,
    },
    // BVM_Sequencer is the address allowed to submit "Sequencer" blocks to the
    // CanonicalTransactionChain.
    {
      name: names.managed.accounts.BVM_Sequencer,
      address: hre.deployConfig.bvmSequencerAddress,
    },
    // BVM_Proposer is the address allowed to submit state roots (transaction results) to the
    // StateCommitmentChain.
    {
      name: names.managed.accounts.BVM_Proposer,
      address: hre.deployConfig.bvmProposerAddress,
    },
    // L1_BIT_ADDRESS indicate l1 bit token erc20 contract address
    {
      name: names.managed.configs.L1_BIT_ADDRESS,
      address: hre.deployConfig.l1BitAddress,
    },
  ]

  // Filter out all addresses that will not change, so that the log statement is maximally
  // verifiable and readable.
  const existingAddresses = {}
  for (const pair of namesAndAddresses) {
    existingAddresses[pair.name] = await Lib_AddressManager.getAddress(
      pair.name
    )
  }
  namesAndAddresses = namesAndAddresses.filter(({ name, address }) => {
    return !hexStringEquals(existingAddresses[name], address)
  })

  await deployAndVerifyAndThen({
    hre,
    name: names.unmanaged.AddressDictator,
    args: [
      Lib_AddressManager.address,
      hre.deployConfig.bvmAddressManagerOwner,
      namesAndAddresses.map((pair) => {
        return pair.name
      }),
      namesAndAddresses.map((pair) => {
        return pair.address
      }),
    ],
  })
}

deployFn.tags = ['upgrade', 'AddressDictator']

export default deployFn

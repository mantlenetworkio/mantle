import { DeployFunction } from 'hardhat-deploy/dist/types'
import {
  deployAndVerifyAndThen,
  getContractFromArtifact,
} from '../src/deploy-utils'
import { names } from '../src/address-names'

const deployFn: DeployFunction = async (hre) => {
  const { deployer } = await hre.getNamedAccounts()

  await deployAndVerifyAndThen({
    hre,
    name: names.managed.contracts.TssStakingSlashing,
    contract: 'TssStakingSlashing',
    args: [],
  })
  console.log('deploy new tss staking slashing success')

  const Impl__TssStakingSlashing = await getContractFromArtifact(
    hre,
    names.managed.contracts.TssStakingSlashing,
    {
      iface: 'TssStakingSlashing',
      signerOrProvider: deployer,
    }
  )
  console.log('new tss staking slashing address',Impl__TssStakingSlashing.address)
}

// This is kept during an upgrade. So no upgrade tag.
deployFn.tags = ['TssContracts', 'upgrade']

export default deployFn

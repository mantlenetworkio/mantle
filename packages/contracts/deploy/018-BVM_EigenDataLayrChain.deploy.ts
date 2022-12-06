/* Imports: External */
import { DeployFunction } from 'hardhat-deploy/dist/types'

import { names } from '../src/address-names'
import {
  deployAndVerifyAndThen,
  getContractFromArtifact,
} from '../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const { deployer } = await hre.getNamedAccounts()

  const owner = hre.deployConfig.bvmAddressManagerOwner
  const eigenSequencerAddress = hre.deployConfig.bvmSequencerAddress
  const dataManagerAddress = hre.deployConfig.dataManagerAddress

  // deploy impl
  await deployAndVerifyAndThen({
    hre,
    name: names.managed.da.BVM_EigenDataLayrChain,
    contract: 'BVM_EigenDataLayrChain',
    args: [],
  })
  console.log('deploy tss group manager success')

  // deploy proxy
  const Impl_BVM_EigenDataLayrChain = await getContractFromArtifact(
    hre,
    names.managed.da.BVM_EigenDataLayrChain,
    {
      iface: 'BVM_EigenDataLayrChain',
      signerOrProvider: deployer,
    }
  )

  const args = [eigenSequencerAddress, dataManagerAddress, 10, 100]
  const callData = Impl_BVM_EigenDataLayrChain.interface.encodeFunctionData(
    'initialize',
    args
  )

  await deployAndVerifyAndThen({
    hre,
    name: names.managed.da.Proxy__BVM_EigenDataLayrChain,
    contract: 'TransparentUpgradeableProxy',
    iface: 'BVM_EigenDataLayrChain',
    args: [Impl_BVM_EigenDataLayrChain.address, owner, callData],
  })
  console.log('deploy eigen da proxy success')
}

deployFn.tags = ['BVM_EigenDataLayrChain', 'upgrade']

export default deployFn

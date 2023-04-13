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
  // const eigenSequencerAddress = hre.deployConfig.bvmSequencerAddress
  const eigenSequencerAddress = hre.deployConfig.bvmEigenSequencerAddress
  const dataManagerAddress = hre.deployConfig.dataManagerAddress
  const reSubmitterAddress = hre.deployConfig.bvmEigenSequencerAddress
  const blockStaleMeasure = hre.deployConfig.blockStaleMeasure
  const daFraudProofPeriod = hre.deployConfig.daFraudProofPeriod
  const l2SubmittedBlockNumber = hre.deployConfig.l2SubmittedBlockNumber

  const args = [eigenSequencerAddress, dataManagerAddress, reSubmitterAddress, blockStaleMeasure, daFraudProofPeriod, l2SubmittedBlockNumber]
  await deployAndVerifyAndThen({
    hre,
    name: names.managed.da.BVM_EigenDataLayrChain,
    contract: 'BVM_EigenDataLayrChain',
    args: [],
  })
  console.log('deploy eigen datalayr chain success')

  const Impl_BVM_EigenDataLayrChain = await getContractFromArtifact(
    hre,
    names.managed.da.BVM_EigenDataLayrChain,
    {
      iface: 'BVM_EigenDataLayrChain',
      signerOrProvider: deployer,
    }
  )

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

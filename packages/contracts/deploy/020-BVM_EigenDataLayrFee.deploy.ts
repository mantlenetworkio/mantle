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

  await deployAndVerifyAndThen({
    hre,
    name: names.managed.da.BVM_EigenDataLayrFee,
    contract: 'BVM_EigenDataLayrFee',
    args: [],
  })
  console.log('deploy eigen datalayr fee success')

  const Impl_BVM_EigenDataLayrFee = await getContractFromArtifact(
    hre,
    names.managed.da.BVM_EigenDataLayrFee,
    {
      iface: 'BVM_EigenDataLayrFee',
      signerOrProvider: deployer,
    }
  )

  await deployAndVerifyAndThen({
    hre,
    name: names.managed.da.Proxy__BVM_EigenDataLayrFee,
    contract: 'TransparentUpgradeableProxy',
    iface: 'BVM_EigenDataLayrFee',
    args: [Impl_BVM_EigenDataLayrFee.address, owner],
  })
  console.log('deploy eigen da fee proxy success')
}

deployFn.tags = ['BVM_EigenDataLayrFee', 'upgrade']

export default deployFn

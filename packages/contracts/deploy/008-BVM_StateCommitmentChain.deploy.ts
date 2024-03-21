/* Imports: External */
import { DeployFunction } from 'hardhat-deploy/dist/types'

/* Imports: Internal */
import {
  deployAndVerifyAndThen,
  getContractFromArtifact,
} from '../src/deploy-utils'
import { names } from '../src/address-names'

const deployFn: DeployFunction = async (hre) => {
  const Lib_AddressManager = await getContractFromArtifact(
    hre,
    names.unmanaged.Lib_AddressManager
  )
  const Proxy__BVM_L1CrossDomainMessenger = await getContractFromArtifact(
    hre,
    names.managed.contracts.Proxy__BVM_L1CrossDomainMessenger
  )

  await deployAndVerifyAndThen({
    hre,
    name: names.managed.contracts.StateCommitmentChain,
    args: [
      Lib_AddressManager.address,
      Proxy__BVM_L1CrossDomainMessenger.address,
      hre.deployConfig.sccFaultProofWindowSeconds,
      hre.deployConfig.sccSequencerPublishWindowSeconds,
    ],
  })
}

deployFn.tags = ['StateCommitmentChain', 'upgrade']

export default deployFn

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
  // const messengerSlotKey = await ChugSplashDictator.messengerSlotKey()
  // const messengerSlotVal = await ChugSplashDictator.messengerSlotVal()
  await deployAndVerifyAndThen({
    hre,
    name: names.managed.contracts.StateCommitmentChain,
    args: [
      Lib_AddressManager.address,
      '0xc48078a734c2e22D43F54B47F7a8fB314Fa5A601',
      hre.deployConfig.sccFaultProofWindowSeconds,
      hre.deployConfig.sccSequencerPublishWindowSeconds,
    ],
  })
}

deployFn.tags = ['StateCommitmentChain', 'upgrade']

export default deployFn

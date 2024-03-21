import { DeployFunction } from 'hardhat-deploy/dist/types'
import {
  deployAndVerifyAndThen,
  getContractFromArtifact,
} from '../src/deploy-utils'
import { names } from '../src/address-names'
import { getContractFactory } from '../src'
import { ethers } from 'ethers'

const deployFn: DeployFunction = async (hre) => {
  const { deployer } = await hre.getNamedAccounts()

  const Impl__TssStakingSlashing = await getContractFromArtifact(
    hre,
    names.managed.contracts.TssStakingSlashing,
    {
      iface: 'TssStakingSlashing',
      signerOrProvider: deployer,
    }
  )
  const owner = hre.deployConfig.bvmAddressManagerOwner
  const l1MantleAddress = hre.deployConfig.proxyL1MantleAddress
  const tssManagerAddress = hre.deployConfig.tssManagerAddress

  // const from = deployer
  const provider = new ethers.providers.JsonRpcBatchProvider(
    hre.deployConfig.contractsRpcUrl
  )
  const deployerWallet = new ethers.Wallet(
    hre.deployConfig.contractsDeployerKey,
    provider
  )
  console.log('deploy privete key', hre.deployConfig.contractsDeployerKey)
  console.log('privder', provider)

  const Proxy__TSS_GroupManager = await getContractFromArtifact(
    hre,
    names.managed.contracts.Proxy__TSS_GroupManager,
    {
      iface: 'TssGroupManager',
      signerOrProvider: deployer,
    }
  )
  console.log(
    'Proxy__TSS_GroupManager Address',
    Proxy__TSS_GroupManager.address
  )

  const Proxy_TssDelegationManager = await getContractFromArtifact(
    hre,
    names.managed.delegation.tss.Proxy__TssDelegationManager,
    {
      iface: 'TssDelegationManager',
      signerOrProvider: deployer,
    }
  )
  console.log(
    'Proxy__TssDelegationManager Address',
    Proxy_TssDelegationManager.address
  )

  const Proxy_TssDelegation = await getContractFromArtifact(
    hre,
    names.managed.delegation.tss.Proxy__TssDelegation,
    {
      iface: 'TssDelegation',
      signerOrProvider: deployer,
    }
  )
  console.log('Proxy__TssDelegation Address', Proxy_TssDelegation.address)

  const Proxy__BVM_L1CrossDomainMessenger = await getContractFromArtifact(
    hre,
    names.managed.contracts.Proxy__BVM_L1CrossDomainMessenger
  )

  const Proxy__TSS_StakingSlashing = await getContractFromArtifact(
    hre,
    names.managed.contracts.Proxy__TSS_StakingSlashing,
    {
      iface: 'TssStakingSlashing',
      signerOrProvider: deployer,
    }
  )
  console.log(
    'Proxy__TSS_StakingSlashing Address',
    Proxy__TSS_StakingSlashing.address
  )


  const delegationProxy = getContractFactory(
    'TransparentUpgradeableProxy'
  ).attach(Proxy__TSS_StakingSlashing.address)


  let args: unknown[]

  args = [
    l1MantleAddress,
    Proxy__TSS_GroupManager.address,
    Proxy_TssDelegationManager.address,
    Proxy_TssDelegation.address,
    Proxy__BVM_L1CrossDomainMessenger.address,
    owner,
    tssManagerAddress,
  ]

  let callData = Impl__TssStakingSlashing.interface.encodeFunctionData(
    'initialize',
    args
  )

  await delegationProxy
    .connect(deployerWallet)
    .upgradeToAndCall(Impl__TssStakingSlashing.address, callData, {
      gasLimit: 2_000_000,
    })
  console.log('update Tss Delegation Proxy__TssDelegation success')
}

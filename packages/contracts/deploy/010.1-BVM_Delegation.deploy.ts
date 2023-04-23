/* Imports: External */
import { DeployFunction } from 'hardhat-deploy/dist/types'

// @ts-ignore
import { names } from '../src/address-names'
import {
  deployAndVerifyAndThen,
  getContractFromArtifact,
} from '../src/deploy-utils'
import {awaitCondition, hexStringEquals} from "@mantleio/core-utils";
import {deploy} from "../test/helpers";
import {address} from "hardhat/internal/core/config/config-validation";

// eslint-disable-next-line @typescript-eslint/no-var-requires
const { getCreate2Address } = require('@ethersproject/address');
const { ethers, upgrades, getContractFactory, getNamedAccounts } = require('hardhat');
const { getImplementationAddress, getProxyAddress } = require('@openzeppelin/upgrades');

const deployFn: DeployFunction = async (hre) => {
  // @ts-ignore
  const { deployer } = await hre.getNamedAccounts()

  const Lib_AddressManager = await getContractFromArtifact(
    hre,
    names.unmanaged.Lib_AddressManager
  )
  // @ts-ignore
  const owner = hre.deployConfig.bvmAddressManagerOwner
  // @ts-ignore
  const l1BitAddress = hre.deployConfig.l1BitAddress

  // @ts-ignore
  let DelegationProxyAddress = address(0)
  let DelegationManagerProxyAddress = address(0)
  let DelegationSlasherProxyAddress = address(0)

  // deploy Delegation impl
  await deployAndVerifyAndThen({
    hre,
    name: names.managed.delegation.Delegation,
    contract: 'Delegation',
    args: [DelegationManagerProxyAddress],
  })
  const Impl__Delegation = await getContractFromArtifact(
    hre,
    names.managed.delegation.Delegation,
    {
      iface: 'Delegation',
      signerOrProvider: deployer,
    }
  )
  console.log('Delegation Implementation Address', Impl__Delegation.address)
  console.log('deploy Delegation success')

  // deploy DelegationManager impl
  await deployAndVerifyAndThen({
    hre,
    name: names.managed.delegation.DelegationManager,
    contract: 'DelegationManager',
    args: [DelegationProxyAddress, DelegationSlasherProxyAddress],
  })
  const Impl__DelegationManager = await getContractFromArtifact(
    hre,
    names.managed.delegation.DelegationManager,
    {
      iface: 'DelegationManager',
      signerOrProvider: deployer,
    }
  )
  console.log(
    'DelegationManager Implementation Address',
    Impl__DelegationManager.address
  )
  console.log('deploy DelegationManager success')

  // deploy DelegationSlasher impl
  await deployAndVerifyAndThen({
    hre,
    name: names.managed.delegation.DelegationSlasher,
    contract: 'DelegationSlasher',
    args: [DelegationManagerProxyAddress, DelegationProxyAddress],
  })
  const Impl__DelegationSlasher = await getContractFromArtifact(
    hre,
    names.managed.delegation.DelegationSlasher,
    {
      iface: 'DelegationSlasher',
      signerOrProvider: deployer,
    }
  )
  console.log(
    'DelegationSlasher Implementation Address',
    Impl__DelegationSlasher.address
  )
  console.log('deploy DelegationSlasher rollup success')

  // deploy Delegation proxy
  let callData = Impl__Delegation.interface.encodeFunctionData('initialize', [
    owner,
  ])
  await deployAndVerifyAndThen({
    hre,
    name: names.managed.delegation.Proxy__Delegation,
    contract: 'TransparentUpgradeableProxy',
    iface: 'Delegation',
    args: [Impl__Delegation.address, owner, callData],
  })
  console.log('deploy Delegation Proxy success')
  const Proxy__Delegation = await getContractFromArtifact(
    hre,
    names.managed.delegation.Proxy__Delegation,
    {
      iface: 'Delegation',
      signerOrProvider: deployer,
    }
  )
  console.log('Proxy__Delegation Address', Proxy__Delegation.address)
  console.log('deploy delegation Proxy__Delegation success')

  callData = Impl__DelegationManager.interface.encodeFunctionData(
    'initialize',
    [owner]
  )
  await deployAndVerifyAndThen({
    hre,
    name: names.managed.delegation.Proxy__DelegationManager,
    contract: 'TransparentUpgradeableProxy',
    iface: 'DelegationManager',
    args: [Impl__DelegationManager.address, owner, callData],
  })

  const Proxy__DelegationSlasher = await getContractFromArtifact(
    hre,
    names.managed.delegation.Proxy__DelegationSlasher,
    {
      iface: 'DelegationSlasher',
      signerOrProvider: deployer,
    }
  )
  console.log(
    'Proxy__DelegationSlasher Address',
    Proxy__DelegationSlasher.address
  )
  console.log('deploy DelegationSlasher Proxy success')
}

// This is kept during an upgrade. So no upgrade tag.
deployFn.tags = ['FraudProof', 'upgrade']

export default deployFn

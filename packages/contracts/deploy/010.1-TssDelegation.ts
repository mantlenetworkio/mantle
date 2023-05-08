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
  // @ts-ignore
  const owner = hre.deployConfig.bvmAddressManagerOwner
  // @ts-ignore
  let DelegationProxyAddress = "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
  let DelegationManagerProxyAddress = "0xD6f15EAC1Cb3B4131Ab4899a52E711e19DEeA73f"
  let DelegationSlasherProxyAddress = "0x82d2984a3A2137634300c0D6DDEf7D3C851EcEa8"

  // deploy Delegation impl
  await deployAndVerifyAndThen({
    hre,
    name: names.managed.delegation.tss.TssDelegation,
    contract: 'TssDelegation',
    args: [DelegationManagerProxyAddress],
  })
  const Impl__TssDelegation = await getContractFromArtifact(
    hre,
    names.managed.delegation.tss.TssDelegation,
    {
      iface: 'TssDelegation',
      signerOrProvider: deployer,
    }
  )
  console.log('TssDelegation Implementation Address', Impl__TssDelegation.address)
  console.log('deploy Tss Delegation success')

  // deploy DelegationManager impl
  await deployAndVerifyAndThen({
    hre,
    name: names.managed.delegation.tss.TssDelegationManager,
    contract: 'TssDelegationManager',
    args: [DelegationProxyAddress, DelegationSlasherProxyAddress],
  })
  const Impl__TssDelegationManager = await getContractFromArtifact(
    hre,
    names.managed.delegation.tss.TssDelegationManager,
    {
      iface: 'TssDelegationManager',
      signerOrProvider: deployer,
    }
  )
  console.log('TssDelegationManager Implementation Address', Impl__TssDelegationManager.address)
  console.log('deploy Tss DelegationManager success')

  // deploy DelegationSlasher impl
  await deployAndVerifyAndThen({
    hre,
    name: names.managed.delegation.tss.TssDelegationSlasher,
    contract: 'TssDelegationSlasher',
    args: [DelegationManagerProxyAddress, DelegationProxyAddress],
  })
  const Impl__TssDelegationSlasher = await getContractFromArtifact(
    hre,
    names.managed.delegation.tss.TssDelegationSlasher,
    {
      iface: 'TssDelegationSlasher',
      signerOrProvider: deployer,
    }
  )
  console.log('Tss DelegationSlasher Implementation Address', Impl__TssDelegationSlasher.address)
  console.log('deploy Tss DelegationSlasher rollup success')

  // deploy Delegation proxy
  let callData = Impl__TssDelegation.interface.encodeFunctionData('initialize', [
    owner,
  ])
  await deployAndVerifyAndThen({
    hre,
    name: names.managed.delegation.tss.Proxy__TssDelegation,
    contract: 'TransparentUpgradeableProxy',
    iface: 'TssDelegation',
    args: [Impl__TssDelegation.address, owner, callData],
  })
  console.log('deploy Tss Delegation Proxy success')
  const Proxy__TssDelegation = await getContractFromArtifact(
    hre,
    names.managed.delegation.tss.Proxy__TssDelegation,
    {
      iface: 'TssDelegation',
      signerOrProvider: deployer,
    }
  )
  console.log('Proxy__TssDelegation Address', Proxy__TssDelegation.address)
  console.log('deploy Tss Delegation Proxy__TssDelegation success')

  callData = Impl__TssDelegationSlasher.interface.encodeFunctionData(
    'initialize',
    [owner]
  )
  await deployAndVerifyAndThen({
    hre,
    name: names.managed.delegation.tss.Proxy__TssDelegationSlasher,
    contract: 'TransparentUpgradeableProxy',
    iface: 'TssDelegationSlasher',
    args: [Impl__TssDelegationSlasher.address, owner, callData],
  })

  const Proxy__TssDelegationSlasher = await getContractFromArtifact(
    hre,
    names.managed.delegation.tss.Proxy__TssDelegationSlasher,
    {
      iface: 'TssDelegationSlasher',
      signerOrProvider: deployer,
    }
  )
  console.log('Proxy__TssDelegationSlasher Address', Proxy__TssDelegationSlasher.address)
  console.log('deploy Tss DelegationSlasher Proxy success')

  callData = Impl__TssDelegationManager.interface.encodeFunctionData(
    'initialize',
    [owner]
  )
  await deployAndVerifyAndThen({
    hre,
    name: names.managed.delegation.tss.Proxy__TssDelegationManager,
    contract: 'TransparentUpgradeableProxy',
    iface: 'TssDelegationManager',
    args: [Impl__TssDelegationManager.address, owner, callData],
  })

  const Proxy__TssDelegationManager = await getContractFromArtifact(
    hre,
    names.managed.delegation.tss.Proxy__TssDelegationManager,
    {
      iface: 'TssDelegationManager',
      signerOrProvider: deployer,
    }
  )
  console.log('Proxy__TssDelegationManager Address', Proxy__TssDelegationManager.address)
  console.log('deploy Tss DelegationManager Proxy success')
}

// This is kept during an upgrade. So no upgrade tag.
deployFn.tags = ['FraudProofDelegation', 'upgrade']

export default deployFn

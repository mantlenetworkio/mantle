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
  let DelegationProxyAddress = "0x88e7b36Ae38D502e4Df97e7240560B77F6B4a58E"
  let DelegationManagerProxyAddress = "0x90c9E4f7F11C93BCec2e1F272b2063640d3bb264"
  let DelegationSlasherProxyAddress = "0x8eC353337d109C0f8eef2200021D28C542CAe8E6"

  // deploy Delegation impl
  await deployAndVerifyAndThen({
    hre,
    name: names.managed.delegation.fraud_proof.FraudProofDelegation,
    contract: 'FraudProofDelegation',
    args: [DelegationManagerProxyAddress],
  })
  const Impl__FraudProofDelegation = await getContractFromArtifact(
    hre,
    names.managed.delegation.fraud_proof.FraudProofDelegation,
    {
      iface: 'FraudProofDelegation',
      signerOrProvider: deployer,
    }
  )
  console.log('FraudProofDelegation Implementation Address', Impl__FraudProofDelegation.address)
  console.log('deploy FraudProof Delegation success')

  // deploy DelegationManager impl
  await deployAndVerifyAndThen({
    hre,
    name: names.managed.delegation.fraud_proof.FraudProofDelegationManager,
    contract: 'FraudProofDelegationManager',
    args: [DelegationProxyAddress, DelegationSlasherProxyAddress],
  })
  const Impl__FraudProofDelegationManager = await getContractFromArtifact(
    hre,
    names.managed.delegation.fraud_proof.FraudProofDelegationManager,
    {
      iface: 'FraudProofDelegationManager',
      signerOrProvider: deployer,
    }
  )
  console.log('FraudProofDelegationManager Implementation Address', Impl__FraudProofDelegationManager.address)
  console.log('deploy FraudProof DelegationManager success')

  // deploy DelegationSlasher impl
  await deployAndVerifyAndThen({
    hre,
    name: names.managed.delegation.fraud_proof.FraudProofDelegationSlasher,
    contract: 'FraudProofDelegationSlasher',
    args: [DelegationManagerProxyAddress, DelegationProxyAddress],
  })
  const Impl__FraudProofDelegationSlasher = await getContractFromArtifact(
    hre,
    names.managed.delegation.fraud_proof.FraudProofDelegationSlasher,
    {
      iface: 'FraudProofDelegationSlasher',
      signerOrProvider: deployer,
    }
  )
  console.log('FraudProof DelegationSlasher Implementation Address', Impl__FraudProofDelegationSlasher.address)
  console.log('deploy FraudProof DelegationSlasher rollup success')

  // deploy Delegation proxy
  let callData = Impl__FraudProofDelegation.interface.encodeFunctionData('initialize', [
    deployer,
  ])
  await deployAndVerifyAndThen({
    hre,
    name: names.managed.delegation.fraud_proof.Proxy__FraudProofDelegation,
    contract: 'TransparentUpgradeableProxy',
    iface: 'FraudProofDelegation',
    args: [Impl__FraudProofDelegation.address, owner, callData],
  })
  console.log('deploy FraudProof Delegation Proxy success')
  const Proxy__FraudProofDelegation = await getContractFromArtifact(
    hre,
    names.managed.delegation.fraud_proof.Proxy__FraudProofDelegation,
    {
      iface: 'FraudProofDelegation',
      signerOrProvider: deployer,
    }
  )
  console.log('Proxy__FraudProofDelegation Address', Proxy__FraudProofDelegation.address)
  console.log('deploy FraudProof Delegation Proxy__FraudProofDelegation success')

  callData = Impl__FraudProofDelegationSlasher.interface.encodeFunctionData(
    'initialize',
    [deployer]
  )
  await deployAndVerifyAndThen({
    hre,
    name: names.managed.delegation.fraud_proof.Proxy__FraudProofDelegationSlasher,
    contract: 'TransparentUpgradeableProxy',
    iface: 'FraudProofDelegationSlasher',
    args: [Impl__FraudProofDelegationSlasher.address, owner, callData],
  })

  const Proxy__FraudProofDelegationSlasher = await getContractFromArtifact(
    hre,
    names.managed.delegation.fraud_proof.Proxy__FraudProofDelegationSlasher,
    {
      iface: 'FraudProofDelegationSlasher',
      signerOrProvider: deployer,
    }
  )
  console.log('Proxy__FraudProofDelegationSlasher Address', Proxy__FraudProofDelegationSlasher.address)
  console.log('deploy FraudProof DelegationSlasher Proxy success')

  callData = Impl__FraudProofDelegationManager.interface.encodeFunctionData(
    'initialize',
    [deployer]
  )
  await deployAndVerifyAndThen({
    hre,
    name: names.managed.delegation.fraud_proof.Proxy__FraudProofDelegationManager,
    contract: 'TransparentUpgradeableProxy',
    iface: 'FraudProofDelegationManager',
    args: [Impl__FraudProofDelegationManager.address, owner, callData],
  })

  const Proxy__FraudProofDelegationManager = await getContractFromArtifact(
    hre,
    names.managed.delegation.fraud_proof.Proxy__FraudProofDelegationManager,
    {
      iface: 'FraudProofDelegationManager',
      signerOrProvider: deployer,
    }
  )
  console.log('Proxy__FraudProofDelegationManager Address', Proxy__FraudProofDelegationManager.address)
  console.log('deploy DelegationManager Proxy success')
}

// This is kept during an upgrade. So no upgrade tag.
deployFn.tags = ['FraudProofDelegation', 'upgrade']

export default deployFn

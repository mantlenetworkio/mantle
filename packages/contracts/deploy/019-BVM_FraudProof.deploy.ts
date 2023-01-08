/* Imports: External */
import { DeployFunction } from 'hardhat-deploy/dist/types'

// @ts-ignore
import { names } from '../src/address-names'
import {
  deployAndVerifyAndThen,
  getContractFromArtifact,
} from '../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  // @ts-ignore
  const { deployer } = await hre.getNamedAccounts()

  // @ts-ignore
  const owner = hre.deployConfig.bvmAddressManagerOwner
  // @ts-ignore
  const l1BitAddress = hre.deployConfig.l1BitAddress
  // deploy verifier impl
  await deployAndVerifyAndThen({
    hre,
    name: names.managed.fraud_proof.Verifier,
    contract: 'Verifier',
    args: [],
  })
  const Impl__Verifier = await getContractFromArtifact(
    hre,
    names.managed.fraud_proof.Verifier,
    {
      iface: 'Verifier',
      signerOrProvider: deployer,
    }
  )
  console.log('Verifier Implementation Address', Impl__Verifier.address)
  console.log('deploy fraud proof verifier success')

  // deploy rollup impl
  await deployAndVerifyAndThen({
    hre,
    name: names.managed.fraud_proof.Rollup,
    contract: 'Rollup',
    args: [owner],
  })
  const Impl__Rollup = await getContractFromArtifact(
    hre,
    names.managed.fraud_proof.Rollup,
    {
      iface: 'Rollup',
      signerOrProvider: deployer,
    }
  )
  console.log('Rollup Implementation Address', Impl__Rollup.address)
  console.log('deploy fraud proof assertion rollup success')

  // deploy verifier proxy
  let callData = Impl__Verifier.interface.encodeFunctionData('initialize', [])
  await deployAndVerifyAndThen({
    hre,
    name: names.managed.fraud_proof.Proxy__Verifier,
    contract: 'TransparentUpgradeableProxy',
    iface: 'Verifier',
    args: [Impl__Verifier.address, owner, callData],
  })
  console.log('deploy fraud proof verifier proxy success')

  // deploy rollup proxy
  const rollupArgs = [
    owner, // address _owner
    Impl__Verifier.address, // address _verifier,
    l1BitAddress, // address _stakeToken,
    5, // uint256 _confirmationPeriod,
    0, // uint256 _challengePeriod,
    0, // uint256 _minimumAssertionPeriod,
    1000000000000, // uint256 _maxGasPerAssertion,
    0, // uint256 _baseStakeAmount
    "0x744c19d2e8593c97867b3b6a3588f51cd9dbc5010a395cf199be4bbb353848b8", // bytes32 _initialVMhash //TODO-FIXME
  ]
  callData = Impl__Rollup.interface.encodeFunctionData('initialize', rollupArgs)
  await deployAndVerifyAndThen({
    hre,
    name: names.managed.fraud_proof.Proxy__Rollup,
    contract: 'TransparentUpgradeableProxy',
    iface: 'Rollup',
    args: [Impl__Rollup.address, owner, callData],
  })
  console.log('deploy fraud proof rollup proxy success')
}

// This is kept during an upgrade. So no upgrade tag.
deployFn.tags = ['FraudProof', 'upgrade']

export default deployFn

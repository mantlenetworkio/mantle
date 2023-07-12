import {Contract, ethers} from 'ethers'
import {Provider} from '@ethersproject/abstract-provider'
import {Signer} from '@ethersproject/abstract-signer'
import {awaitCondition, getChainId, sleep} from '@mantleio/core-utils'
import {HttpNetworkConfig} from 'hardhat/types'
import assert from "assert";
import { HardhatRuntimeEnvironment } from 'hardhat/types'
import { Deployment, DeployResult } from 'hardhat-deploy/dist/types'


/**
 * Wrapper around hardhat-deploy with some extra features.
 *
 * @param opts Options for the deployment.
 * @param opts.hre HardhatRuntimeEnvironment.
 * @param opts.contract Name of the contract to deploy.
 * @param opts.name Name to use for the deployment file.
 * @param opts.iface Interface to use for the returned contract.
 * @param opts.args Arguments to pass to the contract constructor.
 * @param opts.postDeployAction Action to perform after the contract is deployed.
 * @returns Deployed contract object.
 */
export const deploy = async ({
                               hre,
                               name,
                               iface,
                               args,
                               contract,
                               postDeployAction,
                             }: {
  hre: HardhatRuntimeEnvironment
  name: string
  args: any[]
  contract?: string
  iface?: string
  postDeployAction?: (contract: Contract) => Promise<void>
}): Promise<Contract> => {
  const { deployer } = await hre.getNamedAccounts()

  // Hardhat deploy will usually do this check for us, but currently doesn't also consider
  // external deployments when doing this check. By doing the check ourselves, we also get to
  // consider external deployments. If we already have the deployment, return early.
  let result: Deployment | DeployResult = await hre.deployments.getOrNull(name)

  // Wrap in a try/catch in case there is not a deployConfig for the current network.
  let numDeployConfirmations: number
  try {
    numDeployConfirmations = hre.deployConfig.numDeployConfirmations
  } catch (e) {
    numDeployConfirmations = 1
  }

  if (result) {
    console.log(`skipping ${name}, using existing at ${result.address}`)
  } else {
    result = await hre.deployments.deploy(name, {
      contract,
      from: deployer,
      args,
      log: true,
      waitConfirmations: numDeployConfirmations,
    })
    console.log(`Deployed ${name} at ${result.address}`)
    // Only wait for the transaction if it was recently deployed in case the
    // result was deployed a long time ago and was pruned from the backend.
    await hre.ethers.provider.waitForTransaction(result.transactionHash)
  }

  // Check to make sure there is code
  const code = await hre.ethers.provider.getCode(result.address)
  if (code === '0x') {
    throw new Error(`no code for ${result.address}`)
  }

  // Create the contract object to return.
  const created = asAdvancedContract({
    confirmations: numDeployConfirmations,
    contract: new Contract(
      result.address,
      iface !== undefined
        ? (await hre.ethers.getContractFactory(iface)).interface
        : result.abi,
      hre.ethers.provider.getSigner(deployer)
    ),
  })

  // Run post-deploy actions if necessary.
  if ((result as DeployResult).newlyDeployed) {
    if (postDeployAction) {
      await postDeployAction(created)
    }
  }

  return created
}

/**
 * @param  {Any} hre Hardhat runtime environment
 * @param  {String} name Contract name from the names object
 * @param  {Any[]} args Constructor arguments
 * @param  {String} contract Name of the solidity contract
 * @param  {String} iface Alternative interface for calling the contract
 * @param  {Function} postDeployAction Called after deployment
 */

export const deployAndVerifyAndThen = async ({
                                               hre,
                                               name,
                                               args,
                                               contract,
                                               iface,
                                               postDeployAction,
                                             }: {
  hre: any
  name: string
  args: any[]
  contract?: string
  iface?: string
  postDeployAction?: (contract: Contract) => Promise<void>
}) => {
  const {deploy: deployments} = hre.deployments
  const {deployer} = await hre.getNamedAccounts()

  const result = await deployments(name, {
    contract,
    from: deployer,
    args,
    log: true,
    waitConfirmations: hre.deployConfig.numDeployConfirmations,
  })

  await hre.ethers.provider.waitForTransaction(result.transactionHash)

  if (result.newlyDeployed) {
    if (!(await isHardhatNode(hre))) {
      // Verification sometimes fails, even when the contract is correctly deployed and eventually
      // verified. Possibly due to a race condition. We don't want to halt the whole deployment
      // process just because that happens.
      try {
        console.log('Verifying on Etherscan...')
        await hre.run('verify:verify', {
          address: result.address,
          constructorArguments: args,
        })
        console.log('Successfully verified on Etherscan')
      } catch (error) {
        console.log('Error when verifying bytecode on Etherscan:')
        console.log(error)
      }

      try {
        console.log('Verifying on Sourcify...')
        await hre.run('sourcify')
        console.log('Successfully verified on Sourcify')
      } catch (error) {
        console.log('Error when verifying bytecode on Sourcify:')
        console.log(error)
      }
    }
    if (postDeployAction) {
      const signer = hre.ethers.provider.getSigner(deployer)
      let abi = result.abi
      if (iface !== undefined) {
        const factory = await hre.ethers.getContractFactory(iface)
        abi = factory.interface
      }
      await postDeployAction(
        getAdvancedContract({
          hre,
          contract: new Contract(result.address, abi, signer),
        })
      )
    }
  }
}

// Returns a version of the contract object which modifies all of the input contract's methods to:
// 1. Waits for a confirmed receipt with more than deployConfig.numDeployConfirmations confirmations.
// 2. Include simple resubmission logic, ONLY for Kovan, which appears to drop transactions.
export const getAdvancedContract = (opts: {
  hre: any
  contract: Contract
}): Contract => {
  // Temporarily override Object.defineProperty to bypass ether's object protection.
  const def = Object.defineProperty
  Object.defineProperty = (obj, propName, prop) => {
    prop.writable = true
    return def(obj, propName, prop)
  }

  const contract = new Contract(
    opts.contract.address,
    opts.contract.interface,
    opts.contract.signer || opts.contract.provider
  )

  // Now reset Object.defineProperty
  Object.defineProperty = def

  // Override each function call to also `.wait()` so as to simplify the deploy scripts' syntax.
  for (const fnName of Object.keys(contract.functions)) {
    const fn = contract[fnName].bind(contract)
    ;(contract as any)[fnName] = async (...args: any) => {
      // We want to use the gas price that has been configured at the beginning of the deployment.
      // However, if the function being triggered is a "constant" (static) function, then we don't
      // want to provide a gas price because we're prone to getting insufficient balance errors.
      let gasPrice = opts.hre.deployConfig.gasPrice || undefined
      if (contract.interface.getFunction(fnName).constant) {
        gasPrice = 0
      }

      const tx = await fn(...args, {
        gasPrice,
      })

      if (typeof tx !== 'object' || typeof tx.wait !== 'function') {
        return tx
      }

      // Special logic for:
      // (1) handling confirmations
      // (2) handling an issue on Kovan specifically where transactions get dropped for no
      //     apparent reason.
      const maxTimeout = 120
      let timeout = 0
      while (true) {
        await sleep(1000)
        const receipt = await contract.provider.getTransactionReceipt(tx.hash)
        if (receipt === null) {
          timeout++
          if (timeout > maxTimeout && opts.hre.network.name === 'kovan') {
            // Special resubmission logic ONLY required on Kovan.
            console.log(
              `WARNING: Exceeded max timeout on transaction. Attempting to submit transaction again...`
            )
            return contract[fnName](...args)
          }
        } else if (
          receipt.confirmations >= opts.hre.deployConfig.numDeployConfirmations
        ) {
          return tx
        }
      }
    }
  }

  return contract
}

export const fundAccount = async (
  hre: any,
  address: string,
  amount: ethers.BigNumber
) => {
  if (!hre.deployConfig.isForkedNetwork) {
    throw new Error('this method can only be used against a forked network')
  }

  console.log(`Funding account ${address}...`)
  await hre.ethers.provider.send('hardhat_setBalance', [
    address,
    amount.toHexString(),
  ])

  console.log(`Waiting for balance to reflect...`)
  await awaitCondition(
    async () => {
      const balance = await hre.ethers.provider.getBalance(address)
      return balance.gte(amount)
    },
    5000,
    100
  )

  console.log(`Account successfully funded.`)
}

export const sendImpersonatedTx = async (opts: {
  hre: any
  contract: ethers.Contract
  fn: string
  from: string
  gas: string
  args: any[]
}) => {
  if (!opts.hre.deployConfig.isForkedNetwork) {
    throw new Error('this method can only be used against a forked network')
  }

  console.log(`Impersonating account ${opts.from}...`)
  await opts.hre.ethers.provider.send('hardhat_impersonateAccount', [opts.from])

  console.log(`Funding account ${opts.from}...`)
  await fundAccount(opts.hre, opts.from, BIG_BALANCE)

  console.log(`Sending impersonated transaction...`)
  const tx = await opts.contract.populateTransaction[opts.fn](...opts.args)
  const provider = new opts.hre.ethers.providers.JsonRpcProvider(
    (opts.hre.network.config as HttpNetworkConfig).url
  )
  await provider.send('eth_sendTransaction', [
    {
      ...tx,
      from: opts.from,
      gas: opts.gas,
    },
  ])

  console.log(`Stopping impersonation of account ${opts.from}...`)
  await opts.hre.ethers.provider.send('hardhat_stopImpersonatingAccount', [
    opts.from,
  ])
}

/**
 * Returns a version of the contract object which modifies all of the input contract's methods to
 * automatically await transaction receipts and confirmations. Will also throw if we timeout while
 * waiting for a transaction to be included in a block.
 *
 * @param opts Options for the contract.
 * @param opts.hre HardhatRuntimeEnvironment.
 * @param opts.contract Contract to wrap.
 * @returns Wrapped contract object.
 */
export const asAdvancedContract = (opts: {
  contract: Contract
  confirmations?: number
  gasPrice?: number
}): Contract => {
  // Temporarily override Object.defineProperty to bypass ether's object protection.
  const def = Object.defineProperty
  Object.defineProperty = (obj, propName, prop) => {
    prop.writable = true
    return def(obj, propName, prop)
  }

  const contract = new Contract(
    opts.contract.address,
    opts.contract.interface,
    opts.contract.signer || opts.contract.provider
  )

  // Now reset Object.defineProperty
  Object.defineProperty = def

  for (const fnName of Object.keys(contract.functions)) {
    const fn = contract[fnName].bind(contract)
    ;(contract as any)[fnName] = async (...args: any) => {
      // We want to use the configured gas price but we need to set the gas price to zero if we're
      // triggering a static function.
      let gasPrice = opts.gasPrice
      if (contract.interface.getFunction(fnName).constant) {
        gasPrice = 0
      }

      // Now actually trigger the transaction (or call).
      const tx = await fn(...args, {
        gasPrice,
      })

      // Meant for static calls, we don't need to wait for anything, we get the result right away.
      if (typeof tx !== 'object' || typeof tx.wait !== 'function') {
        return tx
      }

      // Wait for the transaction to be included in a block and wait for the specified number of
      // deployment confirmations.
      const maxTimeout = 120
      let timeout = 0
      while (true) {
        await sleep(1000)
        const receipt = await contract.provider.getTransactionReceipt(tx.hash)
        if (receipt === null) {
          timeout++
          if (timeout > maxTimeout) {
            throw new Error('timeout exceeded waiting for txn to be mined')
          }
        } else if (receipt.confirmations >= (opts.confirmations || 0)) {
          return tx
        }
      }
    }
  }

  return contract
}

export const getContractFromArtifact = async (
  hre: any,
  name: string,
  options: {
    iface?: string
    signerOrProvider?: Signer | Provider | string
  } = {}
): Promise<ethers.Contract> => {
  const artifact = await hre.deployments.get(name)
  await hre.ethers.provider.waitForTransaction(artifact.receipt.transactionHash)

  // Get the deployed contract's interface.
  let iface = new hre.ethers.utils.Interface(artifact.abi)
  // Override with optional iface name if requested.
  if (options.iface) {
    const factory = await hre.ethers.getContractFactory(options.iface)
    iface = factory.interface
  }

  let signerOrProvider: Signer | Provider = hre.ethers.provider
  if (options.signerOrProvider) {
    if (typeof options.signerOrProvider === 'string') {
      signerOrProvider = hre.ethers.provider.getSigner(options.signerOrProvider)
    } else {
      signerOrProvider = options.signerOrProvider
    }
  }

  return getAdvancedContract({
    hre,
    contract: new hre.ethers.Contract(
      artifact.address,
      iface,
      signerOrProvider
    ),
  })
}

/**
 * Gets multiple contract objects from their respective deployed artifacts.
 *
 * @param hre HardhatRuntimeEnvironment.
 * @param configs Array of contract names and options.
 * @returns Array of contract objects.
 */
export const getContractsFromArtifacts = async (
  hre: HardhatRuntimeEnvironment,
  configs: Array<{
    name: string
    iface?: string
    signerOrProvider?: Signer | Provider | string
  }>
): Promise<ethers.Contract[]> => {
  const contracts = []
  for (const config of configs) {
    contracts.push(await getContractFromArtifact(hre, config.name, config))
  }
  return contracts
}

/**
 * Helper function for asserting that a contract variable is set to the expected value.
 *
 * @param contract Contract object to query.
 * @param variable Name of the variable to query.
 * @param expected Expected value of the variable.
 */
export const assertContractVariable = async (
  contract: ethers.Contract,
  variable: string,
  expected: any
) => {
  // Need to make a copy that doesn't have a signer or we get the error that contracts with
  // signers cannot override the from address.
  const temp = new ethers.Contract(
    contract.address,
    contract.interface,
    contract.provider
  )

  const actual = await temp.callStatic[variable]({
    from: ethers.constants.AddressZero,
  })

  if (ethers.utils.isAddress(expected)) {
    assert(
      actual.toLowerCase() === expected.toLowerCase(),
      `[FATAL] ${variable} is ${actual} but should be ${expected}`
    )
    return
  }

  assert(
    actual === expected || (actual.eq && actual.eq(expected)),
    `[FATAL] ${variable} is ${actual} but should be ${expected}`
  )
}



export const isHardhatNode = async (hre) => {
  return (await getChainId(hre.ethers.provider)) === 31337
}

// Large balance to fund accounts with.
export const BIG_BALANCE = ethers.BigNumber.from(`0xFFFFFFFFFFFFFFFFFFFF`)

export const HexToBytes = async (hex: string) => {
  const bytes = []
  for (let c = 0; c < hex.length; c += 2) {
    bytes.push(parseInt(hex.substr(c, 2), 16))
  }
  return bytes
}

import { getContractInterface, predeploys } from '@mantleio/contracts'
import { ethers, Contract } from 'ethers'

import { toAddress } from './coercion'
import { DeepPartial } from './type-utils'
import {
  OEContracts,
  OEL1Contracts,
  OEL2Contracts,
  OEContractsLike,
  OEL2ContractsLike,
  AddressLike,
  BridgeAdapters,
  BridgeAdapterData,
  ICrossChainMessenger,
  L2ChainID,
} from '../interfaces'
import {
  StandardBridgeAdapter,
  ETHBridgeAdapter,
  ERC20BridgeAdapter,
} from '../adapters'

/**
 * Full list of default L2 contract addresses.
 */
export const DEFAULT_L2_CONTRACT_ADDRESSES: OEL2ContractsLike = {
  L2CrossDomainMessenger: predeploys.L2CrossDomainMessenger,
  L2StandardBridge: predeploys.L2StandardBridge,
  BVM_L1BlockNumber: predeploys.BVM_L1BlockNumber,
  BVM_L2ToL1MessagePasser: predeploys.BVM_L2ToL1MessagePasser,
  BVM_DeployerWhitelist: predeploys.BVM_DeployerWhitelist,
  BVM_ETH: predeploys.BVM_ETH,
  BVM_MANTLE: predeploys.BVM_MANTLE,
  BVM_GasPriceOracle: predeploys.BVM_GasPriceOracle,
  BVM_SequencerFeeVault: predeploys.BVM_SequencerFeeVault,
  WETH: predeploys.WETH9,
  TssRewardContract: predeploys.TssRewardContract,
}

/**
 * We've changed some contract names in this SDK to be a bit nicer. Here we remap these nicer names
 * back to the original contract names so we can look them up.
 */
const NAME_REMAPPING = {
  AddressManager: 'Lib_AddressManager' as const,
  BVM_L1BlockNumber: 'iBVM_L1BlockNumber' as const,
  WETH: 'WETH9' as const,
}

/**
 * Mapping of L2 chain IDs to MNT contract addresses to detect correct token in standard-bridge
 */
export const L1_MNT_ADDRESS: {
  [ChainID in L2ChainID]: string
} = {
  [L2ChainID.MANTLE]: "0x3c3a81e81dc49A522A592e7622A7E711c06bf354",
  [L2ChainID.MANTLE_KOVAN]: "0x3c3a81e81dc49A522A592e7622A7E711c06bf354",
  [L2ChainID.MANTLE_GOERLIQA]: "0x3c3a81e81dc49A522A592e7622A7E711c06bf354",
  [L2ChainID.MANTLE_TESTNET]: "0xc1dC2d65A2243c22344E725677A3E3BEBD26E604",
  [L2ChainID.MANTLE_HARDHAT_LOCAL]: "0x3c3a81e81dc49A522A592e7622A7E711c06bf354",
  [L2ChainID.MANTLE_HARDHAT_DEVNET]: "0x3c3a81e81dc49A522A592e7622A7E711c06bf354"
}

/**
 * Mapping of L1 chain IDs to the appropriate contract addresses for the OE deployments to the
 * given network. Simplifies the process of getting the correct contract addresses for a given
 * contract name.
 */
export const CONTRACT_ADDRESSES: {
  [ChainID in L2ChainID]: OEContractsLike
} = {
  [L2ChainID.MANTLE]: {
    l1: {
      AddressManager: '0x6968f3F16C3e64003F02E121cf0D5CCBf5625a42' as const,
      L1CrossDomainMessenger:
        '0x676A795fe6E43C17c668de16730c3F690FEB7120' as const,
      L1StandardBridge: '0x95fC37A27a2f68e3A647CDc081F0A89bb47c3012' as const,
      StateCommitmentChain:
        '0x89E9D387555AF0cDE22cb98833Bae40d640AD7fa' as const,
      CanonicalTransactionChain:
        '0x291dc3819b863e19b0a9b9809F8025d2EB4aaE93' as const,
      BondManager: '0x31aBe1c466C2A8b95fd84258dD1471472979B650' as const,
      Rollup:
        process.env.Rollup ||
        ('0xD1328C9167e0693B689b5aa5a024379d4e437858' as const),
    },
    l2: DEFAULT_L2_CONTRACT_ADDRESSES,
  },
  [L2ChainID.MANTLE_KOVAN]: {
    l1: {
      AddressManager: '0x100Dd3b414Df5BbA2B542864fF94aF8024aFdf3a' as const,
      L1CrossDomainMessenger:
        '0x4361d0F75A0186C05f971c566dC6bEa5957483fD' as const,
      L1StandardBridge: '0x22F24361D548e5FaAfb36d1437839f080363982B' as const,
      StateCommitmentChain:
        '0xD7754711773489F31A0602635f3F167826ce53C5' as const,
      CanonicalTransactionChain:
        '0xf7B88A133202d41Fe5E2Ab22e6309a1A4D50AF74' as const,
      BondManager: '0xc5a603d273E28185c18Ba4d26A0024B2d2F42740' as const,
      Rollup:
        process.env.Rollup ||
        ('0x9faB987C9C469EB23Da31B7848B28aCf30905eA8' as const),
    },
    l2: DEFAULT_L2_CONTRACT_ADDRESSES,
  },
  [L2ChainID.MANTLE_GOERLIQA]: {
    l1: {
      AddressManager:
        process.env.ADDRESS_MANAGER_ADDRESS ||
        ('0x327903410307971Ca7Ba8A6CB2291D3b8825d7F5' as const),
      L1CrossDomainMessenger:
        process.env.L1_CROSS_DOMAIN_MESSENGER_ADDRESS ||
        ('0x3f41DAcb2dB659e45826126d004ad3E0C8eA680e' as const),
      L1StandardBridge:
        process.env.L1_STANDARD_BRIDGE_ADDRESS ||
        ('0x4cf99b9BC9B2Da64033D1Fb65146Ea60fbe8AD4B' as const),
      StateCommitmentChain:
        process.env.STATE_COMMITMENT_CHAIN_ADDRESS ||
        ('0x88EC574e2ef0EcF9043373139099f7E535F94dBC' as const),
      CanonicalTransactionChain:
        process.env.CANONICAL_TRANSACTION_CHAIN_ADDRESS ||
        ('0x258e80D5371fD7fFdDFE29E60b366f9FC44844c8' as const),
      BondManager:
        process.env.BOND_MANAGER_ADDRESS ||
        ('0xc723Cb5f3337c2F6Eab9b29E78CE42a28B8661d1' as const),
      Rollup:
        process.env.Rollup ||
        ('0x9faB987C9C469EB23Da31B7848B28aCf30905eA8' as const),
    },
    l2: DEFAULT_L2_CONTRACT_ADDRESSES,
  },
  [L2ChainID.MANTLE_TESTNET]: {
    l1: {
      AddressManager: '0xA647F5947C50248bc4b2eF773791c9C2bc01C65A' as const,
      L1CrossDomainMessenger:
        '0x7Bfe603647d5380ED3909F6f87580D0Af1B228B4' as const,
      L1StandardBridge: '0xc92470D7Ffa21473611ab6c6e2FcFB8637c8f330' as const,
      StateCommitmentChain:
        '0x91A5D806BA73d0AA4bFA9B318126dDE60582e92a' as const,
      CanonicalTransactionChain:
        '0x654e6dF111F98374d9e5d908D7a5392C308aA18D' as const,
      BondManager: '0xeBE3f28BbFa7bB8f2C066C1A792073203B985e27' as const,
      Rollup:
        process.env.Rollup ||
        ('0x9faB987C9C469EB23Da31B7848B28aCf30905eA8' as const),
    },
    l2: DEFAULT_L2_CONTRACT_ADDRESSES,
  },
  [L2ChainID.MANTLE_HARDHAT_LOCAL]: {
    l1: {
      AddressManager:
        process.env.ADDRESS_MANAGER_ADDRESS ||
        ('0x92aBAD50368175785e4270ca9eFd169c949C4ce1' as const),
      L1CrossDomainMessenger:
        process.env.L1_CROSS_DOMAIN_MESSENGER_ADDRESS ||
        ('0x7959CF3b8ffC87Faca8aD8a1B5D95c0f58C0BEf8' as const),
      L1StandardBridge:
        process.env.L1_STANDARD_BRIDGE_ADDRESS ||
        ('0x8BAccFF561FDe61D6bC8B6f299fFBa561d2189B9' as const),
      StateCommitmentChain:
        process.env.STATE_COMMITMENT_CHAIN_ADDRESS ||
        ('0xd9e2F450525079e1e29fB23Bc7Caca6F61f8fD4a' as const),
      CanonicalTransactionChain:
        process.env.CANONICAL_TRANSACTION_CHAIN_ADDRESS ||
        ('0x0090171f848B2aa86918E5Ef2406Ab3d424fdd83' as const),
      BondManager:
        process.env.BOND_MANAGER_ADDRESS ||
        ('0x9faB987C9C469EB23Da31B7848B28aCf30905eA8' as const),
      Rollup:
        process.env.Rollup ||
        ('0x9faB987C9C469EB23Da31B7848B28aCf30905eA8' as const),
    },
    l2: DEFAULT_L2_CONTRACT_ADDRESSES,
  },
  [L2ChainID.MANTLE_HARDHAT_DEVNET]: {
    l1: {
      AddressManager:
        process.env.ADDRESS_MANAGER_ADDRESS ||
        ('0x92aBAD50368175785e4270ca9eFd169c949C4ce1' as const),
      L1CrossDomainMessenger:
        process.env.L1_CROSS_DOMAIN_MESSENGER_ADDRESS ||
        ('0x7959CF3b8ffC87Faca8aD8a1B5D95c0f58C0BEf8' as const),
      L1StandardBridge:
        process.env.L1_STANDARD_BRIDGE_ADDRESS ||
        ('0x8BAccFF561FDe61D6bC8B6f299fFBa561d2189B9' as const),
      StateCommitmentChain:
        process.env.STATE_COMMITMENT_CHAIN_ADDRESS ||
        ('0xd9e2F450525079e1e29fB23Bc7Caca6F61f8fD4a' as const),
      CanonicalTransactionChain:
        process.env.CANONICAL_TRANSACTION_CHAIN_ADDRESS ||
        ('0x0090171f848B2aa86918E5Ef2406Ab3d424fdd83' as const),
      BondManager:
        process.env.BOND_MANAGER_ADDRESS ||
        ('0x9faB987C9C469EB23Da31B7848B28aCf30905eA8' as const),
      Rollup:
        process.env.Rollup ||
        ('0x9faB987C9C469EB23Da31B7848B28aCf30905eA8' as const),
    },
    l2: DEFAULT_L2_CONTRACT_ADDRESSES,
  },
}

/**
 * Mapping of L1 chain IDs to the list of custom bridge addresses for each chain.
 */
export const BRIDGE_ADAPTER_DATA: {
  [ChainID in L2ChainID]?: BridgeAdapterData
} = {
  [L2ChainID.MANTLE]: {
    BitBTC: {
      Adapter: StandardBridgeAdapter,
      l1Bridge: '0xaBA2c5F108F7E820C049D5Af70B16ac266c8f128' as const,
      l2Bridge: '0x158F513096923fF2d3aab2BcF4478536de6725e2' as const,
    },
    DAI: {
      Adapter: ERC20BridgeAdapter,
      l1Bridge: '0x10E6593CDda8c58a1d0f14C5164B376352a55f2F' as const,
      l2Bridge: '0x467194771dAe2967Aef3ECbEDD3Bf9a310C76C65' as const,
    },
  },
  [L2ChainID.MANTLE_TESTNET]: {
    wstETH: {
      Adapter: ERC20BridgeAdapter,
      l1Bridge: '0x2fD573Ace456904709444d04AdCa189fB19e725a' as const,
      l2Bridge: '0x08C2EE913D3cb544D182bCC7632cB0B382A2933e' as const,
    },
  },
  [L2ChainID.MANTLE_KOVAN]: {
    wstETH: {
      Adapter: ERC20BridgeAdapter,
      l1Bridge: '0xa88751C0a08623E11ff38c6B70F2BbEe7865C17c' as const,
      l2Bridge: '0xF9C842dE4381a70eB265d10CF8D43DceFF5bA935' as const,
    },
    BitBTC: {
      Adapter: StandardBridgeAdapter,
      l1Bridge: '0x0b651A42F32069d62d5ECf4f2a7e5Bd3E9438746' as const,
      l2Bridge: '0x0CFb46528a7002a7D8877a5F7a69b9AaF1A9058e' as const,
    },
    USX: {
      Adapter: StandardBridgeAdapter,
      l1Bridge: '0x40E862341b2416345F02c41Ac70df08525150dC7' as const,
      l2Bridge: '0xB4d37826b14Cd3CB7257A2A5094507d701fe715f' as const,
    },
    DAI: {
      Adapter: ERC20BridgeAdapter,
      l1Bridge: '0xb415e822C4983ecD6B1c1596e8a5f976cf6CD9e3' as const,
      l2Bridge: '0x467194771dAe2967Aef3ECbEDD3Bf9a310C76C65' as const,
    },
  },
}

/**
 * Returns an ethers.Contract object for the given name, connected to the appropriate address for
 * the given L2 chain ID. Users can also provide a custom address to connect the contract to
 * instead. If the chain ID is not known then the user MUST provide a custom address or this
 * function will throw an error.
 *
 * @param contractName Name of the contract to connect to.
 * @param l2ChainId Chain ID for the L2 network.
 * @param opts Additional options for connecting to the contract.
 * @param opts.address Custom address to connect to the contract.
 * @param opts.signerOrProvider Signer or provider to connect to the contract.
 * @returns An ethers.Contract object connected to the appropriate address and interface.
 */
export const getOEContract = (
  contractName: keyof OEL1Contracts | keyof OEL2Contracts,
  l2ChainId: number,
  opts: {
    address?: AddressLike
    signerOrProvider?: ethers.Signer | ethers.providers.Provider
  } = {}
): Contract => {
  const addresses = CONTRACT_ADDRESSES[l2ChainId]
  if (addresses === undefined && opts.address === undefined) {
    throw new Error(
      `cannot get contract ${contractName} for unknown L2 chain ID ${l2ChainId}, you must provide an address`
    )
  }

  const name = NAME_REMAPPING[contractName] || contractName
  let iface: ethers.utils.Interface

  // eslint-disable-next-line prefer-const
  iface = getContractInterface(name)

  return new Contract(
    toAddress(
      opts.address || addresses.l1[contractName] || addresses.l2[contractName]
    ),
    iface,
    opts.signerOrProvider
  )
}

/**
 * Automatically connects to all contract addresses, both L1 and L2, for the given L2 chain ID. The
 * user can provide custom contract address overrides for L1 or L2 contracts. If the given chain ID
 * is not known then the user MUST provide custom contract addresses for ALL L1 contracts or this
 * function will throw an error.
 *
 * @param l2ChainId Chain ID for the L2 network.
 * @param opts Additional options for connecting to the contracts.
 * @param opts.l1SignerOrProvider: Signer or provider to connect to the L1 contracts.
 * @param opts.l2SignerOrProvider: Signer or provider to connect to the L2 contracts.
 * @param opts.overrides Custom contract address overrides for L1 or L2 contracts.
 * @returns An object containing ethers.Contract objects connected to the appropriate addresses on
 * both L1 and L2.
 */
export const getAllOEContracts = (
  l2ChainId: number,
  opts: {
    l1SignerOrProvider?: ethers.Signer | ethers.providers.Provider
    l2SignerOrProvider?: ethers.Signer | ethers.providers.Provider
    overrides?: DeepPartial<OEContractsLike>
  } = {}
): OEContracts => {
  const addresses = CONTRACT_ADDRESSES[l2ChainId] || {
    l1: {
      AddressManager: undefined,
      L1CrossDomainMessenger: undefined,
      L1StandardBridge: undefined,
      StateCommitmentChain: undefined,
      CanonicalTransactionChain: undefined,
      BondManager: undefined,
    },
    l2: DEFAULT_L2_CONTRACT_ADDRESSES,
  }

  // Attach all L1 contracts.
  const l1Contracts = {} as OEL1Contracts
  for (const [contractName, contractAddress] of Object.entries(addresses.l1)) {
    l1Contracts[contractName] = getOEContract(
      contractName as keyof OEL1Contracts,
      l2ChainId,
      {
        address: opts.overrides?.l1?.[contractName] || contractAddress,
        signerOrProvider: opts.l1SignerOrProvider,
      }
    )
  }

  // Attach all L2 contracts.
  const l2Contracts = {} as OEL2Contracts
  for (const [contractName, contractAddress] of Object.entries(addresses.l2)) {
    l2Contracts[contractName] = getOEContract(
      contractName as keyof OEL2Contracts,
      l2ChainId,
      {
        address: opts.overrides?.l2?.[contractName] || contractAddress,
        signerOrProvider: opts.l2SignerOrProvider,
      }
    )
  }

  return {
    l1: l1Contracts,
    l2: l2Contracts,
  }
}

/**
 * Gets a series of bridge adapters for the given L2 chain ID.
 *
 * @param l2ChainId Chain ID for the L2 network.
 * @param messenger Cross chain messenger to connect to the bridge adapters
 * @param opts Additional options for connecting to the custom bridges.
 * @param opts.overrides Custom bridge adapters.
 * @returns An object containing all bridge adapters
 */
export const getBridgeAdapters = (
  l2ChainId: number,
  messenger: ICrossChainMessenger,
  opts?: {
    overrides?: BridgeAdapterData
  }
): BridgeAdapters => {
  const adapterData: BridgeAdapterData = {
    ...(CONTRACT_ADDRESSES[l2ChainId]
      ? {
          Standard: {
            Adapter: StandardBridgeAdapter,
            l1Bridge: CONTRACT_ADDRESSES[l2ChainId].l1.L1StandardBridge,
            l2Bridge: predeploys.L2StandardBridge,
          },
          ETH: {
            Adapter: ETHBridgeAdapter,
            l1Bridge: CONTRACT_ADDRESSES[l2ChainId].l1.L1StandardBridge,
            l2Bridge: predeploys.L2StandardBridge,
          },
        }
      : {}),
    ...(BRIDGE_ADAPTER_DATA[l2ChainId] || {}),
    ...(opts?.overrides || {}),
  }

  const adapters: BridgeAdapters = {}
  for (const [bridgeName, bridgeData] of Object.entries(adapterData)) {
    adapters[bridgeName] = new bridgeData.Adapter({
      messenger,
      l1Bridge: bridgeData.l1Bridge,
      l2Bridge: bridgeData.l2Bridge,
    })
  }

  return adapters
}

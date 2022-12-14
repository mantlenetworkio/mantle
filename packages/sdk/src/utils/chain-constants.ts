import { L1ChainID, L2ChainID } from '../interfaces'

export const DEPOSIT_CONFIRMATION_BLOCKS: {
  [ChainID in L2ChainID]: number
} = {
  [L2ChainID.MANTLE]: 50 as const,
  [L2ChainID.MANTLE_KOVAN]: 12 as const,
  [L2ChainID.MANTLE_GOERLIQA]: 12 as const,
  [L2ChainID.MANTLE_TESTNET]: 12 as const,
  [L2ChainID.MANTLE_HARDHAT_LOCAL]: 2 as const,
  [L2ChainID.MANTLE_HARDHAT_DEVNET]: 2 as const,
}

export const CHAIN_BLOCK_TIMES: {
  [ChainID in L1ChainID]: number
} = {
  [L1ChainID.MAINNET]: 13 as const,
  [L1ChainID.GOERLI]: 15 as const,
  [L1ChainID.KOVAN]: 4 as const,
  [L1ChainID.HARDHAT_LOCAL]: 1 as const,
}

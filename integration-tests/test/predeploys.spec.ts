/* Imports: External */
import { ethers } from 'ethers'
import { predeploys, getContractInterface } from '@mantleio/contracts'

/* Imports: Internal */
import { expect } from './shared/setup'
import { MantleEnv } from './shared/env'

describe('predeploys', () => {
  let env: MantleEnv
  before(async () => {
    env = await MantleEnv.new()
  })

  describe('WETH9', () => {
    let weth9: ethers.Contract
    before(() => {
      weth9 = new ethers.Contract(
        predeploys.WETH9,
        getContractInterface('WETH9'),
        env.l2Wallet
      )
    })

    it('should have the correct name', async () => {
      expect(await weth9.name()).to.equal('Wrapped Ether')
    })

    it('should have the correct symbol', async () => {
      expect(await weth9.symbol()).to.equal('WETH')
    })

    it('should have the correct decimals', async () => {
      expect(await weth9.decimals()).to.equal(18)
    })
  })

  describe('BVM_ETH', () => {
    let BVMEth: ethers.Contract
    before(() => {
      BVMEth = new ethers.Contract(
        predeploys.BVM_ETH,
        getContractInterface('BVM_ETH'),
        env.l2Wallet
      )
    })

    it('should have the correct name', async () => {
      expect(await BVMEth.name()).to.equal('Ether')
    })

    it('should have the correct symbol', async () => {
      expect(await BVMEth.symbol()).to.equal('ETH')
    })

    it('should have the correct decimals', async () => {
      expect(await BVMEth.decimals()).to.equal(18)
    })
  })

  describe('L2CrossDomainMessenger', () => {
    let l2CrossDomainMessenger: ethers.Contract
    before(() => {
      l2CrossDomainMessenger = new ethers.Contract(
        predeploys.L2CrossDomainMessenger,
        getContractInterface('L2CrossDomainMessenger'),
        env.l2Wallet
      )
    })

    it('should throw when calling xDomainMessageSender', async () => {
      await expect(
        l2CrossDomainMessenger.xDomainMessageSender()
      ).to.be.revertedWith('xDomainMessageSender is not set')
    })
  })

  describe('L2StandardBridge', () => {
    let l2StandardBridge: ethers.Contract
    before(() => {
      l2StandardBridge = new ethers.Contract(
        predeploys.L2StandardBridge,
        getContractInterface('L2StandardBridge'),
        env.l2Wallet
      )
    })

    it('should have the correct messenger address', async () => {
      expect(await l2StandardBridge.messenger()).to.equal(
        predeploys.L2CrossDomainMessenger
      )
    })

    it('should have a nonzero bridge address', async () => {
      expect(await l2StandardBridge.l1TokenBridge()).to.not.equal(
        ethers.constants.AddressZero
      )
    })
  })

  describe('BVM_SequencerFeeVault', () => {
    let BVMSequencerFeeVault: ethers.Contract
    before(() => {
      BVMSequencerFeeVault = new ethers.Contract(
        predeploys.BVM_SequencerFeeVault,
        getContractInterface('BVM_SequencerFeeVault'),
        env.l2Wallet
      )
    })

    it('should have a nonzero l1FeeWallet', async () => {
      expect(await BVMSequencerFeeVault.l1FeeWallet()).to.not.equal(
        ethers.constants.AddressZero
      )
    })
  })
})

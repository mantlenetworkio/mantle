/* External Imports */
import { ethers } from 'hardhat'
import { Contract } from 'ethers'
import { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers'

import { expect } from '../../../setup'
import { deploy } from '../../../helpers'

describe('BVM_ETH', () => {
  let signer1: SignerWithAddress
  let signer2: SignerWithAddress
  before(async () => {
    ;[signer1, signer2] = await ethers.getSigners()
  })

  let BVM_ETH: Contract
  beforeEach(async () => {
    BVM_ETH = await deploy('BVM_ETH')
  })

  describe('transfer', () => {
    it('should revert', async () => {
      await expect(BVM_ETH.transfer(signer2.address, 100)).to.be.revertedWith(
        'BVM_ETH: transfer is disabled pending further community discussion.'
      )
    })
  })

  describe('approve', () => {
    it('should revert', async () => {
      await expect(BVM_ETH.approve(signer2.address, 100)).to.be.revertedWith(
        'BVM_ETH: approve is disabled pending further community discussion.'
      )
    })
  })

  describe('transferFrom', () => {
    it('should revert', async () => {
      await expect(
        BVM_ETH.transferFrom(signer1.address, signer2.address, 100)
      ).to.be.revertedWith(
        'BVM_ETH: transferFrom is disabled pending further community discussion.'
      )
    })
  })

  describe('increaseAllowance', () => {
    it('should revert', async () => {
      await expect(
        BVM_ETH.increaseAllowance(signer2.address, 100)
      ).to.be.revertedWith(
        'BVM_ETH: increaseAllowance is disabled pending further community discussion.'
      )
    })
  })

  describe('decreaseAllowance', () => {
    it('should revert', async () => {
      await expect(
        BVM_ETH.decreaseAllowance(signer2.address, 100)
      ).to.be.revertedWith(
        'BVM_ETH: decreaseAllowance is disabled pending further community discussion.'
      )
    })
  })
})

import { ethers } from 'hardhat'
import { Contract } from 'ethers'
import { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers'

import { expect } from '../../../setup'
import {deploy} from "../../../helpers";

describe('TssRewardContract', () => {
  let signer1: SignerWithAddress
  let signer2: SignerWithAddress
  before(async () => {
    ;[signer1,signer2] = await ethers.getSigners()
  })

  let tssRewardContract: Contract
  beforeEach(async () => {
    tssRewardContract = await deploy('TssRewardContract', {
      signer: signer1,
      args: [signer1.address,signer2.address],
    })
  })

  describe('owner', () => {
    it('should have an owner', async () => {
      expect(await tssRewardContract.owner()).to.equal(signer2.address)
    })
  })

  describe('deadAddress', () => {
    it('should have an deadAddress', async () => {
      expect(await tssRewardContract.deadAddress()).to.equal(signer1.address)
    })
  })

})

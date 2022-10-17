import { Signer, Contract, constants } from 'ethers'
import chai from 'chai'

import { deploy } from '../../../helpers'

const { expect } = chai
const { ethers } = require('hardhat')

describe('Sequencer', () => {
  let accounts: Signer[]
  let sequencer: Contract
  let bitToken: Contract
  const mintAddress = '0x0000000000000000000000000000000000000001'
  const nodeID = '0x10'

  before('deploy stakingSlashing contracts', async () => {
    accounts = await ethers.getSigners()
    await deployBitToken()
    await deploySequencer()
  })

  it('Change bit address', async () => {
    await sequencer.changeBitAddress(bitToken.address)
    expect(await sequencer.bitToken()).to.eq(bitToken.address.toString())
  })

  it('CreateSequencer', async () => {
    await bitToken.connect(accounts[0]).approve(sequencer.address, 100)
    // err case: Invild amount
    await expect(
      sequencer.connect(accounts[0]).createSequencer(0, mintAddress, nodeID)
    ).to.be.revertedWith('Invild amount')
    // err case: Invild address
    await expect(
      sequencer
        .connect(accounts[0])
        .createSequencer(100, constants.AddressZero, nodeID)
    ).to.be.revertedWith('Invild address, address can not be 0')
    // create sequencer by account[0]
    await sequencer
      .connect(accounts[0])
      .createSequencer(100, mintAddress, nodeID)
    const owners = await sequencer.getOwners()
    expect(owners.length).to.eq(1)
    expect(owners[0]).to.eq(await accounts[0].getAddress())
    expect(await sequencer.owners(0)).to.eq(await accounts[0].getAddress())
    // err case: Already have deposit
    await expect(
      sequencer.connect(accounts[0]).createSequencer(100, mintAddress, nodeID)
    ).to.be.revertedWith('Already have deposit')
    // err case: Already have deposit
    await expect(
      sequencer.connect(accounts[1]).createSequencer(100, mintAddress, nodeID)
    ).to.be.revertedWith('This mint address already have sender')

    // create sequencer by account[1]
    await bitToken.connect(accounts[1]).approve(sequencer.address, 100)
    await sequencer
      .connect(accounts[1])
      .createSequencer(
        100,
        '0x0000000000000000000000000000000000000002',
        nodeID
      )
  })

  it('SequencerQuery', async () => {
    expect((await sequencer.getSequencers()).length).to.eq(2)
    expect((await sequencer.getOwners()).length).to.eq(2)

    // query sequencer by account[0]
    const sequencerInfo0 = await sequencer.getSequencer(
      await accounts[0].getAddress()
    )
    expect(sequencerInfo0.owner).to.eq(await accounts[0].getAddress())
    expect(sequencerInfo0.mintAddress).to.eq(mintAddress)
    expect(sequencerInfo0.nodeID).to.eq(nodeID)
    expect(sequencerInfo0.amount).to.eq(100)
    expect(sequencerInfo0.keyIndex).to.eq(0)

    // query sequencer by account[1]
    const sequencerInfo1 = await sequencer.getSequencer(
      await accounts[1].getAddress()
    )
    expect(sequencerInfo1.owner).to.eq(await accounts[1].getAddress())
    expect(sequencerInfo1.mintAddress).to.eq(
      '0x0000000000000000000000000000000000000002'
    )
    expect(sequencerInfo1.nodeID).to.eq(nodeID)
    expect(sequencerInfo1.amount).to.eq(100)
    expect(sequencerInfo1.keyIndex).to.eq(1)
    // check sequencer exist
    expect(await sequencer.isSequencer(await accounts[0].getAddress())).to.eq(
      true
    )
    expect(await sequencer.isSequencer(await accounts[1].getAddress())).to.eq(
      true
    )
    expect(await sequencer.isSequencer(await accounts[2].getAddress())).to.eq(
      false
    )
  })

  it('Deposit', async () => {
    await bitToken.connect(accounts[0]).approve(sequencer.address, 200)
    // err case: Invild amount
    await expect(sequencer.connect(accounts[0]).deposit(0)).to.be.revertedWith(
      'Invild amount'
    )
    await expect(
      sequencer.connect(accounts[2]).deposit(200)
    ).to.be.revertedWith('Sequencer not exist')

    // deposit 200 to account[0]
    await sequencer.connect(accounts[0]).deposit(200)
    // query sequencer by account[0] then check the amount is 100+200 = 300 and balance is 700
    const sequencerInfo0 = await sequencer.getSequencer(
      await accounts[0].getAddress()
    )
    expect(sequencerInfo0.amount).to.eq(300)
    expect(await bitToken.balanceOf(accounts[0].getAddress())).to.eq(700)
  })

  it('Withdraw', async () => {
    // err case: Invild amount
    await expect(sequencer.connect(accounts[0]).withdraw(0)).to.be.revertedWith(
      'Invild amount'
    )
    // err case: Sequencer not exist
    await expect(
      sequencer.connect(accounts[2]).withdraw(200)
    ).to.be.revertedWith('Sequencer not exist')
    // withdraw a little
    await sequencer.connect(accounts[0]).withdraw(100)
    // check deposit and balance
    const sequencerInfo0 = await sequencer.getSequencer(
      await accounts[0].getAddress()
    )
    expect(sequencerInfo0.amount).to.eq(200)
    expect(await bitToken.balanceOf(await accounts[0].getAddress())).to.eq(800)
    // withdraw all, 300 > 200
    await sequencer.connect(accounts[0]).withdraw(300)
    // check sequencer delete and balance
    expect(await sequencer.isSequencer(await accounts[0].getAddress())).to.eq(
      false
    )
    expect(await bitToken.balanceOf(await accounts[0].getAddress())).to.eq(1000)
    expect((await sequencer.getOwners()).length).to.eq(1)
    expect(await sequencer.owners(0)).to.eq(await accounts[1].getAddress())
  })

  it('Withdraw all', async () => {
    // err case: Sequencer not exist
    await expect(
      sequencer.connect(accounts[2]).withdraw(200)
    ).to.be.revertedWith('Sequencer not exist')
    // withdraw all
    await sequencer.connect(accounts[1]).withdrawAll()
    // check sequencer delete and balance
    expect(await sequencer.isSequencer(await accounts[1].getAddress())).to.eq(
      false
    )
    expect(await bitToken.balanceOf(await accounts[1].getAddress())).to.eq(1000)
    expect((await sequencer.getOwners()).length).to.eq(0)
    expect((await sequencer.getSequencers()).length).to.eq(0)
  })

  const deployBitToken = async () => {
    bitToken = await deploy('TestERC20')

    // mint then check balance
    bitToken.mint(accounts[0].getAddress(), 1000)
    bitToken.mint(accounts[1].getAddress(), 1000)
    expect(await bitToken.balanceOf(accounts[0].getAddress())).to.eq(1000)
    expect(await bitToken.balanceOf(accounts[1].getAddress())).to.eq(1000)
  }

  const deploySequencer = async () => {
    sequencer = await deploy('Sequencer')
    await sequencer.initialize(ethers.constants.AddressZero)
  }
})

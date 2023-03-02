import { Signer, Contract, constants } from 'ethers'
import chai from 'chai'

// @ts-ignore
import { deploy } from '../../../helpers'

const { expect } = chai
// eslint-disable-next-line @typescript-eslint/no-var-requires
const { ethers } = require('hardhat')

describe('RollUp', () => {
  let accounts: Signer[]
  let rollUp: Contract
  let verifier: Contract
  let token: Contract
  let assertionMap: Contract
  let addressManager: Contract

  before('setup', async () => {
    accounts = await ethers.getSigners()
    await deployAddressManager()
    await deploySCC()
    await deployToken()
    await deployVerifier()
    await deployRollUp()
  })

  it('stake', async () => {
    const amountStake = 1000
    expect(await rollUp.isStaked(await accounts[0].getAddress())).to.eq(false)
    expect(await rollUp.numStakers()).to.eq(0)

    await rollUp.stake({ value: amountStake })
    expect(await rollUp.isStaked(await accounts[0].getAddress())).to.eq(true)
    expect(await rollUp.numStakers()).to.eq(1)
    const staker = await rollUp.stakers(await accounts[0].getAddress())
    expect(staker.isStaked).to.eq(true)
    expect(staker.amountStaked.toNumber()).to.eq(amountStake)
    expect(staker.assertionID.toNumber()).to.eq(0)
    expect(staker.currentChallenge).to.eq(constants.AddressZero)
  })

  it('unstake', async () => {
    const amountUnStake = 100
    let staker = await rollUp.stakers(await accounts[0].getAddress())
    const amountStake = staker.amountStaked.toNumber()
    const baseStakeAmount = await rollUp.baseStakeAmount()

    await rollUp.unstake(amountUnStake)

    staker = await rollUp.stakers(await accounts[0].getAddress())
    expect(staker.amountStaked.toNumber()).to.eq(amountStake - baseStakeAmount)
  })

  it('removeStake', async () => {
    expect(await rollUp.isStaked(await accounts[0].getAddress())).to.eq(true)
    await rollUp.removeStake(await accounts[0].getAddress())
    expect(await rollUp.isStaked(await accounts[0].getAddress())).to.eq(false)
  })

  it('createAssertion', async () => {
    const amountStake = 10000
    await rollUp.stake({ value: amountStake })

    await rollUp.createAssertion(
      '0x0000000000000000000000000000000000000000000000000000000000000001',
      1
    )
    expect(await rollUp.lastCreatedAssertionID()).to.eq(1)

    await rollUp.createAssertion(
      '0x0000000000000000000000000000000000000000000000000000000000000001',
      2
    )
    expect(await rollUp.lastCreatedAssertionID()).to.eq(2)

    await rollUp.createAssertion(
      '0x0000000000000000000000000000000000000000000000000000000000000001',
      3
    )

    expect(await rollUp.lastCreatedAssertionID()).to.eq(3)

    await rollUp.createAssertion(
      '0x0000000000000000000000000000000000000000000000000000000000000001',
      4
    )
    expect(await rollUp.lastCreatedAssertionID()).to.eq(4)
    await rollUp.connect(await accounts[3]).stake({ value: 1000 })
    await rollUp.connect(await accounts[3]).advanceStake(1)
    await rollUp.connect(await accounts[3]).advanceStake(2)
    await rollUp
      .connect(await accounts[3])
      .createAssertion(
        '0x0000000000000000000000000000000000000000000000000000000000000002',
        3
      )
    expect(await rollUp.lastCreatedAssertionID()).to.eq(5)
  })

  it('confirmFirstUnresolvedAssertion', async () => {
    expect(await rollUp.lastResolvedAssertionID()).to.eq(0)
    expect(await rollUp.lastConfirmedAssertionID()).to.eq(0)

    await rollUp.confirmFirstUnresolvedAssertion()
    expect(await rollUp.lastResolvedAssertionID()).to.eq(1)
    expect(await rollUp.lastConfirmedAssertionID()).to.eq(1)

    await rollUp.confirmFirstUnresolvedAssertion()
    expect(await rollUp.lastResolvedAssertionID()).to.eq(2)
    expect(await rollUp.lastConfirmedAssertionID()).to.eq(2)
  })

  it('challengeAssertion', async () => {
    const players = [
      await accounts[0].getAddress(),
      await accounts[3].getAddress(),
    ]
    const assertionIDs = [3, 5]
    await rollUp.challengeAssertion(players, assertionIDs)
    expect(
      (await rollUp.stakers(await accounts[0].getAddress())).currentChallenge
    ).to.eq(
      (await rollUp.stakers(await accounts[3].getAddress())).currentChallenge
    )
  })

  it('verifyOneStepProof', async () => {
    const winnerAddr = await accounts[3].getAddress()
    const loserAddr = await accounts[0].getAddress()
    const ownerAddr = await accounts[1].getAddress()

    const challengeImp = await deploy('Challenge')
    const challengeAddr = (await rollUp.stakers(winnerAddr)).currentChallenge
    const challenge = new Contract(
      challengeAddr,
      challengeImp.interface,
      accounts[0]
    )

    await challenge.initializeChallengeLength(1)
    const winnerAmount = (
      await rollUp.stakers(winnerAddr)
    ).amountStaked.toNumber()
    const loserAmount = (
      await rollUp.stakers(loserAddr)
    ).amountStaked.toNumber()
    const withdrawLoser = loserAmount - winnerAmount
    // todo change to mock data
    await challenge.connect(accounts[3]).verifyOneStepProof(
      '0x0000', // proof
      1, // challengedStepIndex
      [await assertionMap.getStateHash(2), await assertionMap.getStateHash(3)], // prevBisection
      0, // prevChallengedSegmentStart
      1 // prevChallengedSegmentLength
    )
    // check amount
    expect(await rollUp.withdrawableFunds(loserAddr)).to.eq(withdrawLoser)
    expect(await rollUp.withdrawableFunds(ownerAddr)).to.eq(
      loserAmount - withdrawLoser - (loserAmount - withdrawLoser) / 2
    )
    expect((await rollUp.stakers(winnerAddr)).amountStaked.toNumber()).to.eq(
      (loserAmount - withdrawLoser) / 2 + winnerAmount
    )
    expect(await rollUp.isStaked(loserAddr)).to.eq(false)
    expect((await rollUp.zombies(0)).stakerAddress).to.eq(loserAddr)
    expect((await rollUp.zombies(0)).lastAssertionID.toNumber()).to.eq(4)
  })

  const deployToken = async () => {
    token = await deploy('TestERC20')
  }

  const deployRollUp = async () => {
    const rollUpImp = await deploy('Rollup')
    assertionMap = await deploy('AssertionMap')

    const rollupArgs = [
      await accounts[1].getAddress(), // roll up owner
      verifier.address, // verifier
      token.address, // stake token
      addressManager.address, // address manager
      assertionMap.address, // assertionMap
      0, // confirmation period
      0, // challenge period
      0, // minimum assertion period
      100, // baseStakeAmount
      '0x0000000000000000000000000000000000000000000000000000000000000000', // initialVMhash
    ]

    const callData = rollUpImp.interface.encodeFunctionData(
      'initialize',
      rollupArgs
    )
    const rollUpProxy = await deploy('TransparentUpgradeableProxy', {
      args: [
        rollUpImp.address, // logic
        await accounts[2].getAddress(), // admin
        callData, // call data
      ],
    })

    expect(await assertionMap.rollupAddress()).to.eq(rollUpProxy.address)

    rollUp = new Contract(rollUpProxy.address, rollUpImp.interface, accounts[0])
    expect(await rollUp.owner()).to.eq(rollupArgs[0])
    expect(await rollUp.verifier()).to.eq(rollupArgs[1])
    expect(await rollUp.stakeToken()).to.eq(rollupArgs[2])
    expect(await rollUp.libAddressManager()).to.eq(rollupArgs[3])
    expect(await rollUp.assertions()).to.eq(rollupArgs[4])
    expect(await rollUp.confirmationPeriod()).to.eq(rollupArgs[5])
    expect(await rollUp.challengePeriod()).to.eq(rollupArgs[6])
    expect(await rollUp.minimumAssertionPeriod()).to.eq(rollupArgs[7])
    expect(await rollUp.baseStakeAmount()).to.eq(rollupArgs[8])
  }

  const deployVerifier = async () => {
    verifier = await deploy('VerifierEntry')
  }

  const deployAddressManager = async () => {
    addressManager = await deploy('Lib_AddressManager')
    await addressManager.setAddress(
      'BVM_Rolluper',
      await accounts[0].getAddress()
    )
  }

  const deploySCC = async () => {
    const scc = await deploy('MockStateCommitmentChain')
    await addressManager.setAddress('StateCommitmentChain', scc.address)
  }
})

import { Signer, Wallet, BytesLike, utils, Contract, BigNumber } from "ethers"
import chai from "chai"
import { deploy } from '../../../helpers'

const { expect } = chai
const { ethers, waffle } = require("hardhat")

describe('StakingSlashing', () => {
  let accounts: Signer[]
  let tssGroup: Contract
  let stakingSlashing: Contract
  let bitToken: Contract
  let myWallet: Wallet
  let tssNodes: Wallet[] = []
  let newTssNodes: Wallet[] = []
  let cpk: Wallet
  let testTokenAddress = "0x1000000000000000000000000000000000000000"
  let testGroupAddress = "0x1000000000000000000000000000000000000001"

  before('deploy stakingSlashing contracts', async () => {
    accounts = await ethers.getSigners()
    await deployBitToken()
    await deployTssGroup()
    await deployStakingSlashing()
    await initAccount()
  })

  it("setAddress test", async () => {
    await stakingSlashing.setAddress(testTokenAddress, testGroupAddress)
    expect(await stakingSlashing.BitToken()).to.eq(testTokenAddress)
    expect(await stakingSlashing.tssGroupContract()).to.eq(testGroupAddress)
    await stakingSlashing.setAddress(bitToken.address, tssGroup.address)
    expect(await stakingSlashing.BitToken()).to.eq(bitToken.address)
    expect(await stakingSlashing.tssGroupContract()).to.eq(tssGroup.address)
  })

  it("setSlashingParams", async () => {
    await expect(stakingSlashing.setSlashingParams([1000, 100], [1, 2])).to.be.revertedWith("invalid param slashAmount, animus <= uptime")
    await expect(stakingSlashing.setSlashingParams([100, 1000], [2, 1])).to.be.revertedWith("invalid param exIncome, animus <= uptime")

    await expect(stakingSlashing.setSlashingParams([100, 1000], [101, 999])).to.be.revertedWith("slashAmount need bigger than exIncome")
    await expect(stakingSlashing.setSlashingParams([100, 1000], [99, 10001])).to.be.revertedWith("slashAmount need bigger than exIncome")
    await expect(stakingSlashing.setSlashingParams([0, 100], [0, 99])).to.be.revertedWith("invalid amount")

    let slashAmounts = [100, 1000]
    let exIncomes = [10, 100]
    await stakingSlashing.setSlashingParams(slashAmounts, exIncomes)
    let res = await stakingSlashing.getSlashingParams()
    expect(res[0].toString()).to.eq(slashAmounts.toString())
    expect(res[1].toString()).to.eq(exIncomes.toString())
  })

  it("staking", async () => {
    let pubKey = "0x" + myWallet.publicKey.substring(4)
    await bitToken.mint(myWallet.address, 100000)
    expect(await bitToken.balanceOf(myWallet.address)).to.eq(100000)
    await bitToken.connect(myWallet).approve(stakingSlashing.address, 10000)

    // error case amount invalid
    await expect(stakingSlashing.connect(myWallet).staking(0, pubKey)).to.be.revertedWith("invalid amount")
    // error case public key invalid : length 65
    await expect(stakingSlashing.connect(myWallet).staking(1000, myWallet.publicKey)).to.be.revertedWith("public key length must 64 bytes")
    // error case public key invalid : invalid equal
    await expect(stakingSlashing.connect(myWallet).staking(1000, "0x1" + myWallet.publicKey.substring(5))).to.be.revertedWith("invalid pubKey")

    // staking 1000 for myWallet
    await stakingSlashing.connect(myWallet).staking(10000, pubKey)
    expect(await bitToken.balanceOf(myWallet.address)).to.eq(90000)
    // check
    let deposit = await stakingSlashing.getDeposits((await myWallet.getAddress()).toString())
    expect(deposit.pubKey).to.eq(pubKey)
    expect(deposit.amount.toNumber()).to.eq(10000)

    // error case public key invalid : not equal
    await expect(stakingSlashing.connect(myWallet).staking(1000, myWallet.publicKey)).to.be.revertedWith("pubKey not equal")
    // error case erc20 transfer failed : need approve first
    await expect(stakingSlashing.connect(myWallet).staking(10000, pubKey)).to.be.revertedWith("ERC20: transfer amount exceeds allowance")

    // add staking for myWallet
    await bitToken.connect(myWallet).approve(stakingSlashing.address, 10000)
    await stakingSlashing.connect(myWallet).staking(10000, pubKey)
    expect(await bitToken.balanceOf(myWallet.address)).to.eq(80000)
    // check
    deposit = await stakingSlashing.getDeposits((await myWallet.getAddress()).toString())
    expect(deposit.pubKey).to.eq(pubKey)
    expect(deposit.amount.toNumber()).to.eq(20000)
  })

  it("withdrawToken", async () => {
    await initTssGroup()
    await expect(stakingSlashing.withdrawToken()).to.be.revertedWith("do not have deposit")
    let pubKey = "0x" + tssNodes[0].publicKey.substring(4)
    expect(await tssGroup.memberExistActive(pubKey)).to.eq(true)
    expect(await tssGroup.memberExistInActive(pubKey)).to.eq(false)
    await expect(stakingSlashing.connect(tssNodes[0]).withdrawToken()).to.be.revertedWith("not at the right tim")

    // withdraw deposit token
    await stakingSlashing.connect(myWallet).withdrawToken()
    let deposit = await stakingSlashing.getDeposits((await myWallet.getAddress()).toString())
    expect(deposit.pubKey.toString()).to.eq("0x")
    expect(deposit.amount.toNumber()).to.eq(0)
    // check balance
    expect(await bitToken.balanceOf(myWallet.address)).to.eq(100000)
  })

  it("batch GetDeposits test", async () => {
    let address = [tssNodes[0].address, tssNodes[1].address, tssNodes[2].address]
    let deposits = await stakingSlashing.batchGetDeposits(address)

    for (let i = 0; i < deposits.length; i++) {
      let pubKey = "0x" + tssNodes[i].publicKey.substring(4)
      expect(deposits[i].pledgor).to.eq(tssNodes[i].address)
      expect(deposits[i].pubKey).to.eq(pubKey)
      expect(deposits[i].amount).to.eq(BigNumber.from(20000))
    }
  })

  it("quitRequest", async () => {
    await stakingSlashing.connect(tssNodes[1]).quitRequest()
    let quitList = await stakingSlashing.getQuitRequestList()
    expect(quitList[0]).to.eq(tssNodes[1].address)
    await expect(stakingSlashing.connect(tssNodes[1]).quitRequest()).to.be.revertedWith("already in quitRequestList")

    let tssNodesPubKey: BytesLike[] = []
    // tssnodes staking first
    for (let i = 0; i < newTssNodes.length; i++) {
      // approve bit tokens
      let pubKey = "0x" + newTssNodes[i].publicKey.substring(4)
      await bitToken.mint(newTssNodes[i].address, 20000)
      expect(await bitToken.balanceOf(newTssNodes[i].address)).to.eq(20000)
      await bitToken.connect(newTssNodes[i]).approve(stakingSlashing.address, 20000)
      // staking
      // console.log(pubKey)
      await stakingSlashing.connect(newTssNodes[i]).staking(20000, pubKey)
      // check
      let deposit = await stakingSlashing.getDeposits((await newTssNodes[i].address).toString())
      expect(deposit.pubKey).to.eq(pubKey)
      expect(deposit.amount.toNumber()).to.eq(20000)
      tssNodesPubKey[i] = pubKey
    }

    await expect(stakingSlashing.connect(newTssNodes[1]).quitRequest()).to.be.revertedWith("not at the inactive group or active group")

    // set inactive tss group members
    await tssGroup.setTssGroupMember(
      4,
      tssNodesPubKey
    )
    let info = await tssGroup.getTssGroupInfo()
    expect(info[0]).to.eq(1)
    expect(info[1]).to.eq(4)
    await stakingSlashing.connect(tssNodes[2]).quitRequest()
  })

  it("slashing", async () => {
    // slash uptime
    let signers = [tssNodes[1].address, tssNodes[2].address, tssNodes[3].address, tssNodes[4].address]
    let message = utils.defaultAbiCoder.encode(
      ["tuple(uint256,address,address[],uint256)"],
      [[
        1,
        tssNodes[0].address,
        signers,
        1
      ]]
    )
    let messageHash = ethers.utils.solidityKeccak256(['bytes'], [message]);
    let signature = await cpk.signMessage(ethers.utils.arrayify(messageHash));

    // slash case uptime : deposit > slashAmount
    await stakingSlashing.connect(tssNodes[1]).slashing(message, signature)
    // check status
    let info = await tssGroup.getTssMember("0x" + tssNodes[0].publicKey.substring(4))
    expect(info.publicKey).to.eq("0x" + tssNodes[0].publicKey.substring(4))
    expect(info.nodeAddress).to.eq(tssNodes[0].address)
    expect(info.status).to.eq(1)
    // check deposit
    let deposit = await stakingSlashing.getDeposits(tssNodes[0].address)
    let deductedAmount = 100
    let exIncome = 10
    let extraAmount = deductedAmount - exIncome
    let remainder = extraAmount % signers.length
    let gain = (extraAmount - remainder) / signers.length

    // check slasher`s deposit
    expect(deposit.amount).to.eq(20000 - deductedAmount)
    deposit = await stakingSlashing.getDeposits(signers[0])
    // check sender`s deposit
    expect(deposit.amount).to.eq(20000 + gain + remainder + exIncome)
    // check signer`s deposit
    for (let i = 1; i < signers.length; i++) {
      deposit = await stakingSlashing.getDeposits(signers[i])
      expect(deposit.amount).to.eq(20000 + gain)
    }

    // unjail
    await stakingSlashing.connect(tssNodes[0]).unJail()
    info = await tssGroup.getTssMember("0x" + tssNodes[0].publicKey.substring(4))
    expect(info.publicKey).to.eq("0x" + tssNodes[0].publicKey.substring(4))
    expect(info.nodeAddress).to.eq(tssNodes[0].address)
    expect(info.status).to.eq(0)

    // err case slash animus : already slashed
    message = utils.defaultAbiCoder.encode(
      ["tuple(uint256,address,address[],uint256)"],
      [[
        1,
        tssNodes[0].address,
        signers,
        2
      ]]
    )
    messageHash = ethers.utils.solidityKeccak256(['bytes'], [message])
    signature = await cpk.signMessage(ethers.utils.arrayify(messageHash))
    await expect(stakingSlashing.connect(tssNodes[1]).slashing(message, signature)).to.be.revertedWith("already slashed")

    // slash case animus
    message = utils.defaultAbiCoder.encode(
      ["tuple(uint256,address,address[],uint256)"],
      [[
        2,
        tssNodes[0].address,
        signers,
        2
      ]]
    )
    messageHash = ethers.utils.solidityKeccak256(['bytes'], [message]);
    signature = await cpk.signMessage(ethers.utils.arrayify(messageHash));
    await stakingSlashing.connect(tssNodes[1]).slashing(message, signature)

    let tssMemberInfo = await tssGroup.getTssGroupInfo()
    expect(tssMemberInfo[3].length).to.eq(4)

    // slash case nothing
    message = utils.defaultAbiCoder.encode(
      ["tuple(uint256,address,address[],uint256)"],
      [[
        3,
        tssNodes[2].address,
        signers,
        0
      ]]
    )
    messageHash = ethers.utils.solidityKeccak256(['bytes'], [message]);
    signature = await cpk.signMessage(ethers.utils.arrayify(messageHash));
    await expect(stakingSlashing.connect(tssNodes[1]).slashing(message, signature)).to.be.revertedWith("err type for slashing")
  })

  const deployBitToken = async () => {
    bitToken = await deploy('TestERC20')
  }

  const deployStakingSlashing = async () => {
    stakingSlashing = await deploy('TssStakingSlashing')
    await stakingSlashing.initialize(bitToken.address, tssGroup.address)
    await tssGroup.setStakingSlash(stakingSlashing.address)
  }

  const deployTssGroup = async () => {
    tssGroup = await deploy('TssGroupManager')
    await tssGroup.initialize()
  }

  const initAccount = async () => {
    const provider = waffle.provider;
    cpk = new Wallet("604adb683b14896ecc16eabead607e9517a4ca0eb6946ce537fdeaa49cdcf3e0", provider)
    myWallet = new Wallet("804adb683b14896ecc16eabead607e9517a4ca0eb6946ce537fdeaa49cdcf3e6", provider)
    await accounts[0].sendTransaction({
      to: myWallet.address,
      value: ethers.utils.parseEther("1")
    })
    let tssNodesPriviteKeys = [
      "104adb683b14896ecc16eabead607e9517a4ca0eb6946ce537fdeaa49cdcf3e1",
      "204adb683b14896ecc16eabead607e9517a4ca0eb6946ce537fdeaa49cdcf3e2",
      "304adb683b14896ecc16eabead607e9517a4ca0eb6946ce537fdeaa49cdcf3e3",
      "404adb683b14896ecc16eabead607e9517a4ca0eb6946ce537fdeaa49cdcf3e4",
      "504adb683b14896ecc16eabead607e9517a4ca0eb6946ce537fdeaa49cdcf3e5",
    ]
    let newTssNodesPriviteKeys = [
      "114adb683b14896ecc16eabead607e9517a4ca0eb6946ce537fdeaa49cdcf3e1",
      "214adb683b14896ecc16eabead607e9517a4ca0eb6946ce537fdeaa49cdcf3e2",
      "314adb683b14896ecc16eabead607e9517a4ca0eb6946ce537fdeaa49cdcf3e3",
      "414adb683b14896ecc16eabead607e9517a4ca0eb6946ce537fdeaa49cdcf3e4",
      "514adb683b14896ecc16eabead607e9517a4ca0eb6946ce537fdeaa49cdcf3e5",
    ]
    for (let i = 0; i < tssNodesPriviteKeys.length; i++) {
      tssNodes[i] = new Wallet(tssNodesPriviteKeys[i], provider)
      await accounts[0].sendTransaction({
        to: tssNodes[i].getAddress(),
        value: ethers.utils.parseEther("1")
      })
    }
    for (let i = 0; i < newTssNodesPriviteKeys.length; i++) {
      newTssNodes[i] = new Wallet(newTssNodesPriviteKeys[i], provider)
      await accounts[0].sendTransaction({
        to: newTssNodes[i].getAddress(),
        value: ethers.utils.parseEther("1")
      })
    }
  }

  const initTssGroup = async () => {
    let tssNodesPubKey: BytesLike[] = []
    // tssnodes staking first
    for (let i = 0; i < tssNodes.length; i++) {
      // approve bit tokens
      let pubKey = "0x" + tssNodes[i].publicKey.substring(4)
      await bitToken.mint(tssNodes[i].address, 20000)
      expect(await bitToken.balanceOf(tssNodes[i].address)).to.eq(20000)
      await bitToken.connect(tssNodes[i]).approve(stakingSlashing.address, 20000)
      // staking
      await stakingSlashing.connect(tssNodes[i]).staking(20000, pubKey)
      // check
      let deposit = await stakingSlashing.getDeposits((await tssNodes[i].address).toString())
      expect(deposit.pubKey).to.eq(pubKey)
      expect(deposit.amount.toNumber()).to.eq(20000)
      tssNodesPubKey[i] = pubKey
    }
    let testGroupKey = "0x" + cpk.publicKey.substring(4)
    // set inactive tss group members
    await tssGroup.setTssGroupMember(
      3,
      tssNodesPubKey
    )

    // commmit their group pub key
    for (let i = 0; i < tssNodes.length; i++) {
      await tssGroup.connect(tssNodes[i]).setGroupPublicKey(
        tssNodesPubKey[i],
        testGroupKey
      )
    }

    let info = await tssGroup.getTssGroupInfo()
    expect(info[0]).to.eq(1)
    expect(info[1]).to.eq(3)
    expect(info[2]).to.eq(testGroupKey)
    expect(info[3].length).to.eq(tssNodesPubKey.length)
  }
})



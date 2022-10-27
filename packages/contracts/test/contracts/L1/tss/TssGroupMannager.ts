import { Signer, Wallet, BytesLike, Contract } from "ethers"
import chai from "chai"
import { deploy } from '../../../helpers'

const { expect } = chai
const { ethers, waffle } = require("hardhat")

describe('TssGroupManager', () => {
  let accounts: Signer[]
  let tssGroup: Contract
  let myWallet: Wallet
  let tssNodes: Wallet[] = []
  let newTssNodes: Wallet[] = []
  let cpk: Wallet
  let tssNodesPublicKeys: BytesLike[] = []
  let cpkPubKey: BytesLike
  let addressManager: Contract

  before('deploy tssgroup contracts', async () => {
    accounts = await ethers.getSigners()
    await deployAddressManager()
    await deployTssGroup()
    await initAccount()
  })

  it(`recover`, async () => {
    const digestHash = '0x8a23d321039f259c854d8746210f87f8d4ab7be011e3b14050ba1998e559a303'
    const sig = '0x10573d194d947211720a9294f2d47a4bd4a97926f040ef34e034d75d12fd19551c3dca19ee6f0cf4f0ecd3b4663972b9afef12691de522214d1fd543dc0ce9cc01'
    const expected = '0xB3ffC97f0AcEe8e1D85D42952d81499777746422'
    const actual = await tssGroup.recover(digestHash, sig)
    expect(expected).to.eq(actual)
  })

  it("public key to address", async () => {
    const provider = waffle.provider;
    let myWallet = new Wallet("a1724a3be3134c9e64d9243428926088c1f4a236e777c19e3c7a974e0da6dba3", provider)
    let pubKey = "0x" + myWallet.publicKey.substring(4)
    let toAddress = await tssGroup.publicKeyToAddress(pubKey)

    expect(toAddress).to.eq(myWallet.address)
  })

  it("set TssGroupMembers", async () => {
    // set member empty public key
    await expect(
      tssGroup.setTssGroupMember(
        0,
        []
      )
    ).to.be.revertedWith('batch public key is empty')
    // set member public key
    await tssGroup.setTssGroupMember(
      3,
      tssNodesPublicKeys
    )
    const msg = await tssGroup.getTssGroupInfo()
    expect(msg[0]).to.eq(0)
    expect(msg[1]).to.eq(3)
  })

  it("set GroupPublicKey", async () => {
    // set empty group public key sender public key not Inactive Member
    await expect(
      tssGroup.connect(newTssNodes[0]).setGroupPublicKey(
        "0x" + newTssNodes[0].publicKey.substring(4),
        cpkPubKey
      )
    ).to.be.revertedWith('your public key is not in InActiveMember')
    // set empty group public key sender public not match
    await expect(
      tssGroup.connect(newTssNodes[0]).setGroupPublicKey(
        "0x" + tssNodes[1].publicKey.substring(4),
        cpkPubKey
      )
    ).to.be.revertedWith('public key not match')

    // commmit their group pub key
    for (let i = 0; i < tssNodes.length; i++) {
      await tssGroup.connect(tssNodes[i]).setGroupPublicKey(
        tssNodesPublicKeys[i],
        cpkPubKey
      )
    }

    let info = await tssGroup.getTssGroupInfo()
    expect(info[0]).to.eq(1)
    expect(info[1]).to.eq(3)
    expect(info[2]).to.eq(cpkPubKey)
    expect(info[3].length).to.eq(tssNodesPublicKeys.length)
  })

  const deployAddressManager = async () => {
    addressManager = await deploy('Lib_AddressManager')
  }

  const deployTssGroup = async () => {
    tssGroup = await deploy('TssGroupManager')
    await tssGroup.initialize()
    await addressManager.setAddress(
      'TssGroupManager',
      tssGroup.address
    )
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
    cpkPubKey = "0x" + cpk.publicKey.substring(4)
    // tssnodes staking first
    for (let i = 0; i < tssNodes.length; i++) {
      // approve bit tokens
      let pubKey = "0x" + tssNodes[i].publicKey.substring(4)
      tssNodesPublicKeys[i] = pubKey
    }
  }
})

/* External Imports */
import { ethers } from 'hardhat'
import { Signer, ContractFactory, Contract } from 'ethers'
import {
  smock,
  MockContractFactory,
  MockContract,
} from '@defi-wonderland/smock'

/* Internal Imports */
import { expect } from '../../setup'

const DUMMY_L2_BRIDGE_ADDRESS: string = ethers.utils.getAddress(
  '0x' + 'acdc'.repeat(10)
)

describe('BitnetworkMintableERC721Factory', () => {
  let signer: Signer
  let Factory__L1ERC721: MockContractFactory<ContractFactory>
  let L1ERC721: MockContract<Contract>
  let BitnetworkMintableERC721Factory: Contract
  let baseURI: string
  const remoteChainId = 100

  beforeEach(async () => {
    ;[signer] = await ethers.getSigners()

    // deploy an ERC721 contract on L1
    Factory__L1ERC721 = await smock.mock(
      '@openzeppelin/contracts/token/ERC721/ERC721.sol:ERC721'
    )
    L1ERC721 = await Factory__L1ERC721.deploy('L1ERC721', 'ERC')

    BitnetworkMintableERC721Factory = await (
      await ethers.getContractFactory('BitnetworkMintableERC721Factory')
    ).deploy(DUMMY_L2_BRIDGE_ADDRESS, remoteChainId)

    baseURI = ''.concat(
      'ethereum:',
      L1ERC721.address.toLowerCase(),
      '@',
      remoteChainId.toString(),
      '/tokenURI?uint256='
    )
  })

  it('should be deployed with the correct constructor argument', async () => {
    expect(await BitnetworkMintableERC721Factory.bridge()).to.equal(
      DUMMY_L2_BRIDGE_ADDRESS
    )
  })

  it('should be able to create a standard ERC721 contract', async () => {
    const tx =
      await BitnetworkMintableERC721Factory.createStandardBitnetworkMintableERC721(
        L1ERC721.address,
        'L2ERC721',
        'ERC'
      )
    const receipt = await tx.wait()

    // Get the BitnetworkMintableERC721Created event
    const erc721CreatedEvent = receipt.events[0]

    // Expect there to be an event emitted for the standard token creation
    expect(erc721CreatedEvent.event).to.be.eq('BitnetworkMintableERC721Created')

    // Get the L2 ERC721 address from the emitted event and check it was created correctly
    const l2ERC721Address = erc721CreatedEvent.args.localToken
    const BitnetworkMintableERC721 = new Contract(
      l2ERC721Address,
      (await ethers.getContractFactory('BitnetworkMintableERC721')).interface,
      signer
    )

    expect(await BitnetworkMintableERC721.bridge()).to.equal(
      DUMMY_L2_BRIDGE_ADDRESS
    )
    expect(await BitnetworkMintableERC721.remoteToken()).to.equal(
      L1ERC721.address
    )
    expect(await BitnetworkMintableERC721.name()).to.equal('L2ERC721')
    expect(await BitnetworkMintableERC721.symbol()).to.equal('ERC')
    expect(await BitnetworkMintableERC721.baseTokenURI()).to.equal(baseURI)

    expect(
      await BitnetworkMintableERC721Factory.isStandardBitnetworkMintableERC721(
        BitnetworkMintableERC721.address
      )
    ).to.equal(true)
  })

  it('should not be able to create a standard token with a 0 address for l1 token', async () => {
    await expect(
      BitnetworkMintableERC721Factory.createStandardBitnetworkMintableERC721(
        ethers.constants.AddressZero,
        'L2ERC721',
        'ERC'
      )
    ).to.be.revertedWith(
      'BitnetworkMintableERC721Factory: L1 token address cannot be address(0)'
    )
  })
})

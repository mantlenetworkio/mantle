/* External Imports */
import { ethers } from 'hardhat'
import { ContractFactory, Contract } from 'ethers'
import { smock, FakeContract } from '@defi-wonderland/smock'
import { remove0x } from '@mantleio/core-utils'
import { keccak256 } from 'ethers/lib/utils'

/* Internal Imports */
import { expect } from '../../../setup'
import { NON_ZERO_ADDRESS } from '../../../helpers/constants'

const ELEMENT_TEST_SIZES = [1, 2, 4, 8, 16]

const callPredeploy = async (
  Helper_PredeployCaller: Contract,
  predeploy: Contract,
  functionName: string,
  functionParams?: any[]
): Promise<any> => {
  return Helper_PredeployCaller.callPredeploy(
    predeploy.address,
    predeploy.interface.encodeFunctionData(functionName, functionParams || [])
  )
}

// TODO: rewrite this test to bypass the execution manager
describe.skip('BVM_L2ToL1MessagePasser', () => {
  let Fake__BVM_ExecutionManager: FakeContract
  before(async () => {
    Fake__BVM_ExecutionManager = await smock.fake<Contract>(
      'BVM_ExecutionManager'
    )
  })

  let Helper_PredeployCaller: Contract
  before(async () => {
    Helper_PredeployCaller = await (
      await ethers.getContractFactory('Helper_PredeployCaller')
    ).deploy()

    Helper_PredeployCaller.setTarget(Fake__BVM_ExecutionManager.address)
  })

  let Factory__BVM_L2ToL1MessagePasser: ContractFactory
  before(async () => {
    Factory__BVM_L2ToL1MessagePasser = await ethers.getContractFactory(
      'BVM_L2ToL1MessagePasser'
    )
  })

  let BVM_L2ToL1MessagePasser: Contract
  beforeEach(async () => {
    BVM_L2ToL1MessagePasser = await Factory__BVM_L2ToL1MessagePasser.deploy()
  })

  describe('passMessageToL1', () => {
    before(async () => {
      Fake__BVM_ExecutionManager.bvmCALLER.returns(NON_ZERO_ADDRESS)
    })

    for (const size of ELEMENT_TEST_SIZES) {
      it(`should be able to pass ${size} messages`, async () => {
        for (let i = 0; i < size; i++) {
          const message = '0x' + '12' + '34'.repeat(i)

          await callPredeploy(
            Helper_PredeployCaller,
            BVM_L2ToL1MessagePasser,
            'passMessageToL1',
            [message]
          )

          expect(
            await BVM_L2ToL1MessagePasser.sentMessages(
              keccak256(message + remove0x(Helper_PredeployCaller.address))
            )
          ).to.equal(true)
        }
      })
    }
  })
})

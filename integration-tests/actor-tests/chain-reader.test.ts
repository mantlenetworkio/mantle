import { actor, run, setupActor } from './lib/convenience'
import { MantleEnv } from '../test/shared/env'

actor('Chain reader', () => {
  let env: MantleEnv

  setupActor(async () => {
    env = await MantleEnv.new()
  })

  run(async (b) => {
    const blockNumber = await b.bench('get block number', () =>
      env.l2Provider.getBlockNumber()
    )
    await b.bench('get random block', () =>
      env.l2Provider.getBlock(Math.floor(blockNumber * Math.random()))
    )
  })
})

import { actor, run, setupActor } from './lib/convenience'
import { BitnetworkEnv } from '../test/shared/env'

actor('Chain reader', () => {
  let env: BitnetworkEnv

  setupActor(async () => {
    env = await BitnetworkEnv.new()
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

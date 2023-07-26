import { ethers } from 'ethers'

import { CrossChainMessenger } from './src';
import {sleep} from "@mantleio/mt-core-utils";

(async () => {
  const l1RpcProvider = new ethers.providers.JsonRpcProvider(process.env.L1_URL)
  // @ts-ignore
  const l1Wallet = new ethers.Wallet(process.env.PRIVATE_KEY, l1RpcProvider)
  const etherBalance =ethers.utils.formatEther(await l1Wallet.getBalance());

  const tokenAddress = process.env.TestMantleToken_ADDRESS;
  const tokenAbi = ["function balanceOf(address account) view returns (uint256)"];
  const tokenContract = new ethers.Contract(tokenAddress!, tokenAbi, l1Wallet);
  const balance = await tokenContract.balanceOf(l1Wallet.address);
  const balanceFormatted = ethers.utils.formatUnits(balance, 18); // 18 是代币的小数位数

  const l2RpcProvider = new ethers.providers.JsonRpcProvider(process.env.L2_URL)
  // @ts-ignore
  const l2Wallet = new ethers.Wallet(process.env.PRIVATE_KEY, l2RpcProvider)

  let withdrawETHtx: any

  while (true) {
    try {
      withdrawETHtx = await l2RpcProvider.getTransaction(
        process.env.TX_HASH || ''
      )
      if (withdrawETHtx.blockNumber < 1) {
        console.log("重新查询 withdrawETHtx, 等待 10 秒! ")
        await wait(10);
        continue
      } else {
        break;
      }
    } catch (error) {
      // 捕获到异常情况
      console.error('请求异常:', error);
    }
  }

  console.log("等待通过 FRAUD_PROOF_WINDOW: ", process.env.FRAUD_PROOF_WINDOW, "秒! ");
  for (let i = 0; i < Math.floor(Number(process.env.FRAUD_PROOF_WINDOW!) / 10) + 1; i++) {
    await wait(10);
    console.log("已等待 10 秒.");
  }
  
  const crossChainMessenger = new CrossChainMessenger({
    l1ChainId: process.env.L1_CHAINID!,
    l2ChainId: process.env.L2_CHAINID!,
    l1SignerOrProvider: l1Wallet,
    l2SignerOrProvider: l2Wallet,
  })

  let proof:any

  while (true) {
    try {
      proof = await crossChainMessenger.getMessageProof(withdrawETHtx)

      if (proof.stateRoot.length == 66){
        console.log("proof.stateRoot:  ", proof.stateRoot)
        break;
      }
    } catch (error) {
      console.error('请求 proof 异常:', error.message);
      console.log("等待 10 秒后, 再次请求...");
      await wait(10);
      continue
    }
  }

  console.log("查询到 proof 之后继续等待 30 秒。");
  await wait(30);

  let finalizeMessageResponse:any

  while (true) {
    try {
      finalizeMessageResponse = await crossChainMessenger.finalizeMessage(
        withdrawETHtx
      )

      if (finalizeMessageResponse.hash.length ==66 ){
        console.log("finalizeMessageResponse txHash: ", finalizeMessageResponse.hash)
        break;
      }
    } catch (error) {
      const err1s = error.toString().split('error={"reason":"');
      const err2s = err1s[1].split('"')

      console.error('请求 finalizeMessage 异常:', err2s[0]);

      console.log("等待 10 秒后, 再次请求...");
      await wait(10);
      continue
    }
  }

})()

function wait(seconds: number): Promise<void> {
  return new Promise((resolve) => {
    setTimeout(() => {
      resolve();
    }, seconds * 1000);
  });
}
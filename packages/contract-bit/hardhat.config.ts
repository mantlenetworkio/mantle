import {HardhatUserConfig} from 'hardhat/types';
import 'hardhat-deploy';
import 'hardhat-deploy-ethers';

const config: HardhatUserConfig = {
  solidity: {
    version: '0.8.9',
  },
  networks: {
    "local": {
      //172.17.0.1: the default docker bridge ip
      url: 'http://172.17.0.1:9545',
      accounts: {mnemonic: 'test test test test test test test test test test test junk' }
    },
  },
  namedAccounts: {
    deployer: 0,
    tokenOwner: 1,
  },
  paths: {
    sources: 'src',
  },
};
export default config;

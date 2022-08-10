module.exports = {
  skipFiles: [
    './test-helpers',
    './test-libraries',
    './L2/predeploys/BVM_DeployerWhitelist.sol'
  ],
  mocha: {
    grep: "@skip-on-coverage",
    invert: true
  }
};

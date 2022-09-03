# Tutorial for hardhat-deploy based on https://hardhat.org/tutorial/

# 1. Setting up the environment

Most Ethereum libraries and tools are written in JavaScript, and so is **Hardhat**. If you're not familiar with Node.js, it's a JavaScript runtime built on Chrome's V8 JavaScript engine. It's the most popular solution to run JavaScript outside of a web browser and **Hardhat** is built on top of it.

## Installing Node.js

You can [skip](./creating-a-new-hardhat-project.md) this section if you already have a working Node.js `>=12.0` installation. If not, here's how to install it on Ubuntu, MacOS and Windows.


### Linux

#### Ubuntu

Copy and paste these commands in a terminal:

```
sudo apt update
sudo apt install curl git
curl -sL https://deb.nodesource.com/setup_12.x | sudo -E bash -
sudo apt install nodejs
```

### MacOS

Make sure you have `git` installed. Otherwise, follow [these instructions](https://www.atlassian.com/git/tutorials/install-git).

There are multiple ways of installing Node.js on MacOS. We will be using [Node Version Manager (nvm)](http://github.com/creationix/nvm).

Copy and paste these commands in a terminal:

```
curl -o- https://raw.githubusercontent.com/creationix/nvm/v0.35.2/install.sh | bash
nvm install 12
nvm use 12
nvm alias default 12
npm install npm --global # Upgrade npm to the latest version
```

### Windows

Installing Node.js on Windows requires a few manual steps. We'll install git, Node.js 12.x and npm.

Download and run these:
1. [Git's installer for Windows](https://git-scm.com/download/win)
2. `node-v12.XX.XX-x64.msi` from [here](https://nodejs.org/dist/latest-v12.x)


## Upgrading your Node.js installation

If your version of Node.js is older than `12.0` follow the instructions below to upgrade.

### Linux

#### Ubuntu

1. Run `sudo apt remove nodejs` in a terminal to remove Node.js.
2. Find the version of Node.js that you want to install [here](https://github.com/nodesource/distributions#debinstall) and follow the instructions.
3. Run `sudo apt update && sudo apt install nodejs` in a terminal to install Node.js again.

### MacOS

You can change your Node.js version using [nvm](http://github.com/creationix/nvm). To upgrade to Node.js `12.x` run these in a terminal:

```
nvm install 12
nvm use 12
nvm alias default 12
npm install npm --global # Upgrade npm to the latest version
```

### Windows

You need to follow the [same installation instructions](#windows) as before but choose a different version. You can check the list of all available versions [here](https://nodejs.org/en/download/releases/).

## Installing yarn

For this tutorial we are going to use [yarn](yarnpkg.com)

To install it do the following:

```
npm install -g yarn
```


# 2. Creating a new Hardhat project

We'll install **Hardhat** using the npm CLI. The **N**ode.js **p**ackage **m**anager is a package manager and an online repository for JavaScript code.

Open a new terminal and run these commands:

```
mkdir hardhat-deploy-tutorial
cd hardhat-deploy-tutorial
yarn init --yes
yarn add -D hardhat
```

::: tip
Installing **Hardhat** will install some Ethereum JavaScript dependencies, so be patient.
:::

In the same directory where you installed **Hardhat** add a `hardhat.config.ts` (we are going to use typescript and the Solidity 0.7.6 compiler)

```typescript
import {HardhatUserConfig} from 'hardhat/types';
const config: HardhatUserConfig = {
  solidity: {
    version: '0.7.6',
  }
};
export default config;

```

## Hardhat's architecture

**Hardhat** is designed around the concepts of **tasks** and **plugins**. The bulk of **Hardhat**'s functionality comes from plugins, which as a developer [you're free to choose](/plugins/) the ones you want to use.

### Tasks
Every time you run **Hardhat** from the CLI you're running a task. e.g. `npx hardhat compile` is running the `compile` task. To see the currently available tasks in your project, run `npx hardhat`. Feel free to explore any task by running `npx hardhat help [task]`.

::: tip
You can create your own tasks. Check out the [Creating a task](/guides/create-task.md) guide.
:::

### Plugins
**Hardhat** is unopinionated in terms of what tools you end up using, but it does come with some built-in defaults, all of which can be overriden. Most of the time the way to use a given tool is by consuming a plugin that integrates it into **Hardhat**.

For this tutorial we are going to use the `hardhat-deploy` and `hardhat-deploy-ethers` plugins. They'll allow you to interact with Ethereum and to test your contracts. We'll explain how they're used later on. We also install `ethers`, `chai`, `mocha` and `typescript` and extra dependencies. To install them, run the following command in your project directory:

```
yarn add -D hardhat-deploy hardhat-deploy-ethers ethers chai chai-ethers mocha @types/chai @types/mocha @types/node typescript ts-node dotenv
```

Edit `hardhat.config.ts` so that it looks like this:

```typescript {1}
import {HardhatUserConfig} from 'hardhat/types';
import 'hardhat-deploy';
import 'hardhat-deploy-ethers';

const config: HardhatUserConfig = {
  solidity: {
    version: '0.7.6',
  },
  namedAccounts: {
    deployer: 0,
  },
};
export default config;

```

We also create the following `tsconfig.json` :

```json
{
  "compilerOptions": {
    "target": "es5",
    "module": "commonjs",
    "strict": true,
    "esModuleInterop": true,
    "moduleResolution": "node",
    "forceConsistentCasingInFileNames": true,
    "outDir": "dist"
  },
  "include": [
    "hardhat.config.ts",
    "./deploy",
    "./test",
  ]
}
```


# 3. Writing and compiling smart contracts

We're going to create a simple smart contract that implements a token that can be transferred. Token contracts are most frequently used to exchange or store value. We won't go in depth into the Solidity code of the contract on this tutorial, but there's some logic we implemented that you should know:

- There is a fixed total supply of tokens that can't be changed.
- The entire supply is assigned to the address that deploys the contract.
- Anyone can receive tokens.
- Anyone with at least one token can transfer tokens.
- The token is non-divisible. You can transfer 1, 2, 3 or 37 tokens but not 2.5.

::: tip
You might have heard about ERC20, which is a token standard in Ethereum. Tokens such as DAI, USDC, MKR and ZRX follow the ERC20 standard which allows them all to be compatible with any software that can deal with ERC20 tokens. **For simplicity's sake the token we're going to build is _not_ an ERC20.**
:::

## Writing smart contracts

While by default hardhat uses `contracts` as the source folder, we prefer to change it to `src`.

You then need to edit your `hardhat.config.ts` file with the new config:

```typescript
import {HardhatUserConfig} from 'hardhat/types';
import 'hardhat-deploy';
import 'hardhat-deploy-ethers';

const config: HardhatUserConfig = {
  solidity: {
    version: '0.7.6',
  },
  namedAccounts: {
    deployer: 0,
  },
  paths: {
    sources: 'src',
  },
};
export default config;

```

Start by creating a new directory called `src` and create a file inside the directory called `Token.sol`.

Paste the code below into the file and take a minute to read the code. It's simple and it's full of comments explaining the basics of Solidity.

::: tip
To get syntax highlighting you should add Solidity support to your text editor. Just look for Solidity or Ethereum plugins. We recommend using Visual Studio Code or Sublime Text 3.
:::

```solidity
// SPDX-License-Identifier: MIT
// The line above is recommended and let you define the license of your contract
// Solidity files have to start with this pragma.
// It will be used by the Solidity compiler to validate its version.
pragma solidity ^0.7.0;


// This is the main building block for smart contracts.
contract Token {
    // Some string type variables to identify the token.
    // The `public` modifier makes a variable readable from outside the contract.
    string public name = "My Hardhat Token";
    string public symbol = "MBT";

    // The fixed amount of tokens stored in an unsigned integer type variable.
    uint256 public totalSupply = 1000000;

    // An address type variable is used to store ethereum accounts.
    address public owner;

    // A mapping is a key/value map. Here we store each account balance.
    mapping(address => uint256) balances;

    /**
     * Contract initialization.
     *
     * The `constructor` is executed only once when the contract is created.
     */
    constructor(address _owner) {
        // The totalSupply is assigned to transaction sender, which is the account
        // that is deploying the contract.
        balances[_owner] = totalSupply;
        owner = _owner;
    }

    /**
     * A function to transfer tokens.
     *
     * The `external` modifier makes a function *only* callable from outside
     * the contract.
     */
    function transfer(address to, uint256 amount) external {
        // Check if the transaction sender has enough tokens.
        // If `require`'s first argument evaluates to `false` then the
        // transaction will revert.
        require(balances[msg.sender] >= amount, "Not enough tokens");

        // Transfer the amount.
        balances[msg.sender] -= amount;
        balances[to] += amount;
    }

    /**
     * Read only function to retrieve the token balance of a given account.
     *
     * The `view` modifier indicates that it doesn't modify the contract's
     * state, which allows us to call it without executing a transaction.
     */
    function balanceOf(address account) external view returns (uint256) {
        return balances[account];
    }
}
```

::: tip
`*.sol` is used for Solidity files. We recommend matching the file name to the contract it contains, which is a common practice.
:::

## Compiling contracts

To compile the contract run `yarn hardhat compile` in your terminal. The `compile` task is one of the built-in tasks.

```
$ yarn hardhat compile
Compiling 1 file with 0.7.3
Compilation finished successfully
```

The contract has been successfully compiled and is ready to be used.

# 4. Deployment Scripts

Before you will be able to test or deploy your contract, you must set up the deployment process that can then be used both in testing as well as deployment to various live networks.
This allow you to focus on what the contracts will be in their final form, setup their parameters and dependencies, and ensure your tests are running against the exact code that will be deployed.

This also removes the need to duplicate the deployment procedures. This is made possible thanks to the `hardhat-deploy` plugin.

## Writing deployment scripts
Create a new directory called `deploy` in the project root, and in that directory create a new file called `001_deploy_token.ts`.

Let's start with the code below. We'll explain it soon, but for now paste this code into `001_deploy_token.ts`:

```typescript
import {HardhatRuntimeEnvironment} from 'hardhat/types';
import {DeployFunction} from 'hardhat-deploy/types';

const func: DeployFunction = async function (hre: HardhatRuntimeEnvironment) {
  const {deployments, getNamedAccounts} = hre;
  const {deploy} = deployments;

  const {deployer, tokenOwner} = await getNamedAccounts();

  await deploy('Token', {
    from: deployer,
    args: [tokenOwner],
    log: true,
  });
};
export default func;
func.tags = ['Token'];

```

Notice the mention of `getNamedAccounts`?

The plugin `hardhat-deploy` allows you to name your accounts. Here there are 2 named accounts:
- `deployer` will be the account used to deploy the contract.
- `tokenOwner` which is passed to the constructor of Token.sol and which will receive the initial supply.

These accounts need to be setup in hardhat.config.ts

Modifiy it so it looks like this:

```typescript
import {HardhatUserConfig} from 'hardhat/types';
import 'hardhat-deploy';
import 'hardhat-deploy-ethers';

const config: HardhatUserConfig = {
  solidity: {
    version: '0.7.6',
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

```

`deployer` was already there and is setup to use the first account (index = 0).

`tokenOwner` is the second account.

Note that instead of index you can use hard-coded addresses or even references other named accounts. You can also have different addresses based on each network. See `hardhat-deploy` documentation [here](https://github.com/wighawag/hardhat-deploy#1-namedaccounts-ability-to-name-addresses)

In your terminal, run `yarn hardhat deploy`. You should see the following output:

```
Nothing to compile
deploying "Token" (tx: 0x259d19f33819ec8d3bd994f82912aec6af1a18ec5d74303cfb28d793a10ff683)...: deployed at 0x5FbDB2315678afecb367f032d93F642f64180aa3 with 592983 gas
Done in 3.66s.
```

Your contract was deployed to the `in-memory` Hardhat network and the output indicates that deployment was successful.

We can now write tests against this contract.

First we will add comments to the deploy script above to explain each line that matters:

```typescript
import {HardhatRuntimeEnvironment} from 'hardhat/types'; // This adds the type from hardhat runtime environment.
import {DeployFunction} from 'hardhat-deploy/types'; // This adds the type that a deploy function is expected to fulfill.

const func: DeployFunction = async function (hre: HardhatRuntimeEnvironment) { // the deploy function receives the hardhat runtime env as an argument
  const {deployments, getNamedAccounts} = hre; // we get the deployments and getNamedAccounts which are provided by hardhat-deploy.
  const {deploy} = deployments; // The deployments field itself contains the deploy function.

  const {deployer, tokenOwner} = await getNamedAccounts(); // Fetch the accounts. These can be configured in hardhat.config.ts as explained above.

  await deploy('Token', { // This will create a deployment called 'Token'. By default it will look for an artifact with the same name. The 'contract' option allows you to use a different artifact.
    from: deployer, // Deployer will be performing the deployment transaction.
    args: [tokenOwner], // tokenOwner is the address used as the first argument to the Token contract's constructor.
    log: true, // Display the address and gas used in the console (not when run in test though).
  });
};
export default func;
func.tags = ['Token']; // This sets up a tag so you can execute the script on its own (and its dependencies).

```


Not as mentioned in the comment, the name of the deployed contract is set to be the same name as the contract name: `Token`. You can deploy different version of it by simply using a different name for it, like so:


```typescript
await deploy('MyToken_1', { // name of the deployed contract
  contract: 'Token', // name of the token source
  from: deployer,
  args: [tokenOwner],
  log: true,
});
```

# 5. Testing contracts

Writing automated tests when building smart contracts is of crucial importance, as your user's money is what's at stake. For this we're going to use **Hardhat Network**, a local Ethereum network designed for development that is built-in and acts as the default network in **Hardhat**. You don't need to set anything up to use it. In our tests we're going to use ethers.js to interact with the Ethereum contract we built in the previous section, and [Mocha](https://mochajs.org/) will be our test runner.

## Writing tests
Create a new directory called `test` in the project root directory and in that `test` directory, create a new file called `Test.test.ts`.

Let's start with the code below. We'll explain it shortly, but for now just paste the following code into `Test.test.ts`:

```typescript
import {expect} from "./chai-setup";

import {ethers, deployments, getNamedAccounts} from 'hardhat';

describe("Token contract", function() {
  it("Deployment should assign the total supply of tokens to the owner", async function() {
    await deployments.fixture(["Token"]);
    const {tokenOwner} = await getNamedAccounts();
    const Token = await ethers.getContract("Token");
    const ownerBalance = await Token.balanceOf(tokenOwner);
    const supply = await Token.totalSupply();
    expect(ownerBalance).to.equal(supply);
  });
});

```

We also create a new file called `chai-setup.ts` in the test folder:

```typescript
import chaiModule from 'chai';
import {chaiEthers} from 'chai-ethers';
chaiModule.use(chaiEthers);
export = chaiModule;

```

This will use chai matchers from `chai-ethers` but also allows you to easily add more.


Then in your terminal run `npx hardhat test`. You should see the following output:

```
$ npx hardhat test

  Token contract
    ✓ Deployment should assign the total supply of tokens to the owner (654ms)


  1 passing (663ms)
```

This means the test passed sucessfully. Now Let's examine each line.

```typescript
await deployments.fixture(["Token"]);
```

Remember the deploy script we wrote earlier? This line allow to execute it prior to the test. It also generates an evm_snapshot automatically so if you write many tests, and they all refer to that fixture, the deployment will not be reexecuted. Indeed, behind the scene it does not redeploy it again and again, instead it automatically reverts to a previous state, speeding up your tests significantly!


```typescript
const {tokenOwner} = await getNamedAccounts();
```

This gives you access to the tokenOwner address, the same address that was used in the deploy script.


```typescript
const Token = await ethers.getContract("Token");
```

Since we already ran the deploy script, we can easily access the deployed contract by name. This is what this line does, and thanks to `hardhat-deploy-ethers` plugin, you get an ethers contract ready to be invoked. If you needed that contract to be associated to a specific signer, you can pass the address as the extra argument like `const TokenAsOwner = await ethers.getContract('Token', tokenOwner);`


```typescript
const ownerBalance = await Token.balanceOf(tokenOwner);
```

Now we can call contract methods on `Token`. To get the balance of the owner account, we can call `balanceOf()`.

```typescript
const supply = await Token.totalSupply();
```

Here we will again use our `Contract` instance to call a smart contract function. `totalSupply()` returns the token's supply amount.

```typescript
expect(ownerBalance).to.equal(supply);
```

Finally we're checking that it's equal to `ownerBalance`, which it should be.


To do this we're using [Chai](https://www.chaijs.com/) which is an assertions library. These assertion functions are called "matchers", and the ones we're using here actually come from `chai-ethers` npm package (which itself is a fork of [Waffle chai matchers](https://getwaffle.io/) without unecessary dependencies).

### Using a different account

 While testing your code, you may need to send a transaction from an account other than the default one. To do this you can use the second argument to `getContract` :

```typescript
import {expect} from "./chai-setup";

import {ethers, deployments, getNamedAccounts, getUnnamedAccounts} from 'hardhat';

describe("Token contract", function() {
  it("Deployment should assign the total supply of tokens to the owner", async function() {
    await deployments.fixture(["Token"]);
    const {tokenOwner} = await getNamedAccounts();
    const users = await getUnnamedAccounts();
    const TokenAsOwner = await ethers.getContract("Token", tokenOwner);
    await TokenAsOwner.transfer(users[0], 50);
    expect(await TokenAsOwner.balanceOf(users[0])).to.equal(50);

    const TokenAsUser0 = await ethers.getContract("Token", users[0]);
    await TokenAsUser0.transfer(users[1], 50);
    expect(await TokenAsOwner.balanceOf(users[1])).to.equal(50);
  });
});
```

### Full coverage

Now that we've covered the basics you'll need for testing your contracts, here's a full test suite for the token with a lot of additional information about Mocha and how to structure your tests. We recommend reading through.


But first we add some utility functions that we will use in the test suite.

Create a folder called `utils` in the `test` folder and inside it, create a file called `index.ts` with the following content:

```typescript
import {Contract} from 'ethers';
import {ethers} from 'hardhat';

export async function setupUsers<T extends {[contractName: string]: Contract}>(
  addresses: string[],
  contracts: T
): Promise<({address: string} & T)[]> {
  const users: ({address: string} & T)[] = [];
  for (const address of addresses) {
    users.push(await setupUser(address, contracts));
  }
  return users;
}

export async function setupUser<T extends {[contractName: string]: Contract}>(
  address: string,
  contracts: T
): Promise<{address: string} & T> {
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const user: any = {address};
  for (const key of Object.keys(contracts)) {
    user[key] = contracts[key].connect(await ethers.getSigner(address));
  }
  return user as {address: string} & T;
}

```

This approach will allow you to have succinct and easy to read tests as you can see from the following example.

Here is the test suite. Overwrite Test.test.ts with the following content:

```typescript
// We import Chai to use its assertion functions here.
import {expect} from "./chai-setup";

// we import our utilities
import {setupUsers, setupUser} from './utils';

// We import the hardhat environment field we are planning to use
import {ethers, deployments, getNamedAccounts, getUnnamedAccounts} from 'hardhat';

// we create a setup function that can be called by every test and setup variable for easy to read tests
async function setup () {
  // it first ensures the deployment is executed and reset (use of evm_snapshot for faster tests)
  await deployments.fixture(["Token"]);

  // we get an instantiated contract in the form of a ethers.js Contract instance:
  const contracts = {
    Token: (await ethers.getContract('Token')),
  };

  // we get the tokenOwner
  const {tokenOwner} = await getNamedAccounts();

  // Get the unnammedAccounts (which are basically all accounts not named in the config,
  // This is useful for tests as you can be sure they have noy been given tokens for example)
  // We then use the utilities function to generate user objects
  // These object allow you to write things like `users[0].Token.transfer(....)`
  const users = await setupUsers(await getUnnamedAccounts(), contracts);
  // finally we return the whole object (including the tokenOwner setup as a User object)
  return {
    ...contracts,
    users,
    tokenOwner: await setupUser(tokenOwner, contracts),
  };
}

// `describe` is a Mocha function that allows you to organize your tests. It's
// not actually needed, but having your tests organized makes debugging them
// easier. All Mocha functions are available in the global scope.

// `describe` receives the name of a section of your test suite, and a callback.
// The callback must define the tests of that section. This callback can't be
// an async function.
describe("Token contract", function() {

  // You can nest describe calls to create subsections.
  describe("Deployment", function () {
    // `it` is another Mocha function. This is the one you use to define your
    // tests. It receives the test name, and a callback function.

    // If the callback function is async, Mocha will `await` it.
    it("Should set the right owner", async function () {
      // Expect receives a value, and wraps it in an Assertion object. These
      // objects have a lot of utility methods to assert values.

      // before the test, we call the fixture function.
      // while mocha have hooks to perform these automatically, they force you to declare the variable in greater scope which can introduce subttle errors
      // as such we prefers to have the setup called right at the beginning of the test. this also allow yout o name it accordingly for easier to read tests.
      const {Token} = await setup();


      // This test expects the owner variable stored in the contract to be equal to our configured owner
      const {tokenOwner} = await getNamedAccounts();
      expect(await Token.owner()).to.equal(tokenOwner);
    });

    it("Should assign the total supply of tokens to the owner", async function () {
      const {Token, tokenOwner} = await setup();
      const ownerBalance = await Token.balanceOf(tokenOwner.address);
      expect(await Token.totalSupply()).to.equal(ownerBalance);
    });
  });

  describe("Transactions", function () {
    it("Should transfer tokens between accounts", async function () {
      const {Token, users, tokenOwner} = await setup();
      // Transfer 50 tokens from owner to users[0]
      await tokenOwner.Token.transfer(users[0].address, 50);
      const users0Balance = await Token.balanceOf(users[0].address);
      expect(users0Balance).to.equal(50);

      // Transfer 50 tokens from users[0] to users[1]
      // We use .connect(signer) to send a transaction from another account
      await users[0].Token.transfer(users[1].address, 50);
      const users1Balance = await Token.balanceOf(users[1].address);
      expect(users1Balance).to.equal(50);
    });

    it("Should fail if sender doesn’t have enough tokens", async function () {
      const {Token, users, tokenOwner} = await setup();
      const initialOwnerBalance = await Token.balanceOf(tokenOwner.address);

      // Try to send 1 token from users[0] (0 tokens) to owner (1000 tokens).
      // `require` will evaluate false and revert the transaction.
      await expect(users[0].Token.transfer(tokenOwner.address, 1)
      ).to.be.revertedWith("Not enough tokens");

      // Owner balance shouldn't have changed.
      expect(await Token.balanceOf(tokenOwner.address)).to.equal(
        initialOwnerBalance
      );
    });

    it("Should update balances after transfers", async function () {
      const {Token, users, tokenOwner} = await setup();
      const initialOwnerBalance = await Token.balanceOf(tokenOwner.address);

      // Transfer 100 tokens from owner to users[0].
      await tokenOwner.Token.transfer(users[0].address, 100);

      // Transfer another 50 tokens from owner to users[1].
      await tokenOwner.Token.transfer(users[1].address, 50);

      // Check balances.
      const finalOwnerBalance = await Token.balanceOf(tokenOwner.address);
      expect(finalOwnerBalance).to.equal(initialOwnerBalance - 150);

      const users0Balance = await Token.balanceOf(users[0].address);
      expect(users0Balance).to.equal(100);

      const users1Balance = await Token.balanceOf(users[1].address);
      expect(users1Balance).to.equal(50);
    });
  });
});

```

This is what the output of `yarn hardhat test` should look like after running the full test suite:

```
$ yarn hardhat test

  Token contract
    Deployment
      ✓ Should set the right owner
      ✓ Should assign the total supply of tokens to the owner
    Transactions
      ✓ Should transfer tokens between accounts (199ms)
      ✓ Should fail if sender doesn’t have enough tokens
      ✓ Should update balances after transfers (111ms)


  5 passing (1s)
```

Keep in mind that when you run `yarn hardhat test`, your contracts will be compiled if they've changed since the last time you ran your tests.


# 6. Debugging with Hardhat Network
**Hardhat** comes built-in with **Hardhat Network**, a local Ethereum network designed for development. It allows you to deploy your contracts, run your tests, and debug your code. It's the default network **Hardhat** connects to, so you don't need to set anything up for it to work, you just run your tests.

## Solidity `console.log`
When running your contracts and tests on **Hardhat Network** you can print logging messages and contract variables calling `console.log()` from your Solidity code. To use it you have to import **Hardhat**'s`console.log` in your contract code.

This is what it looks like:

```solidity
pragma solidity 0.7.6;

import "hardhat/console.sol";

contract Token {
  //...
}
```

Try adding some `console.log` statments to the `transfer()` function as if you were using it in JavaScript:

```solidity {2,3}
function transfer(address to, uint256 amount) external {
  console.log("Sender balance is %s tokens", balances[msg.sender]);
  console.log("Trying to send %s tokens to %s", amount, to);

  require(balances[msg.sender] >= amount, "Not enough tokens");

  balances[msg.sender] -= amount;
  balances[to] += amount;
}
```

The logging output will show when you run your tests:

```{8-11,14-17}
$ yarn hardhat test

  Token contract
    Deployment
      ✓ Should set the right owner
      ✓ Should assign the total supply of tokens to the owner
    Transactions
Sender balance is 1000 tokens
Trying to send 50 tokens to 0xead9c93b79ae7c1591b1fb5323bd777e86e150d4
Sender balance is 50 tokens
Trying to send 50 tokens to 0xe5904695748fe4a84b40b3fc79de2277660bd1d3
      ✓ Should transfer tokens between accounts (373ms)
      ✓ Should fail if sender doesn’t have enough tokens
Sender balance is 1000 tokens
Trying to send 100 tokens to 0xead9c93b79ae7c1591b1fb5323bd777e86e150d4
Sender balance is 900 tokens
Trying to send 100 tokens to 0xe5904695748fe4a84b40b3fc79de2277660bd1d3
      ✓ Should update balances after transfers (187ms)


  5 passing (2s)
```
Check out the [documentation](/hardhat-network/README.md#console-log) to learn more about this feature.

# 7. Deploying to a live network
Once you're ready to share your app with other people, you may want to deploy it to a live network! This way others can access an instance that's not running locally on your system.

The Ethereum network that deals with real money is called "mainnet", and then there are other live networks that don't deal with real money but do mimic the real world scenario well, and can be used by others as a shared staging environment. These are called "testnets" and Ethereum has several of them: *Ropsten*, *Kovan*, *Rinkeby* and *Goerli*.

At the software level, deploying to a testnet is the same as deploying to mainnet. The only difference is which network you connect to.

Since we use `hardhat-deploy` plugin and we already set up our deployment procedures for the tests, we are ready to deploy to a live network, we just need to add some configuration for the network we intend to deploy to.

As explained in our deployment section you can execute `yarn hardhat deploy` which will give you the following output, but does not actually deploy your contract anywhere except the default "in-memory" network (`hardhat`)

```
Nothing to compile
deploying "Token" (tx: 0x259d19f33819ec8d3bd994f82912aec6af1a18ec5d74303cfb28d793a10ff683)...: deployed at 0x5FbDB2315678afecb367f032d93F642f64180aa3 with 592983 gas
Done in 3.79s.
```

To deploy to a specific network, you need to add `--network <network-name>` like this:

```
yarn hardhat --network <network-name> deploy
```

## Deploying to remote networks
To deploy to a remote network such as mainnet or any testnet, you need to add a `network` entry to your `hardhat.config.js` file. We’ll use Rinkeby for this example, but you can add any network.

To make it easier to handle the private keys and network configuration, we create a new folder at the root of your project `utils`

In it we create a file `network.ts` with the following content:

```typescript
import 'dotenv/config';
export function node_url(networkName: string): string {
  if (networkName) {
    const uri = process.env['ETH_NODE_URI_' + networkName.toUpperCase()];
    if (uri && uri !== '') {
      return uri;
    }
  }

  let uri = process.env.ETH_NODE_URI;
  if (uri) {
    uri = uri.replace('{{networkName}}', networkName);
  }
  if (!uri || uri === '') {
    if (networkName === 'localhost') {
      return 'http://localhost:8545';
    }
    return '';
  }
  if (uri.indexOf('{{') >= 0) {
    throw new Error(
      `invalid uri or network not supported by node provider : ${uri}`
    );
  }
  return uri;
}

export function getMnemonic(networkName?: string): string {
  if (networkName) {
    const mnemonic = process.env['MNEMONIC_' + networkName.toUpperCase()];
    if (mnemonic && mnemonic !== '') {
      return mnemonic;
    }
  }

  const mnemonic = process.env.MNEMONIC;
  if (!mnemonic || mnemonic === '') {
    return 'test test test test test test test test test test test junk';
  }
  return mnemonic;
}

export function accounts(networkName?: string): {mnemonic: string} {
  return {mnemonic: getMnemonic(networkName)};
}
```

Then we can modifiy our `hardhat.config.ts` file to contain the following:

```typescript {5,11,15-20}
import {HardhatUserConfig} from 'hardhat/types';
import 'hardhat-deploy';
import 'hardhat-deploy-ethers';
import {node_url, accounts} from './utils/network';

const config: HardhatUserConfig = {
  solidity: {
    version: '0.7.6',
  },
  networks: {
    rinkeby: {
      url: node_url('rinkeby'),
      accounts: accounts('rinkeby'),
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

```

Finally we need to setup the environment variable that `utils/networks.ts` reads automatically from `.env`

create a `.env` with the following content. This is where you write your own alchemy api key and mnemonic for rinkeby

```
ETH_NODE_URI_RINKEBY=https://eth-rinkeby.alchemyapi.io/v2/<alchmey api key>
MNEMONIC_RINKEBY=<mnemonic for rinkeby>
```

We're using [Alchemy](https://www.alchemyapi.io), but pointing `url` to any Ethereum node or gateway would work. Go grab your api key and come back.

To deploy on Rinkeby you need to send rinkeby-ETH into the address that's going to be making the deployment. You can get some ETH for testnets from a faucet, a service that distributes testing ETH for free. [Here's the one for Rinkeby](https://faucet.metamask.io/). You'll have to change Metamask's network to Rinkeby before transacting.

::: tip
You can get some ETH for other testnets following these links:

* [Kovan faucet](https://faucet.kovan.network/)
* [Rinkeby faucet](https://faucet.rinkeby.io/)
* [Goerli faucet](https://goerli-faucet.slock.it/)
:::

Finally, run:
```
yarn hardhat --network rinkeby deploy
```

If everything went well, you should see something like:

```
Nothing to compile
deploying "Token" (tx: 0xb40879c3162e6a924cfadfc1027c4629dd57ee4ba08a5f8af575be1c751cd515)...: deployed at 0x8bDFEf5f67685725BC0eD9f54f20A2A4d3FEDA98 with 475842 gas
```

You will also see that some files have been created in the `deployments/rinkeby` folder.

Most notably you'll see `deployments/rinkeby/Token.json` which contains useful information about your deployed contract, including the address, abi, and the solidity input used to create it.

You can then verify it using sourcify or etherscan.

For sourcify you can execute the following:

```
yarn hardhat --network rinkeby sourcify
```

this should give you the following output (with different address) :

```
verifying Token (0x8bDFEf5f67685725BC0eD9f54f20A2A4d3FEDA98 on chain 4) ...
 => contract Token is now verified
```

For etherscan you can do the following:

(Note you can also specify the api key via the env variable ETHERSCAN_API_KEY)

```
yarn hardhat --network rinkeby etherscan-verify --api-key <api-key>
```

You should then see:

```
verifying Token (0x8bDFEf5f67685725BC0eD9f54f20A2A4d3FEDA98) ...
waiting for result...
 => contract Token is now verified
```

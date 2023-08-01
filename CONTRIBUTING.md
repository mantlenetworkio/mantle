# Mantle repo contributing guide

🎈 Thanks for your help improving the project! We are so happy to have you!

There are plenty of ways to contribute, in particular we appreciate support in the following areas:

- Reporting issues. For security issues see [Security policy](https://github.com/mantlenetworkio/mantle/blob/main/SECURITY.md).
- Fixing and responding to existing issues. You can start off with those tagged ["good first issue"](https://github.com/mantlenetworkio/mantle/contribute) which are meant as introductory issues for external contributors.
- Improving the [documentation](https://github.com/mantlenetworkio/documents) and [tutorials](https://docs.mantle.xyz/network/for-devs/tutorials).
- Become a "Mantler" and answer questions in the [mantle Discord](https://discord.com/invite/0xMantle).
- Get involved in the protocol design process by proposing changes or new features or write parts of the spec yourself in the [mantle-specs](https://docs.mantle.xyz/network/intro/overview).

Note that we have a [Code of Conduct](https://github.com/mantlenetworkio/.github/blob/main/CODE_OF_CONDUCT.md), please follow it in all your interactions with the project.


## Contribution Guidelines

We believe one of the things that makes Mantle special is its coherent design and we seek to retain this defining characteristic. From the outset we defined some guidelines to ensure new contributions only ever enhance the project:

- Quality: Code in the Mantle project should meet the style guidelines, with sufficient test-cases, descriptive commit messages, evidence that the contribution does not break any compatibility commitments or cause adverse feature interactions, and evidence of high-quality peer-review
- Size: The Mantle project’s culture is one of small pull-requests, regularly submitted. The larger a pull-request, the more likely it is that you will be asked to resubmit as a series of self-contained and individually reviewable smaller PRs
- Maintainability: If the feature will require ongoing maintenance (eg support for a particular brand of database), we may ask you to accept responsibility for maintaining this feature



## Workflow for Pull Requests

🚨 Before making any non-trivial change, please first open an issue describing the change to solicit feedback and guidance. This will increase the likelihood of the PR getting merged.

In general, the smaller the diff the easier it will be for us to review quickly.

In order to contribute, fork the appropriate branch. Our development branch is `develop`, and the latest release that is normally `release/X.X.X` branch, see [details about our branching model](https://github.com/mantlenetworkio/mantle/blob/main/CONTRIBUTING.md#branching-model-and-releases).

Additionally, if you are writing a new feature, please ensure you add appropriate test cases.

We recommend using the [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) format on commit messages.

Unless your PR is ready for immediate review and merging, please mark it as 'draft' (or simply do not open a PR yet).

**Bonus:** Add comments to the diff under the "Files Changed" tab on the PR page to clarify any sections where you think we might have questions about the approach taken.

### Response time:
We aim to provide a meaningful response to all PRs and issues from external contributors within 2 business days.


## Branching Model

### Active Branches

| Branch                                                              | Status                                                                                                    |
|---------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------|
| [main](https://github.com/mantlenetworkio/mantle/tree/main/)        | Accepts PRs from `release/X.X.X` when we intend to deploy to mainnet/testnet.                             |
| [develop](https://github.com/mantlenetworkio/mantle/tree/develop/)  | Accepts PRs for all changes, particularly those not backwards compatible with `release/X.X.X` and `main`. |
| release/X.X.X                                                       | Create a release branch from the `develop` branch, and the release branch only receives related bugs.     |


### Overview

We generally follow [this Git branching model](https://nvie.com/posts/a-successful-git-branching-model/).
Please read the linked post if you're planning to make frequent PRs into this repository (e.g., people working at/with Mantle).

### Production branch

Our production branch is `main`.
The `main` branch contains the code for our latest "stable" releases.
Updates from `main` **always** come from the `release/X.X.X` branch.
We only ever update the `main` branch when we intend to deploy code within the `release/X.X.X` to the Mantle mainnet/testnet.
Our update process takes the form of a PR merging the `release/X.X.X` branch into the `main` branch.

### Development branch

Our development branch is [`develop`](https://github.com/mantlenetworkio/mantle/tree/develop/).
Changes that are not backwards compatible and all changes to contracts within `packages/contracts/contracts` MUST be directed towards `develop`.

**Changes to contracts within `packages/contracts/contracts` are usually NOT considered backwards compatible and SHOULD be made against towards `develop` branch**.
If you're changing or adding a contract and you're unsure about which branch to make a PR into, default to using the `develop` branch.

### Release candidate branches

Branches marked `release/X.X.X` are **release candidate branches**.
Release candidates are merged into `main` once they've been fully deployed.
We may sometimes have more than one active `release/X.X.X` branch if we're in the middle of a deployment.
`release/X.X.X` contains the most up-to-date software that remains backwards compatible with our latest [experimental network](https://docs.mantle.xyz/network/intro/quick-start).
See table in the **Active Branches** section above to find the right branch to target.

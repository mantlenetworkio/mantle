# Mantle repo contributing guide

🎈 Thanks for your help improving the project! We are so happy to have you!

There are plenty of ways to contribute, in particular we appreciate support in the following areas:

- Reporting issues. For security issues see [Security policy](https://github.com/mantlenetworkio/.github/blob/master/SECURITY.md).
- Fixing and responding to existing issues. You can start off with those tagged ["good first issue"](https://github.com/mantlenetworkio/mantle/contribute) which are meant as introductory issues for external contributors.
- Improving the [community site](https://community.mantle.io/)[documentation](https://github.com/mantlenetworkio/documents) and [tutorials](https://github.com/mantlenetworkio/tutorial).
- Become an "Optimizer" and answer questions in the [mantle Discord](https://discord.com/invite/jrnFEvq).
- Get involved in the protocol design process by proposing changes or new features or write parts of the spec yourself in the [optimistic-specs repo](https://github.com/mantlenetworkio/optimistic-specs).

Note that we have a [Code of Conduct](https://github.com/mantlenetworkio/.github/blob/master/CODE_OF_CONDUCT.md), please follow it in all your interactions with the project.


## Contribution Guidelines

We believe one of the things that makes Mantle special is its coherent design and we seek to retain this defining characteristic. From the outset we defined some guidelines to ensure new contributions only ever enhance the project:

- Quality: Code in the Mantle project should meet the style guidelines, with sufficient test-cases, descriptive commit messages, evidence that the contribution does not break any compatibility commitments or cause adverse feature interactions, and evidence of high-quality peer-review
- Size: The Mantle project’s culture is one of small pull-requests, regularly submitted. The larger a pull-request, the more likely it is that you will be asked to resubmit as a series of self-contained and individually reviewable smaller PRs
- Maintainability: If the feature will require ongoing maintenance (eg support for a particular brand of database), we may ask you to accept responsibility for maintaining this feature



## Workflow for Pull Requests

🚨 Before making any non-trivial change, please first open an issue describing the change to solicit feedback and guidance. This will increase the likelihood of the PR getting merged.

In general, the smaller the diff the easier it will be for us to review quickly.

In order to contribute, fork the appropriate branch, for non-breaking changes to production that is `develop` and for the next release that is normally `release/X.X.X` branch, see [details about our branching model](https://github.com/mantlenetworkio/mantle/blob/develop/README.md#branching-model-and-releases).

Additionally, if you are writing a new feature, please ensure you add appropriate test cases.

Follow the [Development Quick Start](#development-quick-start) to set up your local development environment.

We recommend using the [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) format on commit messages.

Unless your PR is ready for immediate review and merging, please mark it as 'draft' (or simply do not open a PR yet).

**Bonus:** Add comments to the diff under the "Files Changed" tab on the PR page to clarify any sections where you think we might have questions about the approach taken.

### Response time:
We aim to provide a meaningful response to all PRs and issues from external contributors within 2 business days.

### Changesets

We use [changesets](https://github.com/atlassian/changesets) to manage releases of our various packages.
You *must* include a `changeset` file in your PR when making a change that would require a new package release.

Adding a `changeset` file is easy:

1. Navigate to the root of the monorepo.
2. Run `yarn changeset`. You'll be prompted to select packages to include in the changeset. Use the arrow keys to move the cursor up and down, hit the `spacebar` to select a package, and hit `enter` to confirm your selection. Select *all* packages that require a new release as a result of your PR.
3. Once you hit `enter` you'll be prompted to decide whether your selected packages need a `major`, `minor`, or `patch` release. We follow the [Semantic Versioning](https://semver.org/) scheme. Please avoid using `major` releases for any packages that are still in version `0.y.z`.
4. Commit your changeset and push it into your PR. The changeset bot will notice your changeset file and leave a little comment to this effect on GitHub.
5. Voilà, c'est fini!

### Rebasing

We use the `git rebase` command to keep our commit history tidy.
Rebasing is an easy way to make sure that each PR includes a series of clean commits with descriptive commit messages
See [this tutorial](https://docs.gitlab.com/ee/topics/git/git_rebase.html) for a detailed explanation of `git rebase` and how you should use it to maintain a clean commit history.


##Branch Rule



### Active Branches

| Branch          | Status                                                                           |
| --------------- | -------------------------------------------------------------------------------- |
| [master](https://github.com/mantlenetworkio/bitnetwork/tree/master/)    | Accepts PRs from other branch  when we intend to deploy to testnet.                                      |
| {username}/{module-name}-{develop-feature}  | Contributors need to build the branch follow the below rules,just like `alice/l2geth-xxx`.Actually,they means the `develop` branch.
| release/X.X.X                                                                          | For now ,we do not support the release branch util the launch of Mantle TestNet |

#### The `master` branch




The `master` branch contains the code for our latest contributions for Mantle.
Updates from `master` always come from the `develop` branch just like `alice/l2geth-xxx`.
Our update process takes the form of a PR merging the `develop` branch into the `master` branch.

#### The `develop` branch

Mantle is the project only

Our primary development branch is [`develop`](https://github.com/mantlenetworkio/bitnetwork/tree/develop/).
`develop` contains the most up-to-date software that remains backwards compatible with our latest experimental [network deployments](https://community.bitnetwork.io/docs/useful-tools/networks/).
If you're making a backwards compatible change, please direct your pull request towards `develop`.

**Changes to contracts within `packages/contracts/contracts` are usually NOT considered backwards compatible and SHOULD be made against a release candidate branch**.
Some exceptions to this rule exist for cases in which we absolutely must deploy some new contract after a release candidate branch has already been fully deployed.
If you're changing or adding a contract and you're unsure about which branch to make a PR into, default to using the latest release candidate branch.
See below for info about release candidate branches.

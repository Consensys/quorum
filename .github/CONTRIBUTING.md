# Contributing

Thank you for your interest in contributing to GoQuorum!

We welcome contributions from anyone on the internet, and are grateful for even the
smallest of fixes!

If you'd like to contribute to quorum please fork, fix, commit and
send a pull request. Commits which do not comply with the coding standards
are ignored.

## Coding guidelines

Please make sure your contributions adhere to our coding guidelines:

 * Code must adhere to the official Go
[formatting](https://golang.org/doc/effective_go.html#formatting) guidelines
(i.e. uses [gofmt](https://golang.org/cmd/gofmt/)).
 * Code must be documented adhering to the official Go
[commentary](https://golang.org/doc/effective_go.html#commentary) guidelines.
 * Pull requests need to be based on and opened against the `master` branch.
 * Commit messages should be prefixed with the package(s) they modify.
   * E.g. "eth, rpc: make trace configs optional"

## Can I have feature X

Before you submit a feature request, please check and make sure that it isn't
possible through some other means. The JavaScript-enabled console is a powerful
feature in the right hands. Please check our
[developer documentation](https://docs.goquorum.consensys.net/en/latest/) for more info
and help.

## Configuration, dependencies, and tests

Please see the [Developers' Guide](https://geth.ethereum.org/docs/developers/devguide)
for more details on configuring your environment, managing project dependencies
and testing procedures.

## Issue reproductibility

Before you create an issue, please try to reproduce it with [quorum quick dev start framework](https://github.com/ConsenSys/quorum-dev-quickstart) with the latest release.
Many issues have been creating under an environment we are not able to reproduce quickly.

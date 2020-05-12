## Quorum Wizard
[Quorum Wizard](https://github.com/jpmorganchase/quorum-wizard) is a command line tool that allows users to set up a development Quorum network on their local machine in less than 2 minutes.

![](docs/quorum-wizard.gif)

## Installation

`quorum-wizard` is written in Javascript and designed to be installed as a global NPM module and run from the command line. Make sure you have [Node.js/NPM](https://docs.npmjs.com/downloading-and-installing-node-js-and-npm) installed.

Using npm:

```Bash
npm install -g quorum-wizard
```

Using [Yarn](https://yarnpkg.com/):

```Bash
yarn global add quorum-wizard
```

### Using Quorum Wizard

Once the global module is installed, just run:

```Bash
quorum-wizard
```

The wizard will then walk you through setting up a network, either using our quickstart settings (a simple 3-node Quorum network using Raft consensus), or customizing the options to fit your needs.

## Options

You can also provide these flags when running quorum-wizard:

* `-q`, `--quickstart`  create 3 node raft network with tessera and cakeshop (no user-input required)
* `-v`, `--verbose`     Turn on additional logs for debugging
* `--version`           Show version number
* `-h`, `--help`        Show help


Note: `npx` is also way to run npm modules without the need to actually install the module. Due to quorum-wizard needing to download and cache the quorum binaries during network setup, using `npx quorum-wizard` will not work at this time.

## Interacting with the Network

To explore the features of Quorum and deploy a private contract, follow the instructions on [Interacting with the Network](./Interacting.md)

## Troubleshooting

**EACCES error when doing global npm install**:

- Sometimes npm is installed in a location where the user doesn't have write permissions. On Mac, installing via [Homebrew](https://brew.sh) usually works better than the standalone installer.
- [Here is the recommended solution from NPM](https://docs.npmjs.com/resolving-eacces-permissions-errors-when-installing-packages-globally)

## Developing
Clone this repo to your local machine.

`yarn install` to get all the dependencies.

`yarn test:watch` to automatically run tests on changes

`yarn start` to automatically build on changes to any files in the src directory

`yarn link` to use your development build when you run the global npm command

`quorum-wizard` to run (alternatively, you can run `node build/index.js`)

## Contributing
[Quorum Wizard](https://github.com/jpmorganchase/quorum-wizard) is built on open source and we invite you to contribute enhancements. Upon review you will be required to complete a Contributor License Agreement (CLA) before we are able to merge. If you have any questions about the contribution process, please feel free to send an email to [info@goquorum.com](mailto:info@goquorum.com).

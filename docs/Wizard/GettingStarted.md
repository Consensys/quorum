## Quorum Wizard
[Quorum Wizard](https://github.com/jpmorganchase/quorum-wizard) is a command line tool that allows users to set up a development Quorum network on their local machine in less than 2 minutes.

![](docs/quorum-wizard.gif)

## Using Quorum Wizard

Quorum Wizard is written in Javascript and designed to be run as a global NPM module from the command line. Make sure you have [Node.js/NPM](https://docs.npmjs.com/downloading-and-installing-node-js-and-npm) installed.

Using npx to run the wizard without the need to install:

```
npx quorum-wizard
```

You can also install the wizard globally with npm:

```Bash
npm install -g quorum-wizard

# Once the global module is installed, run:
quorum-wizard
```

Note: Many installations of npm don't have permission to install global modules and will throw an EACCES error. [Here is the recommended solution from NPM](https://docs.npmjs.com/resolving-eacces-permissions-errors-when-installing-packages-globally)

## Dependencies

Here are the dependencies (in addition to NodeJS) that are required depending on the mode that you run the wizard in:

Bash:

- Java (when running Tessera and/or Cakeshop)

Docker Compose:

- Docker
- docker-compose

Kubernetes:

- Docker (for generating resources during network creation)
- kubectl
- minikube, Docker Desktop with Kubernetes enabled, or some other kubernetes context

## Options

You can also provide these flags when running quorum-wizard:

| Flags | Effect |
| - | - |
| `-q`, `--quickstart` | Create 3 node raft network with tessera and cakeshop (no user-input required) |
| `-v`, `--verbose` | Turn on additional logs for debugging |
| `--version` | Show version number |
| `-h`, `--help` | Show help |

## Interacting with the Network

To explore the features of Quorum and deploy a private contract, follow the instructions on [Interacting with the Network](./Interacting.md)

## Developing
Clone this repo to your local machine.

`npm install` to get all the dependencies.

`npm run test:watch` to automatically run tests on changes

`npm run start` to automatically build on changes to any files in the src directory

`npm link` to use your development build when you run the global npm command

`quorum-wizard` to run (alternatively, you can run `node build/index.js`)

## Contributing
Quorum Wizard is built on open source and we invite you to contribute enhancements. Upon review you will be required to complete a Contributor License Agreement (CLA) before we are able to merge. If you have any questions about the contribution process, please feel free to send an email to [info@goquorum.com](mailto:info@goquorum.com).

## Getting Help
Stuck at some step? Please join our <a href="https://www.goquorum.com/slack-inviter" target="_blank" rel="noopener">Slack community</a> for support.

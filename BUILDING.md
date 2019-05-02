
# Building Quorum

Note: Building Quorum requires both a Go (version 1.9 or later) and a C compiler. You can install them using your favourite package manager. 

Clone the repository and build the source:

```
git clone https://github.com/jpmorganchase/quorum.git
cd quorum
make all
make test
```

Binaries are placed within `./build/bin`, most notably `geth` and `bootnode`. Either add this directory to your `$PATH` or copy those two bins into your PATH:

```sh
# assumes that /usr/local/bin is in your PATH
cp ./build/bin/geth ./build/bin/bootnode /usr/local/bin/
```

# Building on Windows
It is possible to build and run Quorum on Windows. Below are the steps required, please use Slack for any questions or support. Keep in mind that original Go-Ethereum provides a number of helper scripts for environment configuration and build execution. We're not planning ot provide this ourselves, but steps below explain what you may need to set up on your system to create such scripts.

1. Install Go version 1.10 or 1.11 for Windows
2. Create a folder that you will bind to be GOPATH
   * Set this folder as GOPATH in your environment: `set GOPATH=your_new_folder`
   * Go will build binaries to a sub-path of GOPATH, so add it into PATH: `set "PATH=%PATH%;%GOPATH%\bin\"`
3. In GOPATH folder, create `src\github.com\ethereum` set of sub paths
4. Checkout Quorum into newly created GOPATH structure like so: `git clone https://github.com/jpmorganchase/quorum.git GOPATH\src\github.com\ethereum\go-ethereum`. Note: **go-ethereum** added at the end is required since Quorum occupies the same namespace as go-ethereum project. For this reason, on windows, a separate GOPATH is recommended if you need to build both projects
5. Change directory into checked out project and build it with `go install -v ./cmd/geth`

The steps mimic the documentation from [go-ethereum project](https://github.com/ethereum/go-ethereum/wiki/Installation-instructions-for-Windows), please reference as needed

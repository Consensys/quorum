
# Building Quorum

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

# Building Quorum on Ubuntu

"make all" may result in the following message on Ubuntu:

build/env.sh go run build/ci.go install
build/env.sh: 30: exec: go: not found
Makefile:29: recipe for target 'all' failed
make: *** [all] Error 127

Solution:
apt-get install golang-go -y

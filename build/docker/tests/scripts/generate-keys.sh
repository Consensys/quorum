#!/bin/bash

#
# calls constellation-node in the docker instance to generate
# key pairs for the constellation nodes and move them to the
# configuration folder
#

/usr/local/bin/constellation-node --generatekeys=tm < /dev/null > /dev/null
mv tm.* /qdata/constellation/
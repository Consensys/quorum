# Changes

- #21504 trie sync -> path based operation
- #21491 memory leak: trie
- #21366 change in download size: renaming, new initial value
- #21334 can't change time on non-empty blocks
- #21501 getTypeSize
- #21455 download queue stats per minute, was 10 sec
- #21517 usb wallet, divided derivation
- #21514 renaming stateDB, unused params (ctx -> _), comments

# UPGRADES

- #21448 go mod, update some deps
- #21432 goja

# CONF

- #21534 some net checkpoints changed

# FIXES

- #21497 fix null JSON-RPC
- #21501 fix getTypeSize
- #21503 fix personal.sign() -> console bridge -> goja -> getjeth -> unlockAccount->sign

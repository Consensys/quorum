# Hacking on Quorum / various notes

## How does private state work?

Original commit from Jeff explains the dual public and private state with INITIAL restrictions:
```
commit 763f939f4725daa136161868d3b01fa7a84eb71e
Author: Jeffrey Wilcke <jeffrey@ethereum.org>
Date:   Mon Oct 31 12:46:40 2016 +0100

   core, core/vm: dual state & read only EVM

   This commit implements a dual state approach. The dual state approach
   separates public and private state by making the core vm environment
   context aware.

   Although not currently implemented it will need to prohibit value
   transfers and it must initialise all transactions from accounts on the
   public state. This means that sending transactions increments the
   account nonce on the public state and contract addresses are derived
   from the public state when initialised by a transaction. For obvious
   reasons, contract created by private contracts are still derived from
   public state.

   This is required in order to have consensus over the public state at all
   times as non-private participants would still process the transaction on
   the public state even though private payload can not be decrypted. This
   means that participants of a private group must do the same in order to
   have public consensus. However the creation of the contract and
   interaction still occurs on the private state.

   It implements support for the following calling model:

   S: sender, (X): private, X: public, ->: direction, [ ]: read only mode

   1. S -> A -> B
   2. S -> (A) -> (B)
   3. S -> (A) -> [ B -> C ]

   It does not support

   1. (S) -> A
   2. (S) -> (A)
   3. S -> (A) -> B

   Implemented "read only" mode for the EVM. Read only mode is checked
   during any opcode that could potentially modify the state. If such an
   opcode is encountered during "read only", it throws an exception.

   The EVM is flagged "read only" when a private contract calls in to
   public state.
```


Some things have changed since, let's look at the EVM structure in some more detail:

```go
type EVM struct {
	...
	// StateDB gives access to the underlying state
	StateDB StateDB
	// Depth is the current call stack
	depth int
	...

	publicState       PublicState
	privateState      PrivateState
	states            [1027]*state.StateDB
	currentStateDepth uint
	readOnly          bool
	readOnlyDepth     uint
}
```

The vanilla EVM has a call depth limit of 1024. Our `states` parallel the EVM call stack, recording as contracts in the public and private state call back and forth to each other. Note it doesn't have to be a "public -> private -> public -> private" back-and-forth chain. It can be any sequence of { public, private }.

The interface for calling is this `Push` / `Pop` sequence:

```go
evm.Push(getDualState(evm, addr))
defer func() { evm.Pop() }()
// ... do work in the pushed state
```

The definitions of `Push` and `Pop` are simple and important enough to duplicate here:

```go
func (env *EVM) Push(statedb StateDB) {
	if env.privateState != statedb {
		env.readOnly = true
		env.readOnlyDepth = env.currentStateDepth
	}

	if castedStateDb, ok := statedb.(*state.StateDB); ok {
		env.states[env.currentStateDepth] = castedStateDb
		env.currentStateDepth++
	}

	env.StateDB = statedb
}
func (env *EVM) Pop() {
	env.currentStateDepth--
	if env.readOnly && env.currentStateDepth == env.readOnlyDepth {
		env.readOnly = false
	}
	env.StateDB = env.states[env.currentStateDepth-1]
}
```

Note the invariant that `StateDB` always points to the current state db.

The other interesting note is read only mode. Any time we call from the private state into the public state (`env.privateState != statedb`), we require anything deeper to be *read only*. Private state transactions can't affect public state, so we throw an EVM exception on any mutating operation (`SELFDESTRUCT, CREATE, SSTORE, LOG0, LOG1, LOG2, LOG3, LOG4`). Question: have any more mutating operations been added? Question: could we not mutate deeper private state?

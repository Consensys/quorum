# Hacking on Quorum / various notes

## How does private state work?

Let's look at the EVM structure:

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

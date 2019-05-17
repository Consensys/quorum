package core

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/params"
)

// callHelper makes it easier to do proper calls and use the state transition object.
// It also manages the nonces of the caller and keeps private and public state, which
// can be freely modified outside of the helper.
type callHelper struct {
	db ethdb.Database

	nonces map[common.Address]uint64
	header types.Header
	gp     *GasPool

	PrivateState, PublicState *state.StateDB
	BC                        *BlockChain
}

// TxNonce returns the pending nonce
func (cg *callHelper) TxNonce(addr common.Address) uint64 {
	return cg.nonces[addr]
}

// MakeCall makes does a call to the recipient using the given input. It can switch between private and public
// by setting the private boolean flag. It returns an error if the call failed.
func (cg *callHelper) MakeCall(private bool, key *ecdsa.PrivateKey, to common.Address, input []byte) (common.Address, error) {
	var (
		from = crypto.PubkeyToAddress(key.PublicKey)
		err  error
	)

	// TODO(joel): these are just stubbed to the same values as in dual_state_test.go
	cg.header.Number = new(big.Int)
	cg.header.Time = new(big.Int).SetUint64(43)
	cg.header.Difficulty = new(big.Int).SetUint64(1000488)
	cg.header.GasLimit = 4700000

	signer := types.MakeSigner(params.QuorumTestChainConfig, cg.header.Number)
	var transaction *types.Transaction
	nonce := cg.TxNonce(from)
	contractAddr := crypto.CreateAddress(from, nonce)
	log.Trace("contract address", "addr", contractAddr.Hex())
	if to == common.HexToAddress("0x0") {
		transaction = types.NewContractCreation(nonce, new(big.Int), 1000000, new(big.Int), input)
	} else {
		transaction = types.NewTransaction(nonce, to, new(big.Int), 1000000, new(big.Int), input)
	}
	tx, err := types.SignTx(transaction, signer, key)
	if err != nil {
		return common.Address{}, err
	}
	defer func() { cg.nonces[from]++ }()
	msg, err := tx.AsMessage(signer)
	if err != nil {
		return common.Address{}, err
	}

	publicState, privateState := cg.PublicState, cg.PrivateState
	if !private {
		privateState = publicState
	} else {
		tx.SetPrivate()
	}

	// TODO(joel): can we just pass nil instead of bc?
	if cg.BC == nil {
		cg.BC, _ = NewBlockChain(cg.db, nil, params.QuorumTestChainConfig, ethash.NewFaker(), vm.Config{})
	}
	context := NewEVMContext(msg, &cg.header, cg.BC, &from)
	vmenv := vm.NewEVM(context, publicState, privateState, params.QuorumTestChainConfig, vm.Config{})
	_, _, _, err = ApplyMessage(vmenv, msg, cg.gp)
	if err != nil {
		return common.Address{}, err
	}
	return contractAddr, nil
}

// MakeCallHelper returns a new callHelper
func MakeCallHelper() *callHelper {
	memdb := ethdb.NewMemDatabase()
	db := state.NewDatabase(memdb)

	publicState, err := state.New(common.Hash{}, db)
	if err != nil {
		panic(err)
	}
	privateState, err := state.New(common.Hash{}, db)
	if err != nil {
		panic(err)
	}
	publicState.SetPersistentEthdb(memdb)
	privateState.SetPersistentEthdb(memdb)
	cg := &callHelper{
		db:           memdb,
		nonces:       make(map[common.Address]uint64),
		gp:           new(GasPool).AddGas(5000000),
		PublicState:  publicState,
		PrivateState: privateState,
	}
	return cg
}

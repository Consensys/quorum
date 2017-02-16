package core

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethdb"
)

// privateTestTx stubs transaction so that it can be flagged as private or not
type privateTestTx struct {
	*types.Transaction
	private bool
}

// IsPrivate returns whether the transaction should be considered privat.
func (ptx *privateTestTx) IsPrivate() bool { return ptx.private }

// callHelper makes it easier to do proper calls and use the state transition object.
// It also manages the nonces of the caller and keeps private and public state, which
// can be freely modified outside of the helper.
type callHelper struct {
	db ethdb.Database

	nonces map[common.Address]uint64
	header types.Header
	gp     *GasPool

	PrivateState, PublicState *state.StateDB
}

// TxNonce returns the pending nonce
func (cg *callHelper) TxNonce(addr common.Address) uint64 {
	return cg.nonces[addr]
}

// MakeCall makes does a call to the recipient using the given input. It can switch between private and public
// by setting the private boolean flag. It returns an error if the call failed.
func (cg *callHelper) MakeCall(private bool, key *ecdsa.PrivateKey, to common.Address, input []byte) error {
	var (
		from = crypto.PubkeyToAddress(key.PublicKey)
		ptx  = privateTestTx{private: private}
		err  error
	)
	ptx.Transaction, err = types.NewTransaction(cg.TxNonce(from), to, new(big.Int), big.NewInt(1000000), new(big.Int), input).SignECDSA(key)
	if err != nil {
		return err
	}
	defer func() { cg.nonces[from]++ }()

	publicState, privateState := cg.PublicState, cg.PrivateState
	if !private {
		privateState = publicState
	}

	cg.header.Number = new(big.Int)
	_, _, err = ApplyMessage(NewEnv(publicState, privateState, &ChainConfig{}, nil, ptx.Transaction, &cg.header, vm.Config{}), ptx, cg.gp)
	if err != nil {
		return err
	}
	return nil
}

// MakeCallHelper returns a new callHelper
func MakeCallHelper() *callHelper {
	db, _ := ethdb.NewMemDatabase()

	publicState, err := state.New(common.Hash{}, db)
	if err != nil {
		panic(err)
	}
	privateState, err := state.New(common.Hash{}, db)
	if err != nil {
		panic(err)
	}
	cg := &callHelper{
		nonces:       make(map[common.Address]uint64),
		gp:           new(GasPool).AddGas(big.NewInt(5000000)),
		PublicState:  publicState,
		PrivateState: privateState,
	}
	return cg
}

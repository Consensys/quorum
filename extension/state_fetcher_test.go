package extension

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
)

func TestDumpAddressWhenFound(t *testing.T) {
	//db := ethdb.NewMemDatabase()
	db := rawdb.NewMemoryDatabase()

	statedb, _ := state.New(common.Hash{}, state.NewDatabase(db))
	address := common.HexToAddress("0x2222222222222222222222222222222222222222")

	stateFetcher := NewStateFetcher(nil)

	// generate a few entries and write them out to the db
	statedb.SetBalance(address, big.NewInt(22))
	statedb.SetCode(address, []byte{3, 3, 3, 3, 3, 3, 3})
	statedb.Commit(false)

	out, _ := stateFetcher.addressStateAsJson(statedb, address)

	want := `{"0x2222222222222222222222222222222222222222":{"state":{"balance":"22","nonce":0,"root":"56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421","codeHash":"87874902497a5bb968da31a2998d8f22e949d1ef6214bcdedd8bae24cca4b9e3","code":"03030303030303"}}}`

	if string(out) != want {
		t.Errorf("dump mismatch:\ngot: %s\nwant: %s\n", string(out), want)
	}
}

func TestDumpAddressWhenNotFound(t *testing.T) {
	//db := ethdb.NewMemDatabase()
	db := rawdb.NewMemoryDatabase()
	statedb, _ := state.New(common.Hash{}, state.NewDatabase(db))
	statedb.Commit(false)

	stateFetcher := NewStateFetcher(nil)

	address := common.HexToAddress("0x2222222222222222222222222222222222222222")
	out, _ := stateFetcher.addressStateAsJson(statedb, address)

	if out != nil {
		t.Errorf("dump mismatch:\ngot: %s\nwant: nil\n", string(out))
	}
}

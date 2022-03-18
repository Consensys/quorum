package bind

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

var (
	tmPrivatePayloadHash common.EncryptedPayloadHash
	tmPrivateTxHash      common.EncryptedPayloadHash
)

func init() {
	payloadHash := crypto.Keccak512([]byte("encrypted-private-payload"))
	privateTxHash := crypto.Keccak512([]byte("encrypted-private-tx"))
	for i := 0; i < 64; i++ {
		tmPrivatePayloadHash[i] = payloadHash[i]
		tmPrivateTxHash[i] = privateTxHash[i]
	}
}

func TestBoundContract_Transact_ContractCreation_PrivateTransaction(t *testing.T) {
	transactor := &mockTransactor{}

	c := NewBoundContract(common.Address{}, abi.ABI{}, nil, transactor, nil)

	senderNonce := 1

	opts := &TransactOpts{
		Nonce:      big.NewInt(int64(senderNonce)),
		PrivateFor: []string{"tm1"},

		// arbitrary values to skip the logic we're not testing
		GasPrice: big.NewInt(0),
		GasLimit: uint64(1),
		Signer:   passthroughSigner,
	}

	tx, err := c.transact(opts, nil, nil)

	wantNonce := uint64(senderNonce)
	wantTo := (*common.Address)(nil)
	wantData := tmPrivatePayloadHash.Bytes()

	require.NoError(t, err)
	require.NotNil(t, tx)
	require.Equal(t, wantNonce, tx.Nonce())
	require.Equal(t, wantTo, tx.To())
	require.Equal(t, wantData, tx.Data())
	require.True(t, tx.IsPrivate())
}

func TestBoundContract_Transact_ContractCreation_PrivacyPrecompile(t *testing.T) {
	transactor := &mockTransactor{}

	c := NewBoundContract(common.Address{}, abi.ABI{}, nil, transactor, nil)

	senderNonce := 1

	opts := &TransactOpts{
		Nonce:                    big.NewInt(int64(senderNonce)),
		PrivateFor:               []string{"tm1"},
		IsUsingPrivacyPrecompile: true,

		// arbitrary values to skip the logic we're not testing
		GasPrice: big.NewInt(0),
		GasLimit: uint64(1),
		Signer:   passthroughSigner,
	}

	pmt, err := c.transact(opts, nil, nil)

	require.NoError(t, err)
	require.NotNil(t, pmt)

	// verify the privacy marker transaction
	wantPMTNonce := uint64(senderNonce)
	wantPMTTo := common.QuorumPrivacyPrecompileContractAddress()
	wantPMTData := tmPrivateTxHash.Bytes()

	require.Equal(t, wantPMTNonce, pmt.Nonce())
	require.Equal(t, &wantPMTTo, pmt.To())
	require.Equal(t, wantPMTData, pmt.Data())
	require.False(t, pmt.IsPrivate())

	// verify the captured internal private transaction
	pvtTx := transactor.capturedInternalPrivateTransaction
	pvtTxArgs := transactor.capturedInternalPrivateTransactionArgs

	wantPvtTxNonce := uint64(senderNonce)
	wantPvtTxTo := (*common.Address)(nil)
	wantPvtTxData := tmPrivatePayloadHash.Bytes()

	require.NotNil(t, pvtTx)
	require.Equal(t, wantPvtTxNonce, pvtTx.Nonce())
	require.Equal(t, wantPvtTxTo, pvtTx.To())
	require.Equal(t, wantPvtTxData, pvtTx.Data())
	require.True(t, pvtTx.IsPrivate())

	require.Equal(t, []string{"tm1"}, pvtTxArgs.PrivateFor)
}

func TestBoundContract_Transact_Transaction_PrivateTransaction(t *testing.T) {
	transactor := &mockTransactor{}

	contractAddr := common.HexToAddress("0x1932c48b2bf8102ba33b4a6b545c32236e342f34")
	c := NewBoundContract(contractAddr, abi.ABI{}, nil, transactor, nil)

	senderNonce := 1

	opts := &TransactOpts{
		Nonce:      big.NewInt(int64(senderNonce)),
		PrivateFor: []string{"tm1"},

		// arbitrary values to skip the logic we're not testing
		GasPrice: big.NewInt(0),
		GasLimit: uint64(1),
		Signer:   passthroughSigner,
	}

	tx, err := c.transact(opts, &contractAddr, nil)

	wantNonce := uint64(senderNonce)
	wantTo := &contractAddr
	wantData := tmPrivatePayloadHash.Bytes()

	require.NoError(t, err)
	require.NotNil(t, tx)
	require.Equal(t, wantNonce, tx.Nonce())
	require.Equal(t, wantTo, tx.To())
	require.Equal(t, wantData, tx.Data())
	require.True(t, tx.IsPrivate())
}

func TestBoundContract_Transact_Transaction_PrivacyPrecompile(t *testing.T) {
	transactor := &mockTransactor{}

	contractAddr := common.HexToAddress("0x1932c48b2bf8102ba33b4a6b545c32236e342f34")
	c := NewBoundContract(contractAddr, abi.ABI{}, nil, transactor, nil)

	senderNonce := 1

	opts := &TransactOpts{
		Nonce:                    big.NewInt(int64(senderNonce)),
		PrivateFor:               []string{"tm1"},
		IsUsingPrivacyPrecompile: true,

		// arbitrary values to skip the logic we're not testing
		GasPrice: big.NewInt(0),
		GasLimit: uint64(1),
		Signer:   passthroughSigner,
	}

	pmt, err := c.transact(opts, &contractAddr, nil)

	require.NoError(t, err)
	require.NotNil(t, pmt)

	// verify the privacy marker transaction
	wantPMTNonce := uint64(senderNonce)
	wantPMTTo := common.QuorumPrivacyPrecompileContractAddress()
	wantPMTData := tmPrivateTxHash.Bytes()

	require.Equal(t, wantPMTNonce, pmt.Nonce())
	require.Equal(t, &wantPMTTo, pmt.To())
	require.Equal(t, wantPMTData, pmt.Data())
	require.False(t, pmt.IsPrivate())

	// verify the captured internal private transaction
	pvtTx := transactor.capturedInternalPrivateTransaction
	pvtTxArgs := transactor.capturedInternalPrivateTransactionArgs

	wantPvtTxNonce := uint64(senderNonce)
	wantPvtTxTo := &contractAddr
	wantPvtTxData := tmPrivatePayloadHash.Bytes()

	require.NotNil(t, pvtTx)
	require.Equal(t, wantPvtTxNonce, pvtTx.Nonce())
	require.Equal(t, wantPvtTxTo, pvtTx.To())
	require.Equal(t, wantPvtTxData, pvtTx.Data())
	require.True(t, pvtTx.IsPrivate())

	require.Equal(t, []string{"tm1"}, pvtTxArgs.PrivateFor)
}

func passthroughSigner(_ common.Address, tx *types.Transaction) (*types.Transaction, error) {
	return tx, nil
}

type mockTransactor struct {
	capturedInternalPrivateTransaction     *types.Transaction
	capturedInternalPrivateTransactionArgs PrivateTxArgs
}

func (s *mockTransactor) PreparePrivateTransaction(_ []byte, _ string) (common.EncryptedPayloadHash, error) {
	return tmPrivatePayloadHash, nil
}

func (s *mockTransactor) DistributeTransaction(_ context.Context, tx *types.Transaction, args PrivateTxArgs) (string, error) {
	s.capturedInternalPrivateTransaction = tx
	s.capturedInternalPrivateTransactionArgs = args
	return tmPrivateTxHash.Hex(), nil
}

func (s *mockTransactor) SendTransaction(_ context.Context, _ *types.Transaction, _ PrivateTxArgs) error {
	return nil
}

func (s *mockTransactor) PendingCodeAt(_ context.Context, _ common.Address) ([]byte, error) {
	panic("implement me")
}

func (s *mockTransactor) PendingNonceAt(_ context.Context, _ common.Address) (uint64, error) {
	panic("implement me")
}

func (s *mockTransactor) SuggestGasPrice(_ context.Context) (*big.Int, error) {
	panic("implement me")
}

func (s *mockTransactor) EstimateGas(_ context.Context, _ ethereum.CallMsg) (gas uint64, err error) {
	panic("implement me")
}

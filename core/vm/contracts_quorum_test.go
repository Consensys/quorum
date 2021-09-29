package vm

import (
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/private"
	"github.com/ethereum/go-ethereum/private/mock_private"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

var (
	sender           common.Address
	senderPrivateKey *ecdsa.PrivateKey
	tmPrivateTxHash  common.EncryptedPayloadHash
	pmtData          []byte
)

func init() {
	sender = common.HexToAddress("0xed9d02e382b34818e88b88a309c7fe71e65f419d")
	senderPrivateKey, _ = crypto.HexToECDSA("e6181caaffff94a09d7e332fc8da9884d99902c7874eb74354bdcadf411929f1")

	privateTxHash := crypto.Keccak512([]byte("encrypted-private-tx"))
	for i := 0; i < 64; i++ {
		tmPrivateTxHash[i] = privateTxHash[i]
	}

	pmtData = append(sender.Bytes(), privateTxHash...)
}

func TestPrivacyMarker_Run_UnsupportedTransaction_DoesNothing(t *testing.T) {
	publicContractCreationTx := types.NewContractCreation(0, nil, 0, nil, []byte{})
	privatePrivacyMarkerTx := types.NewTransaction(0, common.QuorumPrivacyPrecompileContractAddress(), nil, 0, nil, []byte{})
	privatePrivacyMarkerTx.SetPrivate()
	require.True(t, privatePrivacyMarkerTx.IsPrivacyMarker())
	require.True(t, privatePrivacyMarkerTx.IsPrivate())

	tests := []struct {
		name      string
		currentTx *types.Transaction
	}{
		{
			name:      "is-nil",
			currentTx: nil,
		},
		{
			name:      "is-not-privacy-marker-tx",
			currentTx: publicContractCreationTx,
		},
		{
			name:      "is-private-privacy-marker-tx",
			currentTx: privatePrivacyMarkerTx,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			pm := privacyMarker{}

			publicState := NewMockStateDB(ctrl)
			innerApplier := &stubInnerApplier{}

			evm := &EVM{
				currentTx:   tt.currentTx,
				publicState: publicState,
				InnerApply:  innerApplier.InnerApply,
			}

			publicState.EXPECT().SetNonce(gomock.Any(), gomock.Any()).Times(0)

			gotByt, gotErr := pm.Run(evm, []byte{})

			require.False(t, innerApplier.wasCalled())
			require.Nil(t, gotByt)
			require.Nil(t, gotErr)

			ctrl.Finish()
		})
	}
}

func TestPrivacyMarker_Run_NonZeroEVMDepth_DoesNothing(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	privacyMarkerTx := types.NewTransaction(0, common.QuorumPrivacyPrecompileContractAddress(), nil, 0, nil, []byte{})
	require.True(t, privacyMarkerTx.IsPrivacyMarker())

	pm := privacyMarker{}

	publicState := NewMockStateDB(ctrl)
	innerApplier := &stubInnerApplier{}

	depth := 1

	evm := &EVM{
		depth:       depth,
		currentTx:   privacyMarkerTx,
		publicState: publicState,
		InnerApply:  innerApplier.InnerApply,
	}

	publicState.EXPECT().SetNonce(gomock.Any(), gomock.Any()).Times(0)

	gotByt, gotErr := pm.Run(evm, []byte{})

	require.False(t, innerApplier.wasCalled())
	require.Nil(t, gotByt)
	require.Nil(t, gotErr)
}

func TestPrivacyMarker_Run_InvalidTransaction_NonceUnchanged(t *testing.T) {
	var (
		publicTx                      *types.Transaction
		publicTxByt                   []byte
		unsignedPrivateTx             *types.Transaction
		incorrectlySignedPrivateTx    *types.Transaction
		incorrectlySignedPrivateTxByt []byte
		setupErr                      error
	)

	publicTx = types.NewTransaction(0, common.QuorumPrivacyPrecompileContractAddress(), nil, 0, nil, []byte{})
	require.False(t, publicTx.IsPrivate())
	publicTxByt, setupErr = json.Marshal(publicTx)
	if setupErr != nil {
		t.Fatalf("unable to marshal tx to json, err = %v", setupErr)
	}

	unsignedPrivateTx = types.NewTransaction(0, common.QuorumPrivacyPrecompileContractAddress(), nil, 0, nil, []byte{})
	unsignedPrivateTx.SetPrivate()

	invalidSig := common.Hex2Bytes("9bea4c4daac7c7c52e093e6a4c35dbbcf8856f1af7b059ba20253e70848d094f8a8fae537ce25ed8cb5af9adac3f141af69bd515bd2ba031522df09b97dd72b100")
	incorrectlySignedPrivateTx, setupErr = unsignedPrivateTx.WithSignature(
		types.QuorumPrivateTxSigner{},
		invalidSig,
	)
	if setupErr != nil {
		t.Fatalf("unable to sign tx, err = %v", setupErr)
	}
	require.True(t, incorrectlySignedPrivateTx.IsPrivate())
	incorrectlySignedPrivateTxByt, setupErr = json.Marshal(incorrectlySignedPrivateTx)
	if setupErr != nil {
		t.Fatalf("unable to marshal tx to json, err = %v", setupErr)
	}

	tests := []struct {
		name               string
		privacyManagerResp []byte // decrypted data from privacy manager
		privacyManagerErr  error
	}{
		{
			name:               "privacy-manager-error",
			privacyManagerResp: nil,
			privacyManagerErr:  errors.New("some error like node is down"),
		},
		{
			name:               "non-participant",
			privacyManagerResp: nil,
			privacyManagerErr:  nil,
		},
		{
			name:               "internal-tx-is-not-private",
			privacyManagerResp: publicTxByt,
			privacyManagerErr:  nil,
		},
		{
			name:               "internal-private-tx-has-invalid-signature",
			privacyManagerResp: incorrectlySignedPrivateTxByt,
			privacyManagerErr:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			unsignedPrivacyMarkerTx := types.NewTransaction(0, common.QuorumPrivacyPrecompileContractAddress(), nil, 0, nil, pmtData)
			signer := types.HomesteadSigner{}
			txHash := signer.Hash(unsignedPrivacyMarkerTx)
			validSig, setupErr := crypto.Sign(txHash.Bytes(), senderPrivateKey)
			if setupErr != nil {
				t.Fatalf("unable to sign tx, err = %v", setupErr)
			}
			privacyMarkerTx, setupErr := unsignedPrivacyMarkerTx.WithSignature(
				signer,
				validSig,
			)
			if setupErr != nil {
				t.Fatalf("unable to sign tx, err = %v", setupErr)
			}

			require.True(t, privacyMarkerTx.IsPrivacyMarker())

			pm := privacyMarker{}

			privacyManager := mock_private.NewMockPrivateTransactionManager(ctrl)
			private.P = privacyManager
			publicState := NewMockStateDB(ctrl)
			innerApplier := nonceIncrementingInnerApplier{
				incrementNonceFunc: func() {
					// this should not be called
					publicState.SetNonce(sender, 1)
				},
			}

			evm := &EVM{
				currentTx:   privacyMarkerTx,
				publicState: publicState,
				InnerApply:  innerApplier.InnerApply,
			}

			privacyManager.EXPECT().Receive(tmPrivateTxHash).Return("", []string{}, tt.privacyManagerResp, nil, tt.privacyManagerErr)

			publicState.EXPECT().GetNonce(gomock.Any()).Times(0)
			publicState.EXPECT().SetNonce(gomock.Any(), gomock.Any()).Times(0)

			gotByt, gotErr := pm.Run(evm, []byte{})

			require.False(t, innerApplier.wasCalled())
			require.Nil(t, gotByt)
			require.Nil(t, gotErr)

			defer ctrl.Finish()
		})
	}
}

func TestPrivacyMarker_Run_SupportedTransaction_ExecutionFails_NonceUnchanged(t *testing.T) {
	var (
		unsignedPrivateTx  *types.Transaction
		signedPrivateTx    *types.Transaction
		signedPrivateTxByt []byte
		setupErr           error
	)

	unsignedPrivateTx = types.NewTransaction(0, common.QuorumPrivacyPrecompileContractAddress(), nil, 0, nil, []byte{})
	unsignedPrivateTx.SetPrivate()

	signer := types.QuorumPrivateTxSigner{}
	txHash := signer.Hash(unsignedPrivateTx)

	validSig, setupErr := crypto.Sign(txHash.Bytes(), senderPrivateKey)
	if setupErr != nil {
		t.Fatalf("unable to sign tx, err = %v", setupErr)
	}

	signedPrivateTx, setupErr = unsignedPrivateTx.WithSignature(
		types.QuorumPrivateTxSigner{},
		validSig,
	)
	if setupErr != nil {
		t.Fatalf("unable to sign tx, err = %v", setupErr)
	}
	require.True(t, signedPrivateTx.IsPrivate())
	signedPrivateTxByt, setupErr = json.Marshal(signedPrivateTx)
	if setupErr != nil {
		t.Fatalf("unable to marshal tx to json, err = %v", setupErr)
	}

	tests := []struct {
		name         string
		innerApplier innerApplier
	}{
		{
			name:         "internal-private-tx-execution-fails",
			innerApplier: &failingInnerApplier{},
		},
		{
			name:         "internal-private-tx-execution-does-not-increment-nonce",
			innerApplier: &stubInnerApplier{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			unsignedPrivacyMarkerTx := types.NewTransaction(0, common.QuorumPrivacyPrecompileContractAddress(), nil, 0, nil, pmtData)
			signer := types.HomesteadSigner{}
			txHash := signer.Hash(unsignedPrivacyMarkerTx)
			validSig, setupErr := crypto.Sign(txHash.Bytes(), senderPrivateKey)
			if setupErr != nil {
				t.Fatalf("unable to sign tx, err = %v", setupErr)
			}
			privacyMarkerTx, setupErr := unsignedPrivacyMarkerTx.WithSignature(
				signer,
				validSig,
			)
			if setupErr != nil {
				t.Fatalf("unable to sign tx, err = %v", setupErr)
			}

			require.True(t, privacyMarkerTx.IsPrivacyMarker())

			pm := privacyMarker{}

			privacyManager := mock_private.NewMockPrivateTransactionManager(ctrl)
			private.P = privacyManager
			publicState := NewMockStateDB(ctrl)

			evm := &EVM{
				currentTx:   privacyMarkerTx,
				publicState: publicState,
				InnerApply:  tt.innerApplier.InnerApply,
			}

			var (
				senderCurrentNonce  uint64 = 10
				senderPreviousNonce uint64 = 9
			)

			privacyManager.EXPECT().Receive(tmPrivateTxHash).Return("", []string{}, signedPrivateTxByt, nil, nil)

			gomock.InOrder(
				publicState.EXPECT().GetNonce(sender).Return(senderCurrentNonce).Times(1), // getting startingNonce
				publicState.EXPECT().SetNonce(sender, senderPreviousNonce).Times(1),       // decrementing nonce to prepare for pvt tx execution
				publicState.EXPECT().SetNonce(sender, senderCurrentNonce).Times(1),        // resetting nonce to startingNonce
			)

			gotByt, gotErr := pm.Run(evm, []byte{})

			require.True(t, tt.innerApplier.wasCalled())
			require.Nil(t, gotByt)
			require.Nil(t, gotErr)

			executedTx := tt.innerApplier.innerTx()

			// we only want to compare the values that matter in the embedded txdata - this is unexported so we resort to
			// using the string representation of the txs for comparison
			require.EqualValues(t, signedPrivateTx.String(), executedTx.String())

			defer ctrl.Finish()
		})
	}
}

func TestPrivacyMarker_Run_SupportedTransaction_ExecutionSucceeds_NonceUnchanged(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		unsignedPrivateTx  *types.Transaction
		signedPrivateTx    *types.Transaction
		signedPrivateTxByt []byte
		setupErr           error
	)

	unsignedPrivateTx = types.NewTransaction(0, common.QuorumPrivacyPrecompileContractAddress(), nil, 0, nil, []byte{})
	unsignedPrivateTx.SetPrivate()

	signer := types.QuorumPrivateTxSigner{}
	txHash := signer.Hash(unsignedPrivateTx)

	validSig, setupErr := crypto.Sign(txHash.Bytes(), senderPrivateKey)
	if setupErr != nil {
		t.Fatalf("unable to sign tx, err = %v", setupErr)
	}

	signedPrivateTx, setupErr = unsignedPrivateTx.WithSignature(
		types.QuorumPrivateTxSigner{},
		validSig,
	)
	if setupErr != nil {
		t.Fatalf("unable to sign tx, err = %v", setupErr)
	}
	require.True(t, signedPrivateTx.IsPrivate())
	signedPrivateTxByt, setupErr = json.Marshal(signedPrivateTx)
	if setupErr != nil {
		t.Fatalf("unable to marshal tx to json, err = %v", setupErr)
	}

	unsignedPrivacyMarkerTx := types.NewTransaction(0, common.QuorumPrivacyPrecompileContractAddress(), nil, 0, nil, pmtData)
	ptmSigner := types.HomesteadSigner{}
	ptmHash := ptmSigner.Hash(unsignedPrivacyMarkerTx)
	validSig, setupErr = crypto.Sign(ptmHash.Bytes(), senderPrivateKey)
	if setupErr != nil {
		t.Fatalf("unable to sign tx, err = %v", setupErr)
	}
	privacyMarkerTx, setupErr := unsignedPrivacyMarkerTx.WithSignature(
		ptmSigner,
		validSig,
	)
	if setupErr != nil {
		t.Fatalf("unable to sign tx, err = %v", setupErr)
	}

	require.True(t, privacyMarkerTx.IsPrivacyMarker())

	pm := privacyMarker{}

	privacyManager := mock_private.NewMockPrivateTransactionManager(ctrl)
	private.P = privacyManager
	publicState := NewMockStateDB(ctrl)

	var (
		senderCurrentNonce  uint64 = 10
		senderPreviousNonce uint64 = 9
	)

	innerApplier := nonceIncrementingInnerApplier{
		incrementNonceFunc: func() {
			publicState.SetNonce(sender, senderPreviousNonce+1)
		},
	}

	evm := &EVM{
		currentTx:   privacyMarkerTx,
		publicState: publicState,
		InnerApply:  innerApplier.InnerApply,
	}

	privacyManager.EXPECT().Receive(tmPrivateTxHash).Return("", []string{}, signedPrivateTxByt, nil, nil)

	gomock.InOrder(
		publicState.EXPECT().GetNonce(sender).Return(senderCurrentNonce).Times(1), // getting startingNonce
		publicState.EXPECT().SetNonce(sender, senderPreviousNonce).Times(1),       // decrementing nonce to prepare for pvt tx execution
		publicState.EXPECT().SetNonce(sender, senderCurrentNonce).Times(1),        // the call in nonceIncrementingInnerApplier
		publicState.EXPECT().SetNonce(sender, senderCurrentNonce).Times(1),        // resetting nonce to startingNonce
	)

	gotByt, gotErr := pm.Run(evm, []byte{})

	require.True(t, innerApplier.wasCalled())
	require.Nil(t, gotByt)
	require.Nil(t, gotErr)

	executedTx := innerApplier.innerTx()

	// we only want to compare the values the matter in the embedded txdata - this is unexported so we resort to
	// using the string representation of the txs for comparison
	require.EqualValues(t, signedPrivateTx.String(), executedTx.String())
}

type innerApplier interface {
	InnerApply(innerTx *types.Transaction) error
	wasCalled() bool
	innerTx() *types.Transaction
}

type stubInnerApplier struct {
	called bool
	tx     *types.Transaction
}

func (m *stubInnerApplier) InnerApply(innerTx *types.Transaction) error {
	m.called = true
	m.tx = innerTx
	return nil
}

func (m *stubInnerApplier) wasCalled() bool {
	return m.called
}

func (m *stubInnerApplier) innerTx() *types.Transaction {
	return m.tx
}

type failingInnerApplier struct {
	called bool
	tx     *types.Transaction
}

func (m *failingInnerApplier) InnerApply(innerTx *types.Transaction) error {
	m.called = true
	m.tx = innerTx
	return errors.New("some error")
}

func (m *failingInnerApplier) wasCalled() bool {
	return m.called
}

func (m *failingInnerApplier) innerTx() *types.Transaction {
	return m.tx
}

type nonceIncrementingInnerApplier struct {
	called             bool
	tx                 *types.Transaction
	incrementNonceFunc func()
}

func (m *nonceIncrementingInnerApplier) InnerApply(innerTx *types.Transaction) error {
	m.called = true
	m.tx = innerTx

	m.incrementNonceFunc()

	return nil
}

func (m *nonceIncrementingInnerApplier) wasCalled() bool {
	return m.called
}

func (m *nonceIncrementingInnerApplier) innerTx() *types.Transaction {
	return m.tx
}

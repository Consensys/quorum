package pluggable

import (
	"math/big"
	"math/rand"
	"testing"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/pluggable/internal/testutils/mock_plugin"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	scheme = "scheme"

	wltUrl = accounts.URL{
		Scheme: scheme,
		Path:   "uripath",
	}

	acct1 = accounts.Account{
		Address: common.HexToAddress("0x4d6d744b6da435b5bbdde2526dc20e9a41cb72e5"),
		URL:     wltUrl,
	}

	acct2 = accounts.Account{
		Address: common.HexToAddress("0x2332f90a329c2c55ba120b1449d36a144d1f9fe4"),
		URL:     accounts.URL{Scheme: scheme, Path: "path/to/file2.json"},
	}
	acct3 = accounts.Account{
		Address: common.HexToAddress("0x992d7a8fca612c963796ecbfe78b300370b9545a"),
		URL:     accounts.URL{Scheme: scheme, Path: "path/to/file3.json"},
	}
	acct4 = accounts.Account{
		Address: common.HexToAddress("0x39ac8f3ae3681b4422fdf808ae18ba4365e37da8"),
		URL:     accounts.URL{Scheme: scheme, Path: "path/to/file4.json"},
	}
)

func validWallet(m *mock_plugin.MockService) *wallet {
	return &wallet{
		url:           wltUrl,
		pluginService: m,
	}
}

func TestWallet_Url(t *testing.T) {
	w := validWallet(nil)
	got := w.URL()
	assert.Equal(t, wltUrl, got)
}

func TestWallet_Status(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	want := "status"

	mockClient := mock_plugin.NewMockService(ctrl)
	mockClient.
		EXPECT().
		Status(gomock.Any()).
		Return(want, nil)

	w := validWallet(mockClient)
	status, err := w.Status()

	assert.NoError(t, err)
	assert.Equal(t, want, status)
}

func TestWallet_Open(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_plugin.NewMockService(ctrl)
	mockClient.
		EXPECT().
		Open(gomock.Any(), "pwd").
		Return(nil)

	w := validWallet(mockClient)
	err := w.Open("pwd")

	assert.NoError(t, err)
}

func TestWallet_Close(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_plugin.NewMockService(ctrl)
	mockClient.
		EXPECT().
		Close(gomock.Any()).
		Return(nil)

	w := validWallet(mockClient)
	err := w.Close()

	assert.NoError(t, err)
}

func TestWallet_Accounts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	want := []accounts.Account{acct1, acct2, acct3, acct4}

	mockClient := mock_plugin.NewMockService(ctrl)
	mockClient.
		EXPECT().
		Accounts(gomock.Any()).
		Return(want)

	w := validWallet(mockClient)
	got := w.Accounts()

	assert.Equal(t, want, got)
}

func TestWallet_Contains(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_plugin.NewMockService(ctrl)
	mockClient.
		EXPECT().
		Contains(gomock.Any(), acct1).
		Return(true)

	w := validWallet(mockClient)
	got := w.Contains(acct1)

	assert.True(t, got)
}

func TestWallet_Derive(t *testing.T) {
	w := validWallet(nil)
	_, err := w.Derive(accounts.DerivationPath{}, true)
	if assert.Error(t, err) {
		assert.Equal(t, accounts.ErrNotSupported, err)
	}
}

func TestWallet_SelfDerive(t *testing.T) {
	w := validWallet(nil)
	// does nothing
	w.SelfDerive([]accounts.DerivationPath{}, nil)
}

func TestWallet_SignData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	toSign := []byte("somedata")
	want := []byte("signeddata")

	mockClient := mock_plugin.NewMockService(ctrl)
	mockClient.
		EXPECT().
		Sign(gomock.Any(), acct1, crypto.Keccak256(toSign)).
		Return(want, nil)

	w := validWallet(mockClient)
	got, err := w.SignData(acct1, "", toSign)

	assert.NoError(t, err)
	assert.Equal(t, want, got)
}

func TestWallet_SignDataWithPassphrase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	toSign := []byte("somedata")
	want := []byte("signeddata")

	mockClient := mock_plugin.NewMockService(ctrl)
	mockClient.
		EXPECT().
		UnlockAndSign(gomock.Any(), acct1, crypto.Keccak256(toSign), "pwd").
		Return(want, nil)

	w := validWallet(mockClient)
	got, err := w.SignDataWithPassphrase(acct1, "pwd", "", toSign)

	assert.NoError(t, err)
	assert.Equal(t, want, got)
}

func TestWallet_SignText(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	toSign := []byte("somedata")
	want := []byte("signeddata")

	mockClient := mock_plugin.NewMockService(ctrl)
	mockClient.
		EXPECT().
		Sign(gomock.Any(), acct1, accounts.TextHash(toSign)).
		Return(want, nil)

	w := validWallet(mockClient)
	got, err := w.SignText(acct1, toSign)

	assert.NoError(t, err)
	assert.Equal(t, want, got)
}

func TestWallet_SignTextWithPassphrase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	toSign := []byte("somedata")
	want := []byte("signeddata")

	mockClient := mock_plugin.NewMockService(ctrl)
	mockClient.
		EXPECT().
		UnlockAndSign(gomock.Any(), acct1, accounts.TextHash(toSign), "pwd").
		Return(want, nil)

	w := validWallet(mockClient)
	got, err := w.SignTextWithPassphrase(acct1, "pwd", toSign)

	assert.NoError(t, err)
	assert.Equal(t, want, got)
}

func TestWallet_SignTx(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name      string
		isPrivate bool
		chainID   *big.Int
		signer    types.Signer
	}{
		{
			name:      "Public EIP155 tx",
			isPrivate: false,
			chainID:   big.NewInt(20),
			signer:    types.NewEIP155Signer(big.NewInt(20)),
		},
		{
			name:      "Public Homestead tx",
			isPrivate: false,
			chainID:   nil,
			signer:    types.HomesteadSigner{},
		},
		{
			name:      "Private tx",
			isPrivate: true,
			chainID:   nil,
			signer:    types.QuorumPrivateTxSigner{},
		},
	}

	toSign := types.NewTransaction(
		1,
		common.HexToAddress("0x2332f90a329c2c55ba120b1449d36a144d1f9fe4"),
		big.NewInt(1),
		0,
		big.NewInt(1),
		nil,
	)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.isPrivate {
				toSign.SetPrivate()
			}

			hashToSign := tt.signer.Hash(toSign)

			mockSig := make([]byte, 65)
			rand.Read(mockSig)

			mockClient := mock_plugin.NewMockService(ctrl)
			mockClient.
				EXPECT().
				Sign(gomock.Any(), acct1, hashToSign.Bytes()).
				Return(mockSig, nil)

			w := validWallet(mockClient)
			got, err := w.SignTx(acct1, toSign, tt.chainID)
			require.NoError(t, err)

			gotV, gotR, gotS := got.RawSignatureValues()

			wantR, wantS, wantV, err := tt.signer.SignatureValues(&types.Transaction{}, mockSig) // tx param is unused by method
			require.NoError(t, err)

			// assert the correct signature is added to the tx
			assert.Equal(t, wantV, gotV)
			assert.Equal(t, wantR, gotR)
			assert.Equal(t, wantS, gotS)

			// assert the rest of the tx is unchanged
			assert.Equal(t, toSign.Nonce(), got.Nonce())
			assert.Equal(t, toSign.GasPrice(), got.GasPrice())
			assert.Equal(t, toSign.Gas(), got.Gas())
			assert.Equal(t, toSign.To(), got.To())
			assert.Equal(t, toSign.Value(), got.Value())
			assert.Equal(t, toSign.Data(), got.Data())
		})
	}
}

func TestWallet_SignTxWithPassphrase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name      string
		isPrivate bool
		chainID   *big.Int
		signer    types.Signer
	}{
		{
			name:      "Public EIP155 tx",
			isPrivate: false,
			chainID:   big.NewInt(20),
			signer:    types.NewEIP155Signer(big.NewInt(20)),
		},
		{
			name:      "Public Homestead tx",
			isPrivate: false,
			chainID:   nil,
			signer:    types.HomesteadSigner{},
		},
		{
			name:      "Private tx",
			isPrivate: true,
			chainID:   nil,
			signer:    types.QuorumPrivateTxSigner{},
		},
	}

	toSign := types.NewTransaction(
		1,
		common.HexToAddress("0x2332f90a329c2c55ba120b1449d36a144d1f9fe4"),
		big.NewInt(1),
		0,
		big.NewInt(1),
		nil,
	)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.isPrivate {
				toSign.SetPrivate()
			}

			hashToSign := tt.signer.Hash(toSign)

			mockSig := make([]byte, 65)
			rand.Read(mockSig)

			mockClient := mock_plugin.NewMockService(ctrl)
			mockClient.
				EXPECT().
				UnlockAndSign(gomock.Any(), acct1, hashToSign.Bytes(), "pwd").
				Return(mockSig, nil)

			w := validWallet(mockClient)
			got, err := w.SignTxWithPassphrase(acct1, "pwd", toSign, tt.chainID)
			require.NoError(t, err)

			gotV, gotR, gotS := got.RawSignatureValues()

			wantR, wantS, wantV, err := tt.signer.SignatureValues(&types.Transaction{}, mockSig) // tx param is unused by method
			require.NoError(t, err)

			// assert the correct signature is added to the tx
			assert.Equal(t, wantV, gotV)
			assert.Equal(t, wantR, gotR)
			assert.Equal(t, wantS, gotS)

			// assert the rest of the tx is unchanged
			assert.Equal(t, toSign.Nonce(), got.Nonce())
			assert.Equal(t, toSign.GasPrice(), got.GasPrice())
			assert.Equal(t, toSign.Gas(), got.Gas())
			assert.Equal(t, toSign.To(), got.To())
			assert.Equal(t, toSign.Value(), got.Value())
			assert.Equal(t, toSign.Data(), got.Data())
		})
	}
}

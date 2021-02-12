package account

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/plugin/account/internal/testutils"
	"github.com/golang/mock/gomock"
	"github.com/jpmorganchase/quorum-account-plugin-sdk-go/mock_proto"
	"github.com/jpmorganchase/quorum-account-plugin-sdk-go/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	scheme = "scheme"

	acct1 = accounts.Account{
		Address: common.HexToAddress("0x4d6d744b6da435b5bbdde2526dc20e9a41cb72e5"),
		URL:     accounts.URL{Scheme: scheme, Path: "acctUri1"},
	}

	acct2 = accounts.Account{
		Address: common.HexToAddress("0x2332f90a329c2c55ba120b1449d36a144d1f9fe4"),
		URL:     accounts.URL{Scheme: scheme, Path: "acctUri2"},
	}
)

func TestPluginGateway_Status(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	resp := &proto.StatusResponse{Status: "some status"}

	wantReq := &proto.StatusRequest{}
	wantStatus := resp.Status

	mockClient := mock_proto.NewMockAccountServiceClient(ctrl)
	mockClient.
		EXPECT().
		Status(gomock.Any(), testutils.StatusRequestMatcher{R: wantReq}).
		Return(resp, nil)

	g := &service{client: mockClient}
	got, err := g.Status(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, wantStatus, got)
}

func TestPluginGateway_Open(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	resp := &proto.OpenResponse{}

	wantReq := &proto.OpenRequest{Passphrase: "pwd"}

	mockClient := mock_proto.NewMockAccountServiceClient(ctrl)
	mockClient.
		EXPECT().
		Open(gomock.Any(), testutils.OpenRequestMatcher{R: wantReq}).
		Return(resp, nil)

	g := &service{client: mockClient}
	err := g.Open(context.Background(), "pwd")

	assert.NoError(t, err)
}

func TestPluginGateway_Close(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	resp := &proto.CloseResponse{}

	wantReq := &proto.CloseRequest{}

	mockClient := mock_proto.NewMockAccountServiceClient(ctrl)
	mockClient.
		EXPECT().
		Close(gomock.Any(), testutils.CloseRequestMatcher{R: wantReq}).
		Return(resp, nil)

	g := &service{client: mockClient}
	err := g.Close(context.Background())

	assert.NoError(t, err)
}

func TestPluginGateway_Accounts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	resp := &proto.AccountsResponse{
		Accounts: []*proto.Account{
			{Address: acct1.Address.Bytes(), Url: acct1.URL.String()},
			{Address: acct2.Address.Bytes(), Url: acct2.URL.String()},
		},
	}

	wantReq := &proto.AccountsRequest{}
	wantAccts := []accounts.Account{acct1, acct2}

	mockClient := mock_proto.NewMockAccountServiceClient(ctrl)
	mockClient.
		EXPECT().
		Accounts(gomock.Any(), testutils.AccountsRequestMatcher{R: wantReq}).
		Return(resp, nil)

	g := &service{client: mockClient}
	got := g.Accounts(context.Background())

	assert.Equal(t, wantAccts, got)
}

func TestPluginGateway_Contains(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	resp := &proto.ContainsResponse{IsContained: true}

	wantReq := &proto.ContainsRequest{
		Address: acct1.Address.Bytes(),
	}

	mockClient := mock_proto.NewMockAccountServiceClient(ctrl)
	mockClient.
		EXPECT().
		Contains(gomock.Any(), testutils.ContainsRequestMatcher{R: wantReq}).
		Return(resp, nil)

	g := &service{client: mockClient}
	got := g.Contains(context.Background(), acct1)

	assert.True(t, got)
}

func TestPluginGateway_Sign(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	want := []byte("signed data")
	resp := &proto.SignResponse{Sig: want}

	toSign := []byte("to sign")
	wantReq := &proto.SignRequest{
		Address: acct1.Address.Bytes(),
		ToSign:  toSign,
	}

	mockClient := mock_proto.NewMockAccountServiceClient(ctrl)
	mockClient.
		EXPECT().
		Sign(gomock.Any(), testutils.SignRequestMatcher{R: wantReq}).
		Return(resp, nil)

	g := &service{client: mockClient}
	got, err := g.Sign(context.Background(), acct1, toSign)

	assert.NoError(t, err)
	assert.Equal(t, want, got)
}

func TestPluginGateway_UnlockAndSign(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	want := []byte("signed data")
	resp := &proto.SignResponse{Sig: want}

	toSign := []byte("to sign")
	wantReq := &proto.UnlockAndSignRequest{
		Address:    acct1.Address.Bytes(),
		ToSign:     toSign,
		Passphrase: "pwd",
	}

	mockClient := mock_proto.NewMockAccountServiceClient(ctrl)
	mockClient.
		EXPECT().
		UnlockAndSign(gomock.Any(), testutils.UnlockAndSignRequestMatcher{R: wantReq}).
		Return(resp, nil)

	g := &service{client: mockClient}
	got, err := g.UnlockAndSign(context.Background(), acct1, toSign, "pwd")

	assert.NoError(t, err)
	assert.Equal(t, want, got)
}

func TestPluginGateway_TimedUnlock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	resp := &proto.TimedUnlockResponse{}

	pwd := "pwd"

	wantReq := &proto.TimedUnlockRequest{
		Address:  acct1.Address.Bytes(),
		Password: pwd,
		Duration: int64(1),
	}

	mockClient := mock_proto.NewMockAccountServiceClient(ctrl)
	mockClient.
		EXPECT().
		TimedUnlock(gomock.Any(), testutils.TimedUnlockRequestMatcher{R: wantReq}).
		Return(resp, nil)

	g := &service{client: mockClient}
	err := g.TimedUnlock(context.Background(), acct1, pwd, time.Nanosecond)

	assert.NoError(t, err)
}

func TestPluginGateway_Lock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	resp := &proto.LockResponse{}

	wantReq := &proto.LockRequest{
		Address: acct1.Address.Bytes(),
	}

	mockClient := mock_proto.NewMockAccountServiceClient(ctrl)
	mockClient.
		EXPECT().
		Lock(gomock.Any(), testutils.LockRequestMatcher{R: wantReq}).
		Return(resp, nil)

	g := &service{client: mockClient}
	err := g.Lock(context.Background(), acct1)

	assert.NoError(t, err)
}

func TestPluginGateway_NewAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	resp := &proto.NewAccountResponse{
		Account: &proto.Account{
			Address: acct1.Address.Bytes(),
			Url:     acct1.URL.String(),
		},
	}

	newAccountConfig := []byte("newacctconfig")

	b, err := json.Marshal(newAccountConfig)
	require.NoError(t, err)
	wantReq := &proto.NewAccountRequest{
		NewAccountConfig: b,
	}

	mockClient := mock_proto.NewMockAccountServiceClient(ctrl)
	mockClient.
		EXPECT().
		NewAccount(gomock.Any(), testutils.NewAccountRequestMatcher{R: wantReq}).
		Return(resp, nil)

	g := &service{client: mockClient}
	gotAcct, err := g.NewAccount(context.Background(), newAccountConfig)

	assert.Equal(t, acct1, gotAcct)
	assert.NoError(t, err)
}

func TestPluginGateway_ImportRawKey(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	resp := &proto.ImportRawKeyResponse{
		Account: &proto.Account{
			Address: acct1.Address.Bytes(),
			Url:     acct1.URL.String(),
		},
	}

	newAccountConfig := []byte("newacctconfig")
	var rawKey = "1fe8f1ad4053326db20529257ac9401f2e6c769ef1d736b8c2f5aba5f787c72b"

	b, err := json.Marshal(newAccountConfig)
	require.NoError(t, err)
	wantReq := &proto.ImportRawKeyRequest{
		RawKey:           rawKey,
		NewAccountConfig: b,
	}

	mockClient := mock_proto.NewMockAccountServiceClient(ctrl)
	mockClient.
		EXPECT().
		ImportRawKey(gomock.Any(), testutils.ImportRawKeyRequestMatcher{R: wantReq}).
		Return(resp, nil)

	g := &service{client: mockClient}
	gotAcct, err := g.ImportRawKey(context.Background(), rawKey, newAccountConfig)

	assert.Equal(t, acct1, gotAcct)
	assert.NoError(t, err)
}

func TestPluginGateway_ImportRawKey_InvalidRawKey(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	newAccountConfig := []byte("newacctconfig")
	var rawKey = "aaaaaa"

	g := &service{}
	_, err := g.ImportRawKey(context.Background(), rawKey, newAccountConfig)

	require.EqualError(t, err, "invalid length, need 256 bits")
}

func Test_ToUrl(t *testing.T) {
	strUrl := "http://myurl:8000"
	want := accounts.URL{
		Scheme: "http",
		Path:   "myurl:8000",
	}
	got, err := ToUrl(strUrl)
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func Test_ToUrl_Invalid(t *testing.T) {
	strUrl := "://noscheme:8000"
	_, err := ToUrl(strUrl)
	require.Error(t, err)
}

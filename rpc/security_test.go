package rpc

import (
	"context"
	"errors"
	"net/http"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/golang/protobuf/ptypes"
	"github.com/jpmorganchase/quorum-security-plugin-sdk-go/proto"
	testifyassert "github.com/stretchr/testify/assert"
)

func TestVerifyAccess_whenNotMatch(t *testing.T) {
	assert := testifyassert.New(t)

	assert.Error(verifyAccess("xyz", "abc", []*proto.GrantedAuthority{
		{
			Service: "bar",
			Method:  "foo",
		},
	}))
}

func TestVerifyAccess_whenEmpty(t *testing.T) {
	assert := testifyassert.New(t)

	assert.Error(verifyAccess("xyz", "abc", nil))
}

func TestVerifyAccess_whenExactMatch(t *testing.T) {
	assert := testifyassert.New(t)

	assert.NoError(verifyAccess("bar", "foo", []*proto.GrantedAuthority{
		{
			Service: "xyz",
			Method:  "abc",
		},
		{
			Service: "bar",
			Method:  "foo",
		},
	}))
}

func TestVerifyAccess_whenWildcardServiceMatch(t *testing.T) {
	assert := testifyassert.New(t)

	assert.NoError(verifyAccess("bar", "foo", []*proto.GrantedAuthority{
		{
			Service: "xyz",
			Method:  "abc",
		},
		{
			Service: "*",
			Method:  "foo",
		},
	}))
}

func TestVerifyAccess_whenWildcardMethodMatch(t *testing.T) {
	assert := testifyassert.New(t)

	assert.NoError(verifyAccess("bar", "foo", []*proto.GrantedAuthority{
		{
			Service: "xyz",
			Method:  "abc",
		},
		{
			Service: "bar",
			Method:  "*",
		},
	}))
}

func TestVerifyAccess_whenWildcardMatch(t *testing.T) {
	assert := testifyassert.New(t)

	assert.NoError(verifyAccess("bar", "foo", []*proto.GrantedAuthority{
		{
			Service: "*",
			Method:  "*",
		},
	}))
}

func TestVerifyExpiration_whenTypical(t *testing.T) {
	assert := testifyassert.New(t)
	expiredAt, _ := ptypes.TimestampProto(time.Now().Add(1 * time.Minute))
	assert.NoError(verifyExpiration(&proto.PreAuthenticatedAuthenticationToken{
		ExpiredAt: expiredAt,
	}))
}

func TestVerifyExpiration_whenExpired(t *testing.T) {
	assert := testifyassert.New(t)
	expiredAt, _ := ptypes.TimestampProto(time.Now().Add(-1 * time.Minute))
	assert.Error(verifyExpiration(&proto.PreAuthenticatedAuthenticationToken{
		ExpiredAt: expiredAt,
	}))
}

func TestExtractToken_whenTypical(t *testing.T) {
	assert := testifyassert.New(t)
	req, _ := http.NewRequest("POST", "", nil)
	arbitraryValue := "xyz"
	req.Header.Set(HttpAuthorizationHeader, arbitraryValue)

	token, ok := extractToken(req)

	assert.True(ok)
	assert.Equal(arbitraryValue, token)
}

func TestExtractToken_whenEmpty(t *testing.T) {
	assert := testifyassert.New(t)
	req, _ := http.NewRequest("POST", "", nil)

	_, ok := extractToken(req)

	assert.False(ok)
}

func TestSecureCall_whenThereIsAuthenticationError(t *testing.T) {
	assert := testifyassert.New(t)
	arbitraryError := errors.New("arbitrary error")
	stubSecurityContextResolver := newStubSecurityContextResolver([]struct{ k, v interface{} }{
		{ctxAuthenticationError, arbitraryError},
	})

	_, err := SecureCall(stubSecurityContextResolver, "")

	assert.EqualError(err, arbitraryError.Error())
}

func TestSecureCall_whenTokenExpired(t *testing.T) {
	assert := testifyassert.New(t)
	expiredAt, _ := ptypes.TimestampProto(time.Now().Add(-1 * time.Hour))
	stubSecurityContextResolver := newStubSecurityContextResolver([]struct{ k, v interface{} }{
		{CtxPreauthenticatedToken, &proto.PreAuthenticatedAuthenticationToken{
			ExpiredAt: expiredAt,
		}},
	})

	_, err := SecureCall(stubSecurityContextResolver, "")

	assert.EqualError(err, "token expired")
}

func TestSecureCall_whenTypical(t *testing.T) {
	assert := testifyassert.New(t)
	expiredAt, _ := ptypes.TimestampProto(time.Now().Add(1 * time.Hour))
	stubSecurityContextResolver := newStubSecurityContextResolver([]struct{ k, v interface{} }{
		{CtxPreauthenticatedToken, &proto.PreAuthenticatedAuthenticationToken{
			ExpiredAt: expiredAt,
			Authorities: []*proto.GrantedAuthority{
				{
					Service: "eth",
					Method:  "blockNumber",
				},
			},
		}},
	})

	_, err := SecureCall(stubSecurityContextResolver, "eth_blockNumber")

	assert.NoError(err)
}

func TestSecureCall_whenAccessDenied(t *testing.T) {
	assert := testifyassert.New(t)
	expiredAt, _ := ptypes.TimestampProto(time.Now().Add(1 * time.Hour))
	stubSecurityContextResolver := newStubSecurityContextResolver([]struct{ k, v interface{} }{
		{CtxPreauthenticatedToken, &proto.PreAuthenticatedAuthenticationToken{
			ExpiredAt: expiredAt,
			Authorities: []*proto.GrantedAuthority{
				{
					Service: "eth",
					Method:  "blockNumber",
				},
			},
		}},
	})

	_, err := SecureCall(stubSecurityContextResolver, "eth_someMethod")

	assert.EqualError(err, "eth_someMethod - access denied")
}

func TestSecureCall_whenMethodInJSONMessageIsNotSupported(t *testing.T) {
	assert := testifyassert.New(t)
	expiredAt, _ := ptypes.TimestampProto(time.Now().Add(1 * time.Hour))
	stubSecurityContextResolver := newStubSecurityContextResolver([]struct{ k, v interface{} }{
		{CtxPreauthenticatedToken, &proto.PreAuthenticatedAuthenticationToken{
			ExpiredAt: expiredAt,
		}},
	})

	_, err := SecureCall(stubSecurityContextResolver, "arbitrary method")

	assert.NoError(err)
}

type stubSecurityContextResolver struct {
	ctx SecurityContext
}

func newStubSecurityContextResolver(ctx []struct{ k, v interface{} }) *stubSecurityContextResolver {
	sc := SecurityContext(context.Background())
	for _, kv := range ctx {
		sc = context.WithValue(sc, kv.k, kv.v)
	}
	return &stubSecurityContextResolver{sc}
}

func (sr *stubSecurityContextResolver) Resolve() SecurityContext {
	return sr.ctx
}

func TestResolvePSIProvider_whenTypicalEndpoints(t *testing.T) {
	testCases := []struct {
		endpoint    string
		expectedPSI types.PrivateStateIdentifier
	}{
		{
			endpoint:    "http://aritraryhost?PSI=PS1",
			expectedPSI: types.ToPrivateStateIdentifier("PS1"),
		},
		{
			endpoint:    "https://aritraryhost?PSI=PS2",
			expectedPSI: types.ToPrivateStateIdentifier("PS2"),
		},
		{
			endpoint:    "ws://aritraryhost?PSI=PS3",
			expectedPSI: types.ToPrivateStateIdentifier("PS3"),
		},
		{
			endpoint:    "wss://aritraryhost?PSI=PS4",
			expectedPSI: types.ToPrivateStateIdentifier("PS4"),
		},
	}
	for _, tc := range testCases {
		actualCtx := resolvePSIProvider(context.Background(), tc.endpoint)

		testifyassert.NotNil(t, actualCtx.Value(CtxPSIProvider))
		f, ok := actualCtx.Value(CtxPSIProvider).(PSIProviderFunc)
		testifyassert.True(t, ok)
		actualPSI, err := f(context.Background())
		testifyassert.NoError(t, err)
		testifyassert.Equal(t, tc.expectedPSI, actualPSI)
	}
}

func TestResolvePSIProvider_whenEnvVariableTakesPrecedence(t *testing.T) {
	_ = os.Setenv(EnvVarPrivateStateIdentifier, "ENV_PS1")
	defer func() { _ = os.Unsetenv(EnvVarPrivateStateIdentifier) }()

	endpoint := "http://aritraryhost?PSI=PS1"
	actualCtx := resolvePSIProvider(context.Background(), endpoint)

	testifyassert.NotNil(t, actualCtx.Value(CtxPSIProvider))
	f, ok := actualCtx.Value(CtxPSIProvider).(PSIProviderFunc)
	testifyassert.True(t, ok)
	actualPSI, err := f(context.Background())
	testifyassert.NoError(t, err)
	testifyassert.Equal(t, types.ToPrivateStateIdentifier("ENV_PS1"), actualPSI)
}

func TestResolvePSIProvider_whenNoPSI(t *testing.T) {
	endpoint := "data/geth.ipc"
	actualCtx := resolvePSIProvider(context.Background(), endpoint)

	testifyassert.Nil(t, actualCtx.Value(CtxPSIProvider))
}

func TestEncodePSI_whenTypical(t *testing.T) {
	actual := encodePSI(strconv.AppendUint(nil, 32, 10), "ARBITRARY")

	testifyassert.Equal(t, "\"ARBITRARY/32\"", string(actual))
}

func TestEncodePSI_whenNoPSI(t *testing.T) {
	actual := encodePSI(strconv.AppendUint(nil, 32, 10), "")

	testifyassert.Equal(t, "32", string(actual))
}

func TestDecodePSI_whenTypical(t *testing.T) {
	input := "\"ARBITRARY/1\""

	psi := decodePSI([]byte(input))

	testifyassert.Equal(t, types.PrivateStateIdentifier("ARBITRARY"), psi)
}

func TestDecodePSI_whenNoPSI(t *testing.T) {
	inputs := []string{
		"1",
		"\"1",
		"1\"",
		"\"xyz\"",
	}
	for _, input := range inputs {
		psi := decodePSI([]byte(input))

		testifyassert.Equal(t, types.DefaultPrivateStateIdentifier, psi, "input: %s", input)
	}
}

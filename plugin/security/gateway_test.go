package security

import (
	"context"
	"crypto/tls"
	"math"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jpmorganchase/quorum-security-plugin-sdk-go/mock_proto"
	"github.com/jpmorganchase/quorum-security-plugin-sdk-go/proto"
	testifyassert "github.com/stretchr/testify/assert"
)

const (
	rsaCertPem = `-----BEGIN CERTIFICATE-----
MIIB0zCCAX2gAwIBAgIJAI/M7BYjwB+uMA0GCSqGSIb3DQEBBQUAMEUxCzAJBgNV
BAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRlcm5ldCBX
aWRnaXRzIFB0eSBMdGQwHhcNMTIwOTEyMjE1MjAyWhcNMTUwOTEyMjE1MjAyWjBF
MQswCQYDVQQGEwJBVTETMBEGA1UECAwKU29tZS1TdGF0ZTEhMB8GA1UECgwYSW50
ZXJuZXQgV2lkZ2l0cyBQdHkgTHRkMFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBANLJ
hPHhITqQbPklG3ibCVxwGMRfp/v4XqhfdQHdcVfHap6NQ5Wok/4xIA+ui35/MmNa
rtNuC+BdZ1tMuVCPFZcCAwEAAaNQME4wHQYDVR0OBBYEFJvKs8RfJaXTH08W+SGv
zQyKn0H8MB8GA1UdIwQYMBaAFJvKs8RfJaXTH08W+SGvzQyKn0H8MAwGA1UdEwQF
MAMBAf8wDQYJKoZIhvcNAQEFBQADQQBJlffJHybjDGxRMqaRmDhX0+6v02TUKZsW
r5QuVbpQhH6u+0UgcW0jp9QwpxoPTLTWGXEWBBBurxFwiCBhkQ+V
-----END CERTIFICATE-----
`
	rsaKeyPem = `-----BEGIN RSA PRIVATE KEY-----
MIIBOwIBAAJBANLJhPHhITqQbPklG3ibCVxwGMRfp/v4XqhfdQHdcVfHap6NQ5Wo
k/4xIA+ui35/MmNartNuC+BdZ1tMuVCPFZcCAwEAAQJAEJ2N+zsR0Xn8/Q6twa4G
6OB1M1WO+k+ztnX/1SvNeWu8D6GImtupLTYgjZcHufykj09jiHmjHx8u8ZZB/o1N
MQIhAPW+eyZo7ay3lMz1V01WVjNKK9QSn1MJlb06h/LuYv9FAiEA25WPedKgVyCW
SmUwbPw8fnTcpqDWE3yTO3vKcebqMSsCIBF3UmVue8YU3jybC3NxuXq3wNm34R8T
xVLHwDXh/6NJAiEAl2oHGGLz64BuAfjKrqwz7qMYr9HCLIe/YsoWq/olzScCIQDi
D2lWusoe2/nEqfDVVWGWlyJ7yOmqaVm/iNUN9B2N2g==
-----END RSA PRIVATE KEY-----
`
)

var (
	abitraryTLSConfigurationData = &proto.TLSConfiguration_Data{
		CertPem: []byte(rsaCertPem),
		KeyPem:  []byte(rsaKeyPem),
	}
)

func TestTransform_whenTypical(t *testing.T) {
	assert := testifyassert.New(t)

	cfg, err := transform(abitraryTLSConfigurationData)

	assert.NoError(err)
	assert.True(cfg.PreferServerCipherSuites)
	assert.EqualValues(defaultCipherSuites, cfg.CipherSuites)
	assert.Equal(uint16(tls.VersionTLS12), cfg.MinVersion)
	assert.EqualValues([]tls.CurveID{
		tls.CurveP521,
		tls.CurveP384,
		tls.CurveP256,
		tls.X25519,
	}, cfg.CurvePreferences)
}

func TestTransform_whenUsingCustomCipherSuites(t *testing.T) {
	defer func() {
		abitraryTLSConfigurationData.CipherSuites = nil
	}()
	assert := testifyassert.New(t)

	abitraryTLSConfigurationData.CipherSuites = []uint32{uint32(tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA)}

	cfg, err := transform(abitraryTLSConfigurationData)

	assert.NoError(err)
	assert.Contains(cfg.CipherSuites, tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA)
}

func TestTransform_whenCipherSuiteOverflow(t *testing.T) {
	defer func() {
		abitraryTLSConfigurationData.CipherSuites = nil
	}()
	assert := testifyassert.New(t)

	abitraryTLSConfigurationData.CipherSuites = []uint32{math.MaxInt32}

	_, err := transform(abitraryTLSConfigurationData)

	assert.Error(err)
}

func TestTLSConfigurationSourcePluginGateway_Get(t *testing.T) {
	assert := testifyassert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockClient := mock_proto.NewMockTLSConfigurationSourceClient(ctrl)
	mockClient.
		EXPECT().
		Get(gomock.Any(), gomock.Any()).
		Return(&proto.TLSConfiguration_Response{
			Data: abitraryTLSConfigurationData,
		}, nil)

	testObject := &TLSConfigurationSourcePluginGateway{client: mockClient}

	tlsConfig, err := testObject.Get(context.Background())

	assert.NoError(err)
	assert.NotNil(tlsConfig)
}

func TestTLSConfigurationSourcePluginGateway_Get_whenNoConfigurationData(t *testing.T) {
	assert := testifyassert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockClient := mock_proto.NewMockTLSConfigurationSourceClient(ctrl)
	mockClient.
		EXPECT().
		Get(gomock.Any(), gomock.Any()).
		Return(&proto.TLSConfiguration_Response{}, nil)

	testObject := &TLSConfigurationSourcePluginGateway{client: mockClient}

	tlsConfig, err := testObject.Get(context.Background())

	assert.NoError(err)
	assert.Nil(tlsConfig)
}

func TestAuthenticationManagerPluginGateway_IsEnabled_always(t *testing.T) {
	testObject := &AuthenticationManagerPluginGateway{}

	ret, err := testObject.IsEnabled(context.Background())

	testifyassert.NoError(t, err)
	testifyassert.True(t, ret)
}

func TestAuthenticationManagerPluginGateway_Authenticate(t *testing.T) {
	assert := testifyassert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	arbitraryPreauthenticatedToken := &proto.AuthenticationToken{
		RawToken: []byte("arbitrary token"),
	}
	mockClient := mock_proto.NewMockAuthenticationManagerClient(ctrl)
	mockClient.
		EXPECT().
		Authenticate(gomock.Any(), gomock.Eq(arbitraryPreauthenticatedToken)).
		Return(&proto.PreAuthenticatedAuthenticationToken{}, nil)

	testObject := &AuthenticationManagerPluginGateway{client: mockClient}

	_, err := testObject.Authenticate(context.Background(), string(arbitraryPreauthenticatedToken.RawToken))

	assert.NoError(err)
}

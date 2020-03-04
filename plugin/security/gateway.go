package security

import (
	"context"
	"crypto/tls"
	"errors"
	"math"

	"github.com/jpmorganchase/quorum-security-plugin-sdk-go/proto"
)

var (
	defaultCipherSuites = []uint16{
		tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
		tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
	}
)

type TLSConfigurationSourcePluginGateway struct {
	client proto.TLSConfigurationSourceClient
}

func (c *TLSConfigurationSourcePluginGateway) Get(ctx context.Context) (*tls.Config, error) {
	resp, err := c.client.Get(ctx, &proto.TLSConfiguration_Request{})
	if err != nil {
		return nil, err
	}
	if resp == nil || resp.GetData() == nil { // no tls config
		return nil, nil
	}
	return transform(resp.GetData())
}

func transform(tlsData *proto.TLSConfiguration_Data) (*tls.Config, error) {
	tlsConfig := &tls.Config{
		// Ensure Key or DH parameter strength >= 4096 bits
		CurvePreferences: []tls.CurveID{
			tls.CurveP521,
			tls.CurveP384,
			tls.CurveP256,
		},
		// Support only TLS1.2 & Above
		MinVersion: tls.VersionTLS12,
	}
	receivedCipherSuites := tlsData.GetCipherSuites()
	cipherSuites := make([]uint16, len(receivedCipherSuites))
	if len(receivedCipherSuites) > 0 {
		for i, cs := range receivedCipherSuites {
			if cs > math.MaxUint16 {
				return nil, errors.New("cipher suite value overflow")
			}
			cipherSuites[i] = uint16(cs)
		}
	} else {
		cipherSuites = defaultCipherSuites
	}
	tlsConfig.CipherSuites = cipherSuites
	tlsConfig.PreferServerCipherSuites = true

	cer, err := tls.X509KeyPair(tlsData.GetCertPem(), tlsData.GetKeyPem())
	if err != nil {
		return nil, err
	}
	tlsConfig.Certificates = []tls.Certificate{cer}

	return tlsConfig, nil
}

type AuthenticationManagerPluginGateway struct {
	client proto.AuthenticationManagerClient
}

func (a *AuthenticationManagerPluginGateway) Authenticate(ctx context.Context, token string) (*proto.PreAuthenticatedAuthenticationToken, error) {
	return a.client.Authenticate(ctx, &proto.AuthenticationToken{
		RawToken: []byte(token),
	})
}

func (a *AuthenticationManagerPluginGateway) IsEnabled(ctx context.Context) (bool, error) {
	return true, nil
}

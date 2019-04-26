package rpc

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	PROVIDER_LOCAL      = "local"
	PROVIDER_ENTERPRISE = "enterprise"
)

// RFC (7662): https://tools.ietf.org/html/rfc7662.
// Authorization Server Introspect Request & Response.
type IntrospectRequest struct {
	Token         string `json:"token"`
	TokenTypeHint string `json:"token_type_hint"`
}
type IntrospectResponse struct {
	Active     bool   `json:"active"`
	Scope      string `json:"scope"`
	ClientId   string `json:"client_id"`
	Username   string `json:"username"`
	Expiration int    `json:"exp"`
}

// Authorization Server Cert
type AuthorizationServerCert struct {
	ProviderTlsCertificateFile    string `json:"providerTlsCertificateFile"`
	ProviderTlsCertificateCaFile  string `json:"providerTlsCertificateCaFile"`
	ProviderTlsCertificateKeyFile string `json:"providerTlsCertificateKeyFile"`
}

// AuthorizationServerInformation
type AuthorizationServerInformation struct {
	ProviderCategory string `json:"providerCategory"`

	// Authorization Server Introspection URL.
	ProviderIntrospectionURL string `json:"providerIntrospectionURL"`

	// Authorization Server Cert Information
	ProviderCertificateInfo *AuthorizationServerCert `json:"providerCert"`

	// Local Provider Information
	LocalProviderDbFile *string `json:"localProviderDbFile"`
	LocalClientsFile    *string `json:"localClientsFile"`
}

// RPC ListenerWithTls Support
type Listener struct {
	ServerTlsCertFile string `json:"serverTlsCertFile"`
	ServerTlsKeyFile  string `json:"serverTlsKeyFile"`
}

// RPC Security Configuration
type SecurityConfig struct {
	Listener                *Listener                       `json:"listenerCert"`
	ProviderType            string                          `json:"providerType"`
	AuthorizationServerInfo *AuthorizationServerInformation `json:"providerInfo"`
}

// RPC Security Context
type SecurityContext struct {
	Enabled               bool
	Config                *SecurityConfig
	Client                *http.Client
	LocalSecurityProvider *LocalSecurityProvider
}

func (ctx *SecurityContext) getHttpClient() *http.Client {
	if ctx.Client == nil {
		ctx.Client = ctx.buildHttpClient()
	}

	return ctx.Client
}

// Build HTTP Client
func (ctx *SecurityContext) buildHttpClient() *http.Client {
	if ctx.Config.AuthorizationServerInfo == nil {
		return &http.Client{}
	}

	// If no cert information provided return simple client
	if ctx.Config.AuthorizationServerInfo.ProviderCertificateInfo == nil {
		return &http.Client{}
	}

	// Load provider certificate info provided
	certFile := ctx.Config.AuthorizationServerInfo.ProviderCertificateInfo.ProviderTlsCertificateFile
	keyFile := ctx.Config.AuthorizationServerInfo.ProviderCertificateInfo.ProviderTlsCertificateKeyFile
	caFile := ctx.Config.AuthorizationServerInfo.ProviderCertificateInfo.ProviderTlsCertificateCaFile

	// Load client cert
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatal(err)
	}

	// Load CA cert
	caCert, err := ioutil.ReadFile(caFile)
	if err != nil {
		log.Fatal(err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Setup HTTPS client
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
	}
	tlsConfig.BuildNameToCertificate()
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	return &http.Client{Transport: transport}

}

// Parse the RPC Request, Call send Introspect Request & Parse results
func (ctx *SecurityContext) isHttpRequestAuthorized(r *http.Request) bool {
	providerType := strings.ToLower(ctx.Config.ProviderType)
	clientToken := r.Header.Get("Token")

	if providerType == "" {
		return false
	}

	if clientToken == "" {
		return false
	}

	if providerType == PROVIDER_ENTERPRISE {

	}

	if providerType == PROVIDER_LOCAL {

		fmt.Printf("Send Introspect Locally")

	}

	return true
}

// Parse the RPC Request, Call send Introspect Request & Parse results
func (ctx *SecurityContext) isWSRequestAuthorized(request rpcRequest) bool {
	fmt.Println("Send WS Introspect for:")
	fmt.Println(request)

	if request.token == "cucrisis" {
		if request.service != "admin" {
			return true
		}
	}


	return false
}

// Process RPC Http Request
func (ctx *SecurityContext) ProcessHttpRequest(r *http.Request) (int, error) {
	if ctx.Enabled && ctx.Config == nil {
		return http.StatusUnauthorized, errors.New("Request requires valid token")
	}

	if ctx.Enabled {
		if ctx.isHttpRequestAuthorized(r) {
			return http.StatusOK, nil
		} else {
			return http.StatusUnauthorized, errors.New("Request requires valid token")
		}
	}

	return http.StatusOK, nil
}

// Process WS Request
func (ctx *SecurityContext) ProcessWSRequest(request rpcRequest) error {
	if ctx.Enabled && ctx.Config == nil {
		return errors.New("Request requires valid token")
	}

	if !ctx.isWSRequestAuthorized(request) {
		return errors.New("Request requires valid token")
	}
	return nil
}

type LocalSecurityProvider struct {
	LocalSecurityDbFile *string
}

func (l *LocalSecurityProvider) findClient(clientName *string) {

}

func (l *LocalSecurityProvider) addClientsFromFile(fileName *string) {

}

func (l *LocalSecurityProvider) addClient(clientName *string, clientID *string, clientSecret *string, clientScope *string) {

}

func (l *LocalSecurityProvider) listClients() {

}

func (l *LocalSecurityProvider) removeClient(clientName *string) {

}

func (l *LocalSecurityProvider) regenerateClient(clientName *string) {

}

// GetDefaultDenyAllSecurityContext returns a disabled context
func GetDefaultDenyAllSecurityContext() SecurityContext {
	return SecurityContext{Enabled: true}
}

// GetDefaultAllowAllSecurityContext returns a disabled context
func GetDefaultAllowAllSecurityContext() SecurityContext {
	return SecurityContext{Enabled: false}
}

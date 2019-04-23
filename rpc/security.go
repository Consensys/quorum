package rpc

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	PROVIDER_LOCAL      = "local"
	PROVIDER_ENTERPRISE = "local"
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
	// Authorization Server Introspection URL.
	ProviderIntrospectionURL string `json:"providerIntrospectionURL"`

	// Authorization Server Cert Information
	ProviderCertificateInfo *AuthorizationServerCert `json:"providerCert"`
}

// RPC Security Configuration
type SecurityConfig struct {
	ProviderType            string                          `json:"providerType"`
	LocalProviderDbFile     string                          `json:"localProviderDbFile"`
	AuthorizationServerInfo *AuthorizationServerInformation `json:"providerInfo"`
}

// RPC Security Context
type SecurityContext struct {
	Enabled bool
	Config  *SecurityConfig
	Client  *http.Client
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

	if providerType == PROVIDER_ENTERPRISE {
		fmt.Printf("Send Introspect Request")
	}

	if strings.ToLower(ctx.Config.ProviderType) == PROVIDER_LOCAL {

	}

	return false
}

// Parse the RPC Request, Call send Introspect Request & Parse results
func (ctx *SecurityContext) isWSRequestAuthorized(r *http.Request) bool {
	return false
}

// Process RPC Http Request
func (ctx *SecurityContext) ProcessHttpRequest(r *http.Request) (int, error) {
	if ctx.Enabled && ctx.Config == nil {
		return http.StatusUnauthorized, errors.New("Unauthorized")
	}

	if ctx.Enabled && strings.ToLower(ctx.Config.ProviderType) == "enterprise" {
		if ctx.isHttpRequestAuthorized(r) {
			return http.StatusOK, nil
		} else {
			return http.StatusUnauthorized, errors.New("Unauthorized")
		}
	}

	return http.StatusOK, nil
}

// Process WS Request
func (ctx *SecurityContext) ProcessWSRequest(r *http.Request) (int, error) {
	if ctx.Enabled && ctx.Config == nil {
		return http.StatusUnauthorized, errors.New("Unauthorized")
	}

	if ctx.Enabled {
		if ctx.isWSRequestAuthorized(r) {
			return http.StatusOK, nil
		} else {
			return http.StatusUnauthorized, errors.New("Unauthorized")
		}
	}

	return http.StatusOK, nil
}


type LocalSecurityProvider struct {
	LocalSecurityDbFile *string
	clientsDb           *ethdb.LDBDatabase
}

func (l *LocalSecurityProvider) init() {
	if l.clientsDb == nil {
		if l.LocalSecurityDbFile == nil {
			file := os.Getenv("QuorumRpcClientDbFile")
			if  file == "" {
				utils.Fatalf("LocalSecurityDbFile not set in Security Context")
			}else{
				l.LocalSecurityDbFile = &file
			}
		}

		db, err := ethdb.NewLDBDatabase(*l.LocalSecurityDbFile, 0, 0)
		if l.clientsDb = db; err != nil {
			utils.Fatalf("Error with local security provider %v", err)
		}

	}
}

func (l *LocalSecurityProvider) findClient(clientName *string){


}

func (l *LocalSecurityProvider) addClient(clientName *string, clientID *string, clientSecret *string, clientScope *string){

}

func (l *LocalSecurityProvider) listClients(){

}

func (l *LocalSecurityProvider) removeClient(clientName *string){

}

func (l *LocalSecurityProvider) regenerateClient(clientName *string){

}



// GetDefaultDenyAllSecurityContext returns a disabled context
func GetDefaultDenyAllSecurityContext() SecurityContext {
	return SecurityContext{Enabled: true}
}

// GetDefaultAllowAllSecurityContext returns a disabled context
func GetDefaultAllowAllSecurityContext() SecurityContext {
	return SecurityContext{Enabled: false}
}

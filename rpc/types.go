// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package rpc

import (
	"fmt"
	"github.com/hashicorp/golang-lru"
	"math"
	"net/http"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/deckarep/golang-set"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

const (
	PendingBlockNumber  = BlockNumber(-2)
	LatestBlockNumber   = BlockNumber(-1)
	EarliestBlockNumber = BlockNumber(0)

	// Strings to be used in config
	LocalSecProvider      = "local"
	EnterpriseSecProvider = "enterprise"
)

// API describes the set of methods offered over the RPC interface
type API struct {
	Namespace string      // namespace under which the rpc methods of Service are exposed
	Version   string      // api version for DApp's
	Service   interface{} // receiver instance which holds the methods
	Public    bool        // indication if the methods must be considered safe for public use
}

// callback is a method callback which was registered in the server
type callback struct {
	rcvr        reflect.Value  // receiver of method
	method      reflect.Method // callback
	argTypes    []reflect.Type // input argument types
	hasCtx      bool           // method's first argument is a context (not included in argTypes)
	errPos      int            // err return idx, of -1 when method cannot return error
	isSubscribe bool           // indication if the callback is a subscription
}

// service represents a registered object
type service struct {
	name          string        // name for service
	typ           reflect.Type  // receiver type
	callbacks     callbacks     // registered handlers
	subscriptions subscriptions // available subscriptions/notifications
}

// serverRequest is an incoming request
type serverRequest struct {
	id            interface{}
	svcname       string
	callb         *callback
	args          []reflect.Value
	isUnsubscribe bool
	err           Error
}

type serviceRegistry map[string]*service // collection of services
type callbacks map[string]*callback      // collection of RPC callbacks
type subscriptions map[string]*callback  // collection of subscription callbacks

type SecurityProvider interface {
	// setup security provider
	Init() error

	// Check if client is authorized. True if authorized, false otherwise.
	IsClientAuthorized(request rpcRequest) bool

	AddClientsFromFile(path *string) ([]ClientInfo, error)

	SetClientScope(clientName string, scope string) error

	SetClientStatus(clientName string, status bool) error

	NewClient(clientName string, clientId string, secret string, scope string, active bool) (ClientInfo, error)

	AddClient(client *ClientInfo) error

	GetClientsList() []*ClientInfo

	RemoveClient(clientName *string) error

	RegenerateClientSecret(clientName *string) (*ClientInfo, error)

	GetClientByToken(clientSecret *string) *ClientInfo

	GetClientById(clientId *string) *ClientInfo

	GetClientByName(clientName *string) *ClientInfo

	AddClientsToFile(clients []*ClientInfo, path *string) error
}

// Server represents a RPC server
type Server struct {
	services serviceRegistry

	run      int32
	codecsMu sync.Mutex
	codecs   mapset.Set

	securityContext SecurityContext
}

// rpcRequest represents a raw incoming RPC request
type rpcRequest struct {
	service  string
	method   string
	id       interface{}
	isPubSub bool
	params   interface{}
	err      Error // invalid batch element
	token    string
}

// Error wraps RPC errors, which contain an error code in addition to the message.
type Error interface {
	Error() string  // returns the message
	ErrorCode() int // returns the code
}

// ServerCodec implements reading, parsing and writing RPC messages for the server side of
// a RPC session. Implementations must be go-routine safe since the codec can be called in
// multiple go-routines concurrently.
type ServerCodec interface {
	// Read next request
	ReadRequestHeaders() ([]rpcRequest, bool, Error)
	// Parse request argument to the given types
	ParseRequestArguments(argTypes []reflect.Type, params interface{}) ([]reflect.Value, Error)
	// Assemble success response, expects response id and payload
	CreateResponse(id interface{}, reply interface{}) interface{}
	// Assemble error response, expects response id and error
	CreateErrorResponse(id interface{}, err Error) interface{}
	// Assemble error response with extra information about the error through info
	CreateErrorResponseWithInfo(id interface{}, err Error, info interface{}) interface{}
	// Create notification response
	CreateNotification(id, namespace string, event interface{}) interface{}
	// Write msg to client.
	Write(msg interface{}) error
	// Close underlying data stream
	Close()
	// Closed when underlying connection is closed
	Closed() <-chan interface{}
}

// RFC (7662): https://tools.ietf.org/html/rfc7662.
// Authorization Server Introspect Request & Response.
type IntrospectRequest struct {
	Token         string `json:"token"`
	TokenTypeHint string `json:"token_type_hint"`
	ClientId      string `json:"client_id"`
}
type IntrospectResponse struct {
	Active     bool      `json:"active"`
	Scope      string    `json:"scope"`
	ClientId   string    `json:"client_id"`
	Expiration int       `json:"exp"`
	Created    time.Time `json:"created"`
}

// RPC Security Configuration
type SecurityConfig struct {
	Listener            *Listener            `json:"listenerCert"`
	ProviderType        string               `json:"providerType"`
	ProviderInformation *ProviderInformation `json:"providerInfo"`
}

// RPC Security Context
type SecurityContext struct {
	Enabled bool
	Config  *SecurityConfig

	Provider SecurityProvider
}

// Enterprise Server Based Security provider
type EnterpriseSecurityProvider struct {
	SecurityConfig      *SecurityConfig
	IntrospectURL       string
	ProviderCertificate *AuthorizationServerCert
	tokensCache         *lru.Cache
	client              http.Client
}

// Local file Based Security provider
type LocalSecurityProvider struct {
	TokensToClients map[string]ClientInfo
	ClientsToTokens map[string]ClientInfo
	clientsFile     *string
}

// Local client information
type ClientInfo struct {
	ClientId string `json:"clientId"`
	Secret   string `json:"secret"`
	Username string `json:"username"`
	Scope    string `json:"scope"`
	Active   bool   `json:"active"`
}

type ClientToken struct {
	Token string
	Scope Scope
}

type Scope struct {
	Service string
	Method  string
}

// Authorization Server Cert
type AuthorizationServerCert struct {
	ProviderTlsCertificateFile    string `json:"providerTlsCertificateFile"`
	ProviderTlsCertificateCaFile  string `json:"providerTlsCertificateCaFile"`
	ProviderTlsCertificateKeyFile string `json:"providerTlsCertificateKeyFile"`
}

// ProviderInformation
type ProviderInformation struct {
	// Authorization Server Cert Information
	EnterpriseProviderCertificateInfo *AuthorizationServerCert `json:"providerCert"`

	// Authorization Server Introspection URL.
	EnterpriseProviderIntrospectionURL string `json:"providerIntrospectionURL"`
	// Authorization Server Introspection Header Key
	EnterpriseProviderIntrospectionClientIdHeader string `json:"providerIntrospectionClientIdHeader"`
	EnterpriseProviderIntrospectionClientId       string `json:"providerClientId"`

	EnterpriseProviderIntrospectionClientSecretHeader string `json:"providerIntrospectionClientSecretHeader"`
	EnterpriseProviderIntrospectionClientSecret       string `json:"providerClientSecret"`

	// Local Provider Information
	LocalProviderFile *string `json:"localProviderFile"`

	// New Users Local Provider Scope
	LocalProviderDefaultClientScope *string `json:"localProviderDefaultClientScope"`
}

// RPC ListenerWithTls Support
type Listener struct {
	ServerTlsCertFile string `json:"serverTlsCertFile"`
	ServerTlsKeyFile  string `json:"serverTlsKeyFile"`
}

type BlockNumber int64

// UnmarshalJSON parses the given JSON fragment into a BlockNumber. It supports:
// - "latest", "earliest" or "pending" as string arguments
// - the block number
// Returned errors:
// - an invalid block number error when the given argument isn't a known strings
// - an out of range error when the given block number is either too little or too large
func (bn *BlockNumber) UnmarshalJSON(data []byte) error {
	input := strings.TrimSpace(string(data))
	if len(input) >= 2 && input[0] == '"' && input[len(input)-1] == '"' {
		input = input[1 : len(input)-1]
	}

	switch input {
	case "earliest":
		*bn = EarliestBlockNumber
		return nil
	case "latest":
		*bn = LatestBlockNumber
		return nil
	case "pending":
		*bn = PendingBlockNumber
		return nil
	}

	blckNum, err := hexutil.DecodeUint64(input)
	if err != nil {
		return err
	}
	if blckNum > math.MaxInt64 {
		return fmt.Errorf("Blocknumber too high")
	}

	*bn = BlockNumber(blckNum)
	return nil
}

func (bn BlockNumber) Int64() int64 {
	return (int64)(bn)
}

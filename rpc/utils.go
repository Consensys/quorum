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
	"bufio"
	"context"
	crand "crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode"
	"unicode/utf8"
)

var (
	subscriptionIDGenMu sync.Mutex
	subscriptionIDGen   = idGenerator()
)

// Is this an exported - upper case - name?
func isExported(name string) bool {
	rune, _ := utf8.DecodeRuneInString(name)
	return unicode.IsUpper(rune)
}

// Is this type exported or a builtin?
func isExportedOrBuiltinType(t reflect.Type) bool {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	// PkgPath will be non-empty even for an exported type,
	// so we need to check the type name as well.
	return isExported(t.Name()) || t.PkgPath() == ""
}

var contextType = reflect.TypeOf((*context.Context)(nil)).Elem()

// isContextType returns an indication if the given t is of context.Context or *context.Context type
func isContextType(t reflect.Type) bool {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t == contextType
}

var errorType = reflect.TypeOf((*error)(nil)).Elem()

// Implements this type the error interface
func isErrorType(t reflect.Type) bool {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t.Implements(errorType)
}

var subscriptionType = reflect.TypeOf((*Subscription)(nil)).Elem()

// isSubscriptionType returns an indication if the given t is of Subscription or *Subscription type
func isSubscriptionType(t reflect.Type) bool {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t == subscriptionType
}

// isPubSub tests whether the given method has as as first argument a context.Context
// and returns the pair (Subscription, error)
func isPubSub(methodType reflect.Type) bool {
	// numIn(0) is the receiver type
	if methodType.NumIn() < 2 || methodType.NumOut() != 2 {
		return false
	}

	return isContextType(methodType.In(1)) &&
		isSubscriptionType(methodType.Out(0)) &&
		isErrorType(methodType.Out(1))
}

// formatName will convert to first character to lower case
func formatName(name string) string {
	ret := []rune(name)
	if len(ret) > 0 {
		ret[0] = unicode.ToLower(ret[0])
	}
	return string(ret)
}

// suitableCallbacks iterates over the methods of the given type. It will determine if a method satisfies the criteria
// for a RPC callback or a subscription callback and adds it to the collection of callbacks or subscriptions. See server
// documentation for a summary of these criteria.
func suitableCallbacks(rcvr reflect.Value, typ reflect.Type) (callbacks, subscriptions) {
	callbacks := make(callbacks)
	subscriptions := make(subscriptions)

METHODS:
	for m := 0; m < typ.NumMethod(); m++ {
		method := typ.Method(m)
		mtype := method.Type
		mname := formatName(method.Name)
		if method.PkgPath != "" { // method must be exported
			continue
		}

		var h callback
		h.isSubscribe = isPubSub(mtype)
		h.rcvr = rcvr
		h.method = method
		h.errPos = -1

		firstArg := 1
		numIn := mtype.NumIn()
		if numIn >= 2 && mtype.In(1) == contextType {
			h.hasCtx = true
			firstArg = 2
		}

		if h.isSubscribe {
			h.argTypes = make([]reflect.Type, numIn-firstArg) // skip rcvr type
			for i := firstArg; i < numIn; i++ {
				argType := mtype.In(i)
				if isExportedOrBuiltinType(argType) {
					h.argTypes[i-firstArg] = argType
				} else {
					continue METHODS
				}
			}

			subscriptions[mname] = &h
			continue METHODS
		}

		// determine method arguments, ignore first arg since it's the receiver type
		// Arguments must be exported or builtin types
		h.argTypes = make([]reflect.Type, numIn-firstArg)
		for i := firstArg; i < numIn; i++ {
			argType := mtype.In(i)
			if !isExportedOrBuiltinType(argType) {
				continue METHODS
			}
			h.argTypes[i-firstArg] = argType
		}

		// check that all returned values are exported or builtin types
		for i := 0; i < mtype.NumOut(); i++ {
			if !isExportedOrBuiltinType(mtype.Out(i)) {
				continue METHODS
			}
		}

		// when a method returns an error it must be the last returned value
		h.errPos = -1
		for i := 0; i < mtype.NumOut(); i++ {
			if isErrorType(mtype.Out(i)) {
				h.errPos = i
				break
			}
		}

		if h.errPos >= 0 && h.errPos != mtype.NumOut()-1 {
			continue METHODS
		}

		switch mtype.NumOut() {
		case 0, 1, 2:
			if mtype.NumOut() == 2 && h.errPos == -1 { // method must one return value and 1 error
				continue METHODS
			}
			callbacks[mname] = &h
		}
	}

	return callbacks, subscriptions
}

// idGenerator helper utility that generates a (pseudo) random sequence of
// bytes that are used to generate identifiers.
func idGenerator() *rand.Rand {
	if seed, err := binary.ReadVarint(bufio.NewReader(crand.Reader)); err == nil {
		return rand.New(rand.NewSource(seed))
	}
	return rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
}

// NewID generates a identifier that can be used as an identifier in the RPC interface.
// e.g. filter and subscription identifier.
func NewID() ID {
	subscriptionIDGenMu.Lock()
	defer subscriptionIDGenMu.Unlock()

	id := make([]byte, 16)
	for i := 0; i < len(id); i += 7 {
		val := subscriptionIDGen.Int63()
		for j := 0; i+j < len(id) && j < 7; j++ {
			id[i+j] = byte(val)
			val >>= 8
		}
	}

	rpcId := hex.EncodeToString(id)
	// rpc ID's are RPC quantities, no leading zero's and 0 is 0x0
	rpcId = strings.TrimLeft(rpcId, "0")
	if rpcId == "" {
		rpcId = "0"
	}

	return ID("0x" + rpcId)
}

// ParseRpcSecurityConfigFile parses RPC Security configuration file to meet struct.
func ParseRpcSecurityConfigFile(configFilePath string) (*SecurityConfig, error) {
	configContent, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	var securityConfigResult SecurityConfig
	err = json.Unmarshal(configContent, &securityConfigResult)
	if err != nil {
		return nil, err
	}

	return &securityConfigResult, nil
}

// GetHttpListenerBasedOnSecurityContext returns http.listener with tls if required otherwise
// it will return un encrypted listener.
func GetHttpListenerBasedOnSecurityContext(endpoint string, ctx SecurityContext) (net.Listener, error) {
	var listener net.Listener
	var err error

	// Check for tls info in config
	if ctx.Config.Listener == nil {
		if listener, err = net.Listen("tcp", endpoint); err != nil {
			return nil, err
		}
		return listener, nil
	} else {
		if ctx.Config.Listener.ServerTlsKeyFile == "" || ctx.Config.Listener.ServerTlsCertFile == "" {
			return nil, fmt.Errorf("RPC Security listener-tls couldn't load tls files")

		} else {
			cer, err := tls.LoadX509KeyPair(ctx.Config.Listener.ServerTlsCertFile, ctx.Config.Listener.ServerTlsKeyFile)
			if err != nil {
				return nil, fmt.Errorf("RPC Security %v", err)

			} else {
				config := &tls.Config{
					// The Certificate information
					Certificates: []tls.Certificate{cer},

					// Ensure Key or DH parameter strength >= 4096 bits
					CurvePreferences: []tls.CurveID{
						tls.CurveP521,
						tls.CurveP384,
						tls.CurveP256,
					},
					CipherSuites: []uint16{
						tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
						tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
						tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
						tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
					},
					// Ensure we prefer our cipher suites.
					PreferServerCipherSuites: true,

					// Support only TLS1.2 & Above
					MinVersion: tls.VersionTLS12,
				}
				listener, err := tls.Listen("tcp", endpoint, config)
				if err != nil {
					return nil, err
				}
				return listener, nil

			}
		}

	}
}

// Fatalf formats a message to standard error and exits the program.
// The message is also printed to standard output if standard error
// is redirected to a different file.
func Fatalf(format string, args ...interface{}) {
	w := io.MultiWriter(os.Stdout, os.Stderr)
	if runtime.GOOS == "windows" {
		// The SameFile check below doesn't work on Windows.
		// stdout is unlikely to get redirected though, so just print there.
		w = os.Stdout
	} else {
		outf, _ := os.Stdout.Stat()
		errf, _ := os.Stderr.Stat()
		if outf != nil && errf != nil && os.SameFile(outf, errf) {
			w = os.Stderr
		}
	}
	fmt.Fprintf(w, "Fatal: "+format+"\n", args...)
	os.Exit(1)
}

//IsTokenExpired has token expired
func IsTokenExpired(createdTime time.Time, exp int) bool {
	elapsed := time.Now().Sub(createdTime)
	return elapsed.Seconds() > float64(exp)
}

//getRemoteScope issues a remote request and return introspect rsponse.
func getIntrospectResponse(request *IntrospectRequest, client *http.Client, cfg *SecurityConfig) (*IntrospectResponse, error) {
	// Create request
	params := url.Values{}
	params.Add("token", request.Token)
	params.Add("token_type_hint", request.TokenTypeHint)

	if cfg.ProviderInformation.EnterpriseProviderIntrospectionClientIdHeader != "" {
		// support env variables
		providerClientId := os.Getenv(cfg.ProviderInformation.EnterpriseProviderIntrospectionClientIdHeader)
		if providerClientId == "" {
			providerClientId = cfg.ProviderInformation.EnterpriseProviderIntrospectionClientId
		}

		if providerClientId != "" {
			params.Add(
				cfg.ProviderInformation.EnterpriseProviderIntrospectionClientIdHeader,
				providerClientId)
		}
	}

	if cfg.ProviderInformation.EnterpriseProviderIntrospectionClientSecretHeader != "" {

		// support env variables
		providerClientSec := os.Getenv(cfg.ProviderInformation.EnterpriseProviderIntrospectionClientSecretHeader)
		if providerClientSec == "" {
			providerClientSec = cfg.ProviderInformation.EnterpriseProviderIntrospectionClientSecret
		}

		if providerClientSec != "" {
			params.Add(
				cfg.ProviderInformation.EnterpriseProviderIntrospectionClientSecretHeader,
				providerClientSec)
		}
	}

	// Parse the url & build request
	serviceURL, err := url.Parse(cfg.ProviderInformation.EnterpriseProviderIntrospectionURL)
	if err != nil {
		return nil, err
	}

	encodedParams := params.Encode()
	req, err := http.NewRequest("POST", serviceURL.String(), strings.NewReader(encodedParams))

	// Set headers
	req.Header.Add("User-Agent", "Quorum")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(encodedParams)))

	if err != nil {
		return nil, err
	}

	// send request to server
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("un excpected status code")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var introspectResp IntrospectResponse
	err = json.Unmarshal(body, &introspectResp)
	if err != nil {
		return nil, err
	}

	// add time to response
	introspectResp.Created = time.Now()
	return &introspectResp, nil

}

//buildHttpClient Build HTTP ClientId. With tls support if
// required by the security context
func buildHttpClient(cfg *SecurityConfig) (*http.Client, error) {
	if cfg.ProviderInformation == nil {
		return &http.Client{}, nil
	}

	// Return non tls supporting client if cert information not provided
	if cfg.ProviderInformation.EnterpriseProviderCertificateInfo == nil {
		return &http.Client{}, nil
	}

	// Load provider certificate info provided
	certFile := cfg.ProviderInformation.EnterpriseProviderCertificateInfo.ProviderTlsCertificateFile
	keyFile := cfg.ProviderInformation.EnterpriseProviderCertificateInfo.ProviderTlsCertificateKeyFile
	caFile := cfg.ProviderInformation.EnterpriseProviderCertificateInfo.ProviderTlsCertificateCaFile

	// Load client cert
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}

	// Load CA cert
	caCert, err := ioutil.ReadFile(caFile)
	if err != nil {
		return nil, err
	}

	// Create certificate pool
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Setup TLS
	tlsConfig := &tls.Config{
		// Certificate information
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,

		// ensure verification happens
		InsecureSkipVerify: false,
	}
	tlsConfig.BuildNameToCertificate()
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	return &http.Client{Transport: transport}, nil

}

//isRequestAuthorized checks the scope against request info
func isRequestAuthorized(scope *Scope, request rpcRequest) bool {
	// Any method in service
	scopeService := strings.ToLower(scope.Service)
	scopeMethod := strings.ToLower(scope.Method)
	requestService := strings.ToLower(request.service)
	requestMethod := strings.ToLower(request.method)

	// Any method & service
	if scopeService == "*" && scopeMethod == "*" {
		return true
	}

	// Any method in service
	if scopeService == requestService && scopeMethod == "*" {
		return true
	}

	// Exact service & Method
	if scopeService == requestService && scopeMethod == requestMethod {
		return true
	}

	return false
}

// parseScopeStr returns list of scope in well formed struct
func parseScopeStr(scope string, separator string) ([]Scope, error) {
	// remove whitespace & split
	scopeList := strings.Split(scope, separator)
	var result = make([]Scope, len(scopeList))

	// iterate over scope
	for i, s := range scopeList {
		// only alpha numeric & .
		clean, err := cleanScope(s)
		if err != nil {
			return nil, err
		}

		var service string
		var function string

		// support format module.service, module, module.
		if strings.Contains(clean, ".") {
			scopeProp := strings.SplitN(clean, ".", 2)
			service = scopeProp[0]
			if scopeProp[1] == "" {
				scopeProp[1] = "*"
			}
			function = scopeProp[1]
		} else {
			service = clean
			function = "*"

		}

		result[i] = Scope{
			Service: service,
			Method:  function,
		}

	}

	return result, nil

}

//cleanScope removes all non alpha numeric except .
func cleanScope(str string) (string, error) {
	str = strings.Join(strings.Fields(str), "")
	reg, err := regexp.Compile("[^a-zA-Z0-9 .,*]+")
	if err != nil {
		return "", err
	}
	return reg.ReplaceAllString(str, ""), nil
}

// cleanString removes all non alpha numeric
func cleanString(str string) (string, error) {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return "", err
	}
	return reg.ReplaceAllString(str, ""), nil
}

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

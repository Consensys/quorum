// Copyright 2018 The go-ethereum Authors
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
	"context"
	"crypto/tls"
	"fmt"
	"net"

	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/plugin/security"
)

// StartHTTPEndpoint starts the HTTP RPC endpoint, configured with cors/vhosts/modules
// Quorum: tlsConfigSource and authManager are introduced to secure the HTTP endpoint
func StartHTTPEndpoint(endpoint string, apis []API, modules []string, cors []string, vhosts []string, timeouts HTTPTimeouts, tlsConfigSource security.TLSConfigurationSource, authManager security.AuthenticationManager) (net.Listener, *Server, bool, error) {
	// Generate the whitelist based on the allowed modules
	whitelist := make(map[string]bool)
	for _, module := range modules {
		whitelist[module] = true
	}
	// Register all the APIs exposed by the services
	handler := NewProtectedServer(authManager)
	for _, api := range apis {
		if whitelist[api.Namespace] || (len(whitelist) == 0 && api.Public) {
			if err := handler.RegisterName(api.Namespace, api.Service); err != nil {
				return nil, nil, false, err
			}
			log.Debug("HTTP registered", "namespace", api.Namespace)
		}
	}
	// All APIs registered, start the HTTP listener
	var (
		listener     net.Listener
		err          error
		isTlsEnabled bool
	)
	if isTlsEnabled, listener, err = startListener(endpoint, tlsConfigSource); err != nil {
		return nil, nil, isTlsEnabled, err
	}
	go NewHTTPServer(cors, vhosts, timeouts, handler).Serve(listener)
	return listener, handler, isTlsEnabled, err
}

// Quorum
// Produce net.Listener instance with TLS support if tlsConfigSource provides the config
func startListener(endpoint string, tlsConfigSource security.TLSConfigurationSource) (bool, net.Listener, error) {
	var tlsConfig *tls.Config
	var err error
	var listener net.Listener
	isTlsEnabled := true
	if tlsConfigSource != nil {
		if tlsConfig, err = tlsConfigSource.Get(context.Background()); err != nil {
			isTlsEnabled = false
		}
	} else {
		isTlsEnabled = false
		err = fmt.Errorf("no TLSConfigurationSource found")
	}
	if isTlsEnabled {
		if listener, err = tls.Listen("tcp", endpoint, tlsConfig); err != nil {
			return isTlsEnabled, nil, err
		}
	} else {
		log.Info("Security: TLS not enabled", "endpoint", endpoint, "reason", err)
		if listener, err = net.Listen("tcp", endpoint); err != nil {
			return isTlsEnabled, nil, err
		}
	}
	return isTlsEnabled, listener, nil
}

// StartWSEndpoint starts a websocket endpoint
// Quorum: tlsConfigSource and authManager are introduced to secure the WS endpoint
func StartWSEndpoint(endpoint string, apis []API, modules []string, wsOrigins []string, exposeAll bool, tlsConfigSource security.TLSConfigurationSource, authManager security.AuthenticationManager) (net.Listener, *Server, bool, error) {

	// Generate the whitelist based on the allowed modules
	whitelist := make(map[string]bool)
	for _, module := range modules {
		whitelist[module] = true
	}
	// Register all the APIs exposed by the services
	handler := NewProtectedServer(authManager)
	for _, api := range apis {
		if exposeAll || whitelist[api.Namespace] || (len(whitelist) == 0 && api.Public) {
			if err := handler.RegisterName(api.Namespace, api.Service); err != nil {
				return nil, nil, false, err
			}
			log.Debug("WebSocket registered", "service", api.Service, "namespace", api.Namespace)
		}
	}
	// All APIs registered, start the HTTP listener
	var (
		listener     net.Listener
		err          error
		isTlsEnabled bool
	)
	if isTlsEnabled, listener, err = startListener(endpoint, tlsConfigSource); err != nil {
		return nil, nil, isTlsEnabled, err
	}
	go NewWSServer(wsOrigins, handler).Serve(listener)
	return listener, handler, isTlsEnabled, err

}

// StartIPCEndpoint starts an IPC endpoint.
func StartIPCEndpoint(ipcEndpoint string, apis []API) (net.Listener, *Server, error) {
	// Register all the APIs exposed by the services.
	handler := NewServer()
	for _, api := range apis {
		if err := handler.RegisterName(api.Namespace, api.Service); err != nil {
			return nil, nil, err
		}
		log.Debug("IPC registered", "namespace", api.Namespace)
	}
	// All APIs registered, start the IPC listener.
	listener, err := ipcListen(ipcEndpoint)
	if err != nil {
		return nil, nil, err
	}
	go handler.ServeListener(listener)
	return listener, handler, nil
}

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
	"crypto/tls"
	"fmt"
	"net"

	"github.com/ethereum/go-ethereum/log"
)

// StartHTTPEndpoint starts the HTTP RPC endpoint, configured with cors/vhosts/modules
func StartHTTPEndpointWithSecurityContext(endpoint string, apis []API, modules []string, cors []string, vhosts []string, timeouts HTTPTimeouts, ctx SecurityContext) (net.Listener, *Server, error) {
	// Generate the whitelist based on the allowed modules
	whitelist := make(map[string]bool)
	for _, module := range modules {
		whitelist[module] = true
	}
	// Register all the APIs exposed by the services
	handler := NewServerWithSecurityCtx(ctx)
	for _, api := range apis {
		if whitelist[api.Namespace] || (len(whitelist) == 0 && api.Public) {
			if err := handler.RegisterName(api.Namespace, api.Service); err != nil {
				return nil, nil, err
			}
			log.Debug("HTTP registered", "namespace", api.Namespace)
		}
	}

	if ctx.Config.Listener == nil {
		// All APIs registered, start the HTTP listener
		var (
			listener net.Listener
			err      error
		)


		if listener, err = net.Listen("tcp", endpoint); err != nil {
			return nil, nil, err
		}
		go NewHTTPServer(cors, vhosts, timeouts, handler).Serve(listener)
		return listener, handler, err
	} else {
		log.Info("RPC Security","listener-tls-cert",  ctx.Config.Listener.ServerTlsCertFile, "listener-tls-key",ctx.Config.Listener.ServerTlsKeyFile)
		if ctx.Config.Listener.ServerTlsKeyFile == "" ||  ctx.Config.Listener.ServerTlsCertFile == "" {
			return nil, nil, fmt.Errorf("RPC Security listener-tls couldn't load tls files")

		}else{
			cer, err := tls.LoadX509KeyPair(ctx.Config.Listener.ServerTlsCertFile, ctx.Config.Listener.ServerTlsKeyFile)
			if err != nil {
				return nil, nil, fmt.Errorf("RPC Security %v", err)

			}else{
				config := &tls.Config{Certificates: []tls.Certificate{cer}, MinVersion:tls.VersionTLS12}
				listener, err :=tls.Listen("tcp", endpoint, config)
				if err != nil {
					return nil, nil, err
				}
				go NewHTTPServer(cors, vhosts, timeouts, handler).Serve(listener)
				return listener, handler, err

			}
		}

	}


}

// StartHTTPEndpoint starts the HTTP RPC endpoint, configured with cors/vhosts/modules
func StartHTTPEndpoint(endpoint string, apis []API, modules []string, cors []string, vhosts []string, timeouts HTTPTimeouts) (net.Listener, *Server, error) {
	// Generate the whitelist based on the allowed modules
	whitelist := make(map[string]bool)
	for _, module := range modules {
		whitelist[module] = true
	}
	// Register all the APIs exposed by the services
	handler := NewServer()
	for _, api := range apis {
		if whitelist[api.Namespace] || (len(whitelist) == 0 && api.Public) {
			if err := handler.RegisterName(api.Namespace, api.Service); err != nil {
				return nil, nil, err
			}
			log.Debug("HTTP registered", "namespace", api.Namespace)
		}
	}


	// All APIs registered, start the HTTP listener
	var (
		listener net.Listener
		err      error
	)
	if listener, err = net.Listen("tcp", endpoint); err != nil {
		return nil, nil, err
	}
	go NewHTTPServer(cors, vhosts, timeouts, handler).Serve(listener)
	return listener, handler, err
}

// StartWSEndpoint starts a websocket endpoint
func StartWSEndpointWithSecurityContext(endpoint string, apis []API, modules []string, wsOrigins []string, exposeAll bool, ctx SecurityContext) (net.Listener, *Server, error) {

	// Generate the whitelist based on the allowed modules
	whitelist := make(map[string]bool)
	for _, module := range modules {
		whitelist[module] = true
	}
	// Register all the APIs exposed by the services
	handler := NewServerWithSecurityCtx(ctx)
	for _, api := range apis {
		if exposeAll || whitelist[api.Namespace] || (len(whitelist) == 0 && api.Public) {
			if err := handler.RegisterName(api.Namespace, api.Service); err != nil {
				return nil, nil, err
			}
			log.Debug("WebSocket registered", "service", api.Service, "namespace", api.Namespace)
		}
	}
	if ctx.Config.Listener == nil {
		// All APIs registered, start the HTTP listener
		var (
			listener net.Listener
			err      error
		)
		if listener, err = net.Listen("tcp", endpoint); err != nil {
			return nil, nil, err
		}

		return listener, handler, err
	} else {
		log.Info("RPC Security","ws-listener-tls-cert",  ctx.Config.Listener.ServerTlsCertFile, "ws-listener-tls-key",ctx.Config.Listener.ServerTlsKeyFile)
		if ctx.Config.Listener.ServerTlsKeyFile == "" ||  ctx.Config.Listener.ServerTlsCertFile == "" {
			return nil, nil, fmt.Errorf("RPC Security ws-listener-tls couldn't load tls files")

		}else{
			cer, err := tls.LoadX509KeyPair(ctx.Config.Listener.ServerTlsCertFile, ctx.Config.Listener.ServerTlsKeyFile)
			if err != nil {
				return nil, nil, fmt.Errorf("RPC Security %v", err)

			}else{
				config := &tls.Config{Certificates: []tls.Certificate{cer}, MinVersion:tls.VersionTLS12}
				listener, err :=tls.Listen("tcp", endpoint, config)
				if err != nil {
					return nil, nil, err
				}
				go NewWSServer(wsOrigins, handler).Serve(listener)
				return listener, handler, err

			}
		}
	}

}

// StartWSEndpoint starts a websocket endpoint
func StartWSEndpoint(endpoint string, apis []API, modules []string, wsOrigins []string, exposeAll bool) (net.Listener, *Server, error) {

	// Generate the whitelist based on the allowed modules
	whitelist := make(map[string]bool)
	for _, module := range modules {
		whitelist[module] = true
	}
	// Register all the APIs exposed by the services
	handler := NewServer()
	for _, api := range apis {
		if exposeAll || whitelist[api.Namespace] || (len(whitelist) == 0 && api.Public) {
			if err := handler.RegisterName(api.Namespace, api.Service); err != nil {
				return nil, nil, err
			}
			log.Debug("WebSocket registered", "service", api.Service, "namespace", api.Namespace)
		}
	}
	// All APIs registered, start the HTTP listener
	var (
		listener net.Listener
		err      error
	)
	if listener, err = net.Listen("tcp", endpoint); err != nil {
		return nil, nil, err
	}
	go NewWSServer(wsOrigins, handler).Serve(listener)
	return listener, handler, err

}

// StartIPCEndpoint starts an IPC endpoint.
func StartIPCEndpointWithSecurityContext(ipcEndpoint string, apis []API, ctx SecurityContext) (net.Listener, *Server, error) {
	// Register all the APIs exposed by the services.
	handler := NewServerWithSecurityCtx(ctx)
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

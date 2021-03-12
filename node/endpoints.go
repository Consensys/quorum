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

package node

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/plugin/security"
	"github.com/ethereum/go-ethereum/rpc"
)

// StartHTTPEndpoint starts the HTTP RPC endpoint.
func StartHTTPEndpoint(endpoint string, timeouts rpc.HTTPTimeouts, handler http.Handler, tlsConfigSource security.TLSConfigurationSource) (*http.Server, net.Addr, bool, error) {
	// start the HTTP listener
	var (
		listener     net.Listener
		err          error
		isTlsEnabled bool
	)
	if isTlsEnabled, listener, err = startListener(endpoint, tlsConfigSource); err != nil {
		return nil, nil, isTlsEnabled, err
	}
	// make sure timeout values are meaningful
	CheckTimeouts(&timeouts)
	// Bundle and start the HTTP server
	httpSrv := &http.Server{
		Handler:      handler,
		ReadTimeout:  timeouts.ReadTimeout,
		WriteTimeout: timeouts.WriteTimeout,
		IdleTimeout:  timeouts.IdleTimeout,

		// Ensure to Disable HTTP/2
		// this configuration and customized tls.Config is to follow: https://blog.bracebin.com/achieving-perfect-ssl-labs-score-with-go
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
	}
	go httpSrv.Serve(listener)
	return httpSrv, listener.Addr(), isTlsEnabled, err
}

// checkModuleAvailability checks that all names given in modules are actually
// available API services. It assumes that the MetadataApi module ("rpc") is always available;
// the registration of this "rpc" module happens in NewServer() and is thus common to all endpoints.
func checkModuleAvailability(modules []string, apis []rpc.API) (bad, available []string) {
	availableSet := make(map[string]struct{})
	for _, api := range apis {
		if _, ok := availableSet[api.Namespace]; !ok {
			availableSet[api.Namespace] = struct{}{}
			available = append(available, api.Namespace)
		}
	}
	for _, name := range modules {
		if _, ok := availableSet[name]; !ok && name != rpc.MetadataApi {
			bad = append(bad, name)
		}
	}
	return bad, available
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

// CheckTimeouts ensures that timeout values are meaningful
func CheckTimeouts(timeouts *rpc.HTTPTimeouts) {
	if timeouts.ReadTimeout < time.Second {
		log.Warn("Sanitizing invalid HTTP read timeout", "provided", timeouts.ReadTimeout, "updated", rpc.DefaultHTTPTimeouts.ReadTimeout)
		timeouts.ReadTimeout = rpc.DefaultHTTPTimeouts.ReadTimeout
	}
	if timeouts.WriteTimeout < time.Second {
		log.Warn("Sanitizing invalid HTTP write timeout", "provided", timeouts.WriteTimeout, "updated", rpc.DefaultHTTPTimeouts.WriteTimeout)
		timeouts.WriteTimeout = rpc.DefaultHTTPTimeouts.WriteTimeout
	}
	if timeouts.IdleTimeout < time.Second {
		log.Warn("Sanitizing invalid HTTP idle timeout", "provided", timeouts.IdleTimeout, "updated", rpc.DefaultHTTPTimeouts.IdleTimeout)
		timeouts.IdleTimeout = rpc.DefaultHTTPTimeouts.IdleTimeout
	}
}

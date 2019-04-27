// Copyright 2016 The go-ethereum Authors
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
	"errors"
	"os"
	"time"
)

const (
	DefaultTLSCertFile = "tlscert.pem"
	DefaultTLSKeyFile  = "tlskey.pem"
)

var (
	validFor        = 365 * 24 * time.Hour
	rsaBits         = 2048
	ErrCertNotFound = errors.New("cannot find the cert file")
	ErrKeyNotFound  = errors.New("cannot find the cert file")
)

func MakeServerTLSConfig() *tls.Config {
	config := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			// http2 require TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256
			// see: https://github.com/golang/net/blob/master/http2/server.go#L222
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}

	return config
}

// CheckCerts check certificate/key files
func CheckCerts(certPath string, keyPath string) error {
	if _, err := os.Stat(certPath); os.IsNotExist(err) {
		return ErrCertNotFound
	}
	if _, err := os.Stat(keyPath); os.IsNotExist(err) {
		return ErrKeyNotFound
	}
	return nil
}

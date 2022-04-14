package http

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/tv42/httpunix"
)

func unixTransport(cfg Config) *httpunix.Transport {
	// Note that clientTimeout doesn't work when using httpunix.Transport, so we set ResponseHeaderTimeout instead
	t := &httpunix.Transport{
		DialTimeout:           time.Duration(cfg.DialTimeout) * time.Second,
		RequestTimeout:        5 * time.Second,
		ResponseHeaderTimeout: time.Duration(cfg.Timeout) * time.Second,
	}
	t.RegisterLocation("c", filepath.Join(cfg.WorkDir, cfg.Socket))
	return t
}

func httpTransport(cfg Config) *http.Transport {
	t := &http.Transport{
		IdleConnTimeout: time.Duration(cfg.HttpIdleConnTimeout) * time.Second,
		WriteBufferSize: cfg.HttpWriteBufferSize,
		ReadBufferSize:  cfg.HttpReadBufferSize,
	}
	return t
}

func newTLSConfig(cfg Config) (*tls.Config, error) {
	rootCAPool, err := loadRootCaCerts(cfg.TlsRootCA)
	if err != nil {
		return nil, err
	}

	var getClientCertFunc func(*tls.CertificateRequestInfo) (*tls.Certificate, error) = nil
	if len(cfg.TlsClientCert) != 0 && len(cfg.TlsClientKey) != 0 {
		getClientCertFunc = func(info *tls.CertificateRequestInfo) (certificate *tls.Certificate, e error) {
			c, err := tls.LoadX509KeyPair(cfg.TlsClientCert, cfg.TlsClientKey)
			if err != nil {
				return nil, fmt.Errorf("failed to load client key pair from '%v', '%v': %v", cfg.TlsClientCert, cfg.TlsClientKey, err)
			}
			return &c, nil
		}
	}

	return &tls.Config{
		RootCAs:              rootCAPool,
		InsecureSkipVerify:   cfg.TlsInsecureSkipVerify,
		GetClientCertificate: getClientCertFunc,
	}, nil
}

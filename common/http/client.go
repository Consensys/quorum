package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/private/engine"
)

func CreateClient(cfg Config) (*engine.Client, error) {
	var client *engine.Client
	if IsSocketConfigured(cfg) {
		log.Info("Connecting to private tx manager using IPC socket")
		client = &engine.Client{
			HttpClient: &http.Client{
				Transport: unixTransport(cfg),
			},
			BaseURL: "http+unix://c",
		}
	} else {
		transport := httpTransport(cfg)
		if cfg.TlsMode == TlsOff {
			log.Info("Connecting to private tx manager using HTTP")
		} else {
			log.Info("Connecting to private tx manager using HTTPS")
			tlsConfig, err := newTLSConfig(cfg)
			if err != nil {
				return nil, fmt.Errorf("unable to create http.client to private tx manager due to: %s", err)
			}
			transport.TLSClientConfig = tlsConfig
		}

		client = &engine.Client{
			HttpClient: &http.Client{
				Timeout:   time.Duration(cfg.Timeout) * time.Second,
				Transport: transport,
			},
			BaseURL: cfg.HttpUrl,
		}
	}

	return client, nil
}

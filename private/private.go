package private

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/private/engine"
	"github.com/ethereum/go-ethereum/private/engine/constellation"
	"github.com/ethereum/go-ethereum/private/engine/notinuse"
	"github.com/ethereum/go-ethereum/private/engine/tessera"
	"github.com/tv42/httpunix"
)

var (
	// global variable to be accessed by other packages
	// singleton gateway to interact with private transaction manager
	P                PrivateTransactionManager
	isPrivacyEnabled = false
)

type Identifiable interface {
	Name() string
	HasFeature(f engine.PrivateTransactionManagerFeature) bool
}

// Interacting with Private Transaction Manager APIs
type PrivateTransactionManager interface {
	Identifiable

	Send(data []byte, from string, to []string, extra *engine.ExtraMetadata) (common.EncryptedPayloadHash, error)
	StoreRaw(data []byte, from string) (common.EncryptedPayloadHash, error)
	SendSignedTx(data common.EncryptedPayloadHash, to []string, extra *engine.ExtraMetadata) ([]byte, error)
	// Returns nil payload if not found
	Receive(data common.EncryptedPayloadHash) ([]byte, *engine.ExtraMetadata, error)
	// Returns nil payload if not found
	ReceiveRaw(data common.EncryptedPayloadHash) ([]byte, *engine.ExtraMetadata, error)
	IsSender(txHash common.EncryptedPayloadHash) (bool, error)
	GetParticipants(txHash common.EncryptedPayloadHash) ([]string, error)
	EncryptPayload(data []byte, from string, to []string, extra *engine.ExtraMetadata) ([]byte, error)
	DecryptPayload(payload common.DecryptRequest) ([]byte, *engine.ExtraMetadata, error)
}

// This loads any config specified via the legacy environment variable
func GetLegacyEnvironmentConfig() (engine.Config, error) {
	return FromEnvironmentOrNil("PRIVATE_CONFIG")
}

func FromEnvironmentOrNil(name string) (engine.Config, error) {
	cfgPath := os.Getenv(name)
	cfg, err := engine.FetchConfigOrIgnore(cfgPath)
	if err != nil {
		return engine.Config{}, err
	}

	return cfg, nil
}

func InitialiseConnection(cfg engine.Config) error {
	var err error
	P, err = NewPrivateTxManager(cfg)
	return err
}

func IsQuorumPrivacyEnabled() bool {
	return isPrivacyEnabled
}

func NewPrivateTxManager(cfg engine.Config) (PrivateTransactionManager, error) {

	if cfg.ConnectionType == engine.NoConnection {
		log.Info("Running with private transaction manager disabled - quorum private transactions will not be supported")
		return &notinuse.PrivateTransactionManager{}, nil
	}

	var client *engine.Client
	if engine.IsSocketConfigured(cfg) {
		log.Info("Connecting to private tx manager using IPC socket")
		client = createIPCClient(cfg)
	} else if cfg.TlsMode != engine.TlsOff {
		log.Info("Connecting to private tx manager using HTTPS")
		var err error
		client, err = createHTTPClientUsingTLS(cfg)
		if err != nil {
			return nil, fmt.Errorf("unable to create http.client to private tx manager due to: %s", err)
		}
	} else {
		log.Info("Connecting to private tx manager using HTTP")
		client = createHTTPClient(cfg)
	}

	ptm, err := selectPrivateTxManager(client)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to private tx manager due to: %s", err)
	}

	isPrivacyEnabled = true
	return ptm, nil
}

func createIPCClient(cfg engine.Config) *engine.Client {
	client := &engine.Client{
		HttpClient: &http.Client{
			Transport: unixTransport(cfg),
		},
		BaseURL: "http+unix://c",
	}
	return client
}

func unixTransport(cfg engine.Config) *httpunix.Transport {
	// Note that clientTimeout doesn't work when using httpunix.Transport, so we set ResponseHeaderTimeout instead
	t := &httpunix.Transport{
		DialTimeout:           time.Duration(cfg.DialTimeout) * time.Second,
		RequestTimeout:        5 * time.Second,
		ResponseHeaderTimeout: time.Duration(cfg.Timeout) * time.Second,
	}
	t.RegisterLocation("c", filepath.Join(cfg.WorkDir, cfg.Socket))
	return t
}

func createHTTPClient(cfg engine.Config) *engine.Client {
	client := &engine.Client{
		HttpClient: &http.Client{
			Timeout:   time.Duration(cfg.Timeout) * time.Second,
			Transport: httpTransport(cfg),
		},
		BaseURL: cfg.HttpUrl,
	}
	return client
}

func httpTransport(cfg engine.Config) *http.Transport {
	t := &http.Transport{
		IdleConnTimeout: time.Duration(cfg.HttpIdleConnTimeout) * time.Second,
		WriteBufferSize: cfg.HttpWriteBufferSize,
		ReadBufferSize:  cfg.HttpReadBufferSize,
	}
	return t
}

func createHTTPClientUsingTLS(cfg engine.Config) (*engine.Client, error) {
	transport, err := httpTransportUsingTLS(cfg)
	if err != nil {
		return nil, err
	}

	client := &engine.Client{
		HttpClient: &http.Client{
			Timeout:   time.Duration(cfg.Timeout) * time.Second,
			Transport: transport,
		},
		BaseURL: cfg.HttpUrl,
	}
	return client, nil
}

func httpTransportUsingTLS(cfg engine.Config) (*http.Transport, error) {
	rootCAPool := x509.NewCertPool()
	rootCA, err := ioutil.ReadFile(cfg.TlsRootCA)
	if err != nil {
		return nil, fmt.Errorf("reading TlsRootCA certificate from '%v' failed : %v", cfg.TlsRootCA, err)
	}
	if !rootCAPool.AppendCertsFromPEM(rootCA) {
		return nil, fmt.Errorf("failed to add TlsRootCA certificate to pool, check that '%v' contains a valid cert", cfg.TlsRootCA)
	}

	t := &http.Transport{
		IdleConnTimeout: time.Duration(cfg.HttpIdleConnTimeout) * time.Second,
		WriteBufferSize: cfg.HttpWriteBufferSize,
		ReadBufferSize:  cfg.HttpReadBufferSize,
		TLSClientConfig: &tls.Config{
			RootCAs: rootCAPool,
			// Load clients key-pair. This will be sent to server
			GetClientCertificate: func(info *tls.CertificateRequestInfo) (certificate *tls.Certificate, e error) {
				c, err := tls.LoadX509KeyPair(cfg.TlsClientCert, cfg.TlsClientKey)
				if err != nil {
					return nil, fmt.Errorf("failed to load client key pair from '%v', '%v': %v", cfg.TlsClientCert, cfg.TlsClientKey, err)
				}
				return &c, nil
			},
		},
	}

	return t, nil
}

// First call /upcheck to make sure the private tx manager is up
// Then call /version to decide which private tx manager client implementation to be used
func selectPrivateTxManager(client *engine.Client) (PrivateTransactionManager, error) {
	res, err := client.Get("/upcheck")
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, engine.ErrPrivateTxManagerNotReady
	}
	res, err = client.Get("/version")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	version, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var privateTxManager PrivateTransactionManager
	defer func() {
		log.Info("Target Private Tx Manager", "name", privateTxManager.Name(), "distributionVersion", string(version))
	}()
	if res.StatusCode != 200 {
		// Constellation doesn't have /version endpoint
		privateTxManager = constellation.New(client)
	} else {
		privateTxManager = tessera.New(client, []byte(tessera.RetrieveTesseraAPIVersion(client)))
	}
	return privateTxManager, nil
}

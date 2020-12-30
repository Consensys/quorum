package private

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
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
	P PrivateTransactionManager
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

func Init() {
	P = FromEnvironmentOrNil("PRIVATE_CONFIG")
}

func FromEnvironmentOrNil(name string) PrivateTransactionManager {
	cfgPath := os.Getenv(name)
	if cfgPath == "" {
		return nil
	}
	if strings.EqualFold(cfgPath, "ignore") {
		return &notinuse.PrivateTransactionManager{}
	}
	return MustNewPrivateTxManager(cfgPath)
}

func MustNewPrivateTxManager(cfgPath string) PrivateTransactionManager {
	ptm, err := NewPrivateTxManager(cfgPath)
	if err != nil {
		panic(err)
	}
	return ptm
}

func NewPrivateTxManager(cfgPath string) (PrivateTransactionManager, error) {
	cfg, err := engine.FetchConfig(cfgPath)
	if err != nil {
		return nil, fmt.Errorf("error using '%s' due to: %s", cfgPath, err)
	}

	var client *engine.Client
	if engine.IsSocketConfigured(cfg) {
		log.Info("Connecting to private tx manager using IPC socket")
		client = createIPCClient(cfg)
	} else if engine.IsTlsConfigured(cfg) {
		log.Info("Connecting to private tx manager using HTTPS")
		client, err = createHTTPClientUsingTLS(cfg)
		if err != nil {
			return nil, fmt.Errorf("unable to create http.client to private tx manager using '%s' due to: %s", cfgPath, err)
		}
	} else {
		log.Info("Connecting to private tx manager using HTTP")
		client = createHTTPClient(cfg)
	}

	ptm, err := selectPrivateTxManager(client)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to private tx manager using '%s' due to: %s", cfgPath, err)
	}
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
	t := &httpunix.Transport{
		DialTimeout:           time.Duration(cfg.SocketConfig.DialTimeout) * time.Second,
		RequestTimeout:        time.Duration(cfg.SocketConfig.RequestTimeout) * time.Second,
		ResponseHeaderTimeout: time.Duration(cfg.SocketConfig.ResponseHeaderTimeout) * time.Second,
	}
	t.RegisterLocation("c", filepath.Join(cfg.SocketConfig.WorkDir, cfg.SocketConfig.Socket))
	return t
}

func createHTTPClient(cfg engine.Config) *engine.Client {
	client := &engine.Client{
		HttpClient: &http.Client{
			Timeout:   time.Duration(cfg.HttpConfig.ClientTimeout) * time.Second,
			Transport: httpTransport(cfg),
		},
		BaseURL: cfg.HttpConfig.Url,
	}
	return client
}

func httpTransport(cfg engine.Config) *http.Transport {
	t := &http.Transport{
		IdleConnTimeout: time.Duration(cfg.HttpConfig.IdleConnTimeout) * time.Second,
		WriteBufferSize: cfg.HttpConfig.WriteBufferSize,
		ReadBufferSize:  cfg.HttpConfig.ReadBufferSize,
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
			Timeout:   time.Duration(cfg.HttpConfig.ClientTimeout) * time.Second,
			Transport: transport,
		},
		BaseURL: cfg.HttpConfig.Url,
	}
	return client, nil
}

func httpTransportUsingTLS(cfg engine.Config) (*http.Transport, error) {
	rootCAPool := x509.NewCertPool()
	rootCA, err := ioutil.ReadFile(cfg.HttpConfig.RootCA)
	if err != nil {
		return nil, fmt.Errorf("reading RootCA certificate from '%v' failed : %v", cfg.HttpConfig.RootCA, err)
	}
	if !rootCAPool.AppendCertsFromPEM(rootCA) {
		return nil, fmt.Errorf("failed to add RootCA certificate to pool, check that '%v' contains a valid cert", cfg.HttpConfig.RootCA)
	}

	t := &http.Transport{
		IdleConnTimeout: time.Duration(cfg.HttpConfig.IdleConnTimeout) * time.Second,
		WriteBufferSize: cfg.HttpConfig.WriteBufferSize,
		ReadBufferSize:  cfg.HttpConfig.ReadBufferSize,
		TLSClientConfig: &tls.Config{
			RootCAs: rootCAPool,
			// Load clients key-pair. This will be sent to server
			GetClientCertificate: func(info *tls.CertificateRequestInfo) (certificate *tls.Certificate, e error) {
				c, err := tls.LoadX509KeyPair(cfg.HttpConfig.ClientCert, cfg.HttpConfig.ClientKey)
				if err != nil {
					return nil, fmt.Errorf("failed to load client key pair from '%v', '%v': %v", cfg.HttpConfig.ClientCert, cfg.HttpConfig.ClientKey, err)
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

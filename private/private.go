package private

import (
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
	P = FromEnvironmentOrNil("PRIVATE_CONFIG")
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
	cfg, err := engine.LoadConfig(cfgPath)
	if err != nil {
		return nil, fmt.Errorf("unable to read %s due to %s", cfgPath, err)
	}

	client := &engine.Client{
		HttpClient: &http.Client{
			Transport: unixTransport(cfg),
		},
		BaseURL: "http+unix://c",
	}

	ptm, err := selectPrivateTxManager(client)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to private tx manager using %s due to %s", cfgPath, err)
	}
	return ptm, nil
}

func unixTransport(cfg engine.Config) *httpunix.Transport {
	t := &httpunix.Transport{
		DialTimeout:           time.Duration(cfg.DialTimeout) * time.Second,
		RequestTimeout:        time.Duration(cfg.RequestTimeout) * time.Second,
		ResponseHeaderTimeout: time.Duration(cfg.ResponseHeaderTimeout) * time.Second,
	}
	t.RegisterLocation("c", filepath.Join(cfg.WorkDir, cfg.Socket))
	return t
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
		log.Info("Target Private Tx Manager", "name", privateTxManager.Name(), "distributionVersion", version)
	}()
	if res.StatusCode != 200 {
		// Constellation doesn't have /version endpoint
		privateTxManager = constellation.New(client)
	} else {
		privateTxManager = tessera.New(client, []byte(tessera.RetrieveTesseraAPIVersion(client)))
	}
	return privateTxManager, nil
}

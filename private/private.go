package private

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/private/engine/constellation"

	"github.com/ethereum/go-ethereum/private/engine"

	"github.com/ethereum/go-ethereum/log"

	"github.com/tv42/httpunix"

	"github.com/ethereum/go-ethereum/common"
)

var (
	// global variable to be accessed by other packages
	// singleton gateway to interact with private transaction manager
	P = FromEnvironmentOrNil("PRIVATE_CONFIG")
)

type PrivateTransactionManager interface {
	Name() string

	Send(data []byte, from string, to []string, acHashes common.EncryptedPayloadHashes, acMerkleRoot common.Hash) (common.EncryptedPayloadHash, error)
	SendSignedTx(data common.EncryptedPayloadHash, to []string, acHashes common.EncryptedPayloadHashes, acMerkleRoot common.Hash) ([]byte, error)
	Receive(data common.EncryptedPayloadHash) ([]byte, common.EncryptedPayloadHashes, common.Hash, error)
}

func FromEnvironmentOrNil(name string) PrivateTransactionManager {
	cfgPath := os.Getenv(name)
	if cfgPath == "" {
		return nil
	}
	if strings.EqualFold(cfgPath, "ignore") {
		return &engine.NotInUsePrivateTxManager{}
	}
	return mustNewPrivateTxManager(cfgPath)
}

func mustNewPrivateTxManager(cfgPath string) PrivateTransactionManager {
	info, err := os.Lstat(cfgPath)
	if err != nil {
		panic(fmt.Sprintf("unable to read %s due to %s", cfgPath, err))
	}
	// We accept either the socket or a configuration file that points to
	// a socket.
	socketPath := cfgPath
	isSocket := info.Mode()&os.ModeSocket != 0
	if !isSocket {
		cfg, err := engine.LoadConfig(cfgPath)
		if err != nil {
			panic(fmt.Sprintf("unable to load configuration file for private transaction manager from %s due to %s", cfgPath, err))
		}
		socketPath = filepath.Join(cfg.WorkDir, cfg.Socket)
	}

	client := &http.Client{
		Transport: unixTransport(socketPath),
	}
	ptm, err := selectPrivateTxManager(client)
	if err != nil {
		panic(fmt.Sprintf("unable to connect to private tx manager using %s due to %s", socketPath, err))
	}
	return ptm
}

func unixTransport(socketPath string) *httpunix.Transport {
	t := &httpunix.Transport{
		DialTimeout:           1 * time.Second,
		RequestTimeout:        5 * time.Second,
		ResponseHeaderTimeout: 5 * time.Second,
	}
	t.RegisterLocation("c", socketPath)
	return t
}

// First call /upcheck to make sure the private tx manager is up
// Then call /version to decide which private tx manager client implementation to be used
func selectPrivateTxManager(client *http.Client) (PrivateTransactionManager, error) {
	res, err := client.Get("http+unix://c/upcheck")
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, engine.ErrPrivateTxManagerNotReady
	}
	res, err = client.Get("http+unix://c/version")
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
		log.Info("Target Private Tx Manager", "name", privateTxManager.Name(), "version", version)
	}()
	privateTxManager = constellation.New(client) // temporarily until Tessera client is fully implemented
	/*
		if res.StatusCode != 200 {
			// Constellation doesn't have /version endpoint
			privateTxManager = constellation.New(client)
		}
		privateTxManager = tessera.New(client)
	*/
	return privateTxManager, nil
}

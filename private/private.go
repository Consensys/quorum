package private

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ethereum/go-ethereum/common"
	http2 "github.com/ethereum/go-ethereum/common/http"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/private/engine"
	"github.com/ethereum/go-ethereum/private/engine/constellation"
	"github.com/ethereum/go-ethereum/private/engine/notinuse"
	"github.com/ethereum/go-ethereum/private/engine/qlightptm"
	"github.com/ethereum/go-ethereum/private/engine/tessera"
	"github.com/ethereum/go-ethereum/rpc"
)

var (
	// global variable to be accessed by other packages
	// singleton gateway to interact with private transaction manager
	P                PrivateTransactionManager
	isPrivacyEnabled = false
)

type HasRPCClient interface {
	SetRPCClient(client *rpc.Client)
}

type Identifiable interface {
	Name() string
	HasFeature(f engine.PrivateTransactionManagerFeature) bool
}

// Interacting with Private Transaction Manager APIs
type PrivateTransactionManager interface {
	Identifiable

	Send(data []byte, from string, to []string, extra *engine.ExtraMetadata) (string, []string, common.EncryptedPayloadHash, error)
	StoreRaw(data []byte, from string) (common.EncryptedPayloadHash, error)
	SendSignedTx(data common.EncryptedPayloadHash, to []string, extra *engine.ExtraMetadata) (string, []string, []byte, error)
	// Returns nil payload if not found
	Receive(data common.EncryptedPayloadHash) (string, []string, []byte, *engine.ExtraMetadata, error)
	// Returns nil payload if not found
	ReceiveRaw(data common.EncryptedPayloadHash) ([]byte, string, *engine.ExtraMetadata, error)
	IsSender(txHash common.EncryptedPayloadHash) (bool, error)
	GetParticipants(txHash common.EncryptedPayloadHash) ([]string, error)
	GetMandatory(txHash common.EncryptedPayloadHash) ([]string, error)
	EncryptPayload(data []byte, from string, to []string, extra *engine.ExtraMetadata) ([]byte, error)
	DecryptPayload(payload common.DecryptRequest) ([]byte, *engine.ExtraMetadata, error)

	Groups() ([]engine.PrivacyGroup, error)
}

// This loads any config specified via the legacy environment variable
func GetLegacyEnvironmentConfig() (http2.Config, error) {
	return FromEnvironmentOrNil("PRIVATE_CONFIG")
}

func FromEnvironmentOrNil(name string) (http2.Config, error) {
	cfgPath := os.Getenv(name)
	cfg, err := http2.FetchConfigOrIgnore(cfgPath)
	if err != nil {
		return http2.Config{}, err
	}

	return cfg, nil
}

func InitialiseConnection(cfg http2.Config, isLightClient bool) error {
	var err error
	if isLightClient {
		P, err = NewQLightTxManager()
		return err
	}
	P, err = NewPrivateTxManager(cfg)
	return err
}

func IsQuorumPrivacyEnabled() bool {
	return isPrivacyEnabled
}

func NewQLightTxManager() (PrivateTransactionManager, error) {
	isPrivacyEnabled = true
	return qlightptm.New(), nil
}

func NewPrivateTxManager(cfg http2.Config) (PrivateTransactionManager, error) {
	if cfg.ConnectionType == http2.NoConnection {
		log.Info("Running with private transaction manager disabled - quorum private transactions will not be supported")
		return &notinuse.PrivateTransactionManager{}, nil
	}

	client, err := http2.CreateClient(cfg)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection to private tx manager due to: %s", err)
	}

	ptm, err := selectPrivateTxManager(client)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to private tx manager due to: %s", err)
	}

	isPrivacyEnabled = true
	return ptm, nil
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

// Retrieve the private transaction that is associated with a privacy marker transaction
func FetchPrivateTransaction(data []byte) (*types.Transaction, []string, *engine.ExtraMetadata, error) {
	return FetchPrivateTransactionWithPTM(data, P)
}

func FetchPrivateTransactionWithPTM(data []byte, ptm PrivateTransactionManager) (*types.Transaction, []string, *engine.ExtraMetadata, error) {
	txHash := common.BytesToEncryptedPayloadHash(data)

	_, managedParties, txData, metadata, err := ptm.Receive(txHash)
	if err != nil {
		return nil, nil, nil, err
	}
	if txData == nil {
		return nil, nil, nil, nil
	}

	var tx types.Transaction
	err = json.NewDecoder(bytes.NewReader(txData)).Decode(&tx)
	if err != nil {
		log.Trace("failed to deserialize private transaction", "err", err)
		return nil, nil, nil, err
	}

	return &tx, managedParties, metadata, nil
}

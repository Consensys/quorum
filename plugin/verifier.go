package plugin

import (
	"fmt"
	"path"

	"github.com/ethereum/go-ethereum/log"
)

// Plugin Integrity Verifier.
// Verifier works on the assumption an attacker can not compromise the integrity of geth running process.
type Verifier interface {
	// verify plugin signature using checksum & pgp public key
	VerifySignature(definition *PluginDefinition, checksum string) error
}

type NonVerifier struct {
}

func (*NonVerifier) VerifySignature(definition *PluginDefinition, checksum string) error {
	return nil
}

func NewNonVerifier() *NonVerifier {
	return &NonVerifier{}
}

func NewVerifier(pm *PluginManager, localVerify bool, publicKey string) (Verifier, error) {
	log.Debug("using verifier", "local", localVerify)
	pluginBaseDir := pm.pluginBaseDir
	centralClient := pm.centralClient
	// resolve public key
	if publicKey == "" {
		publicKey = fmt.Sprintf("file://%s", path.Join(pluginBaseDir, DefaultPublicKeyFile))
	}
	publicKeyPath, err := resolveFilePath(publicKey)
	if err != nil {
		return nil, err
	}
	if localVerify {
		return NewLocalVerifier(publicKeyPath, pluginBaseDir)
	} else {
		return NewOnlineVerifier(centralClient), nil
	}
}

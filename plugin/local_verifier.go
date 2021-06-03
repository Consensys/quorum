package plugin

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

// For Cloudsmith, this references to the latest GPG key
// being setup in the repo
const DefaultPublicKeyFile = "gpg.key"

// Local Implementation of plugin.Verifier
type LocalVerifier struct {
	PublicKeyPath    string // where to obtain PGP public key
	SignatureBaseDir string // where to obtain plugin signature file
}

// Build a new LocalVerifier
func NewLocalVerifier(publicKeyPath string, pluginSignatureBaseDir string) (*LocalVerifier, error) {
	if _, err := os.Stat(publicKeyPath); os.IsNotExist(err) {
		return nil, err
	}
	stat, err := os.Stat(pluginSignatureBaseDir)
	if os.IsNotExist(err) {
		return nil, err
	}
	if !stat.Mode().IsDir() {
		return nil, fmt.Errorf("pluginSignatureBaseDir is not a directory")
	}
	verifier := &LocalVerifier{
		PublicKeyPath:    publicKeyPath,
		SignatureBaseDir: pluginSignatureBaseDir,
	}
	return verifier, nil
}

// Verify a plugin giving its name from Central
func (v *LocalVerifier) VerifySignature(definition *PluginDefinition, checksum string) error {
	pluginSigPath := path.Join(v.SignatureBaseDir, definition.SignatureFileName())
	if _, err := os.Stat(pluginSigPath); os.IsNotExist(err) {
		return err
	}
	pubkey, err := ioutil.ReadFile(v.PublicKeyPath)
	if err != nil {
		return err
	}
	sig, err := ioutil.ReadFile(pluginSigPath)
	if err != nil {
		return err
	}
	return verify(sig, pubkey, checksum)
}

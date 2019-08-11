package vault

import (
	"errors"
	"fmt"
	"strings"
)

type HashicorpWalletConfig struct {
	Client  HashicorpClientConfig
	Secrets []HashicorpSecretConfig
}

type HashicorpClientConfig struct {
	Url              string `toml:",omitempty"`
	Approle          string `toml:",omitempty"`
	CaCert           string `toml:",omitempty"`
	ClientCert       string `toml:",omitempty"`
	ClientKey        string `toml:",omitempty"`
	StorePrivateKeys bool   `toml:",omitempty"`
	VaultPollingIntervalMillis int `toml:",omitempty"`
}

type HashicorpSecretConfig struct {
	AddressSecret           string `toml:",omitempty"`
	PrivateKeySecret        string `toml:",omitempty"`
	AddressSecretVersion    int    `toml:",omitempty"`
	PrivateKeySecretVersion int    `toml:",omitempty"`
	SecretEngine            string `toml:",omitempty"`
}

// Validate checks that the HashicorpWalletConfig has the minimum fields defined to be a valid configuration.  If the configuration is invalid an error is returned describing which fields have not been defined otherwise nil is returned.
//
// This should be used to validate configs intended to be used for retrieving from a Vault (i.e. in normal node operation).  For configs intended to be used for writing to a Vault use ValidateSkipVersion.
func (w HashicorpWalletConfig) Validate() error {
	return w.validate(false)
}

// ValidateSkipVersion checks that the HashicorpWalletConfig has the minimum fields defined to be a valid configuration, ignoring the version fields.  If the configuration is invalid an error is returned describing which fields have not been defined otherwise nil is returned.
//
// This should be used over Validate when validating configs intended to be used to write to the vault (i.e. in new account creation) as it is not necessary to specify the version number in these cases.
//
// It is not recommended to use ValidateSkipVersion when validating configs intended to retrieve from a Vault as this will allow secrets to be configured with version=0 (i.e. always retrieve the latest version of a secret).  This is to protect against secrets being updated and a node then being unable to access the original accounts it was configured with because the wallet is now only capable of retrieving the latest version of the secret.
func (w HashicorpWalletConfig) ValidateSkipVersion() error {
	return w.validate(true)
}

func (w HashicorpWalletConfig) validate(skipVersion bool) error {
	var errs []string

	if w.Client.Url == "" {
		errs = append(errs, fmt.Sprint("Invalid vault client config: Vault url must be provided"))
	}

	for _, s := range w.Secrets {

		if s.AddressSecret == "" {
			errs = append(errs, fmt.Sprintf("Invalid vault secret config, vault=%v: AddressSecret must be provided", w.Client.Url))
		}

		if s.PrivateKeySecret == "" {
			errs = append(errs, fmt.Sprintf("Invalid vault secret config, vault=%v: PrivateKeySecret must be provided", w.Client.Url))
		}

		if s.AddressSecretVersion <= 0 && !skipVersion {
			errs = append(errs, fmt.Sprintf("Invalid vault secret config, vault=%v, secret=%v: AddressSecretVersion must be specified for vault secret and must be greater than zero", w.Client.Url, s.AddressSecret))
		}

		if s.PrivateKeySecretVersion <= 0 && !skipVersion {
			errs = append(errs, fmt.Sprintf("Invalid vault secret config, vault=%v, secret=%v: AddressSecretVersion must be specified for vault secret and must be greater than zero", w.Client.Url, s.PrivateKeySecret))
		}

		if s.SecretEngine == "" {
			errs = append(errs, fmt.Sprintf("Invalid vault secret config, vault=%v, AddressSecret=%v: SecretEngine must be provided", w.Client.Url, s.AddressSecret))
		}
	}

	if len(errs) > 0 {
		return errors.New("\n" + strings.Join(errs, "\n"))
	}

	return nil
}

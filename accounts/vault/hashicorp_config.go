package vault

type hashicorpWalletConfig struct {
	Client hashicorpClientConfig
	Secrets []hashicorpSecretData
}

type hashicorpClientConfig struct {
	Url            string `toml:",omitempty"`
	Approle        string `toml:",omitempty"`
	CaCert         string `toml:",omitempty"`
	ClientCert     string `toml:",omitempty"`
	ClientKey      string `toml:",omitempty"`
	UseSecretCache   bool   `toml:",omitempty"`
}

type hashicorpSecretData struct {
	Name         string `toml:",omitempty"`
	SecretEngine string `toml:",omitempty"`
	Version      int    `toml:",omitempty"`
	AccountID    string `toml:",omitempty"`
	KeyID        string `toml:",omitempty"`
}


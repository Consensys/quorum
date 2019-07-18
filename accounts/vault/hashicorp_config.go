package vault

type hashicorpWalletConfig struct {
	Client  hashicorpClientConfig
	Secrets []hashicorpSecretConfig
}

type hashicorpClientConfig struct {
	Url              string `toml:",omitempty"`
	Approle          string `toml:",omitempty"`
	CaCert           string `toml:",omitempty"`
	ClientCert       string `toml:",omitempty"`
	ClientKey        string `toml:",omitempty"`
	StorePrivateKeys bool   `toml:",omitempty"`
	VaultPollingIntervalSecs int `toml:",omitempty"`
}

type hashicorpSecretConfig struct {
	AddressSecret           string `toml:",omitempty"`
	PrivateKeySecret        string `toml:",omitempty"`
	AddressSecretVersion    int    `toml:",omitempty"`
	PrivateKeySecretVersion int    `toml:",omitempty"`
	SecretEngine            string `toml:",omitempty"`
}

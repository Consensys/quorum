package vault

import "testing"

func getMinimumValidConfig() HashicorpWalletConfig {
	return HashicorpWalletConfig{
		Client: HashicorpClientConfig{
			Url: "someurl",
		},
		Secrets: []HashicorpSecretConfig{
			{
				AddressSecret:         "addr",
				PrivateKeySecret: "key",
				AddressSecretVersion: 1,
				PrivateKeySecretVersion: 1,
				SecretEngine: "kv",
			},
			{
				AddressSecret:         "otherAddr",
				PrivateKeySecret: "otherKey",
				AddressSecretVersion: 1,
				PrivateKeySecretVersion: 1,
				SecretEngine: "kv",
			},
		},
	}
}

func TestHashicorpWalletConfig_Validate_ValidReturnsNil(t *testing.T) {
	w := getMinimumValidConfig()

	err := w.Validate()

	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
}

func TestHashicorpWalletConfig_Validate_NoVaultUrlReturnsError(t *testing.T) {
	w := getMinimumValidConfig()
	w.Client.Url = ""

	err := w.Validate()

	if err == nil {
		t.Error("Wanted error")
	}
}

func TestHashicorpWalletConfig_Validate_NoAddressSecretNameReturnsError(t *testing.T) {
	w := getMinimumValidConfig()
	w.Secrets[0].AddressSecret = ""

	err := w.Validate()

	if err == nil {
		t.Error("Wanted error")
	}
}

func TestHashicorpWalletConfig_Validate_NoPrivateKeySecretNameReturnsError(t *testing.T) {
	w := getMinimumValidConfig()
	w.Secrets[0].PrivateKeySecret = ""

	err := w.Validate()

	if err == nil {
		t.Error("Wanted error")
	}
}

func TestHashicorpWalletConfig_Validate_NoSecretEngineReturnsError(t *testing.T) {
	w := getMinimumValidConfig()
	w.Secrets[0].SecretEngine = ""

	err := w.Validate()

	if err == nil {
		t.Error("Wanted error")
	}
}

func TestHashicorpWalletConfig_Validate_ZeroAddressSecretVersionReturnsError(t *testing.T) {
	w := getMinimumValidConfig()
	w.Secrets[0].AddressSecretVersion = 0

	err := w.Validate()

	if err == nil {
		t.Error("Wanted error")
	}
}

func TestHashicorpWalletConfig_Validate_ZeroPrivateKeySecretVersionReturnsError(t *testing.T) {
	w := getMinimumValidConfig()
	w.Secrets[0].PrivateKeySecretVersion = 0

	err := w.Validate()

	if err == nil {
		t.Error("Wanted error")
	}
}

func TestHashicorpWalletConfig_Validate_NegativeAddressSecretVersionReturnsError(t *testing.T) {
	w := getMinimumValidConfig()
	w.Secrets[0].AddressSecretVersion = -1

	err := w.Validate()

	if err == nil {
		t.Error("Wanted error")
	}
}

func TestHashicorpWalletConfig_Validate_NegativePrivateKeySecretVersionReturnsError(t *testing.T) {
	w := getMinimumValidConfig()
	w.Secrets[0].PrivateKeySecretVersion = -1

	err := w.Validate()

	if err == nil {
		t.Error("Wanted error")
	}
}

func TestHashicorpWalletConfig_Validate_MultipleErrorsAreCombined(t *testing.T) {
	w := getMinimumValidConfig()
	w.Secrets[0].AddressSecret = ""
	w.Secrets[1].AddressSecret = ""

	err := w.Validate()

	if err == nil {
		t.Error("Wanted error")
	}

	want := "\nInvalid vault secret config, vault=someurl: AddressSecret must be provided\nInvalid vault secret config, vault=someurl: AddressSecret must be provided"

	got := err.Error()

	if got != want {
		t.Errorf("Incorrect error\nwant: %v\ngot : %v", want, got)
	}
}

func TestHashicorpWalletConfig_ValidateSkipVersion_DoesNotRequireVersionsToBeGreaterThanZero(t *testing.T) {
	w := getMinimumValidConfig()
	w.Secrets[0].AddressSecretVersion = 0
	w.Secrets[0].PrivateKeySecretVersion = -1

	err := w.ValidateSkipVersion()

	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
}

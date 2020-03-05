package plugin

// Implementation of plugin.Verifier that uses remote server to verify plugins.
type OnlineVerifier struct {
	centralClient *CentralClient
}

func NewOnlineVerifier(centralClient *CentralClient) *OnlineVerifier {
	return &OnlineVerifier{centralClient: centralClient}
}

// Verify a plugin giving its name from Central
func (v *OnlineVerifier) VerifySignature(definition *PluginDefinition, checksum string) error {
	sig, err := v.centralClient.PluginSignature(definition)
	if err != nil {
		return err
	}
	pubkey, err := v.centralClient.PublicKey()
	if err != nil {
		return err
	}
	return verify(sig, pubkey, checksum)
}

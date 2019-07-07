package vault

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/hashicorp/vault/api"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestVaultWallet_URL(t *testing.T) {
	in := accounts.URL{Scheme: "http", Path: "url"}
	w := vaultWallet{url: in}

	got := w.URL()

	if in.Cmp(got) != 0 {
		t.Fatalf("want: %v, got: %v", in, got)
	}
}

// makeMockHashicorpService creates a new httptest.Server which responds with mockResponse for all requests.  A default Hashicorp api.Client with URL updated with the httptest.Server's URL is returned.  The Close() function for the httptest.Server and should be executed before test completion (probably best to defer as soon as it is returned)
func makeMockHashicorpClient(t *testing.T, mockResponse []byte) (*api.Client, func()) {
	vaultServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(mockResponse)
	}))

	//create default client and update URL to use mock vault server
	config := api.DefaultConfig()
	config.Address = vaultServer.URL
	client, err := api.NewClient(config)

	if err != nil {
		t.Fatalf("err creating client: %v", err)
	}

	return client, vaultServer.Close
}

func TestVaultWallet_Status_Hashicorp_ClosedWhenServiceHasNoClient(t *testing.T) {
	w := vaultWallet{vault: &hashicorpService{}}

	status, err := w.Status()

	if err != nil {
		t.Fatal(err)
	}

	if status != closed {
		t.Fatalf("want: %v, got: %v", closed, status)
	}
}

func TestVaultWallet_Status_Hashicorp_HealthcheckSuccessful(t *testing.T) {
	const (
		uninitialised = "uninitialized"
		sealed = "sealed"
		open = "open"
	)

	makeMockHashicorpResponse := func(t *testing.T, vaultStatus string) []byte {
		var vaultResponse api.HealthResponse

		switch vaultStatus {
		case uninitialised:
			vaultResponse.Initialized = false
		case sealed:
			vaultResponse.Initialized = true
			vaultResponse.Sealed = true
		case open:
			vaultResponse.Initialized = true
			vaultResponse.Sealed = false
		}

		b, err := json.Marshal(vaultResponse)

		if err != nil {
			t.Fatalf("err marshalling mock response: %v", err)
		}

		return b
	}

	tests := []struct{
		vaultStatus string
		want string
		wantErr error
	}{
		{vaultStatus: uninitialised, want: hashicorpUninitialized, wantErr: hashicorpUninitializedErr},
		{vaultStatus: sealed, want: hashicorpSealed, wantErr: hashicorpSealedErr},
		{vaultStatus: open, want: open, wantErr: nil},
	}

	for _, tt := range tests {
		t.Run(tt.vaultStatus, func(t *testing.T) {
			b := makeMockHashicorpResponse(t, tt.vaultStatus)
			c, cleanup := makeMockHashicorpClient(t, b)
			defer cleanup()

			w := vaultWallet{
				vault: &hashicorpService{client: c},
			}

			status, err := w.Status()

			if tt.wantErr != err {
				t.Fatalf("want: %v, got: %v", tt.wantErr, err)
			}

			if tt.want != status {
				t.Fatalf("want: %v, got: %v", tt.want, status)
			}
		})
	}
}

func TestVaultWallet_Status_Hashicorp_HealthcheckFailed(t *testing.T) {
	b := []byte("this is not the bytes for an api.HealthResponse and will cause a client error")

	c, cleanup := makeMockHashicorpClient(t, b)
	defer cleanup()

	w := vaultWallet{
		vault: &hashicorpService{client: c},
	}

	status, err := w.Status()

	if _, ok := err.(hashicorpHealthcheckErr); !ok {
		t.Fatal("returned error should be of type hashicorpHealthcheckErr")
	}

	if status != hashicorpHealthcheckFailed {
		t.Fatalf("want: %v, got: %v", hashicorpHealthcheckFailed, status)
	}
}

func TestVaultWallet_Open_Hashicorp_ReturnsErrIfAlreadyOpen(t *testing.T) {
	w := vaultWallet{vault: &hashicorpService{client: &api.Client{}}}

	if err := w.Open(""); err != accounts.ErrWalletAlreadyOpen {
		t.Fatalf("want: %v, got: %v", accounts.ErrWalletAlreadyOpen, err)
	}
}

func TestVaultWallet_Open_Hashicorp_CreatesClientFromConfig(t *testing.T) {
	// create mock server which responds to all requests with the same response
	mockResponse := api.Secret{RequestID: "myrequestid"}
	b, err := json.Marshal(mockResponse)

	if err != nil {
		t.Fatal(err)
	}

	vaultServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(b)
	}))
	defer vaultServer.Close()

	config := hashicorpClientConfig{
		Url: vaultServer.URL,
	}

	w := vaultWallet{vault: &hashicorpService{config: config}}

	if err := w.Open(""); err != nil {
		t.Fatalf("error: %v", err)
	}

	v, ok := w.vault.(*hashicorpService)

	if !ok {
		t.Fatal("type assertion failed")
	}

	got := v.client

	if got == nil {
		t.Fatal("client not created")
	}

	if got.Address() != vaultServer.URL {
		t.Fatalf("address: want: %v, got: %v", vaultServer.URL, got.Address())
	}

	// make a request to the vault server using the client to verify client config was correctly applied
	resp, err := got.Logical().Read("vaultpath/to/secret")

	if err != nil {
		t.Fatalf("error making request using created client: %v", err)
	}

	if !reflect.DeepEqual(mockResponse, *resp) {
		t.Fatalf("response not as expected\nwant: %v\ngot : %v", mockResponse, resp)
	}
}

func TestVaultWallet_Open_Hashicorp_CreatesTLSClientFromConfig(t *testing.T) {
	// create mock server which responds to all requests with an empty secret
	mockResponse := api.Secret{}
	b, err := json.Marshal(mockResponse)

	if err != nil {
		t.Fatal(err)
	}

	vaultServer := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(b)
	}))

	// read TLS certs
	rootCert, err := ioutil.ReadFile("testdata/caRoot.pem")

	if err != nil {
		t.Error(err)
	}

	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(rootCert)

	cert, err := ioutil.ReadFile("testdata/localhost-with-san-chain.pem")

	if err != nil {
		t.Error(err)
	}

	key, err := ioutil.ReadFile("testdata/localhost-with-san.key")

	if err != nil {
		t.Error(err)
	}

	keypair, err := tls.X509KeyPair(cert, key)

	if err != nil {
		t.Error(err)
	}

	serverTlsConfig := &tls.Config{
		Certificates: []tls.Certificate{keypair},
		ClientAuth: tls.RequireAnyClientCert,
		ClientCAs: certPool,
	}

	// add TLS config to server and start
	vaultServer.TLS = serverTlsConfig

	vaultServer.StartTLS()
	defer vaultServer.Close()

	// create wallet with config and open
	config := hashicorpClientConfig{
		Url: vaultServer.URL,
		//Approle: "myapprole",
		CaCert: "testdata/caRoot.pem",
		ClientCert: "testdata/quorum-client-chain.pem",
		ClientKey: "testdata/quorum-client.key",
		//EnvVarPrefix: "prefix",
		//UseSecretCache: false,
	}

	w := vaultWallet{vault: &hashicorpService{config: config}}

	if err := w.Open(""); err != nil {
		t.Fatalf("error: %v", err)
	}

	// verify created client uses config
	v, ok := w.vault.(*hashicorpService)

	if !ok {
		t.Fatal("type assertion failed")
	}

	got := v.client

	if got == nil {
		t.Fatal("client not created")
	}

	if got.Address() != vaultServer.URL {
		t.Fatalf("address: want: %v, got: %v", vaultServer.URL, got.Address())
	}

	// make a request to the vault server using the client - if TLS was applied correctly on the client then the request will be allowed
	// TODO test passes even if vault client config contains no TLS certs
	if _, err := got.Logical().Read("vaultpath/to/secret"); err != nil {
		t.Fatalf("error making request using created client: %v", err)
	}
}


package vault

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/hashicorp/vault/api"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
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

func TestVaultWallet_Open_Hashicorp_CreatesClientUsingConfig(t *testing.T) {
	if err := os.Setenv(api.EnvVaultToken, "mytoken"); err != nil {
		t.Fatal(err)
	}

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

	w := vaultWallet{vault: &hashicorpService{config: config}, updateFeed: &event.Feed{}}

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

func TestVaultWallet_Open_Hashicorp_CreatesTLSClientUsingConfig(t *testing.T) {
	if err := os.Setenv(api.EnvVaultToken, "mytoken"); err != nil {
		t.Fatal(err)
	}

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
		ClientAuth: tls.RequireAndVerifyClientCert,
		ClientCAs: certPool,
	}

	// add TLS config to server and start
	vaultServer.TLS = serverTlsConfig

	vaultServer.StartTLS()
	defer vaultServer.Close()

	// create wallet with config and open
	config := hashicorpClientConfig{
		Url: vaultServer.URL,
		CaCert: "testdata/caRoot.pem",
		ClientCert: "testdata/quorum-client-chain.pem",
		ClientKey: "testdata/quorum-client.key",
	}

	w := vaultWallet{vault: &hashicorpService{config: config}, updateFeed: &event.Feed{}}

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
	if _, err := got.Logical().Read("vaultpath/to/secret"); err != nil {
		t.Fatalf("error making request using created client: %v", err)
	}
}

func TestVaultWallet_Open_Hashicorp_ClientAuthenticatesUsingEnvVars(t *testing.T) {
	const (
		myToken = "myToken"
		myRoleId = "myRoleId"
		mySecretId = "mySecretId"
		myApproleToken = "myApproleToken"
	)

	setAndHandleErrors := func(t *testing.T, env, val string) {
		if err := os.Setenv(env, val); err != nil {
			t.Fatal(err)
		}
	}

	set := func(t *testing.T, env string) {
		switch env {
		case api.EnvVaultToken:
			setAndHandleErrors(t, api.EnvVaultToken, myToken)
		case roleIDEnv:
			setAndHandleErrors(t, roleIDEnv, myRoleId)
		case secretIDEnv:
			setAndHandleErrors(t, secretIDEnv, mySecretId)
		}
	}

	// makeMockApproleVaultServer creates an httptest.Server for handling approle auth requests.  The server and its Close function are returned.  Close must be called to ensure the server is stopped (best to defer the function as soon as it is returned).
	//
	// The server will expose only the path /v1/auth/{approlePath}/login.  If approlePath = "" then the default value of "approle" will be used.  The server will respond with an api.Secret containing the provided token.
	makeMockApproleVaultServer := func (t *testing.T, approlePath string) (*httptest.Server, func()) {

		vaultResponse := &api.Secret{Auth: &api.SecretAuth{ClientToken: myApproleToken}}
		b, err := json.Marshal(vaultResponse)

		if err != nil {
			t.Fatal(err)
		}

		if approlePath == "" {
			approlePath = "approle"
		}

		mux := http.NewServeMux()
		mux.HandleFunc(fmt.Sprintf("/v1/auth/%v/login", approlePath), func(w http.ResponseWriter, r *http.Request) {
			w.Write(b)
		})

		vaultServer := httptest.NewServer(mux)

		return vaultServer, vaultServer.Close
	}

	tests := map[string]struct{
		envVars []string
		approle string
		wantToken string
	}{
		"token auth": {envVars: []string{api.EnvVaultToken}, wantToken: myToken},
		"default approle auth": {envVars: []string{roleIDEnv, secretIDEnv}, wantToken: myApproleToken},
		"custom approle auth": {envVars: []string{roleIDEnv, secretIDEnv}, approle: "nondefault", wantToken: myApproleToken},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			//initialize environment
			os.Clearenv()
			for _, e := range tt.envVars {
				set(t, e)
				defer os.Unsetenv(e)
			}

			vaultServer, cleanup := makeMockApproleVaultServer(t, tt.approle)
			defer cleanup()

			config := hashicorpClientConfig{
				Url: vaultServer.URL,
				Approle: tt.approle,
			}

			w := vaultWallet{vault: &hashicorpService{config: config}, updateFeed: &event.Feed{}}

			if err := w.Open(""); err != nil {
				t.Fatalf("error: %v", err)
			}

			// verify the client is set up as expected
			v, ok := w.vault.(*hashicorpService)

			if !ok {
				t.Fatal("type assertion failed")
			}

			got := v.client

			if got == nil {
				t.Fatal("client not created")
			}

			if tt.wantToken != got.Token() {
				t.Fatalf("incorrect client token: want: %v, got: %v", tt.wantToken, got.Token())
			}
		})
	}
}

func TestVaultWallet_Open_Hashicorp_ErrAuthenticatingClient(t *testing.T) {
	const (
		myToken = "myToken"
		myRoleId = "myRoleId"
		mySecretId = "mySecretId"
	)

	setAndHandleErrors := func(t *testing.T, env, val string) {
		if err := os.Setenv(env, val); err != nil {
			t.Fatal(err)
		}
	}

	set := func(t *testing.T, env string) {
		switch env {
		case api.EnvVaultToken:
			setAndHandleErrors(t, api.EnvVaultToken, myToken)
		case roleIDEnv:
			setAndHandleErrors(t, roleIDEnv, myRoleId)
		case secretIDEnv:
			setAndHandleErrors(t, secretIDEnv, mySecretId)
		}
	}

	tests := map[string]struct{
		envVars []string
		want error
	}{
		"no auth provided": {envVars: []string{}, want: noHashicorpEnvSetErr},
		"only role id": {envVars: []string{roleIDEnv}, want: invalidApproleAuthErr},
		"only secret id": {envVars: []string{secretIDEnv}, want: invalidApproleAuthErr},
		"role id and token": {envVars: []string{api.EnvVaultToken, roleIDEnv}, want: invalidApproleAuthErr},
		"secret id and token": {envVars: []string{api.EnvVaultToken, secretIDEnv}, want: invalidApproleAuthErr},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			//initialize environment
			os.Clearenv()
			for _, e := range tt.envVars {
				set(t, e)
				defer os.Unsetenv(e)
			}

			config := hashicorpClientConfig{
				Url: "http://url:1",
			}

			w := vaultWallet{vault: &hashicorpService{config: config}}

			if err := w.Open(""); err != tt.want {
				t.Fatalf("want error: %v\ngot: %v", tt.want, err)
			}
		})
	}
}

// Note: This is an integration test, as such the scope of the test is large, covering the VaultBackend, VaultWallet and HashicorpService
func TestVaultWallet_Open_Hashicorp_SendsEventToBackendSubscribers(t *testing.T) {
	if err := os.Setenv(api.EnvVaultToken, "mytoken"); err != nil {
		t.Fatal(err)
	}

	walletConfig := hashicorpWalletConfig{
		Client: hashicorpClientConfig{
			Url: "http://url:1",
		},
	}

	b := NewHashicorpBackend([]hashicorpWalletConfig{walletConfig})

	if len(b.wallets) != 1 {
		t.Fatalf("NewHashicorpBackend: incorrect number of wallets created: want 1, got: %v", len(b.wallets))
	}

	subscriber := make(chan accounts.WalletEvent, 1)
	b.Subscribe(subscriber)

	if b.updateScope.Count() != 1 {
		t.Fatalf("incorrect number of subscribers for backend: want: %v, got: %v", 1, b.updateScope.Count())
	}

	if err := b.wallets[0].Open(""); err != nil {
		t.Fatalf("error: %v", err)
	}

	if len(subscriber) != 1 {
		t.Fatal("event not added to subscriber")
	}

	got := <-subscriber

	want := accounts.WalletEvent{Wallet: b.wallets[0], Kind: accounts.WalletOpened}

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want: %v, got: %v", want, got)
	}
}

func TestVaultWallet_Close_Hashicorp_ReturnsStateToBeforeOpen(t *testing.T) {
	if err := os.Setenv(api.EnvVaultToken, "mytoken"); err != nil {
		t.Fatal(err)
	}

	config := hashicorpWalletConfig{
		Client: hashicorpClientConfig{Url: "http://url:1"},
		Secrets: []hashicorpSecretData{{Name: "secret1"}},
	}

	w, err := newHashicorpWallet(config, &event.Feed{})

	if err != nil {
		t.Fatal(err)
	}

	unopened, err := newHashicorpWallet(config, &event.Feed{})

	if err != nil {
		t.Fatal(err)
	}

	cmpOpts := []cmp.Option{cmp.AllowUnexported(vaultWallet{}, hashicorpService{}), cmpopts.IgnoreUnexported(event.Feed{})}

	if diff := cmp.Diff(unopened, w, cmpOpts...); diff != "" {
		t.Fatalf("cmp does not consider the two wallets equal\n%v", diff)
	}

	if err := w.Open(""); err != nil {
		t.Fatalf("error: %v", err)
	}

	if diff := cmp.Diff(unopened, w, cmpOpts...); diff == "" {
		t.Fatalf("cmp does not consider the wallets different after one was opened\n%v", diff)
	}

	if err := w.Close(); err != nil {
		t.Fatalf("error: %v", err)
	}

	if diff := cmp.Diff(unopened, w, cmpOpts...); diff != "" {
		t.Fatalf("cmp does not consider the two wallets equal after one was opened and closed:\n%v", diff)
	}
}

func TestVaultWallet_Accounts_ReturnsCopyOfAccountsInWallet(t *testing.T) {
	w := vaultWallet{
		vault: &hashicorpService{accts: []accounts.Account{{URL: accounts.URL{Scheme: "http", Path: "url:1"}}}},
	}

	got := w.Accounts()

	v := w.vault.(*hashicorpService)

	if !cmp.Equal(v.accts, got) {
		t.Fatalf("want: %v, got: %v", v.accts, got)
	}

	got[0].URL = accounts.URL{Scheme: "http", Path: "changed:1"}

	if cmp.Equal(v.accts, got) {
		t.Fatalf("changes to the returned accounts should not change the wallet's record of accounts")
	}
}

func TestVaultWallet_Contains(t *testing.T) {
	makeAcct := func(addr, url string) accounts.Account {
		var u accounts.URL

		if url != "" {
			//to parse a string url as an accounts.URL it must first be in json format
			toParse := fmt.Sprintf("\"%v\"", url)

			if err := u.UnmarshalJSON([]byte(toParse)); err != nil {
				t.Fatal(err)
			}
		}

		return accounts.Account{Address: common.StringToAddress(addr), URL: u}
	}

	tests := map[string]struct{
		accts []accounts.Account
		toFind accounts.Account
		want bool
	}{
		"same addr and url": {accts: []accounts.Account{makeAcct("addr1", "http://url:1")}, toFind: makeAcct("addr1", "http://url:1"), want: true},
		"same addr no url": {accts: []accounts.Account{makeAcct("addr1", "http://url:1")}, toFind: makeAcct("addr1", ""), want: true},
		"multiple": {accts: []accounts.Account{makeAcct("addr1", "http://url:1"), makeAcct("addr2", "http://url:2")}, toFind: makeAcct("addr2", "http://url:2"), want: true},
		"same addr diff url": {accts: []accounts.Account{makeAcct("addr1", "http://url:1")}, toFind: makeAcct("addr1", "http://url:2"), want: false},
		"diff addr same url": {accts: []accounts.Account{makeAcct("addr1", "http://url:1")}, toFind: makeAcct("addr2", "http://url:1"), want: false},
		"diff addr no url": {accts: []accounts.Account{makeAcct("addr1", "http://url:1")}, toFind: makeAcct("addr2", ""), want: false},
		"diff addr diff url": {accts: []accounts.Account{makeAcct("addr1", "http://url:1")}, toFind: makeAcct("addr2", "http://url:2"), want: false},
		"no accts": {toFind: makeAcct("addr1", "http://url:1"), want: false},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			w := vaultWallet{
				vault: &hashicorpService{accts: tt.accts},
			}

			got := w.Contains(tt.toFind)

			if tt.want != got {
				t.Fatalf("want: %v, got: %v", tt.want, got)
			}
		})
	}
}


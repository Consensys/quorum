package vault

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/event"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/hashicorp/vault/api"
	"io/ioutil"
	"math/big"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"sort"
	"sync"
	"testing"
	"time"
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

// Note: This is an integration test, as such the scope of the test is large.  It covers the VaultBackend, vaultWallet and hashicorpService
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

type accountsByUrl []accounts.Account

func (a accountsByUrl) Len() int {
	return len(a)
}

func (a accountsByUrl) Less(i, j int) bool {
	return (a[i].URL).Cmp(a[j].URL) < 0
}

func (a accountsByUrl) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func acctsEqual(a, b []accounts.Account) bool {
	if len(a) != len(b) {
		return false
	}

	sort.Sort(accountsByUrl(a))
	sort.Sort(accountsByUrl(b))

	equal := func(a, b accounts.Account) bool {
		return a.Address == b.Address && (a.URL == b.URL || a.URL == accounts.URL{} || b.URL == accounts.URL{})
	}

	for i := 0; i < len(a); i++ {
		if !equal(a[i], b[i]) {
			return false
		}
	}

	return true
}

func TestVaultWallet_Open_Hashicorp_AccountsRetrieved(t *testing.T) {
	if err := os.Setenv(api.EnvVaultToken, "mytoken"); err != nil {
		t.Fatal(err)
	}

	makeVaultResponse := func(keyValPairs map[string]string) []byte {
		resp := api.Secret{
			Data: map[string]interface{}{
				"data": keyValPairs,
			},
		}

		b, err := json.Marshal(resp)

		if err != nil {
			t.Fatal(err)
		}

		return b
	}

	makeSecret := func(name string) hashicorpSecretConfig {
		return hashicorpSecretConfig{AddressSecret: name, AddressSecretVersion: 1, SecretEngine: "kv"}
	}

	const (
		secretEngine = "kv"
		secret1 = "sec1"
		secret2 = "sec2"
		multiValSecret = "multiValSec"
	)

	mux := http.NewServeMux()

	mux.HandleFunc(fmt.Sprintf("/v1/%s/data/%s", secretEngine, secret1), func(w http.ResponseWriter, r *http.Request) {
		body := makeVaultResponse(map[string]string{
			"address": "ed9d02e382b34818e88b88a309c7fe71e65f419d",
		})

		w.Write(body)
	})

	mux.HandleFunc(fmt.Sprintf("/v1/%s/data/%s", secretEngine, secret2), func(w http.ResponseWriter, r *http.Request) {
		body := makeVaultResponse(map[string]string{
			"otherAddress": "ca843569e3427144cead5e4d5999a3d0ccf92b8e",
		})

		w.Write(body)
	})

	mux.HandleFunc(fmt.Sprintf("/v1/%s/data/%s", secretEngine, multiValSecret), func(w http.ResponseWriter, r *http.Request) {
		body := makeVaultResponse(map[string]string{
			"address": "ed9d02e382b34818e88b88a309c7fe71e65f419d",
			"otherAddress": "ca843569e3427144cead5e4d5999a3d0ccf92b8e",
		})

		w.Write(body)
	})

	vaultServer := httptest.NewServer(mux)
	defer vaultServer.Close()

	tests := map[string]struct{
		secrets []hashicorpSecretConfig
		wantAccts []accounts.Account
	}{
		"account retrieved": {
			secrets:   []hashicorpSecretConfig{makeSecret(secret1)},
			wantAccts: []accounts.Account{
				{Address: common.HexToAddress("ed9d02e382b34818e88b88a309c7fe71e65f419d")},
			},
		},
		"account not retrieved when vault secret has multiple values": {
			secrets:   []hashicorpSecretConfig{makeSecret(multiValSecret)},
			wantAccts: []accounts.Account{},
		},
		"unretrievable accounts are ignored": {
			secrets:   []hashicorpSecretConfig{makeSecret(multiValSecret), makeSecret(secret1)},
			wantAccts: []accounts.Account{
				{Address: common.HexToAddress("ed9d02e382b34818e88b88a309c7fe71e65f419d")},
			},
		},
		"accounts retrieved regardless of vault secrets keyvalue key": {
			secrets: []hashicorpSecretConfig{makeSecret(secret1), makeSecret(secret2)},
			wantAccts: []accounts.Account{
				{Address: common.HexToAddress("ed9d02e382b34818e88b88a309c7fe71e65f419d")},
				{Address: common.HexToAddress("ca843569e3427144cead5e4d5999a3d0ccf92b8e")},
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			wltConfig := hashicorpWalletConfig{
				Client: hashicorpClientConfig{
					Url: vaultServer.URL,
					VaultPollingIntervalMillis: 1,
				},
				Secrets: tt.secrets,
			}

			w, err := newHashicorpWallet(wltConfig, &event.Feed{})

			if err != nil {
				t.Fatal(err)
			}

			if err := w.Open(""); err != nil {
				t.Fatal(err)
			}

			// need to block to let accountRetrievalLoop do its thing
			time.Sleep(4 * time.Millisecond)

			//TODO wantAccts do not have URLs set so URL equality is not being checked
			if !acctsEqual(tt.wantAccts, w.Accounts()) {
				t.Fatalf("wallet accounts do not equal wanted accounts\nwant: %v\ngot : %v", tt.wantAccts, w.Accounts())
			}
		})
	}
}

// TODO This is a long running test (>5s) so perhaps should be excluded from test suite by default?
func TestVaultWallet_Open_Hashicorp_AccountsRetrievedWhenVaultAvailable(t *testing.T) {
	if err := os.Setenv(api.EnvVaultToken, "mytoken"); err != nil {
		t.Fatal(err)
	}

	makeVaultResponse := func(keyValPairs map[string]string) []byte {
		resp := api.Secret{
			Data: map[string]interface{}{
				"data": keyValPairs,
			},
		}

		b, err := json.Marshal(resp)

		if err != nil {
			t.Fatal(err)
		}

		return b
	}

	body := makeVaultResponse(map[string]string{"address": "ed9d02e382b34818e88b88a309c7fe71e65f419d"})
	vaultServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(body)
	}))
	defer vaultServer.Close()

	// use an incorrect vault url to simulate an inaccessible vault
	wltConfig := hashicorpWalletConfig{
		Client: hashicorpClientConfig{
			Url: "http://incorrecturl:1",
			VaultPollingIntervalMillis: 1,
		},
		Secrets: []hashicorpSecretConfig{
			{AddressSecret: "sec1", AddressSecretVersion: 1, SecretEngine: "kv"},
		},
	}

	w, err := newHashicorpWallet(wltConfig, &event.Feed{})

	if err != nil {
		t.Fatal(err)
	}

	if err := w.Open(""); err != nil {
		t.Fatal(err)
	}

	// need to block to let accountRetrievalLoop do its thing
	// a long sleep is used here to give the vault client time to make its request to the vault and wait for the response before the go scheduler returns focus to this test
	// such a long sleep is required because the vault client retries multiple times before determining that the vault cannot be reached
	time.Sleep(3 * time.Second)

	if len(w.Accounts()) != 0 {
		t.Fatalf("wallet should have no accounts as vault server is inaccessible: got: %v", w.Accounts())
	}

	// update vault client to use correct url to simulate vault becoming accessible
	v := w.vault.(*hashicorpService)
	if err := v.client.SetAddress(vaultServer.URL); err != nil {
		t.Fatal(err)
	}

	// need to block to let accountRetrievalLoop do its thing
	// a long sleep is used here to give the vault client time to make its request to the vault and wait for the response before the go scheduler returns focus to this test
	// such a long sleep is required because we must wait for the vault client to finish any attempted request to the incorrect vault url
	time.Sleep(3 * time.Second)

	wantAccts := []accounts.Account{
		{Address: common.HexToAddress("ed9d02e382b34818e88b88a309c7fe71e65f419d")},
	}

	//TODO wantAccts do not have URLs set so URL equality is not being checked
	if !acctsEqual(wantAccts, w.Accounts()) {
		t.Fatalf("wallet accounts do not equal wanted accounts\nwant: %v\ngot : %v", wantAccts, w.Accounts())
	}
}

type keysByD []*ecdsa.PrivateKey

func (k keysByD) Len() int {
	return len(k)
}

func (k keysByD) Less(i, j int) bool {
	return (k[i].D).Cmp(k[j].D) < 0
}

func (k keysByD) Swap(i, j int) {
	k[i], k[j] = k[j], k[i]
}

func keysEqual(a, b []*ecdsa.PrivateKey) bool {
	if len(a) != len(b) {
		return false
	}

	sort.Sort(keysByD(a))
	sort.Sort(keysByD(b))

	equal := func(a, b *ecdsa.PrivateKey) bool {
		return a.D.Cmp(b.D) == 0
	}

	for i := 0; i < len(a); i++ {
		if !equal(a[i], b[i]) {
			return false
		}
	}

	return true
}

func TestVaultWallet_Open_Hashicorp_PrivateKeysRetrievedWhenEnabled(t *testing.T) {
	if err := os.Setenv(api.EnvVaultToken, "mytoken"); err != nil {
		t.Fatal(err)
	}

	makeVaultResponse := func(keyValPairs map[string]string) []byte {
		resp := api.Secret{
			Data: map[string]interface{}{
				"data": keyValPairs,
			},
		}

		b, err := json.Marshal(resp)

		if err != nil {
			t.Fatal(err)
		}

		return b
	}

	makeSecret := func(addrName, keyName string) hashicorpSecretConfig {
		return hashicorpSecretConfig{AddressSecret: addrName, AddressSecretVersion: 1, PrivateKeySecret: keyName, PrivateKeySecretVersion: 1, SecretEngine: "kv"}
	}

	makeKey := func(hex string) *ecdsa.PrivateKey {
		key, err := crypto.HexToECDSA(hex)

		if err != nil {
			t.Fatal(err)
		}

		return key
	}

	const (
		secretEngine = "kv"
		key1 = "key1"
		key2 = "key2"
		addr1 = "addr1"
		addr2 = "addr2"
		multiValSecret = "multiValSec"
	)

	mux := http.NewServeMux()

	mux.HandleFunc(fmt.Sprintf("/v1/%s/data/%s", secretEngine, addr1), func(w http.ResponseWriter, r *http.Request) {
		body := makeVaultResponse(map[string]string{
			"addr": "ed9d02e382b34818e88b88a309c7fe71e65f419d",
		})

		w.Write(body)
	})

	mux.HandleFunc(fmt.Sprintf("/v1/%s/data/%s", secretEngine, key1), func(w http.ResponseWriter, r *http.Request) {
		body := makeVaultResponse(map[string]string{
			"key": "e6181caaffff94a09d7e332fc8da9884d99902c7874eb74354bdcadf411929f1",
		})

		w.Write(body)
	})

	mux.HandleFunc(fmt.Sprintf("/v1/%s/data/%s", secretEngine, addr2), func(w http.ResponseWriter, r *http.Request) {
		body := makeVaultResponse(map[string]string{
			"addr": "ca843569e3427144cead5e4d5999a3d0ccf92b8e",
		})

		w.Write(body)
	})

	mux.HandleFunc(fmt.Sprintf("/v1/%s/data/%s", secretEngine, key2), func(w http.ResponseWriter, r *http.Request) {
		body := makeVaultResponse(map[string]string{
			"otherKey": "4762e04d10832808a0aebdaa79c12de54afbe006bfffd228b3abcc494fe986f9",
		})

		w.Write(body)
	})

	mux.HandleFunc(fmt.Sprintf("/v1/%s/data/%s", secretEngine, multiValSecret), func(w http.ResponseWriter, r *http.Request) {
		body := makeVaultResponse(map[string]string{
			"key": "e6181caaffff94a09d7e332fc8da9884d99902c7874eb74354bdcadf411929f1",
			"otherKey": "4762e04d10832808a0aebdaa79c12de54afbe006bfffd228b3abcc494fe986f9",
		})

		w.Write(body)
	})

	vaultServer := httptest.NewServer(mux)
	defer vaultServer.Close()

	tests := map[string]struct{
		secrets []hashicorpSecretConfig
		wantKeys []*ecdsa.PrivateKey
	}{
		"key retrieved": {
			secrets:   []hashicorpSecretConfig{makeSecret(addr1, key1)},
			wantKeys: []*ecdsa.PrivateKey{
				makeKey("e6181caaffff94a09d7e332fc8da9884d99902c7874eb74354bdcadf411929f1"),
			},
		},
		"key not retrieved when vault secret has multiple values": {
			secrets:   []hashicorpSecretConfig{makeSecret(addr1, multiValSecret)},
			wantKeys: []*ecdsa.PrivateKey{},
		},
		"unretrievable keys are ignored": {
			secrets:   []hashicorpSecretConfig{makeSecret(addr1, multiValSecret), makeSecret(addr2, key2)},
			wantKeys: []*ecdsa.PrivateKey{
				makeKey("4762e04d10832808a0aebdaa79c12de54afbe006bfffd228b3abcc494fe986f9"),
			},
		},
		"keys retrieved regardless of vault secrets keyvalue key": {
			secrets: []hashicorpSecretConfig{makeSecret(addr1, key1), makeSecret(addr2, key2)},
			wantKeys: []*ecdsa.PrivateKey{
				makeKey("e6181caaffff94a09d7e332fc8da9884d99902c7874eb74354bdcadf411929f1"),				makeKey("4762e04d10832808a0aebdaa79c12de54afbe006bfffd228b3abcc494fe986f9"),
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			wltConfig := hashicorpWalletConfig{
				Client: hashicorpClientConfig{
					Url: vaultServer.URL,
					VaultPollingIntervalMillis: 1,
					StorePrivateKeys: true,
				},
				Secrets: tt.secrets,
			}

			w, err := newHashicorpWallet(wltConfig, &event.Feed{})

			if err != nil {
				t.Fatal(err)
			}

			if err := w.Open(""); err != nil {
				t.Fatal(err)
			}

			// need to block to let accountRetrievalLoop do its thing
			time.Sleep(4 * time.Millisecond)

			keyHandlers := w.vault.(*hashicorpService).keyHandlers

			gotKeys := getRetrievedKeys(keyHandlers)

			if !keysEqual(tt.wantKeys, gotKeys) {
				t.Fatalf("keys in vaultService do not equal wanted keys\nwant: %v\ngot : %v", tt.wantKeys, gotKeys)
			}
		})
	}
}

func getRetrievedKeys(keyHandlers map[common.Address]map[accounts.URL]hashicorpKeyHandler) []*ecdsa.PrivateKey {
	gotKeys := []*ecdsa.PrivateKey{}

	for _, kh := range keyHandlers {
		for _, gg := range kh {
			if gg.key != nil {
				gotKeys = append(gotKeys, gg.key)
			}
		}
	}

	return gotKeys
}

// TODO This is a long running test (>10s) so perhaps should be excluded from test suite by default?
func TestVaultWallet_Open_Hashicorp_PrivateKeysRetrievedWhenEnabledAndVaultAvailable(t *testing.T) {
	if err := os.Setenv(api.EnvVaultToken, "mytoken"); err != nil {
		t.Fatal(err)
	}

	makeKey := func(hex string) *ecdsa.PrivateKey {
		key, err := crypto.HexToECDSA(hex)

		if err != nil {
			t.Fatal(err)
		}

		return key
	}

	makeVaultResponse := func(keyValPairs map[string]string) []byte {
		resp := api.Secret{
			Data: map[string]interface{}{
				"data": keyValPairs,
			},
		}

		b, err := json.Marshal(resp)

		if err != nil {
			t.Fatal(err)
		}

		return b
	}

	body := makeVaultResponse(map[string]string{"key": "e6181caaffff94a09d7e332fc8da9884d99902c7874eb74354bdcadf411929f1"})
	vaultServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(body)
	}))
	defer vaultServer.Close()

	tests := map[string]struct{
		storePrivateKeys bool
		wantKeys []*ecdsa.PrivateKey
	}{
		"disabled": {storePrivateKeys: false, wantKeys: []*ecdsa.PrivateKey{}},
		"enabled": {storePrivateKeys: true, wantKeys: []*ecdsa.PrivateKey{makeKey("e6181caaffff94a09d7e332fc8da9884d99902c7874eb74354bdcadf411929f1")}},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			// use an incorrect vault url to simulate an inaccessible vault
			wltConfig := hashicorpWalletConfig{
				Client: hashicorpClientConfig{
					Url: "http://incorrecturl:1",
					VaultPollingIntervalMillis: 1,
					StorePrivateKeys: tt.storePrivateKeys,
				},
				Secrets: []hashicorpSecretConfig{
					{PrivateKeySecret: "sec1", PrivateKeySecretVersion: 1, SecretEngine: "kv"},
				},
			}

			w, err := newHashicorpWallet(wltConfig, &event.Feed{})

			if err != nil {
				t.Fatal(err)
			}

			if err := w.Open(""); err != nil {
				t.Fatal(err)
			}

			// need to block to let accountRetrievalLoop do its thing
			// a long sleep is used here to give the vault client time to make its request to the vault and wait for the response before the go scheduler returns focus to this test
			// such a long sleep is required because the vault client retries multiple times before determining that the vault cannot be reached
			time.Sleep(3 * time.Second)

			v := w.vault.(*hashicorpService)

			gotKeys := getRetrievedKeys(v.keyHandlers)

			if len(gotKeys) != 0 {
				t.Fatalf("vaultService should have no keys as vault server is inaccessible: got: %v", gotKeys)
			}

			// update vault client to use correct url to simulate vault becoming accessible
			if err := v.client.SetAddress(vaultServer.URL); err != nil {
				t.Fatal(err)
			}

			// need to block to let accountRetrievalLoop do its thing
			// a long sleep is used here to give the vault client time to make its request to the vault and wait for the response before the go scheduler returns focus to this test
			// such a long sleep is required because we must wait for the vault client to finish any attempted request to the incorrect vault url
			time.Sleep(3 * time.Second)

			gotKeys = getRetrievedKeys(v.keyHandlers)

			if !keysEqual(tt.wantKeys, gotKeys) {
				t.Fatalf("keys in vaultService do not equal wanted keys\nwant: %v\ngot : %v", tt.wantKeys, gotKeys)
			}
		})
	}
}

func TestVaultWallet_Open_Hashicorp_RetrievalLoopsStopWhenAllSecretsRetrieved(t *testing.T) {
	if err := os.Setenv(api.EnvVaultToken, "mytoken"); err != nil {
		t.Fatal(err)
	}

	makeVaultResponse := func(keyValPairs map[string]string) []byte {
		resp := api.Secret{
			Data: map[string]interface{}{
				"data": keyValPairs,
			},
		}

		b, err := json.Marshal(resp)

		if err != nil {
			t.Fatal(err)
		}

		return b
	}

	makeSecret := func(addrName, keyName string) hashicorpSecretConfig {
		return hashicorpSecretConfig{AddressSecret: addrName, AddressSecretVersion: 1, PrivateKeySecret: keyName, PrivateKeySecretVersion: 1, SecretEngine: "kv"}
	}

	const (
		secretEngine = "kv"
		addrName = "addr1"
		keyName = "key1"
	)

	var getAddrCount, getKeyCount int

	mux := http.NewServeMux()

	mux.HandleFunc(fmt.Sprintf("/v1/%s/data/%s", secretEngine, addrName), func(w http.ResponseWriter, r *http.Request) {
		getAddrCount++

		body := makeVaultResponse(map[string]string{
			"addr": "ed9d02e382b34818e88b88a309c7fe71e65f419d",
		})

		w.Write(body)
	})

	mux.HandleFunc(fmt.Sprintf("/v1/%s/data/%s", secretEngine, keyName), func(w http.ResponseWriter, r *http.Request) {
		getKeyCount++

		body := makeVaultResponse(map[string]string{
			"key": "e6181caaffff94a09d7e332fc8da9884d99902c7874eb74354bdcadf411929f1",
		})

		w.Write(body)
	})

	vaultServer := httptest.NewServer(mux)
	defer vaultServer.Close()

	wltConfig := hashicorpWalletConfig{
		Client: hashicorpClientConfig{
			Url: vaultServer.URL,
			VaultPollingIntervalMillis: 1,
			StorePrivateKeys: true,
		},
		Secrets: []hashicorpSecretConfig{makeSecret(addrName, keyName)},
	}

	w, err := newHashicorpWallet(wltConfig, &event.Feed{})

	if err != nil {
		t.Fatal(err)
	}

	if err := w.Open(""); err != nil {
		t.Fatal(err)
	}

	// need to block to let accountRetrievalLoop do its thing
	// the sleep length is long enough that multiple calls to the vault would occur if the loop did not stop once all secrets were retrieved
	time.Sleep(10 * time.Millisecond)

	if getAddrCount != 1 || getKeyCount != 1 {
		t.Fatalf("retrieval loops should have made just one call to vault, got secret and then stopped: \naccountRetrievalLoop vault call count: %v\nprivateKeyRetrievalLoop vault call count: %v", getAddrCount, getKeyCount)
	}
}

func TestVaultWallet_Close_Hashicorp_ReturnsStateToBeforeOpen(t *testing.T) {
	if err := os.Setenv(api.EnvVaultToken, "mytoken"); err != nil {
		t.Fatal(err)
	}

	config := hashicorpWalletConfig{
		Client: hashicorpClientConfig{Url: "http://url:1"},
		Secrets: []hashicorpSecretConfig{{AddressSecret: "addr1"}},
	}

	w, err := newHashicorpWallet(config, &event.Feed{})

	if err != nil {
		t.Fatal(err)
	}

	unopened, err := newHashicorpWallet(config, &event.Feed{})

	if err != nil {
		t.Fatal(err)
	}

	cmpOpts := []cmp.Option{
		cmp.AllowUnexported(vaultWallet{}, hashicorpService{}),
		cmpopts.IgnoreUnexported(event.Feed{}, sync.RWMutex{}),
	}

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

func TestVaultWallet_SignHash_Hashicorp_ErrorIfAccountNotKnown(t *testing.T) {
	w := vaultWallet{
		vault: &hashicorpService{
			accts: []accounts.Account{},
		},
	}

	acct := accounts.Account{Address: common.HexToAddress("ed9d02e382b34818e88b88a309c7fe71e65f419d")}

	toSign := crypto.Keccak256([]byte("to_sign"))

	if _, err := w.SignHash(acct, toSign); err != accounts.ErrUnknownAccount {
		t.Fatalf("incorrect error returned:\nwant: %v\ngot : %v", accounts.ErrUnknownAccount, err)
	}
}

func TestVaultWallet_SignHash_Hashicorp_SignsWithInMemoryKeyIfAvailableAndDoesNotZeroKey(t *testing.T) {
	addr := common.HexToAddress("ed9d02e382b34818e88b88a309c7fe71e65f419d")
	url := accounts.URL{Scheme: "http", Path: "url:1"}
	acct := accounts.Account{
		Address: addr,
		URL: url,
	}

	key, err := crypto.HexToECDSA("e6181caaffff94a09d7e332fc8da9884d99902c7874eb74354bdcadf411929f1")

	if err != nil {
		t.Fatal(err)
	}

	w := vaultWallet{
		vault: &hashicorpService{
			accts: []accounts.Account{acct},
			keyHandlers: map[common.Address]map[accounts.URL]hashicorpKeyHandler{
				addr: {
					url: hashicorpKeyHandler{key: key},
				},
			},
		},
	}

	toSign := crypto.Keccak256([]byte("to_sign"))

	got, err := w.SignHash(acct, toSign)

	if err != nil {
		t.Fatalf("error signing hash: %v", err)
	}

	want, err := crypto.Sign(toSign, key)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(want, got) {
		t.Fatalf("incorrect signHash result:\nwant: %v\ngot : %v", want, got)
	}

	vaultServiceKey := w.vault.(*hashicorpService).keyHandlers[acct.Address][acct.URL].key

	if vaultServiceKey == nil || vaultServiceKey.D.Int64() == 0 {
		t.Fatal("unlocked key was zeroed after use")
	}
}

func TestVaultWallet_SignHash_Hashicorp_ErrorIfAmbiguousAccount(t *testing.T) {
	addr := common.HexToAddress("ed9d02e382b34818e88b88a309c7fe71e65f419d")

	url1 := accounts.URL{Scheme: "http", Path: "url:1"}
	url2 := accounts.URL{Scheme: "http", Path: "url:2"}

	acct1 := accounts.Account{Address: addr, URL: url1}
	acct2 := accounts.Account{Address: addr, URL: url2}

	// Two accounts have the same address but different URLs
	w := vaultWallet{
		vault: &hashicorpService{
			accts: []accounts.Account{acct1, acct2},
			keyHandlers: map[common.Address]map[accounts.URL]hashicorpKeyHandler{
				addr: {
					url1: hashicorpKeyHandler{},
					url2: hashicorpKeyHandler{},
				},
			},
		},
	}

	toSign := crypto.Keccak256([]byte("to_sign"))

	// The provided account does not specify the exact account to use as no URL is provided
	acct := accounts.Account{
		Address: addr,
	}

	_, err := w.SignHash(acct, toSign)
	e := err.(*keystore.AmbiguousAddrError)

	want := []accounts.Account{acct1, acct2}

	if diff := cmp.Diff(want, e.Matches); diff != "" {
		t.Fatalf("ambiguous accounts mismatch (-want +got):\n%s", diff)
	}
}

func TestVaultWallet_SignHash_Hashicorp_AmbiguousAccountAllowedIfOnlyOneAccountWithGivenAddress(t *testing.T) {
	addr := common.HexToAddress("ed9d02e382b34818e88b88a309c7fe71e65f419d")
	url := accounts.URL{Scheme: "http", Path: "url:1"}
	acct1 := accounts.Account{Address: addr, URL: url}

	key, err := crypto.HexToECDSA("e6181caaffff94a09d7e332fc8da9884d99902c7874eb74354bdcadf411929f1")

	if err != nil {
		t.Fatal(err)
	}

	w := vaultWallet{
		vault: &hashicorpService{
			accts: []accounts.Account{acct1},
			keyHandlers: map[common.Address]map[accounts.URL]hashicorpKeyHandler{
				addr: {
					url: hashicorpKeyHandler{key: key},
				},
			},
		},
	}

	toSign := crypto.Keccak256([]byte("to_sign"))

	// The provided account does not specify the exact account to use as no URL is provided
	acct := accounts.Account{
		Address: addr,
	}

	got, err := w.SignHash(acct, toSign)

	if err != nil {
		t.Fatalf("error signing hash: %v", err)
	}

	want, err := crypto.Sign(toSign, key)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(want, got) {
		t.Fatalf("incorrect signHash result:\nwant: %v\ngot : %v", want, got)
	}

	vaultServiceKeyHandlers := w.vault.(*hashicorpService).keyHandlers[acct.Address]

	var vaultServiceKey *ecdsa.PrivateKey

	for _, kh := range vaultServiceKeyHandlers {
		vaultServiceKey = kh.key

		if vaultServiceKey == nil || vaultServiceKey.D.Int64() == 0 {
			t.Fatal("unlocked key was zeroed after use")
		}
	}
}

func TestVaultWallet_SignHash_Hashicorp_SignsWithKeyFromVaultAndDoesNotStoreInMemory(t *testing.T) {
	makeMockHashicorpResponse := func(t *testing.T, hexKey string) []byte {
		var vaultResponse api.Secret

		vaultResponse.Data = map[string]interface{}{
			"data": map[string]interface{}{
				"key": hexKey,
			},
		}

		b, err := json.Marshal(vaultResponse)

		if err != nil {
			t.Fatalf("err marshalling mock response: %v", err)
		}

		return b
	}

	acct := accounts.Account{
		Address: common.HexToAddress("ed9d02e382b34818e88b88a309c7fe71e65f419d"),
		URL: accounts.URL{Scheme: "http", Path: "url:1"},
	}

	hexKey := "e6181caaffff94a09d7e332fc8da9884d99902c7874eb74354bdcadf411929f1"
	key, err := crypto.HexToECDSA(hexKey)

	if err != nil {
		t.Fatal(err)
	}

	client, cleanup := makeMockHashicorpClient(t, makeMockHashicorpResponse(t, hexKey))
	defer cleanup()

	secret := hashicorpSecretConfig{
		PrivateKeySecret: "mykey",
		PrivateKeySecretVersion: 1,
		SecretEngine: "kv",
	}

	w := vaultWallet{
		vault: &hashicorpService{
			client: client,
			accts: []accounts.Account{acct},
			keyHandlers: map[common.Address]map[accounts.URL]hashicorpKeyHandler{
				acct.Address: {
					acct.URL: {
						secret: secret,
					},
				},
			},
		},
	}

	toSign := crypto.Keccak256([]byte("to_sign"))

	got, err := w.SignHash(acct, toSign)

	if err != nil {
		t.Fatalf("error signing hash: %v", err)
	}

	want, err := crypto.Sign(toSign, key)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(want, got) {
		t.Fatalf("incorrect signHash result:\nwant: %v\ngot : %v", want, got)
	}

	vaultServiceKey := w.vault.(*hashicorpService).keyHandlers[acct.Address][acct.URL].key

	if vaultServiceKey != nil {
		t.Fatal("unlocked key should not be stored after use")
	}
}

func TestVaultWallet_SignTx_Hashicorp_UsesDifferentSigners(t *testing.T) {
	addr := common.HexToAddress("ed9d02e382b34818e88b88a309c7fe71e65f419d")
	url := accounts.URL{Scheme: "http", Path: "url:1"}
	acct := accounts.Account{
		Address: addr,
		URL: url,
	}

	key, err := crypto.HexToECDSA("e6181caaffff94a09d7e332fc8da9884d99902c7874eb74354bdcadf411929f1")

	if err != nil {
		t.Fatal(err)
	}

	w := vaultWallet{
		vault: &hashicorpService{
			accts: []accounts.Account{acct},
			keyHandlers: map[common.Address]map[accounts.URL]hashicorpKeyHandler{
				addr: {
					url: hashicorpKeyHandler{key: key},
				},
			},
		},
	}

	makePublicTx := func() *types.Transaction {
		return types.NewTransaction(0, common.Address{}, nil, 0, nil, nil)
	}

	makePrivateTx := func() *types.Transaction {
		tx := makePublicTx()
		tx.SetPrivate()
		return tx
	}

	tests := map[string]struct{
		toSign *types.Transaction
		signer types.Signer
		chainID *big.Int
	}{
		"private tx no chainID uses QuorumPrivateTxSigner": {toSign: makePrivateTx(), signer: types.QuorumPrivateTxSigner{}},
		"private tx and chainID uses QuorumPrivateTxSigner": {toSign: makePrivateTx(), signer: types.QuorumPrivateTxSigner{}, chainID: big.NewInt(1)},
		"public tx no chainID uses HomesteadSigner": {toSign: makePublicTx(), signer: types.HomesteadSigner{}},
		"public tx and chainID uses EIP155Signer": {toSign: makePublicTx(), signer: types.NewEIP155Signer(big.NewInt(1)), chainID: big.NewInt(1)},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := w.SignTx(acct, tt.toSign, tt.chainID)

			if err != nil {
				t.Fatalf("error signing tx: %v", err)
			}

			h := tt.signer.Hash(tt.toSign)
			wantSignature, err := crypto.Sign(h[:], key)

			if err != nil {
				t.Fatal(err)
			}

			var toSignCpy types.Transaction
			toSignCpy = *tt.toSign
			want, err := toSignCpy.WithSignature(tt.signer, wantSignature)

			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(want, got) {
				t.Fatalf("incorrect signTx result :\nwant: %v\ngot : %v", want, got)
			}
		})
	}
}

func TestVaultWallet_SignTx_Hashicorp_ErrorIfAccountNotKnown(t *testing.T) {
	w := vaultWallet{
		vault: &hashicorpService{
			accts: []accounts.Account{},
		},
	}

	acct := accounts.Account{Address: common.HexToAddress("ed9d02e382b34818e88b88a309c7fe71e65f419d")}

	toSign := &types.Transaction{}

	if _, err := w.SignTx(acct, toSign, nil); err != accounts.ErrUnknownAccount {
		t.Fatalf("incorrect error returned:\nwant: %v\ngot : %v", accounts.ErrUnknownAccount, err)
	}
}

func TestVaultWallet_SignTx_Hashicorp_SignsWithInMemoryKeyIfAvailableAndDoesNotZeroKey(t *testing.T) {
	addr := common.HexToAddress("ed9d02e382b34818e88b88a309c7fe71e65f419d")
	url := accounts.URL{Scheme: "http", Path: "url:1"}
	acct := accounts.Account{
		Address: addr,
		URL: url,
	}

	key, err := crypto.HexToECDSA("e6181caaffff94a09d7e332fc8da9884d99902c7874eb74354bdcadf411929f1")

	if err != nil {
		t.Fatal(err)
	}

	w := vaultWallet{
		vault: &hashicorpService{
			accts: []accounts.Account{acct},
			keyHandlers: map[common.Address]map[accounts.URL]hashicorpKeyHandler{
				addr: {
					url: hashicorpKeyHandler{key: key},
				},
			},
		},
	}

	toSign := types.NewTransaction(0, common.Address{}, nil, 0, nil, nil)

	got, err := w.SignTx(acct, toSign, nil)

	if err != nil {
		t.Fatalf("error signing hash: %v", err)
	}

	wantSigner := types.HomesteadSigner{}
	h := wantSigner.Hash(toSign)
	wantSignature, err := crypto.Sign(h[:], key)

	if err != nil {
		t.Fatal(err)
	}

	var toSignCpy types.Transaction
	toSignCpy = *toSign
	want, err := toSignCpy.WithSignature(wantSigner, wantSignature)

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("incorrect signTx result :\nwant: %v\ngot : %v", want, got)
	}

	vaultServiceKey := w.vault.(*hashicorpService).keyHandlers[acct.Address][acct.URL].key

	if vaultServiceKey == nil || vaultServiceKey.D.Int64() == 0 {
		t.Fatal("unlocked key was zeroed after use")
	}
}

func TestVaultWallet_SignTx_Hashicorp_ErrorIfAmbiguousAccount(t *testing.T) {
	addr := common.HexToAddress("ed9d02e382b34818e88b88a309c7fe71e65f419d")

	url1 := accounts.URL{Scheme: "http", Path: "url:1"}
	url2 := accounts.URL{Scheme: "http", Path: "url:2"}

	acct1 := accounts.Account{Address: addr, URL: url1}
	acct2 := accounts.Account{Address: addr, URL: url2}

	// Two accounts have the same address but different URLs
	w := vaultWallet{
		vault: &hashicorpService{
			accts: []accounts.Account{acct1, acct2},
			keyHandlers: map[common.Address]map[accounts.URL]hashicorpKeyHandler{
				addr: {
					url1: hashicorpKeyHandler{},
					url2: hashicorpKeyHandler{},
				},
			},
		},
	}

	toSign := types.NewTransaction(0, common.Address{}, nil, 0, nil, nil)

	// The provided account does not specify the exact account to use as no URL is provided
	acct := accounts.Account{
		Address: addr,
	}

	_, err := w.SignTx(acct, toSign, nil)
	e := err.(*keystore.AmbiguousAddrError)

	want := []accounts.Account{acct1, acct2}

	if diff := cmp.Diff(want, e.Matches); diff != "" {
		t.Fatalf("ambiguous accounts mismatch (-want +got):\n%s", diff)
	}
}

func TestVaultWallet_SignTx_Hashicorp_AmbiguousAccountAllowedIfOnlyOneAccountWithGivenAddress(t *testing.T) {
	addr := common.HexToAddress("ed9d02e382b34818e88b88a309c7fe71e65f419d")
	url := accounts.URL{Scheme: "http", Path: "url:1"}
	acct1 := accounts.Account{Address: addr, URL: url}

	key, err := crypto.HexToECDSA("e6181caaffff94a09d7e332fc8da9884d99902c7874eb74354bdcadf411929f1")

	if err != nil {
		t.Fatal(err)
	}

	w := vaultWallet{
		vault: &hashicorpService{
			accts: []accounts.Account{acct1},
			keyHandlers: map[common.Address]map[accounts.URL]hashicorpKeyHandler{
				addr: {
					url: hashicorpKeyHandler{key: key},
				},
			},
		},
	}

	toSign := types.NewTransaction(0, common.Address{}, nil, 0, nil, nil)

	// The provided account does not specify the exact account to use as no URL is provided
	acct := accounts.Account{
		Address: addr,
	}

	got, err := w.SignTx(acct, toSign, nil)

	if err != nil {
		t.Fatalf("error signing hash: %v", err)
	}

	wantSigner := types.HomesteadSigner{}
	h := wantSigner.Hash(toSign)
	wantSignature, err := crypto.Sign(h[:], key)

	if err != nil {
		t.Fatal(err)
	}

	var toSignCpy types.Transaction
	toSignCpy = *toSign
	want, err := toSignCpy.WithSignature(wantSigner, wantSignature)

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("incorrect signTx result :\nwant: %v\ngot : %v", want, got)
	}

	vaultServiceKeyHandlers := w.vault.(*hashicorpService).keyHandlers[acct.Address]

	var vaultServiceKey *ecdsa.PrivateKey

	for _, kh := range vaultServiceKeyHandlers {
		vaultServiceKey = kh.key

		if vaultServiceKey == nil || vaultServiceKey.D.Int64() == 0 {
			t.Fatal("unlocked key was zeroed after use")
		}
	}
}

func TestVaultWallet_SignTx_Hashicorp_SignsWithKeyFromVaultAndDoesNotStoreInMemory(t *testing.T) {
	makeMockHashicorpResponse := func(t *testing.T, hexKey string) []byte {
		var vaultResponse api.Secret

		vaultResponse.Data = map[string]interface{}{
			"data": map[string]interface{}{
				"key": hexKey,
			},
		}

		b, err := json.Marshal(vaultResponse)

		if err != nil {
			t.Fatalf("err marshalling mock response: %v", err)
		}

		return b
	}

	acct := accounts.Account{
		Address: common.HexToAddress("ed9d02e382b34818e88b88a309c7fe71e65f419d"),
		URL: accounts.URL{Scheme: "http", Path: "url:1"},
	}

	hexKey := "e6181caaffff94a09d7e332fc8da9884d99902c7874eb74354bdcadf411929f1"
	key, err := crypto.HexToECDSA(hexKey)

	if err != nil {
		t.Fatal(err)
	}

	client, cleanup := makeMockHashicorpClient(t, makeMockHashicorpResponse(t, hexKey))
	defer cleanup()

	secret := hashicorpSecretConfig{
		PrivateKeySecret: "mykey",
		PrivateKeySecretVersion: 1,
		SecretEngine: "kv",
	}

	w := vaultWallet{
		vault: &hashicorpService{
			client: client,
			accts: []accounts.Account{acct},
			keyHandlers: map[common.Address]map[accounts.URL]hashicorpKeyHandler{
				acct.Address: {
					acct.URL: {
						secret: secret,
					},
				},
			},
		},
	}

	toSign := types.NewTransaction(0, common.Address{}, nil, 0, nil, nil)

	got, err := w.SignTx(acct, toSign, nil)

	if err != nil {
		t.Fatalf("error signing hash: %v", err)
	}

	wantSigner := types.HomesteadSigner{}
	h := wantSigner.Hash(toSign)
	wantSignature, err := crypto.Sign(h[:], key)

	if err != nil {
		t.Fatal(err)
	}

	var toSignCpy types.Transaction
	toSignCpy = *toSign
	want, err := toSignCpy.WithSignature(wantSigner, wantSignature)

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("incorrect signTx result :\nwant: %v\ngot : %v", want, got)
	}

	vaultServiceKey := w.vault.(*hashicorpService).keyHandlers[acct.Address][acct.URL].key

	if vaultServiceKey != nil {
		t.Fatal("unlocked key should not be stored after use")
	}
}

package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/rpc"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"testing"
)


const targetHost = "http://localhost:22000"

func TestHttpNetUnauthorizedResponse(t *testing.T){
	payloadBytes, err := json.Marshal(
		map[string]interface{}{
			"id":1,
			"jsonrpc": "2.0",
			"params" : []string{},
			"method":  "eth_accounts",
		})

	if err != nil {
		log.Fatalln(err)
	}

	resp , err := http.Post(targetHost, "application/json",bytes.NewBuffer(payloadBytes))
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusUnauthorized {
		t.Fatal("Expected status 401, recv status:" ,resp.StatusCode)
	}
}

func TestHttpRpcHealth(t *testing.T){
	resp , err := http.Get(fmt.Sprintf("%s/%s", targetHost, "health-check"))

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatal("Expected status 200, recv status:" ,resp.StatusCode)
	}
}

func TestRpcSecurityConfigLoading(t *testing.T) {
	// Write config file
	securityConfigFile := path.Join(getTestDirectory(),"/rpc-security.json")
	securityConfig := rpc.SecurityConfig{ProviderType: "local"}
	securityConfigJson, _ := json.Marshal(securityConfig)
	err := ioutil.WriteFile(securityConfigFile,securityConfigJson, 0644)
	if err != nil {
		t.Fatalf(err.Error())
	}

	// Read config file
	config, err  := rpc.ParseRpcSecurityConfigFile(securityConfigFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if config.ProviderType !=  securityConfig.ProviderType {
		t.Fatalf("TestRpcSecurityConfigLoading error")
	}
}


// Return Test Directory
func getTestDirectory() string {
	testFolder, found := os.LookupEnv("RPC_SECURITY_TESTING_FOLDER")
	if found {
		return testFolder
	} else {
		return "testdata"
	}
}
package tessera

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/private/engine"
)

var (
	arbitraryHash                  = common.BytesToEncryptedPayloadHash([]byte("arbitrary"))
	arbitraryHash1                 = common.BytesToEncryptedPayloadHash([]byte("arbitrary1"))
	arbitraryNotFoundHash          = common.BytesToEncryptedPayloadHash([]byte("not found"))
	arbitraryHashNoPrivateMetadata = common.BytesToEncryptedPayloadHash([]byte("no private extra data"))
	arbitraryPrivatePayload        = []byte("arbitrary private payload")
	arbitraryFrom                  = "arbitraryFrom"
	arbitraryTo                    = []string{"arbitraryTo1", "arbitraryTo2"}
	arbitraryExtra                 = &engine.ExtraMetadata{
		ACHashes:     Must(common.Base64sToEncryptedPayloadHashes([]string{arbitraryHash.ToBase64()})).(common.EncryptedPayloadHashes),
		ACMerkleRoot: common.StringToHash("arbitrary root hash"),
	}

	testServer *httptest.Server
	testObject *tesseraPrivateTxManager

	sendRequestCaptor         = make(chan *capturedRequest)
	receiveRequestCaptor      = make(chan *capturedRequest)
	sendSignedTxRequestCaptor = make(chan *capturedRequest)
)

type capturedRequest struct {
	err     error
	request interface{}
	header  http.Header
}

func TestMain(m *testing.M) {
	setup()
	retCode := m.Run()
	teardown()
	os.Exit(retCode)
}

func Must(o interface{}, err error) interface{} {
	if err != nil {
		panic(fmt.Sprintf("%s", err))
	}
	return o
}

func setup() {
	mux := http.NewServeMux()
	mux.HandleFunc("/send", MockSendAPIHandlerFunc)
	mux.HandleFunc("/transaction/", MockReceiveAPIHandlerFunc)
	mux.HandleFunc("/sendsignedtx", MockSendSignedTxAPIHandlerFunc)

	testServer = httptest.NewServer(mux)

	testObject = New(&engine.Client{
		HttpClient: &http.Client{},
		BaseURL:    testServer.URL,
	})
}

func MockSendAPIHandlerFunc(response http.ResponseWriter, request *http.Request) {
	actualRequest := new(sendRequest)
	if err := json.NewDecoder(request.Body).Decode(actualRequest); err != nil {
		go func(o *capturedRequest) { sendRequestCaptor <- o }(&capturedRequest{err: err})
	} else {
		go func(o *capturedRequest) { sendRequestCaptor <- o }(&capturedRequest{request: actualRequest, header: request.Header})
		data, _ := json.Marshal(&sendResponse{
			Key: arbitraryHash.ToBase64(),
		})
		response.Write(data)
	}
}

func MockReceiveAPIHandlerFunc(response http.ResponseWriter, request *http.Request) {
	path := string([]byte(request.RequestURI)[:strings.LastIndex(request.RequestURI, "?")])
	actualRequest, err := url.PathUnescape(strings.TrimPrefix(path, "/transaction/"))
	if err != nil {
		go func(o *capturedRequest) { sendRequestCaptor <- o }(&capturedRequest{err: err})
	} else {
		go func(o *capturedRequest) {
			receiveRequestCaptor <- o
		}(&capturedRequest{request: actualRequest, header: request.Header})
		if actualRequest == arbitraryNotFoundHash.ToBase64() {
			response.WriteHeader(http.StatusNotFound)
		} else {
			var data []byte
			if actualRequest == arbitraryHashNoPrivateMetadata.ToBase64() {
				data, _ = json.Marshal(&receiveResponse{
					Payload: arbitraryPrivatePayload,
				})
			} else {
				data, _ = json.Marshal(&receiveResponse{
					Payload:                      arbitraryPrivatePayload,
					ExecHash:                     base64.StdEncoding.EncodeToString(arbitraryExtra.ACMerkleRoot.Bytes()),
					AffectedContractTransactions: arbitraryExtra.ACHashes.ToBase64s(),
				})
			}
			response.Write(data)
		}
	}
}

func MockSendSignedTxAPIHandlerFunc(response http.ResponseWriter, request *http.Request) {
	actualRequest := new(sendSignedTxRequest)
	if err := json.NewDecoder(request.Body).Decode(actualRequest); err != nil {
		go func(o *capturedRequest) { sendSignedTxRequestCaptor <- o }(&capturedRequest{err: err})
	} else {
		go func(o *capturedRequest) { sendSignedTxRequestCaptor <- o }(&capturedRequest{request: actualRequest, header: request.Header})
		data, _ := json.Marshal(&sendSignedTxResponse{
			Key: arbitraryHash.ToBase64(),
		})
		response.Write(data)
	}
}

func teardown() {
	testServer.Close()
}

func verifyRequetHeader(h http.Header, t *testing.T) {
	if h.Get("Content-type") != "application/json" {
		t.Errorf("expected Content-type header is application/json")
	}

	if h.Get("Accept") != "application/json" {
		t.Errorf("expected Accept header is application/json")
	}
}

func TestSend_whenTypical(t *testing.T) {
	actualHash, err := testObject.Send(arbitraryPrivatePayload, arbitraryFrom, arbitraryTo, arbitraryExtra)
	if err != nil {
		t.Fatalf("%s", err)
	}
	capturedRequest := <-sendRequestCaptor

	if capturedRequest.err != nil {
		t.Fatalf("%s", capturedRequest.err)
	}

	verifyRequetHeader(capturedRequest.header, t)

	actualRequest := capturedRequest.request.(*sendRequest)

	if string(actualRequest.Payload) != string(arbitraryPrivatePayload) {
		t.Errorf("Payload: expected %s but got %s", arbitraryPrivatePayload, actualRequest.Payload)
	}

	if actualRequest.From != arbitraryFrom {
		t.Errorf("From: expected %s but got %s", arbitraryFrom, actualRequest.From)
	}

	if !reflect.DeepEqual(actualRequest.To, arbitraryTo) {
		t.Errorf("To: expected %v but got %v", arbitraryTo, actualRequest.To)
	}

	expectedACHashes := arbitraryExtra.ACHashes.ToBase64s()
	if !reflect.DeepEqual(actualRequest.AffectedContractTransactions, expectedACHashes) {
		t.Errorf("AffectedContractTransactions: expected %v but got %v", expectedACHashes, actualRequest.AffectedContractTransactions)
	}

	expectedMerkleRoot := base64.StdEncoding.EncodeToString(arbitraryExtra.ACMerkleRoot.Bytes())
	if actualRequest.ExecHash != expectedMerkleRoot {
		t.Errorf("ExecHash: expected %s but got %s", actualRequest.ExecHash, expectedMerkleRoot)
	}

	if actualHash.Hex() != arbitraryHash.Hex() {
		t.Errorf("EncryptedPayloadHash: expected %s but got %s", arbitraryHash.Hex(), actualHash.Hex())
	}
}

func TestReceive_whenTypical(t *testing.T) {
	_, actualExtra, err := testObject.Receive(arbitraryHash1)
	if err != nil {
		t.Fatalf("%s", err)
	}
	capturedRequest := <-receiveRequestCaptor

	if capturedRequest.err != nil {
		t.Fatalf("%s", capturedRequest.err)
	}

	verifyRequetHeader(capturedRequest.header, t)

	actualRequest := capturedRequest.request.(string)

	if actualRequest != arbitraryHash1.ToBase64() {
		t.Errorf("Key: expected %s but got %s", arbitraryHash1.ToBase64(), actualRequest)
	}

	if !reflect.DeepEqual(actualExtra.ACHashes, arbitraryExtra.ACHashes) {
		t.Errorf("ACHashes: expected %v but got %v", arbitraryExtra.ACHashes, actualExtra.ACHashes)
	}

	if actualExtra.ACMerkleRoot.Hex() != arbitraryExtra.ACMerkleRoot.Hex() {
		t.Errorf("MerkelRoot: expected %s but got %s", arbitraryExtra.ACMerkleRoot.Hex(), actualExtra.ACMerkleRoot.Hex())
	}
}

func TestReceive_whenPayloadNotFound(t *testing.T) {
	data, _, err := testObject.Receive(arbitraryNotFoundHash)
	if err != nil {
		t.Fatalf("%s", err)
	}
	capturedRequest := <-receiveRequestCaptor

	if capturedRequest.err != nil {
		t.Fatalf("%s", capturedRequest.err)
	}

	verifyRequetHeader(capturedRequest.header, t)

	actualRequest := capturedRequest.request.(string)

	if actualRequest != arbitraryNotFoundHash.ToBase64() {
		t.Errorf("Key: expected %s but got %s", arbitraryNotFoundHash.ToBase64(), actualRequest)
	}

	if data != nil {
		t.Errorf("Payload: expected nil but got %v", data)
	}
}

func TestReceive_whenHavingPayloadButNoPrivateExtraMetadata(t *testing.T) {
	_, actualExtra, err := testObject.Receive(arbitraryHashNoPrivateMetadata)
	if err != nil {
		t.Fatalf("%s", err)
	}
	capturedRequest := <-receiveRequestCaptor

	if capturedRequest.err != nil {
		t.Fatalf("%s", capturedRequest.err)
	}

	verifyRequetHeader(capturedRequest.header, t)

	actualRequest := capturedRequest.request.(string)

	if actualRequest != arbitraryHashNoPrivateMetadata.ToBase64() {
		t.Errorf("Key: expected %s but got %s", arbitraryHashNoPrivateMetadata.ToBase64(), actualRequest)
	}

	if actualExtra.ACHashes == nil || len(actualExtra.ACHashes) > 0 {
		t.Errorf("ACHashes: expected empty and not nil but got %v", actualExtra.ACHashes)
	}

	if !common.EmptyHash(actualExtra.ACMerkleRoot) {
		t.Errorf("MerkelRoot: expected empty hash but got %s", actualExtra.ACMerkleRoot.Hex())
	}
}

func TestSendSignedTx_whenTypical(t *testing.T) {
	_, err := testObject.SendSignedTx(arbitraryHash, arbitraryTo, arbitraryExtra)
	if err != nil {
		t.Fatalf("%s", err)
	}
	capturedRequest := <-sendSignedTxRequestCaptor

	if capturedRequest.err != nil {
		t.Fatalf("%s", capturedRequest.err)
	}

	verifyRequetHeader(capturedRequest.header, t)

	actualRequest := capturedRequest.request.(*sendSignedTxRequest)

	if !reflect.DeepEqual(actualRequest.To, arbitraryTo) {
		t.Errorf("To: expected %v but got %v", arbitraryTo, actualRequest.To)
	}

	expectedACHashes := arbitraryExtra.ACHashes.ToBase64s()
	if !reflect.DeepEqual(actualRequest.AffectedContractTransactions, expectedACHashes) {
		t.Errorf("AffectedContractTransactions: expected %v but got %v", expectedACHashes, actualRequest.AffectedContractTransactions)
	}

	expectedMerkleRoot := base64.StdEncoding.EncodeToString(arbitraryExtra.ACMerkleRoot.Bytes())
	if actualRequest.ExecHash != expectedMerkleRoot {
		t.Errorf("ExecHash: expected %s but got %s", actualRequest.ExecHash, expectedMerkleRoot)
	}
}

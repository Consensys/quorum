package tessera

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/private/engine"
	testifyassert "github.com/stretchr/testify/assert"
)

var (
	emptyHash                      = common.EncryptedPayloadHash{}
	arbitraryHash                  = common.BytesToEncryptedPayloadHash([]byte("arbitrary"))
	arbitraryHash1                 = common.BytesToEncryptedPayloadHash([]byte("arbitrary1"))
	arbitraryNotFoundHash          = common.BytesToEncryptedPayloadHash([]byte("not found"))
	arbitraryHashNoPrivateMetadata = common.BytesToEncryptedPayloadHash([]byte("no private extra data"))
	arbitraryPrivatePayload        = []byte("arbitrary private payload")
	arbitraryFrom                  = "arbitraryFrom"
	arbitraryTo                    = []string{"arbitraryTo1", "arbitraryTo2"}
	arbitraryPrivacyFlag           = engine.PrivacyFlagPartyProtection
	arbitraryExtra                 = &engine.ExtraMetadata{
		ACHashes:     Must(common.Base64sToEncryptedPayloadHashes([]string{arbitraryHash.ToBase64()})).(common.EncryptedPayloadHashes),
		ACMerkleRoot: common.StringToHash("arbitrary root hash"),
		PrivacyFlag:  arbitraryPrivacyFlag,
	}

	testServer *httptest.Server
	testObject *tesseraPrivateTxManager

	sendRequestCaptor                    = make(chan *capturedRequest)
	receiveRequestCaptor                 = make(chan *capturedRequest)
	sendSignedTxRequestCaptor            = make(chan *capturedRequest)
	sendSignedTxOctetStreamRequestCaptor = make(chan *capturedRequest)
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
	}, []byte("2.0.0"))
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
					ExecHash:                     arbitraryExtra.ACMerkleRoot.ToBase64(),
					AffectedContractTransactions: arbitraryExtra.ACHashes.ToBase64s(),
					PrivacyFlag:                  arbitraryPrivacyFlag,
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

func MockSendSignedTxOctetStreamAPIHandlerFunc(response http.ResponseWriter, request *http.Request) {
	actualRequest := new(sendSignedTxRequest)
	reqHash, err := ioutil.ReadAll(request.Body)
	if err != nil {
		go func(o *capturedRequest) { sendSignedTxOctetStreamRequestCaptor <- o }(&capturedRequest{err: err})
		return
	}
	actualRequest.Hash = reqHash
	actualRequest.To = strings.Split(request.Header["C11n-To"][0], ",")

	go func(o *capturedRequest) { sendSignedTxOctetStreamRequestCaptor <- o }(&capturedRequest{request: actualRequest, header: request.Header})
	response.Write([]byte(common.BytesToEncryptedPayloadHash(reqHash).ToBase64()))
}

func teardown() {
	testServer.Close()
}

func verifyRequestHeader(h http.Header, t *testing.T) {
	if h.Get("Content-type") != "application/json" {
		t.Errorf("expected Content-type header is application/json")
	}

	if h.Get("Accept") != "application/json" {
		t.Errorf("expected Accept header is application/json")
	}
}

func TestSend_whenTypical(t *testing.T) {
	assert := testifyassert.New(t)

	actualHash, err := testObject.Send(arbitraryPrivatePayload, arbitraryFrom, arbitraryTo, arbitraryExtra)
	if err != nil {
		t.Fatalf("%s", err)
	}
	capturedRequest := <-sendRequestCaptor

	if capturedRequest.err != nil {
		t.Fatalf("%s", capturedRequest.err)
	}

	verifyRequestHeader(capturedRequest.header, t)

	actualRequest := capturedRequest.request.(*sendRequest)

	assert.Equal(arbitraryPrivatePayload, actualRequest.Payload, "request.payload")
	assert.Equal(arbitraryFrom, actualRequest.From, "request.from")
	assert.Equal(arbitraryTo, actualRequest.To, "request.to")
	assert.Equal(arbitraryPrivacyFlag, actualRequest.PrivacyFlag, "request.privacyFlag")
	assert.Equal(arbitraryExtra.ACHashes.ToBase64s(), actualRequest.AffectedContractTransactions, "request.affectedContractTransactions")
	assert.Equal(arbitraryExtra.ACMerkleRoot.ToBase64(), actualRequest.ExecHash, "request.execHash")
	assert.Equal(arbitraryHash, actualHash, "returned hash")
}

func TestSend_whenTesseraVersionDoesNotSupportPrivacyEnhancements(t *testing.T) {
	assert := testifyassert.New(t)

	testObjectNoPE := New(&engine.Client{
		HttpClient: &http.Client{},
		BaseURL:    testServer.URL,
	}, []byte("0.10-SNAPSHOT"))

	assert.False(testObjectNoPE.HasFeature(engine.PrivacyEnhancements), "the supplied version does not support privacy enhancements")

	// trying to send a party protection transaction
	_, err := testObjectNoPE.Send(arbitraryPrivatePayload, arbitraryFrom, arbitraryTo, arbitraryExtra)
	if err != engine.ErrPrivateTxManagerDoesNotSupportPrivacyEnhancements {
		t.Fatal("Expecting send to raise ErrPrivateTxManagerDoesNotSupportPrivacyEnhancements")
	}
}

func TestSendRaw_whenTesseraVersionDoesNotSupportPrivacyEnhancements(t *testing.T) {
	assert := testifyassert.New(t)

	mux := http.NewServeMux()
	mux.HandleFunc("/send", MockSendAPIHandlerFunc)
	mux.HandleFunc("/transaction/", MockReceiveAPIHandlerFunc)
	mux.HandleFunc("/sendsignedtx", MockSendSignedTxOctetStreamAPIHandlerFunc)

	testServerNoPE := httptest.NewServer(mux)
	defer testServerNoPE.Close()

	testObjectNoPE := New(&engine.Client{
		HttpClient: &http.Client{},
		BaseURL:    testServerNoPE.URL,
	}, []byte("0.10-SNAPSHOT"))

	assert.False(testObjectNoPE.HasFeature(engine.PrivacyEnhancements), "the supplied version does not support privacy enhancements")

	// trying to send a party protection transaction
	_, err := testObjectNoPE.SendSignedTx(arbitraryHash, arbitraryTo, arbitraryExtra)
	if err != engine.ErrPrivateTxManagerDoesNotSupportPrivacyEnhancements {
		t.Fatal("Expecting send to raise ErrPrivateTxManagerDoesNotSupportPrivacyEnhancements")
	}

	// send a standard private transaction and check that the old version of the /sendsignedtx is used (using octetstream content type)

	// caching incomplete item
	_, _, err = testObjectNoPE.ReceiveRaw(arbitraryHashNoPrivateMetadata)
	if err != nil {
		t.Fatalf("%s", err)
	}
	<-receiveRequestCaptor

	// caching complete item
	_, err = testObjectNoPE.SendSignedTx(arbitraryHashNoPrivateMetadata, arbitraryTo, &engine.ExtraMetadata{
		PrivacyFlag: engine.PrivacyFlagStandardPrivate})
	if err != nil {
		t.Fatalf("%s", err)
	}
	req := <-sendSignedTxOctetStreamRequestCaptor
	assert.Equal("application/octet-stream", req.header["Content-Type"][0])

	_, actualExtra, err := testObjectNoPE.Receive(arbitraryHashNoPrivateMetadata)
	if err != nil {
		t.Fatalf("%s", err)
	}
	assert.Equal(engine.PrivacyFlagStandardPrivate, actualExtra.PrivacyFlag, "cached privacy flag")

}

func TestReceive_whenTypical(t *testing.T) {
	assert := testifyassert.New(t)

	_, actualExtra, err := testObject.Receive(arbitraryHash1)
	if err != nil {
		t.Fatalf("%s", err)
	}
	capturedRequest := <-receiveRequestCaptor

	if capturedRequest.err != nil {
		t.Fatalf("%s", capturedRequest.err)
	}

	verifyRequestHeader(capturedRequest.header, t)

	actualRequest := capturedRequest.request.(string)

	assert.Equal(arbitraryHash1.ToBase64(), actualRequest, "requested hash")
	assert.Equal(arbitraryExtra.ACHashes, actualExtra.ACHashes, "returned affected contract transaction hashes")
	assert.Equal(arbitraryExtra.ACMerkleRoot, actualExtra.ACMerkleRoot, "returned merkle root")
	assert.Equal(arbitraryExtra.PrivacyFlag, actualExtra.PrivacyFlag, "returned privacy flag")
}

func TestReceive_whenPayloadNotFound(t *testing.T) {
	assert := testifyassert.New(t)

	data, _, err := testObject.Receive(arbitraryNotFoundHash)
	if err != nil {
		t.Fatalf("%s", err)
	}
	capturedRequest := <-receiveRequestCaptor

	if capturedRequest.err != nil {
		t.Fatalf("%s", capturedRequest.err)
	}

	verifyRequestHeader(capturedRequest.header, t)

	actualRequest := capturedRequest.request.(string)

	assert.Equal(arbitraryNotFoundHash.ToBase64(), actualRequest, "requested hash")
	assert.Nil(data, "returned payload when not found")
}

func TestReceive_whenEncryptedPayloadHashIsEmpty(t *testing.T) {
	assert := testifyassert.New(t)

	data, _, err := testObject.Receive(emptyHash)
	if err != nil {
		t.Fatalf("%s", err)
	}
	assert.Empty(receiveRequestCaptor, "no request is actually sent")
	assert.Nil(data, "returned payload when not found")
}

func TestReceive_whenHavingPayloadButNoPrivateExtraMetadata(t *testing.T) {
	assert := testifyassert.New(t)

	_, actualExtra, err := testObject.Receive(arbitraryHashNoPrivateMetadata)
	if err != nil {
		t.Fatalf("%s", err)
	}
	capturedRequest := <-receiveRequestCaptor

	if capturedRequest.err != nil {
		t.Fatalf("%s", capturedRequest.err)
	}

	verifyRequestHeader(capturedRequest.header, t)

	actualRequest := capturedRequest.request.(string)

	assert.Equal(arbitraryHashNoPrivateMetadata.ToBase64(), actualRequest, "requested hash")
	assert.Empty(actualExtra.ACHashes, "returned affected contract transaction hashes")
	assert.True(common.EmptyHash(actualExtra.ACMerkleRoot), "returned merkle root")
}

func TestSendSignedTx_whenTypical(t *testing.T) {
	assert := testifyassert.New(t)

	_, err := testObject.SendSignedTx(arbitraryHash, arbitraryTo, arbitraryExtra)
	if err != nil {
		t.Fatalf("%s", err)
	}
	capturedRequest := <-sendSignedTxRequestCaptor

	if capturedRequest.err != nil {
		t.Fatalf("%s", capturedRequest.err)
	}

	verifyRequestHeader(capturedRequest.header, t)

	actualRequest := capturedRequest.request.(*sendSignedTxRequest)

	assert.Equal(arbitraryTo, actualRequest.To, "request.to")
	assert.Equal(arbitraryExtra.ACHashes.ToBase64s(), actualRequest.AffectedContractTransactions, "request.affectedContractTransactions")
	assert.Equal(arbitraryExtra.ACMerkleRoot.ToBase64(), actualRequest.ExecHash, "request.execHash")
}

func TestReceive_whenCachingRawPayload(t *testing.T) {
	assert := testifyassert.New(t)

	// caching incomplete item
	_, _, err := testObject.ReceiveRaw(arbitraryHashNoPrivateMetadata)
	if err != nil {
		t.Fatalf("%s", err)
	}
	<-receiveRequestCaptor

	// caching complete item
	_, err = testObject.SendSignedTx(arbitraryHashNoPrivateMetadata, arbitraryTo, arbitraryExtra)
	if err != nil {
		t.Fatalf("%s", err)
	}
	<-sendSignedTxRequestCaptor

	_, actualExtra, err := testObject.Receive(arbitraryHashNoPrivateMetadata)
	if err != nil {
		t.Fatalf("%s", err)
	}

	assert.Equal(arbitraryExtra.ACHashes, actualExtra.ACHashes, "cached affected contract transaction hashes")
	assert.Equal(arbitraryExtra.ACMerkleRoot, actualExtra.ACMerkleRoot, "cached merkle root")
	assert.Equal(arbitraryExtra.PrivacyFlag, actualExtra.PrivacyFlag, "cached privacy flag")
}

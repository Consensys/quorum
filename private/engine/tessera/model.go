package tessera

// request object for /send API
type sendRequest struct {
	Payload []byte `json:"payload"`

	// base64-encoded
	From string `json:"from,omitempty"`

	To []string `json:"to"`

	// Transactions' encrypted payload hashes for affected contracts
	AffectedContractTransactions []string `json:"affectedContractTransactions"`

	// Merkle root for affected contracts
	ExecHash string `json:"execHash"`

	PrivateStateValidation bool `json:"privateStateValidation"`
}

// response object for /send API
type sendResponse struct {
	// Base64-encoded
	Key string `json:"key"`
}

type receiveRequest struct {
	// Base64-encoded
	Key string `json:"key"`

	To string `json:"to"`
}

type receiveResponse struct {
	Payload []byte `json:"payload"`

	// Transactions' encrypted payload hashes for affected contracts
	AffectedContractTransactions []string `json:"affectedContractTransactions"`

	// Merkle root for affected contracts
	ExecHash string `json:"execHash"`

	PrivateStateValidation bool `json:"privateStateValidation"`
}

type sendSignedTxRequest struct {
	Hash []byte   `json:"hash"`
	To   []string `json:"to"`
	// Transactions' encrypted payload hashes for affected contracts
	AffectedContractTransactions []string `json:"affectedContractTransactions"`
	// Merkle root for affected contracts
	ExecHash               string `json:"execHash"`
	PrivateStateValidation bool   `json:"privateStateValidation"`
}

type sendSignedTxResponse struct {
	*sendResponse
}

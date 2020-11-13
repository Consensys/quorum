//go:generate mockgen -source contract_indexer.go -destination mock_contract_indexer.go -package multitenancy

package multitenancy

import (
	"errors"
	"io"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
)

var contractIndexPrefix = []byte("contractIndex")

type ContractIndexWriter interface {
	WriteIndex(contractAddress common.Address, contractParties *ContractParties) error
}

type ContractIndexReader interface {
	ReadIndex(contractAddress common.Address) (*ContractParties, error)
}

type ContractIndex struct {
	db ethdb.Database
}

func NewContractIndex(db ethdb.Database) *ContractIndex {
	return &ContractIndex{
		db: db,
	}
}

type ContractParties struct {
	// EOA address that was used to sign the contract creation transaction
	CreatorAddress common.Address
	// List of Tessera Public Keys
	Parties []string
}

type contractPartiesRLP struct {
	CreatorAddress common.Address
	Parties        []string
}

func (ci ContractIndex) WriteIndex(contractAddress common.Address, contractParties *ContractParties) error {
	data, err := rlp.EncodeToBytes(contractParties)
	if err != nil {
		return err
	}
	if err = ci.db.Put(append(contractIndexPrefix, contractAddress.Bytes()...), data); err != nil {
		log.Error("Error writing contract index", "Contract Address", contractAddress)
		return err
	}
	return nil
}

func (ci ContractIndex) ReadIndex(contractAddress common.Address) (*ContractParties, error) {
	var ca ContractParties
	contractPartiesBytes, err := ci.db.Get(append(contractIndexPrefix, contractAddress.Bytes()...))
	if err != nil {
		log.Error("Error retrieving Contract Addresses from index", "Contract Address", contractAddress)
		return nil, err
	}
	if len(contractPartiesBytes) == 0 {
		log.Error("Empty response returned", "Contract Address", contractAddress)
		return nil, errors.New("empty response querying contract index")
	}
	if err := rlp.DecodeBytes(contractPartiesBytes, &ca); err != nil {
		return nil, err
	}
	log.Trace("Contract index found", "addess", contractAddress, "creatorEOA", ca.CreatorAddress.Hex(), "parties", ca.Parties)
	return &ca, nil
}

func (cp *ContractParties) DecodeRLP(s *rlp.Stream) error {
	var partiesRLP contractPartiesRLP
	if err := s.Decode(&partiesRLP); err != nil {
		return err
	}
	cp.CreatorAddress, cp.Parties = partiesRLP.CreatorAddress, partiesRLP.Parties
	return nil
}

func (cp *ContractParties) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, contractPartiesRLP{
		CreatorAddress: cp.CreatorAddress,
		Parties:        cp.Parties,
	})
}

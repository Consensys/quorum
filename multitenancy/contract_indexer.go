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
	WriteIndex(contractAddress common.Address, contractParties *ContractIndexItem) error
}

type ContractIndexReader interface {
	ReadIndex(contractAddress common.Address) (*ContractIndexItem, error)
}

// write index direct to eth DB
type ContractIndex struct {
	db ethdb.Database
}

func NewContractIndex(db ethdb.Database) *ContractIndex {
	return &ContractIndex{
		db: db,
	}
}

type ContractIndexItem struct {
	// EOA address that was used to sign the contract creation transaction
	CreatorAddress common.Address
	IsPrivate      bool
	// List of Tessera Public Keys
	Parties []string
}

type contractIndexItemRLP struct {
	CreatorAddress common.Address
	IsPrivate      bool
	Parties        []string
}

func (ci ContractIndex) WriteIndex(contractAddress common.Address, indexItem *ContractIndexItem) error {
	data, err := rlp.EncodeToBytes(indexItem)
	if err != nil {
		return err
	}
	if err = ci.db.Put(append(contractIndexPrefix, contractAddress.Bytes()...), data); err != nil {
		log.Error("Error writing contract index", "Contract Address", contractAddress)
		return err
	}
	return nil
}

func (ci ContractIndex) ReadIndex(contractAddress common.Address) (*ContractIndexItem, error) {
	var ca ContractIndexItem
	contractIndexItemBytes, err := ci.db.Get(append(contractIndexPrefix, contractAddress.Bytes()...))
	if err != nil {
		log.Error("Error retrieving Contract Addresses from index", "Contract Address", contractAddress)
		return nil, err
	}
	if len(contractIndexItemBytes) == 0 {
		log.Error("Empty response returned", "Contract Address", contractAddress)
		return nil, errors.New("empty response querying contract index")
	}
	if err := rlp.DecodeBytes(contractIndexItemBytes, &ca); err != nil {
		return nil, err
	}
	log.Trace("Contract index found", "addess", contractAddress, "creatorEOA", ca.CreatorAddress.Hex(), "parties", ca.Parties)
	return &ca, nil
}

func (cii *ContractIndexItem) DecodeRLP(s *rlp.Stream) error {
	var indexItemRLP contractIndexItemRLP
	if err := s.Decode(&indexItemRLP); err != nil {
		return err
	}
	cii.CreatorAddress, cii.IsPrivate, cii.Parties = indexItemRLP.CreatorAddress, indexItemRLP.IsPrivate, indexItemRLP.Parties
	return nil
}

func (cii *ContractIndexItem) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, contractIndexItemRLP{
		CreatorAddress: cii.CreatorAddress,
		IsPrivate:      cii.IsPrivate,
		Parties:        cii.Parties,
	})
}

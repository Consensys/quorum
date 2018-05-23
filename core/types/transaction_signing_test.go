// Copyright 2016 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package types

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
)

func TestEIP155Signing(t *testing.T) {
	key, _ := crypto.GenerateKey()
	addr := crypto.PubkeyToAddress(key.PublicKey)

	signer := NewEIP155Signer(big.NewInt(18))
	tx, err := SignTx(NewTransaction(0, addr, new(big.Int), new(big.Int), new(big.Int), nil), signer, key)
	if err != nil {
		t.Fatal(err)
	}

	from, err := Sender(signer, tx)
	if err != nil {
		t.Fatal(err)
	}
	if from != addr {
		t.Errorf("exected from and address to be equal. Got %x want %x", from, addr)
	}
}

func TestEIP155ChainId(t *testing.T) {
	key, _ := crypto.GenerateKey()
	addr := crypto.PubkeyToAddress(key.PublicKey)

	signer := NewEIP155Signer(big.NewInt(18))
	tx, err := SignTx(NewTransaction(0, addr, new(big.Int), new(big.Int), new(big.Int), nil), signer, key)
	if err != nil {
		t.Fatal(err)
	}
	if !tx.Protected() {
		t.Fatal("expected tx to be protected")
	}

	if tx.ChainId().Cmp(signer.chainId) != 0 {
		t.Error("expected chainId to be", signer.chainId, "got", tx.ChainId())
	}

	tx = NewTransaction(0, addr, new(big.Int), new(big.Int), new(big.Int), nil)
	tx, err = SignTx(tx, HomesteadSigner{}, key)
	if err != nil {
		t.Fatal(err)
	}

	if tx.Protected() {
		t.Error("didn't expect tx to be protected")
	}

	if tx.ChainId().Sign() != 0 {
		t.Error("expected chain id to be 0 got", tx.ChainId())
	}
}

func TestEIP155SigningVitalik(t *testing.T) {
	// Test vectors come from http://vitalik.ca/files/eip155_testvec.txt is not available
	// Since we cannot use chainId of 1, new raw transaction and address paris are create as below rule.
	// gasPrice: 20 * 10**9, gas: 21000, data: "", to: 0x3535353535353535353535353535353535353535, value: 10**18, nonce: 9, chainid: 10,
	// privateKeys used are 0x4646464646464646464646464646464646464646464646464646464646464640 to 0x4646464646464646464646464646464646464646464646464646464646464649
	for i, test := range []struct {
		txRlp, addr string
	}{
		{"f86c098504a817c800825208943535353535353535353535353535353535353535880de0b6b3a76400008037a024692193af7d2c93f2b105d8d79f91106c4e44bed054b8056e0cb13d1f48c598a01126dbf7c80423fbd2c8b062812d15ad3144caea7f599eed86a17d17da5e9d5d", "0x5fAA510EB3f838aC398a293b5714ad279f9cECF4"},
		{"f86c098504a817c800825208943535353535353535353535353535353535353535880de0b6b3a76400008038a0262959a9f060bf3e205e28e046a3331036d7824487c21f9c9a234da7bf8d947da078c0623a89396c1e49eb24c53e38c2cb44d39e82b66d3e5c3216f55883eaa18f", "0xDe6AB723c23bba740410129F9edc952fD6fbced4"},
		{"f86c098504a817c800825208943535353535353535353535353535353535353535880de0b6b3a76400008038a0ed95916ac3f361c77d91504f7fd9e582926b96f786669356efc5027c96ea59c1a04f8c52b39d3303df7f577b2087f7a550de117f07c3d7eac8d24b2e58fd2cb722", "0x4107c605Be9cFc6C8c625bCB3f762E963472457E"},
		{"f86c098504a817c800825208943535353535353535353535353535353535353535880de0b6b3a76400008038a0e15270bbd1f981aa17b41dba76c6e244439eab3766ac535cfd914fbe7bd45e56a05cd6920d3ad53c16fbe659450b9ff1e52bb3983998d769db5b10b6d42320e0c1", "0x68B48A376F3158362443Ee0DF16f5C30b4aCE9B7"},
		{"f86c098504a817c800825208943535353535353535353535353535353535353535880de0b6b3a76400008038a04007e77c19387f9d59d3d9e64f2f1c961d89e76a4478e000bbba436b03605d58a00edc33fe0f5a8598cc70f8ac7c24d9d22580f51a1fe6e52872819c35df6cc955", "0x14Fe11894410453c01485a7e337c3F63fC512d14"},
		{"f86c098504a817c800825208943535353535353535353535353535353535353535880de0b6b3a76400008037a0a45e844857928b7309f5607180680887fd98940aadd4cc21126ecbe55c4f363fa07c67cab2c7a53b09fdea0f07bce9a7cacb5edfacf19539472e17fd286af2dcf4", "0xAa9b8181391561bCe5199c5b1762aa26832EE548"},
		{"f86c098504a817c800825208943535353535353535353535353535353535353535880de0b6b3a76400008037a0745aa5ba54d4b1cf6d264556c75abce97bfd2d07b9864a298ebcb6dc5268667fa04d6224e1036d4820ac36dfe768c3230b0450f657a2b8b459689da069f2c1a383", "0x9d8A62f656a8d1615C1294fd71e9CFb3E4855A4F"},
		{"f86c098504a817c800825208943535353535353535353535353535353535353535880de0b6b3a76400008038a007d494d57c46ffb9a7d560497751e4f75c10083b1f703b57f155193c34837a07a006f58ebd3d8df59b8d2d8ea484a861592b5c18dc9652075a080758787d2db7ae", "0x5a17650BE84F28Ed583e93E6ed0C99b1D1FC1b34"},
		{"f86c098504a817c800825208943535353535353535353535353535353535353535880de0b6b3a76400008038a0878b3ea61b8135c5f6c6907ea6669fc0f9eb9c90e4b648923498627c4cb8b5b6a02e8de6963a8649f16cadfebf23def96cafd89072ad19ed04c29e5f880e5ede98", "0x0EfbD0bEC0dA8dCc0Ad442A7D337E9CDc2dd6a54"},
		{"f86c098504a817c800825208943535353535353535353535353535353535353535880de0b6b3a76400008037a0fa0fd46b84c9d488558da70103960f1c43400b4d92cc9ca37737bd508635acefa01b2b01c2667ce7b48cc3fcfdba75181d097134ae6e57af4b691ba6d5c118d0c1", "0x0E8E18e1A11E6196f6B82426196027d042Fd6812"},
	} {
		signer := NewEIP155Signer(big.NewInt(10))

		var tx *Transaction
		err := rlp.DecodeBytes(common.Hex2Bytes(test.txRlp), &tx)
		if err != nil {
			t.Errorf("%d: %v", i, err)
			continue
		}

		from, err := Sender(signer, tx)
		if err != nil {
			t.Errorf("%d: %v", i, err)
			continue
		}

		addr := common.HexToAddress(test.addr)
		if from != addr {
			t.Errorf("%d: expected %x got %x", i, addr, from)
		}

	}
}

func TestChainId(t *testing.T) {
	key, _ := defaultTestKey()

	tx := NewTransaction(0, common.Address{}, new(big.Int), new(big.Int), new(big.Int), nil)

	var err error
	tx, err = SignTx(tx, NewEIP155Signer(big.NewInt(10)), key)
	if err != nil {
		t.Fatal(err)
	}

	_, err = Sender(NewEIP155Signer(big.NewInt(11)), tx)
	if err != ErrInvalidChainId {
		t.Error("expected error:", ErrInvalidChainId)
	}

	_, err = Sender(NewEIP155Signer(big.NewInt(10)), tx)
	if err != nil {
		t.Error("expected no error")
	}
}

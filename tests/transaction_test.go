// Copyright 2015 The go-ethereum Authors
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

package tests

import (
	"testing"

	"github.com/ethereum/go-ethereum/params"
)

func TestTransaction(t *testing.T) {
	t.Parallel()

	txt := new(testMatcher)
	// These can't be parsed, invalid hex in RLP
	txt.skipLoad("^ttWrongRLP/.*")
	// We don't allow more than uint64 in gas amount
	// This is a pseudo-consensus vulnerability, but not in practice
	// because of the gas limit
	txt.skipLoad("^ttGasLimit/TransactionWithGasLimitxPriceOverflow.json")

	//Quorum - skip the tests below as they have V=37/38 for transactions being tested and it is causing quorum to use quorum private signer for public transactions
	txt.skipLoad("^ttGasLimit/TransactionWithHihghGasLimit63m1.json")
	txt.skipLoad("^ttVValue/V_equals38.json")
	txt.skipLoad("^ttVValue/V_equals37.json")
	txt.skipLoad("^ttSignature/Vitalik_9.json")
	txt.skipLoad("^ttSignature/Vitalik_8.json")
	txt.skipLoad("^ttSignature/Vitalik_7.json")
	txt.skipLoad("^ttSignature/Vitalik_6.json")
	txt.skipLoad("^ttSignature/Vitalik_5.json")
	txt.skipLoad("^ttSignature/Vitalik_4.json")
	txt.skipLoad("^ttSignature/Vitalik_3.json")
	txt.skipLoad("^ttSignature/Vitalik_2.json")
	txt.skipLoad("^ttSignature/Vitalik_11.json")
	txt.skipLoad("^ttSignature/Vitalik_10.json")
	txt.skipLoad("^ttSignature/Vitalik_1.json")

	// We _do_ allow more than uint64 in gas price, as opposed to the tests
	// This is also not a concern, as long as tx.Cost() uses big.Int for
	// calculating the final cozt
	txt.skipLoad(".*TransactionWithGasPriceOverflow.*")

	// The nonce is too large for uint64. Not a concern, it means geth won't
	// accept transactions at a certain point in the distant future
	txt.skipLoad("^ttNonce/TransactionWithHighNonce256.json")

	// The value is larger than uint64, which according to the test is invalid.
	// Geth accepts it, which is not a consensus issue since we use big.Int's
	// internally to calculate the cost
	txt.skipLoad("^ttValue/TransactionWithHighValueOverflow.json")
	txt.walk(t, transactionTestDir, func(t *testing.T, name string, test *TransactionTest) {
		cfg := params.MainnetChainConfig
		if err := txt.checkFailure(t, name, test.Run(cfg)); err != nil {
			t.Error(err)
		}
	})
}

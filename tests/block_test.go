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
	"math/big"
	"path/filepath"
	"testing"
)

func TestBcInvalidHeaderTests(t *testing.T) {
	err := RunBlockTest(big.NewInt(1000000), nil, nil, filepath.Join(blockTestDir, "bcInvalidHeaderTest.json"), BlockSkipTests)
	if err != nil {
		t.Fatal(err)
	}
}

func TestBcInvalidRLPTests(t *testing.T) {
	err := RunBlockTest(big.NewInt(1000000), nil, nil, filepath.Join(blockTestDir, "bcInvalidRLPTest.json"), BlockSkipTests)
	if err != nil {
		t.Fatal(err)
	}
}

func TestBcRPCAPITests(t *testing.T) {
	err := RunBlockTest(big.NewInt(1000000), nil, nil, filepath.Join(blockTestDir, "bcRPC_API_Test.json"), BlockSkipTests)
	if err != nil {
		t.Fatal(err)
	}
}

func TestBcForkBlockTests(t *testing.T) {
	err := RunBlockTest(big.NewInt(1000000), nil, nil, filepath.Join(blockTestDir, "bcForkBlockTest.json"), BlockSkipTests)
	if err != nil {
		t.Fatal(err)
	}
}

func TestBcTotalDifficulty(t *testing.T) {
	// skip because these will fail due to selfish mining fix
	t.Skip()

	err := RunBlockTest(big.NewInt(1000000), nil, nil, filepath.Join(blockTestDir, "bcTotalDifficultyTest.json"), BlockSkipTests)
	if err != nil {
		t.Fatal(err)
	}
}

// TODO: iterate over files once we got more than a few
func TestBcRandom(t *testing.T) {
	err := RunBlockTest(big.NewInt(1000000), nil, big.NewInt(10), filepath.Join(blockTestDir, "RandomTests/bl201507071825GO.json"), BlockSkipTests)
	if err != nil {
		t.Fatal(err)
	}
}

func TestBcMultiChain(t *testing.T) {
	// skip due to selfish mining
	t.Skip()

	err := RunBlockTest(big.NewInt(1000000), nil, big.NewInt(10), filepath.Join(blockTestDir, "bcMultiChainTest.json"), BlockSkipTests)
	if err != nil {
		t.Fatal(err)
	}
}

func TestHomesteadBcInvalidHeaderTests(t *testing.T) {
	err := RunBlockTest(big.NewInt(0), nil, nil, filepath.Join(blockTestDir, "Homestead", "bcInvalidHeaderTest.json"), BlockSkipTests)
	if err != nil {
		t.Fatal(err)
	}
}

func TestHomesteadBcRPCAPITests(t *testing.T) {
	err := RunBlockTest(big.NewInt(0), nil, nil, filepath.Join(blockTestDir, "Homestead", "bcRPC_API_Test.json"), BlockSkipTests)
	if err != nil {
		t.Fatal(err)
	}
}

func TestHomesteadBcSuicideIssue(t *testing.T) {
	err := RunBlockTest(big.NewInt(0), nil, nil, filepath.Join(blockTestDir, "Homestead", "bcSuicideIssue.json"), BlockSkipTests)
	if err != nil {
		t.Fatal(err)
	}
}

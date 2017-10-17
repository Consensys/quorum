// Copyright 2017 Zerocoin Electric Coin Company LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sha256

import (
	"encoding/hex"
	"testing"
)

// Tests sha256 compress function i.e. checksum with no padding.
// https://github.com/zcash/zcash/blob/70db019c6ae989acde0a0affd6a1f1c28ec9a3d2/src/test/sha256compress_tests.cpp

func TestCompress1(t *testing.T) {
	preimage := make([]byte, 64)
	h := NewCompress()
	h.Write(preimage)
	actual := hex.EncodeToString(h.Compress())
	expected := "da5698be17b9b46962335799779fbeca8ce5d491c0d26243bafef9ea1837a9d8"
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got: '%s'", expected, actual)
	}
}

func TestCompress2(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic as expected")
		}
	}()

	preimage := make([]byte, 63)
	h := NewCompress()
	h.Write(preimage)
	h.Compress()
}

func TestCompress3(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic as expected")
		}
	}()

	preimage := make([]byte, 65)
	h := NewCompress()
	h.Write(preimage)
	h.Compress()
}

func TestCompress4(t *testing.T) {
	buf := make([]byte, 1)
	h := NewCompress()
	for i := 0; i < 64; i++ {
		h.Write(buf)
	}
	actual := hex.EncodeToString(h.Compress())
	expected := "da5698be17b9b46962335799779fbeca8ce5d491c0d26243bafef9ea1837a9d8"
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got: '%s'", expected, actual)
	}
}

func TestCompress5(t *testing.T) {
	preimage := []byte{'a', 'b', 'c', 'd',
		'a', 'b', 'c', 'd',
		'a', 'b', 'c', 'd',
		'a', 'b', 'c', 'd',
		'a', 'b', 'c', 'd',
		'a', 'b', 'c', 'd',
		'a', 'b', 'c', 'd',
		'a', 'b', 'c', 'd',
		'a', 'b', 'c', 'd',
		'a', 'b', 'c', 'd',
		'a', 'b', 'c', 'd',
		'a', 'b', 'c', 'd',
		'a', 'b', 'c', 'd',
		'a', 'b', 'c', 'd',
		'a', 'b', 'c', 'd',
		'a', 'b', 'c', 'd',
	}
	h := NewCompress()
	h.Write(preimage)
	actual := hex.EncodeToString(h.Compress())
	expected := "867d9811862dbdab2f8fa343e3e841df7db2ded433172800b0369e8741ec70da"
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got: '%s'", expected, actual)
	}
}

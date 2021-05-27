// Copyright 2014 The go-ethereum Authors
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

package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsAll_whenTypical(t *testing.T) {
	source := []string{"1", "2"}
	target1 := []string{"3"}
	target2 := []string{"1", "2"}

	assert.True(t, ContainsAll(source, target1, target2))
}

func TestContainsAll_whenNot(t *testing.T) {
	source := []string{"1", "2"}
	target := []string{"3", "4"}

	assert.False(t, ContainsAll(source, target))
}

func TestContainsAll_whenTargetIsSubset(t *testing.T) {
	source := []string{"1", "2"}
	target := []string{"1"}

	assert.True(t, ContainsAll(source, target))
}

func TestContainsAll_whenTargetIsSuperSet(t *testing.T) {
	source := []string{"2"}
	target := []string{"1", "2"}

	assert.False(t, ContainsAll(source, target))
}

func TestContainsAll_whenSourceIsEmpty(t *testing.T) {
	var source []string
	target := []string{"1", "2"}

	assert.False(t, ContainsAll(source, target))
}

func TestContainsAll_whenSourceIsNil(t *testing.T) {
	target := []string{"1", "2"}

	assert.False(t, ContainsAll(nil, target))
}

func TestContainsAll_whenTargetIsEmpty(t *testing.T) {
	source := []string{"1", "2"}

	assert.True(t, ContainsAll(source, []string{}))
}

func TestContainsAll_whenTargetIsNil(t *testing.T) {
	source := []string{"1", "2"}

	assert.True(t, ContainsAll(source, nil))
}

func TestAppendSkipDuplicates_whenTypical(t *testing.T) {
	source := []string{"1", "2"}
	additional := []string{"1", "3"}

	assert.Equal(t, []string{"1", "2", "3"}, AppendSkipDuplicates(source, additional...))
}

func TestAppendSkipDuplicates_whenSourceIsNil(t *testing.T) {
	additional := []string{"1", "3"}

	assert.Equal(t, []string{"1", "3"}, AppendSkipDuplicates(nil, additional...))
}

func TestAppendSkipDuplicates_whenElementIsNil(t *testing.T) {
	assert.Equal(t, []string{"1", "3"}, AppendSkipDuplicates([]string{"1", "3"}, nil...))
}

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

// ContainsAll returns true if all elements in one of the targets are in the source,
// false otherwise.
func ContainsAll(source []string, targets ...[]string) bool {
	mark := make(map[string]bool, len(source))
	for _, str := range source {
		mark[str] = true
	}
	for _, target := range targets {
		foundAll := true
		for _, str := range target {
			if _, found := mark[str]; !found {
				foundAll = false
				break
			}
		}
		if foundAll {
			return true
		}
	}
	return false
}

// AppendSkipDuplicates appends source with elements with a condition
// that those elemments must NOT already exist in the source
func AppendSkipDuplicates(slice []string, elems ...string) (result []string) {
	mark := make(map[string]bool, len(slice))
	for _, val := range slice {
		mark[val] = true
	}
	result = slice
	for _, val := range elems {
		if _, ok := mark[val]; !ok {
			result = append(result, val)
		}
	}
	return result
}

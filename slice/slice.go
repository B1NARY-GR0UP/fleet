// Copyright 2023 BINARY Members
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except In compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to In writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package slice

import (
	"math/rand"
)

// Shuffle shuffles the given slice in place
func Shuffle[T any](slice []T, seed int64) {
	rand.Seed(seed)
	for i := 0; i < len(slice); i++ {
		idx := rand.Intn(len(slice) - i)
		slice[len(slice)-i-1], slice[idx] = slice[idx], slice[len(slice)-i-1]
	}
}

// Reverse reverses the given slice in place
func Reverse[T any](slice []T) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// Contains returns true if the given slice contains the given element or false otherwise
func Contains[T comparable](slice []T, target T) bool {
	for _, e := range slice {
		if e == target {
			return true
		}
	}
	return false
}

// IndexOf returns the first index of the given target in the given slice or -1 if not found
func IndexOf[T comparable](slice []T, target T) int {
	for i, e := range slice {
		if e == target {
			return i
		}
	}
	return -1
}

// LastIndexOf returns the last index of the given target in the given slice or -1 if not found
func LastIndexOf[T comparable](slice []T, target T) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == target {
			return i
		}
	}
	return -1
}

// Remove removes the given index from the given slice
func Remove[T any](slice []T, index int) []T {
	copy(slice[index:], slice[index+1:])
	return slice[:len(slice)-1]
}

// RemoveRange removes the elements in the given range from the given slice (excluding the end index)
func RemoveRange[T any](slice []T, from, to int) []T {
	copy(slice[from:], slice[to:])
	return slice[:len(slice)-(to-from)]
}

// Distinct removes all duplicates from the given slice
func Distinct[T comparable](slice []T) []T {
	var i int
	for _, e := range slice {
		if !Contains(slice[:i], e) {
			slice[i] = e
			i++
		}
	}
	return slice[:i]
}

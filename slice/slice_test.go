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
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShuffle(t *testing.T) {
	sli := []int{1, 2, 3, 4, 5}
	Shuffle(sli, 1)
	fmt.Println(sli)
	assert.NotEqual(t, []int{1, 2, 3, 4, 5}, sli)
}

func TestReverse(t *testing.T) {
	sli := []int{1, 2, 3, 4, 5}
	Reverse(sli)
	assert.Equal(t, []int{5, 4, 3, 2, 1}, sli)
}

func TestIndexOf(t *testing.T) {
	sli := []int{1, 2, 3, 3, 4, 5}
	res := IndexOf(sli, 3)
	assert.Equal(t, 2, res)
}

func TestLastIndexOf(t *testing.T) {
	sli := []int{1, 2, 3, 3, 4, 5}
	res := LastIndexOf(sli, 3)
	assert.Equal(t, 3, res)
}

func TestRemoveRange(t *testing.T) {
	sli := []int{1, 2, 3, 4, 5}
	res := RemoveRange(sli, 1, 4)
	assert.Equal(t, []int{1, 5}, res)
}

func TestDistinct(t *testing.T) {
	sli := []int{1, 1, 2, 2, 2, 3, 4, 4, 5, 6, 7, 8, 8, 9, 9}
	res := Distinct(sli)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, res)
}

func TestRemove(t *testing.T) {
	sli := []int{1, 2, 3, 4, 5}
	res := Remove(sli, 2)
	assert.Equal(t, []int{1, 2, 4, 5}, res)
}

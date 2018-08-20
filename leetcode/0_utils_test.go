package leetcode

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestUtils(t *testing.T) {
	as := assert.New(t)

	a := []int{1, 2, 3, 4}
	b := copy_int_slice(a)
	a[0] = 100

	as.Equal([]int{100, 2, 3, 4}, a)
	as.Equal([]int{1, 2, 3, 4}, b)

	a = delete_int_slice_index(a, 1)
	as.Equal([]int{100, 3, 4}, a)

	a = delete_int_slice_index(a, 2)
	as.Equal([]int{100, 3}, a)

	a = delete_int_slice_index(a, 0)
	as.Equal([]int{3}, a)

	a = delete_int_slice_index(a, 0)
	as.Equal([]int{}, a)
}

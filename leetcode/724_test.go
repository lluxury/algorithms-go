package leetcode

import (
	"testing"

	"github.com/Chyroc/algorithms-go/test"
)

func Test_724(t *testing.T) {
	test.Runs(t, pivotIndex, []*test.Case{
		{Input: `[1, 7, 3, 6, 5, 6]`, Output: `3`},
		{Input: `[1, 2, 3]`, Output: `-1`},
	})
}

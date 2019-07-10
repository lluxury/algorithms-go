package leetcode

import (
	"testing"

	"github.com/lluxury/algorithms-go/test"
)

func Test_766(t *testing.T) {
	test.Runs(t, isToeplitzMatrix, []*test.Case{
		{Input: `[  [1,2,3,4],  [5,1,2,3],  [9,5,1,2]]`, Output: `true`},
		{Input: `[  [1,2],  [2,2]]`, Output: `false`},
	})
}

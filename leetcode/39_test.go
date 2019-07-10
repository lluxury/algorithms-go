package leetcode

import (
	"testing"

	"github.com/lluxury/algorithms-go/test"
)

func Test_39(t *testing.T) {
	test.Runs(t, combinationSum, []*test.Case{
		{Input: `[2, 3, 6, 7] \n 7`, Output: `[ [2, 2, 3], [7] ]`},
		{Input: `[2, 3, 5] \n 8`, Output: `[ [2, 2, 2, 2], [2, 3, 3], [3,5] ]`},
	})
}

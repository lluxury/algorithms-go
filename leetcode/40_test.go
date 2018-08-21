package leetcode

import (
	"github.com/Chyroc/algorithms-go/test"
	"testing"
)

func Test_40(t *testing.T) {
	test.Runs(t, combinationSum2, []*test.Case{
		{Input: `[10, 1, 2, 7, 6, 1, 5] \n 8`, Output: `[[1,1,6],[1,2,5],[1,7],[2,6]]`},
		{Input: `[2, 5, 2, 1, 2] \n 5`, Output: `[[1,2,2],[5]]`},
	})
}

package leetcode

import (
	"github.com/Chyroc/algorithms-go/test"
	"testing"
)

func Test_78(t *testing.T) {
	test.Runs(t, subsets, []*test.Case{
		{Input: `[]`, Output: `[[]]`},
		{Input: `[1]`, Output: `[[],[1]]`},
		{Input: `[1, 2]`, Output: `[[],[1],[1,2],[2]]`},
		{Input: `[1, 2, 3]`, Output: `[[],[1],[1,2],[1,2,3],[1,3],[2],[2,3],[3]]`},
	})
}

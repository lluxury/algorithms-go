package leetcode

import (
	"github.com/lluxury/algorithms-go/test"
	"testing"
)

func Test_90(t *testing.T) {
	test.Runs(t, subsetsWithDup, []*test.Case{
		{Input: `[]`, Output: `[[]]`},
		{Input: `[1]`, Output: `[[],[1]]`},
		{Input: `[1,2]`, Output: `[[],[1],[1,2],[2]]`},
		{Input: `[1,2,3]`, Output: `[[],[1],[1,2],[1,2,3],[1,3],[2],[2,3],[3]]`},
		{Input: `[1,2,3,1]`, Output: `[[],[1],[1,1],[1,1,2],[1,1,2,3],[1,1,3],[1,2],[1,2,3],[1,3],[2],[2,3],[3]]`},
	})
}

package leetcode

import (
	"testing"
	"github.com/Chyroc/algorithms-go/test"
)

func Test_47(t *testing.T) {
	test.Runs(t, permuteUnique, []*test.Case{
		{Input: `[]`, Output: `[[]]`},
		{Input: `[1,2,3]`, Output: `[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]`},
		{Input: `[1, 1, 2]`, Output: `[[1,1,2],[1,2,1],[2,1,1]]`},
		{Input: `[3, 3, 0, 3]`, Output: `[[0,3,3,3],[3,0,3,3],[3,3,0,3],[3,3,3,0]]`},
	})
}

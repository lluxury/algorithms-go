package leetcode

import (
	"testing"

	"github.com/Chyroc/algorithms-go/test"
)

func Test_733(t *testing.T) {
	test.Runs(t, floodFill, []*test.Case{
		{Input: `[[1,1,1],[1,1,0],[1,0,1]] \n 1 \n 1 \n 2`, Output: `[[2,2,2],[2,2,0],[2,0,1]]`},
	})
}

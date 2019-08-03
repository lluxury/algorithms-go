package leetcode

import (
	"testing"

	"github.com/lluxury/algorithms-go/test"
)

func Test_994(t *testing.T) {
	test.Runs(t, orangesRotting, []*test.Case{
		{Input: `[[2,1,1],[1,1,0],[0,1,1]]`, Output: `4`},
		{Input: `[[2,1,1],[0,1,1],[1,0,1]]`, Output: `-1`},
	})
}

package leetcode

import (
	"testing"

	"github.com/lluxury/algorithms-go/test"
)

func Test_769(t *testing.T) {
	test.Runs(t, maxChunksToSorted, []*test.Case{
		{Input: `[4,3,2,1,0]`, Output: `1`},
		{Input: `[1,0,2,3,4]`, Output: `4`},
	})
}

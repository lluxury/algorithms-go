package leetcode

import (
	"testing"

	"github.com/lluxury/algorithms-go/test"
)

func Test_849(t *testing.T) {
	test.Runs(t, maxDistToClosest, []*test.Case{
		{Input: `[0,0,0,0,1,0,1]`, Output: `4`},
		{Input: `[1,0,0,0,1,0,1]`, Output: `2`},
		{Input: `[1,0,0,0]`, Output: `3`},
		{Input: `[0,0,1]`, Output: `2`},
	})
}

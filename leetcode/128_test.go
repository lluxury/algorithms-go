package leetcode

import (
	"github.com/Chyroc/algorithms-go/test"
	"testing"
)

func Test_128(t *testing.T) {
	test.Runs(t, longestConsecutive, []*test.Case{
		{Input: `[100, 4, 200, 1, 3, 2]`, Output: "4"},
		{Input: `[1,2,0,1]`, Output: `3`},
		{Input: `[4,0,-4,-2,2,5,2,0,-8,-8,-8,-8,-1,7,4,5,5,-4,6,6,-3]`, Output: `5`},
	})
}

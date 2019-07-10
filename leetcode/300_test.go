package leetcode

import (
	"testing"

	"github.com/lluxury/algorithms-go/test"
)

func Test_300(t *testing.T) {
	test.Runs(t, lengthOfLIS, []*test.Case{
		{Input: `[10,9,2,5,3,7,101,18]`, Output: `4`},
		{Input: `[]`, Output: `0`},
		{Input: `[2,2]`, Output: `1`},
	})
}

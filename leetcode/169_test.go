package leetcode

import (
	"testing"

	"github.com/lluxury/algorithms-go/test"
)

func Test_169(t *testing.T) {
	test.Runs(t, majorityElement, []*test.Case{
		{Input: `[3,2,3]`, Output: `3`},
		{Input: `[2,2,1,1,1,2,2]`, Output: `2`},
	})
}

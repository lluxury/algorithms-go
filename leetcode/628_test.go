package leetcode

import (
	"testing"

	"github.com/lluxury/algorithms-go/test"
)

func Test_628(t *testing.T) {
	test.Runs(t, maximumProduct, []*test.Case{
		{Input: `[1,2,3]`, Output: `6`},
		{Input: `[1,2,3,4]`, Output: `24`},
	})
}

package leetcode

import (
	"github.com/lluxury/algorithms-go/test"
	"testing"
)

func Test_152(t *testing.T) {
	test.Runs(t, maxProduct, []*test.Case{
		{Input: `[2,3,-2,4]`, Output: `6`},
		{Input: `[-2,0,-1]`, Output: `0`},
	})
}

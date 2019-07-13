package leetcode

import (
	"testing"

	"github.com/lluxury/algorithms-go/test"
)

func Test_188(t *testing.T) {
	test.Runs(t, maxProfit_188, []*test.Case{
	    {Input: `2\n[2,4,1]`, Output: `2`},
	    {Input: `2\n[3,2,6,5,0,3]`, Output: `7`},
	})
}


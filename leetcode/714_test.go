package leetcode

import (
	"testing"

	"github.com/lluxury/algorithms-go/test"
)

func Test_714(t *testing.T) {
	test.Runs(t, maxProfit_714, []*test.Case{
	    {Input: `[1,3,2,8,4,9]\n2`, Output: `8`},
	})
}


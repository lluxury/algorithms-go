package leetcode

import (
	"testing"

	"github.com/lluxury/algorithms-go/test"
)

func Test_309(t *testing.T) {
	test.Runs(t, maxProfit_309, []*test.Case{
		{Input: `[1,2,3,0,2]`, Output: `3`},
	})
}

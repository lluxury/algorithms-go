package leetcode

import (
	"github.com/lluxury/algorithms-go/test"
	"testing"
)

func Test_122(t *testing.T) {
	test.Runs(t, maxProfit_122, []*test.Case{
		{Input: `[7,1,5,3,6,4]`, Output: `7`},
		{Input: `[1,2,3,4,5]`, Output: `4`},
		{Input: `[7,6,4,3,1]`, Output: `0`},
	})
}

package leetcode

import (
	"testing"

	"github.com/lluxury/algorithms-go/test"
)

func Test_123(t *testing.T) {
	test.Runs(t, maxProfit_123, []*test.Case{
		{Input: `[3,3,5,0,0,3,1,4]`, Output: `6`},
		{Input: `[1,2,3,4,5]`, Output: `4`},
		{Input: `[7,6,4,3,1]`, Output: `0`},
	})
}

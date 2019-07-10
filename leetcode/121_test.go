package leetcode

import (
	"github.com/lluxury/algorithms-go/test"
	"testing"
)

func Test_121(t *testing.T) {
	test.Runs(t, maxProfit, []*test.Case{
        {Input: `[7,1,5,3,6,4]`, Output: `5`},
		{Input: `[7,6,4,3,1]`, Output: `0`},
	})
}

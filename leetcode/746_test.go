package leetcode

import (
	"testing"

	"github.com/lluxury/algorithms-go/test"
)

func Test_746(t *testing.T) {
	test.Runs(t, minCostClimbingStairs, []*test.Case{
		//{Input: `[10, 15, 20]`, Output: `15`},
		{Input: `[1, 100, 1, 1, 1, 100, 1, 1, 100, 1]`, Output: `6`},
	})
}

package leetcode

import (
	"github.com/lluxury/algorithms-go/test"
	"testing"
)

func Test_70(t *testing.T) {
	test.Runs(t, climbStairs, []*test.Case{
		{Input: `2`, Output: "2"},
		{Input: `3`, Output: "3"},
	})
}

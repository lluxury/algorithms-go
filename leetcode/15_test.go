package leetcode

import (
	"github.com/lluxury/algorithms-go/test"
	"testing"
)

func Test_15(t *testing.T) {
	test.Runs(t, threeSum, []*test.Case{
		{Input: `[-1, 0, 1, 2, -1, -4]`, Output: "[[-1, -1, 2],[-1, 0, 1]]"},
	})
}

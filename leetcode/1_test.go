package leetcode

import (
	"testing"

	"github.com/lluxury/algorithms-go/test"
)

func Test_1(t *testing.T) {
	for _, f := range []func(nums []int, target int) []int{twoSum_1, twoSum_2} {
		test.Runs(t, f, []*test.Case{
			{Input: `[2, 7, 11, 15] \n 9`, Output: "[0, 1]"},
		})
	}
	test.Runs(t, twoSum_3, []*test.Case{
		{Input: `[2, 7, 11, 15] \n 9`, Output: "[1, 0]"},
	})
}

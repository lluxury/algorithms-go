package leetcode

import (
	"testing"
	"github.com/Chyroc/algorithms-go/test"
)

func Test_674(t *testing.T) {
	test.Runs(t, findLengthOfLCIS, []*test.Case{
		{Input: `[1, 3, 5, 7]`, Output: "4"},
		{Input: `[1, 3, 5, 4, 7]`, Output: "3"},
		{Input: `[2, 2, 2]`, Output: "1"},
	})
}

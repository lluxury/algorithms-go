package leetcode

import (
	"testing"

	"github.com/Chyroc/algorithms-go/test"
)

func Test_908(t *testing.T) {
	test.Runs(t, smallestRangeI, []*test.Case{
		{Input: `[1] \n 0`, Output: `0`},
		{Input: `[0,10] \n 2`, Output: `6`},
		{Input: `[1,3,6] \n 3`, Output: `0`},
	})
}

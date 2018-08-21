package leetcode

import (
	"testing"

	"github.com/Chyroc/algorithms-go/test"
)

func Test_786(t *testing.T) {
	test.Runs(t, kthSmallestPrimeFraction, []*test.Case{
		{Input: `[1,2,3,5] \n 3`, Output: `[2,5]`},
		{Input: `[1,7] \n 1`, Output: `[1,7]`},
	})
}

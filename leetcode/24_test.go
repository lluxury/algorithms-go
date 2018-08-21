package leetcode

import (
	"testing"

	"github.com/Chyroc/algorithms-go/test"
)

func Test_24(t *testing.T) {
	test.Runs(t, swapPairs, []*test.Case{
		{Input: `1->2->3->4`, Output: `2->1->4->3`},
		{Input: `1`, Output: `1`},
	})
}

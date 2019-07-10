package leetcode

import (
	"github.com/lluxury/algorithms-go/test"
	"testing"
)

func Test_19(t *testing.T) {
	test.Runs(t, removeNthFromEnd, []*test.Case{
		{Input: `1->2->3->4->5 \n 2`, Output: `1->2->3->5`},
		{Input: `1 \n 1`, Output: ``},
		{Input: `1->2 \n 1`, Output: `1`},
		{Input: `1->2 \n 2`, Output: `2`},
		{Input: `1->2->3 \n 3`, Output: `2->3`},
	})
}

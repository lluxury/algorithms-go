package leetcode

import (
	"github.com/lluxury/algorithms-go/lib"
	"github.com/lluxury/algorithms-go/test"
	"testing"
)

func Test_206(t *testing.T) {
	for _, f := range []func(head *lib.ListNode) *lib.ListNode{reverseList, reverseList2, reverseList3} {
		test.Runs(t, f, []*test.Case{
			{Input: `1->2->3->4->5`, Output: `5->4->3->2->1`},
		})
	}
}

package leetcode

import (
	"github.com/Chyroc/algorithms-go/lib"
	"github.com/Chyroc/algorithms-go/test"
	"testing"
)

/*

> https://leetcode.com/problems/reverse-nodes-in-k-group/description/

简单的一个： https://leetcode.com/problems/swap-nodes-in-pairs/description/


*/

func reverseKGroup_nil(head *lib.ListNode, k int) *lib.ListNode {
	tmp := head
	for i := 0; i < k-1; i++ {
		if &tmp != nil {
			if tmp.Next == nil {
				tmp = nil
			} else {
				tmp = tmp.Next
			}
		}
	}

	return tmp
}

func reverseKGroup(head *lib.ListNode, k int) *lib.ListNode {
	x := reverseKGroup_nil(head, k)
	if x == nil {
		return head
	}

	return nil
}

func Test_25(t *testing.T) {
	t.SkipNow()
	test.Runs(t, reverseKGroup, []*test.Case{
		{Input: `1->2->3->4->5 \n 6`, Output: `1->2->3->4->5`},
		//{Input: `1->2->3->4->5 \n 2`, Output: `2->1->4->3->5`},
		//{Input: `1->2->3->4->5 \n 3`, Output: `3->2->1->4->5`},
	})
}

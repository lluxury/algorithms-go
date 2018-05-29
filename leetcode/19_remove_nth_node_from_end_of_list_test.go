package leetcode

import (
	"github.com/Chyroc/algorithms-go/lib"
	"testing"
	"github.com/Chyroc/algorithms-go/test"
)

/*

> https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/


Given a linked list, remove the n-th node from the end of list and return its head.

Example:

Given linked list: 1->2->3->4->5, and n = 2.

After removing the second node from the end, the linked list becomes 1->2->3->5.
Note:

Given n will always be valid.

Follow up:

Could you do this in one pass?



* 给一个链表，让删除倒数第n个，返回新链表；只能遍历一遍
  * 将每个数组存在数组中 X
  * 从第一个和第n个遍历，第二个到头的时候，第一个也就到了倒数第n个

*/

func removeNthFromEnd(head *lib.ListNode, n int) *lib.ListNode {
	var r1 = head
	var r2 = head

	for i := 0; i < n; i++ {
		r2 = r2.Next
	}

	if r2 == nil {
		return r1.Next
	}

	for r2.Next != nil {
		r1 = r1.Next
		r2 = r2.Next
	}

	if r1.Next != nil {
		r1.Next = r1.Next.Next
	}

	return head
}

func Test_19(t *testing.T) {
	test.Runs(t, removeNthFromEnd, []*test.Case{
		{Input: `1->2->3->4->5 \n 2`, Output: `1->2->3->5`},
		{Input: `1 \n 1`, Output: ``},
		{Input: `1->2 \n 1`, Output: `1`},
		{Input: `1->2 \n 2`, Output: `2`},
		{Input: `1->2->3 \n 3`, Output: `2->3`},
	})
}

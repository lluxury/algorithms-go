package leetcode

import (
	"github.com/Chyroc/algorithms-go/lib"
	"testing"
	"github.com/Chyroc/algorithms-go/test"
)

/*
> https://leetcode.com/problems/reverse-linked-list/description/

Reverse a singly linked list.

click to show more hints.

**Hint:**
A linked list can be reversed either iteratively or recursively. Could you implement both?

- 题目
  - 反转单链表
  - `1 → 2 → 3 → Ø`, we would like to change it to `Ø ← 1 ← 2 ← 3`，最后返回的是最后3的地址
- 思考
  - 两种方法
  - 循环
    - 先设一个循环的当前节点 curr 和前一节点指针 prev
    - 设当前节点 A ,下一节点 B
    - 先斩断 A 指向 B
    - 将 A 指向前一个 prev
    - 将 当前节点 curr 前移
    - 将 前置节点 prev 前移
  - 递归
    - 如果只有一个或者0个节点，返回
    - 将第一个节点之外的链表反正，返回最后一个指针t，因为第一个节点不会是最后一个节点，左右返回的也是t
    - 将 后面反转的那个链表的第一个节点 指向 全链表的第一个节点
    - 将 全链表的第一个节点 指向 null

*/

func reverseList(head *lib.ListNode) *lib.ListNode {
	var prev *lib.ListNode
	var curr = head

	for curr != nil {
		tmp := curr.Next
		curr.Next = prev
		prev = curr
		curr = tmp
	}

	return prev
}

func reverseList2(head *lib.ListNode) *lib.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	t := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil

	return t
}

func Test_206(t *testing.T) {
	test.Runs(t, reverseList, []*test.Case{
		{Input: `1->2->3->4->5`, Output: `5->4->3->2->1`},
	})
	test.Runs(t, reverseList2, []*test.Case{
		{Input: `1->2->3->4->5`, Output: `5->4->3->2->1`},
	})
}

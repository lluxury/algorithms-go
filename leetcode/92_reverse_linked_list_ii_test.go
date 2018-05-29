package leetcode

import (
	"github.com/Chyroc/algorithms-go/lib"
	"testing"
	"github.com/Chyroc/algorithms-go/test"
)

/*

> https://leetcode.com/problems/reverse-linked-list-ii/description/


Reverse a linked list from position m to n. Do it in one-pass.

Note: 1 ≤ m ≤ n ≤ length of list.

Example:

Input: 1->2->3->4->5->NULL, m = 2, n = 4
Output: 1->4->3->2->5->NULL

* 反转链表中指定的第m到n个元素
  * 首先将指针移动到第m个元素的前一个元素那里，假设为pre
  * 然后操作n-m次：将pre的下一个节点和下下一个节点交换
  * 返回第一个节点
*/

func reverseBetween(head *lib.ListNode, m int, n int) *lib.ListNode {
	if head == nil {
		return nil
	}

	dummy := &lib.ListNode{Next: head}

	pre := dummy
	for i := 0; i < m-1; i++ {
		pre = pre.Next
	}

	start := pre.Next
	then := pre.Next.Next
	for i := 0; i < n-m; i++ {
		start.Next = then.Next
		then.Next = pre.Next
		pre.Next = then
		then = start.Next
	}

	return dummy.Next
}

func reverseBetween2(head *lib.ListNode, m int, n int) *lib.ListNode {
	if head == nil {
		return nil
	}

	dummy := &lib.ListNode{Next: head}

	pre := dummy
	for i := 0; i < m-1; i++ {
		pre = pre.Next
	}

	for i := 0; i < n-m; i++ {
		tmp1 := pre.Next
		tmp2 := pre.Next.Next

		pre.Next = pre.Next.Next.Next
		tmp1.Next = pre.Next
		pre.Next = tmp2

		

	}

	return dummy.Next
}

func Test_92(t *testing.T) {
	test.Runs(t, reverseBetween2, []*test.Case{
		{Input: `1->2->3->4->5 \n 2 \n 4`, Output: `1->4->3->2->5`},
	})
}

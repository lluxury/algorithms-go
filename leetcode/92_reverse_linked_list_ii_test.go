package leetcode

import (
	"testing"

	"github.com/Chyroc/algorithms-go/lib"
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
  * ￥￥然后操作n-m次：将first到pre之间的节点移动到first.next节点之后￥￥
  * 返回第一个节点

翻转链表
假设有节点 : N -> A -> B -> C -> D -> E -> F
设定：
  * N是pre节点
  * A是first节点，也就是要翻转的链表的第一个节点
操作
  * 1. 将first到pre之间的节点，即A移动到B之后:   N -> B -> a -> C -> D -> E -> F
  * 2. 将first到pre之间的节点，即BA移动到C之后:  N -> C -> B -> a -> D -> E -> F
  * 1. 将first到pre之间的节点，即CBA移动到D之后: N -> D -> C -> B -> a -> E -> F
解题步骤
  * 首先计算出then
  * 一定要把first和pre之间的当做一个区间，就算真的只有一个，也要当做一个段来思考
  * 然后从first指向谁开始做起，也就是first.next=xxx
保全第一个数据
  * dummy := &lib.ListNode{Next: head}

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

	first := pre.Next
	for i := 0; i < n-m; i++ {
		then := first.Next

		first.Next = then.Next
		then.Next = pre.Next
		pre.Next = then
	}

	return dummy.Next
}

func Test_92(t *testing.T) {
	test.Runs(t, reverseBetween, []*test.Case{
		{Input: `1->2->3->4->5 \n 2 \n 4`, Output: `1->4->3->2->5`},
	})
}

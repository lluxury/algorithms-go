package leetcode

import "github.com/Chyroc/algorithms-go/lib"

/*
 * [24] Swap Nodes in Pairs
 *
 * https://leetcode.com/problems/swap-nodes-in-pairs/description/
 *
 * algorithms
 * Medium (40.35%)
 * Total Accepted:    232.6K
 * Total Submissions: 575.5K
 * Testcase Example:  '[1,2,3,4]'
 *
 * Given a linked list, swap every two adjacent nodes and return its head.
 *
 * Example:
 *
 *
 * Given 1->2->3->4, you should return the list as 2->1->4->3.
 *
 * Note:
 *
 *
 * Your algorithm should use only constant extra space.
 * You may not modify the values in the list's nodes, only nodes itself may be
 * changed.
 *
 *
 */

/*

* 交换链表中的数据，每两个数据交换一次
  * 只能交换数据，而不能改变链表中的节点的数据
  * 不能使用额外空间

更难的一个：https://leetcode.com/problems/reverse-nodes-in-k-group/description/

* 如果没有节点或者只有一个节点，返回这个节点
* 如果有两个节点，交换这两个节点，并且把输入节点的下一个节点的next返回
* 如果有更多个，那么把输入节点的next的next递归处理，作为当前节点的next


*/

func swapPairs(head *lib.ListNode) *lib.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	tmp := head.Next
	head.Next = swapPairs(head.Next.Next)
	tmp.Next = head
	return tmp
}

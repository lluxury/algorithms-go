package leetcode

import "github.com/lluxury/algorithms-go/lib"

/*
 * [203] Remove Linked List Elements
 *
 * https://leetcode.com/problems/remove-linked-list-elements/description/
 *
 * algorithms
 * Easy (33.95%)
 * Total Accepted:    166.1K
 * Total Submissions: 489.1K
 * Testcase Example:  '[1,2,6,3,4,5,6]\n6'
 *
 * Remove all elements from a linked list of integers that have value val.
 *
 * Example:
 *
 *
 * Input:  1->2->6->3->4->5->6, val = 6
 * Output: 1->2->3->4->5
 *
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
 移除一个链表里面给定数值的节点
   * 遍历之
   * 如果为nil，返回
   * 如果当前节点的值等于v，那说明当前节点应该丢弃，所以直接用next迭代
   * 否则，说明当前节点没问题，需要将下一个节点作为输入进行计算，然后放在当前节点的下一个节点上
*/

func removeElements(head *lib.ListNode, val int) *lib.ListNode {
	if head == nil {
		return nil
	}
	if head.Val == val {
		return removeElements(head.Next, val)
	}

	head.Next = removeElements(head.Next, val)

	return head
}

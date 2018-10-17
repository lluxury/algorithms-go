package leetcode

import "github.com/Chyroc/algorithms-go/lib"

/*
 * [83] Remove Duplicates from Sorted List
 *
 * https://leetcode.com/problems/remove-duplicates-from-sorted-list/description/
 *
 * algorithms
 * Easy (40.74%)
 * Total Accepted:    256.2K
 * Total Submissions: 628.5K
 * Testcase Example:  '[1,1,2]'
 *
 * Given a sorted linked list, delete all duplicates such that each element
 * appear only once.
 *
 * Example 1:
 *
 *
 * Input: 1->1->2
 * Output: 1->2
 *
 *
 * Example 2:
 *
 *
 * Input: 1->1->2->3->3
 * Output: 1->2->3
 *
 */

/*

* 给定一个排序的链表（换成数组也是一样的），去掉重复元素
  * 循环链表
    * 如果当前数据和前一个数字相等，那么直接跳过，到下一次循环
    * 否则将当前节点加到新的数组

*/

func deleteDuplicates(head *lib.ListNode) *lib.ListNode {
	if head == nil {
		return nil
	}
	var p = head
	var n = head
	var tmp = head.Val
	for n != nil {
		if tmp == n.Val {
			n = n.Next
		} else {
			tmp = n.Val
			p.Next = n
			p = p.Next
			n = n.Next
		}
	}
	p.Next = nil
	return head
}

package leetcode

import "github.com/lluxury/algorithms-go/lib"

/*
 * [2] Add Two Numbers
 *
 * https://leetcode.com/problems/add-two-numbers/description/
 *
 * algorithms
 * Medium (29.02%)
 * Total Accepted:    562.4K
 * Total Submissions: 1.9M
 * Testcase Example:  '[2,4,3]\n[5,6,4]'
 *
 * You are given two non-empty linked lists representing two non-negative
 * integers. The digits are stored in reverse order and each of their nodes
 * contain a single digit. Add the two numbers and return it as a linked list.
 *
 * You may assume the two numbers do not contain any leading zero, except the
 * number 0 itself.
 *
 * Example:
 *
 *
 * Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
 * Output: 7 -> 0 -> 8
 * Explanation: 342 + 465 = 807.
 *
 *
 */

/*

 就是把十进制数存在链表里面，然后求和（链表前面的数字是最低位）

*/

// 链表节点分别相加
func addTwoNumbers(l1 *lib.ListNode, l2 *lib.ListNode) *lib.ListNode {
	l3 := &lib.ListNode{}
	l := l3

	add := 0
	for l1 != nil || l2 != nil || add != 0 {
		var x, y int
		if l1 == nil {
			x = 0
		} else {
			x = l1.Val
			l1 = l1.Next
		}

		if l2 == nil {
			y = 0
		} else {
			y = l2.Val
			l2 = l2.Next
		}

		z := x + y + add

		l3.Val = z % 10
		add = z / 10

		if l1 != nil || l2 != nil || add != 0 {
			l3.Next = &lib.ListNode{}
			l3 = l3.Next
		} else {
			l3.Next = nil
		}
	}

	return l
}

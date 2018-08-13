package leetcode

import "github.com/Chyroc/algorithms-go/lib"

/*
 * [234] Palindrome Linked List
 *
 * https://leetcode.com/problems/palindrome-linked-list/description/
 *
 * algorithms
 * Easy (33.94%)
 * Total Accepted:    181.1K
 * Total Submissions: 532.3K
 * Testcase Example:  '[1,2]'
 *
 * Given a singly linked list, determine if it is a palindrome.
 *
 * Example 1:
 *
 *
 * Input: 1->2
 * Output: false
 *
 * Example 2:
 *
 *
 * Input: 1->2->2->1
 * Output: true
 *
 * Follow up:
 * Could you do it in O(n) time and O(1) space?
 *
 */

/*

 * 判断单链表是否是回文链表，要求使用 O(n) time and O(1) space

 */

func isPalindrome_234(head *lib.ListNode) bool {
	head2 := head.Clone()
	r := reverseList3(head)

	return lib.EqualsToListNode(head2, r)
}

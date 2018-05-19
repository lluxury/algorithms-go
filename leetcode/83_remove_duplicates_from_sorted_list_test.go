package leetcode

/*
> https://leetcode.com/problems/remove-duplicates-from-sorted-list/description/


Given a sorted linked list, delete all duplicates such that each element appear only once.

For example,
Given `1->1->2`, return `1->2`.
Given `1->1->2->3->3`, return `1->2->3`.

* 给定一个排序的链表（换成数组也是一样的），去掉重复元素
  * 循环链表
    * 如果当前数据和前一个数字相等，那么直接跳过，到下一次循环
    * 否则将当前节点加到新的数组


*/

import (
	"strconv"
	"testing"

	"github.com/Chyroc/algorithms-go/test"
	"github.com/Chyroc/algorithms-go/lib"
)

func (l *ListNode) String() string {
	s := ""
	for l != nil {
		s = s + strconv.Itoa(l.Val)
		l = l.Next
	}
	return s
}

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

func Test_83(t *testing.T) {
	test.Runs(t, deleteDuplicates, []*test.Case{
		{Input: `1 -> 1 -> 2`, Output: "1 -> 2"},
		{Input: ``, Output: ""},
	})
}

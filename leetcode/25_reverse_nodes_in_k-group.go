package leetcode
import "github.com/Chyroc/algorithms-go/lib"

/*
Given a linked list, reverse the nodes of a linked list k at a time and return its modified list.

k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes in the end should remain as it is.

Example:

Given this linked list: 1->2->3->4->5

For k = 2, you should return: 2->1->4->3->5

For k = 3, you should return: 3->2->1->4->5

Note:

Only constant extra memory is allowed.
You may not alter the values in the list's nodes, only nodes itself may be changed.*/





// import (
// 	"strconv"
// 	"strings"
// )

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func (node *ListNode) Add(v int) {
// 	if node == nil {
// 		*node = ListNode{}
// 	}
// 	tmp := node
// 	for tmp.Next != nil {
// 		tmp = tmp.Next
// 	}
// 	tmp.Next = &ListNode{Val: v}
// }

// func (node *ListNode) Clone() *ListNode {
// 	if node == nil {
// 		return nil
// 	}
// 	var l = &ListNode{Val: node.Val}
// 	var tmpL = l
// 	var tmpNode = node

// 	for tmpNode.Next != nil {
// 		tmpNode = tmpNode.Next
// 		tmpL.Next = &ListNode{Val: tmpNode.Val}
// 		tmpL = tmpL.Next
// 	}

// 	return l
// }

// func EqualsToListNode(a, b *ListNode) bool {
// 	if a == nil || b == nil {
// 		return b == b
// 	}

// 	return a.Val == b.Val && EqualsToListNode(a.Next, b.Next)
// }

// func (node *ListNode) Unmarshal(param string) (interface{}, error) {
// 	var list = new(ListNode)
// 	for k, v := range strings.Split(param, "->") {

// 		v = strings.TrimSpace(v)

// 		i, err := strconv.Atoi(v)
// 		if err != nil {
// 			return nil, err
// 		}

// 		if k == 0 {
// 			list = &ListNode{Val: i}
// 		} else {
// 			list.Add(i)
// 		}
// 	}
// 	return list, nil
// }

// func (node *ListNode) Marshal() (string, error) {
// 	var s []string
// 	var tmp = node

// 	if node != nil {
// 		s = append(s, strconv.Itoa(node.Val))
// 	}

// 	for tmp.Next != nil {
// 		tmp = tmp.Next
// 		s = append(s, strconv.Itoa(tmp.Val))
// 	}
// 	return strings.Join(s, "->"), nil
// }

// func (node *ListNode) String() string {
// 	s, _ := node.Marshal()
// 	return s
// }

func reverseKGroup(head *lib.ListNode, k int) *lib.ListNode {
	if head == nil || head.Next == nil || k < 2 {
		return head
	}
	var next_group *lib.ListNode = head

	for i := 0; i < k; i++ {
		if next_group != nil {
			next_group = next_group.Next
		} else {
			return head
		}

	}
	var new_next_group *lib.ListNode = reverseKGroup(next_group, k)
	var prev *lib.ListNode = nil
	var cur *lib.ListNode = head
	for cur != next_group {
		var next *lib.ListNode = cur.Next
		// cur.Next = prev != nil ? prev : new_next_group
		if prev != nil {
			cur.Next = prev
		} else {
			cur.Next = new_next_group
		}
		prev = cur
		cur = next
	}
	return prev
}



package leetcode

import "github.com/lluxury/algorithms-go/lib"

/*
Given a linked list, return the node where the cycle begins. If there is no cycle, return null.

To represent a cycle in the given linked list, we use an integer pos which represents the position (0-indexed) in the linked list where tail connects to. If pos is -1, then there is no cycle in the linked list.

Note: Do not modify the linked list.



Example 1:

Input: head = [3,2,0,-4], pos = 1
Output: tail connects to node index 1
Explanation: There is a cycle in the linked list, where tail connects to the second node.


Example 2:

Input: head = [1,2], pos = 0
Output: tail connects to node index 0
Explanation: There is a cycle in the linked list, where tail connects to the first node.


Example 3:

Input: head = [1], pos = -1
Output: no cycle
Explanation: There is no cycle in the linked list.
*/

// 返回列表有换的地方,没有环返回null

func detectCycle(head *lib.ListNode) *lib.ListNode {
	var slow *lib.ListNode = head
	var fast *lib.ListNode = head

	// if (head == nil || head.Next == nil){
	//     return false
	// }

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			var slow2 *lib.ListNode = head

			for slow2 != slow {
				slow2 = slow2.Next
				slow = slow.Next
			}
			return slow2
		}
	}
	return nil
}

// func detectCycle(head *ListNode) *ListNode {
//     var slow *ListNode = head
//     var fast *ListNode = head

//     // if (head == nil || head.Next == nil){
//     //     return false
//     // }

//     for (fast != nil && fast.Next != nil){
//         fast = fast.Next.Next
//         slow = slow.Next
//         if (slow == fast){
//             var slow2 *ListNode = head

//             for (slow2 != slow){
//                 slow2 = slow2.Next
//                 slow = slow.Next
//             }
//             return slow2
//         }
//     }
// return nil
// }

// 后期整理

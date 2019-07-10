package leetcode

import "github.com/lluxury/algorithms-go/lib"

/*
 * [23] Merge k Sorted Lists
 *
 * https://leetcode.com/problems/merge-k-sorted-lists/description/
 *
 * algorithms
 * Hard (29.66%)
 * Total Accepted:    255.4K
 * Total Submissions: 859.2K
 * Testcase Example:  '[[1,4,5],[1,3,4],[2,6]]'
 *
 * Merge k sorted linked lists and return it as one sorted list. Analyze and
 * describe its complexity.
 *
 * Example:
 *
 *
 * Input:
 * [
 * 1->4->5,
 * 1->3->4,
 * 2->6
 * ]
 * Output: 1->1->2->3->4->4->5->6
 *
 *
 */

/*


* 合并k个有序链表
 * 和这题相似的一个题目是合并2个有序数组
   * 那个简单，直接用i和j指向数组的前面合并就行
   * 时间复杂度 O(m+n)
 * 另一个类似的是找两个有序数组中的第k大的数，和上面的类似
 * 本题
   * 可以考虑和上面使用相同的办法，使用一个长度k的数组存index，那么每次取最小值的时候就需要一个O(k)的比较，整体是需要一个O(km)的遍历（假设平均长度n，k个的话总长度就是km），所以是时间复杂度O(mk^2)
   * 可以将上面那个数组换成最小堆，这样每次比较就是O(logk)，复杂度O(mklogk)
   * 如果用n表示所有数据的总个数，那么他们的复杂度分别是O(nk) O(n logk)

*/

func mergeKLists(lists []*lib.ListNode) *lib.ListNode {
	switch len(lists) {
	case 0:
		return nil
	case 1:
		return lists[0]
	}

	var input []interface{}
	for _, v := range lists {
		if v != nil {
			input = append(input, v)
		}
	}
	heap := lib.NewHeap(func(a, b interface{}) bool { return a.(*lib.ListNode).Val < b.(*lib.ListNode).Val }, input...)

	if heap.Len() == 0 {
		return nil
	}

	var tmp *lib.ListNode
	var result *lib.ListNode
	for {
		top, ok := heap.Pop()
		if !ok || top == nil {
			break
		}
		t := top.(*lib.ListNode)
		if tmp == nil {
			tmp = &lib.ListNode{Val: t.Val}
			result = tmp
		} else {
			tmp.Next = &lib.ListNode{Val: t.Val}
			tmp = tmp.Next
		}
		if t.Next != nil {
			heap.Add(t.Next)
		}
	}

	return result
}

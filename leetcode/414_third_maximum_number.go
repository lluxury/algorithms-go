package leetcode

import (
	"github.com/Chyroc/algorithms-go/lib"
)

/*
 * [414] Third Maximum Number
 *
 * https://leetcode.com/problems/third-maximum-number/description/
 *
 * algorithms
 * Easy (28.19%)
 * Total Accepted:    67.9K
 * Total Submissions: 240.8K
 * Testcase Example:  '[3,2,1]'
 *
 * Given a non-empty array of integers, return the third maximum number in this
 * array. If it does not exist, return the maximum number. The time complexity
 * must be in O(n).
 *
 * Example 1:
 *
 * Input: [3, 2, 1]
 *
 * Output: 1
 *
 * Explanation: The third maximum is 1.
 *
 *
 *
 * Example 2:
 *
 * Input: [1, 2]
 *
 * Output: 2
 *
 * Explanation: The third maximum does not exist, so the maximum (2) is
 * returned instead.
 *
 *
 *
 * Example 3:
 *
 * Input: [2, 2, 3, 1]
 *
 * Output: 1
 *
 * Explanation: Note that the third maximum here means the third maximum
 * distinct number.
 * Both numbers with value 2 are both considered as second maximum.
 *
 *
 */

/*

 给一个数组，qui里面第三大的数（相同的数字只计算一遍）

 第x大的数：用最小堆，长度为x，然后返回最小值

 本题还有一个额外的条件：相同的数字只计算一遍
 所以还加了一个map，记录什么数字访问过，已经访问过的数字，不加入到最小堆

*/

func thirdMax(nums []int) int {
	h := lib.NewHeap(func(a, b interface{}) bool { return a.(int) < b.(int) })

	m := make(map[int]bool)

	for _, v := range nums {
		if !m[v] {
			m[v] = true
			h.Add(v)

			if h.Len() > 3 {
				y, _ := h.Pop()
				delete(m, y.(int))
			}
		}

	}

	if len(m) >= 3 {
		return h.Peek().(int)
	}

	max := h.Peek().(int)
	for {
		x, ok := h.Pop()
		if !ok {
			break
		}
		max = x.(int)
	}

	return max
}

package leetcode

/*
 * [169] Majority Element
 *
 * https://leetcode.com/problems/majority-element/description/
 *
 * algorithms
 * Easy (49.75%)
 * Total Accepted:    304.5K
 * Total Submissions: 612K
 * Testcase Example:  '[3,2,3]'
 *
 * Given an array of size n, find the majority element. The majority element is
 * the element that appears more than ⌊ n/2 ⌋ times.
 *
 * You may assume that the array is non-empty and the majority element always
 * exist in the array.
 *
 * Example 1:
 *
 *
 * Input: [3,2,3]
 * Output: 3
 *
 * Example 2:
 *
 *
 * Input: [2,2,1,1,1,2,2]
 * Output: 2
 *
 */

func majorityElement(nums []int) int {
	n := len(nums) / 2
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
		if m[v] > n {
			return v
		}
	}
	return 0
}

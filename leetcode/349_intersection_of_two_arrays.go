package leetcode

/*
 * [349] Intersection of Two Arrays
 *
 * https://leetcode.com/problems/intersection-of-two-arrays/description/
 *
 * algorithms
 * Easy (48.70%)
 * Total Accepted:    148.8K
 * Total Submissions: 304.3K
 * Testcase Example:  '[1,2,2,1]\n[2,2]'
 *
 * Given two arrays, write a function to compute their intersection.
 *
 * Example 1:
 *
 *
 * Input: nums1 = [1,2,2,1], nums2 = [2,2]
 * Output: [2]
 *
 *
 *
 * Example 2:
 *
 *
 * Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
 * Output: [9,4]
 *
 *
 * Note:
 *
 *
 * Each element in the result must be unique.
 * The result can be in any order.
 *
 *
 *
 *
 */

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	for _, v := range nums1 {
		m[v] = 1
	}

	data := []int{}
	for _, v := range nums2 {
		if m[v] == 1 {
			data = append(data, v)
			m[v]--
		}
	}

	return data
}

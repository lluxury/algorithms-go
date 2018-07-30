package leetcode

/*
 * [300] Longest Increasing Subsequence
 *
 * https://leetcode.com/problems/longest-increasing-subsequence/description/
 *
 * algorithms
 * Medium (39.07%)
 * Total Accepted:    140.2K
 * Total Submissions: 358.9K
 * Testcase Example:  '[10,9,2,5,3,7,101,18]'
 *
 * Given an unsorted array of integers, find the length of longest increasing
 * subsequence.
 *
 * Example:
 *
 *
 * Input: [10,9,2,5,3,7,101,18]
 * Output: 4
 * Explanation: The longest increasing subsequence is [2,3,7,101], therefore
 * the length is 4.
 *
 * Note:
 *
 *
 * There may be more than one LIS combination, it is only necessary for you to
 * return the length.
 * Your algorithm should run in O(n2) complexity.
 *
 *
 * Follow up: Could you improve it to O(n log n) time complexity?
 *
 */

/*

最大增序子序列的长度

动态规划算法

	https://www.zhihu.com/question/23995189
	https://www.cnblogs.com/steven_oyj/archive/2010/05/22/1741374.html

	首先需要定义 状态 和 状态转移方程

	状态
	Fk = 数组前k项的 最长递增子序列的长度

	状态转移方程
	F1 = 1
	Fk = 在i = 0-k-1的情况下，如果第i项比第k项小，Fk = max(Fi + 1)

*/

func max_all(s ...int) int {
	switch len(s) {
	case 0:
		panic("max all")
	case 1:
		return s[0]
	case 2:
		if s[0] > s[1] {
			return s[0]
		}
		return s[1]
	default:
		return max_all(append([]int{max_all(s[0], s[1])}, s[2:]...)...)
	}
}

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	m := make([]int, len(nums))
	for k := range m {
		m[k] = 1
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				m[i] = max_all(m[i], m[j]+1)
			}
		}
	}

	return max_all(m...)
}

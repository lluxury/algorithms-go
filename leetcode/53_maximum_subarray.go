package leetcode

/*
 * [53] Maximum Subarray
 *
 * https://leetcode.com/problems/maximum-subarray/description/
 *
 * algorithms
 * Easy (40.84%)
 * Total Accepted:    341.2K
 * Total Submissions: 835.3K
 * Testcase Example:  '[-2,1,-3,4,-1,2,1,-5,4]'
 *
 * Given an integer array nums, find the contiguous subarray (containing at
 * least one number) which has the largest sum and return its sum.
 *
 * Example:
 *
 *
 * Input: [-2,1,-3,4,-1,2,1,-5,4],
 * Output: 6
 * Explanation: [4,-1,2,1] has the largest sum = 6.
 *
 *
 * Follow up:
 *
 * If you have figured out the O(n) solution, try coding another solution using
 * the divide and conquer approach, which is more subtle.
 *
 */

/*

给一个数组，求连续子序列和的最大值

动态规划算法

	首先需要定义 状态 和 状态转移方程

	状态
	Fk = 包含第k项以及他之前的连续子序列的和的最大值

	状态转移
	Fx 要么等于第k项本身（前面的连续和小于0），要么是他加上F(k-1)
	结果是所有Fk的最大值

本题是动态规划的一种类型，用局部值和全局值做状态，局部值指包含有当前状态的值

*/

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	local := nums[0]
	global := nums[0]

	for i := 1; i < len(nums); i++ {
		local = max(nums[i], nums[i]+local)
		global = max(local, global)
	}
	return global
}

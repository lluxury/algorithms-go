package leetcode

/*
 * [724] Find Pivot Index
 *
 * https://leetcode.com/problems/find-pivot-index/description/
 *
 * algorithms
 * Easy (39.26%)
 * Total Accepted:    28.1K
 * Total Submissions: 71.7K
 * Testcase Example:  '[1,7,3,6,5,6]'
 *
 * Given an array of integers nums, write a method that returns the "pivot"
 * index of this array.
 *
 * We define the pivot index as the index where the sum of the numbers to the
 * left of the index is equal to the sum of the numbers to the right of the
 * index.
 *
 * If no such index exists, we should return -1. If there are multiple pivot
 * indexes, you should return the left-most pivot index.
 *
 *
 * Example 1:
 *
 * Input:
 * nums = [1, 7, 3, 6, 5, 6]
 * Output: 3
 * Explanation:
 * The sum of the numbers to the left of index 3 (nums[3] = 6) is equal to the
 * sum of numbers to the right of index 3.
 * Also, 3 is the first index where this occurs.
 *
 *
 *
 * Example 2:
 *
 * Input:
 * nums = [1, 2, 3]
 * Output: -1
 * Explanation:
 * There is no index that satisfies the conditions in the problem statement.
 *
 *
 *
 * Note:
 * The length of nums will be in the range [0, 10000].
 * Each element nums[i] will be an integer in the range [-1000, 1000].
 *
 */

/*

 给一个数组，求一个index，使得他左边的数等于右边的数
 * 遍历一遍，求每个数到起点的和
 * 遍历第二遍，可以通过计算知道当前节点之前和之和的和，判断是否相等

*/

func pivotIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	sum := make([]int, len(nums))
	for k, v := range nums {
		if k == 0 {
			sum[k] = v
		} else {
			sum[k] = v + sum[k-1]
		}
	}

	last := sum[len(sum)-1]

	for k, v := range sum {
		if k == 0 {
			if last-v == 0 {
				return 0
			}
		} else {
			left := sum[k-1]
			right := last - v
			if left == right {
				return k
			}
		}
	}

	return -1
}

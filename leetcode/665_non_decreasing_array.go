package leetcode

/*
 * [665] Non-decreasing Array
 *
 * https://leetcode.com/problems/non-decreasing-array/description/
 *
 * algorithms
 * Easy (19.90%)
 * Total Accepted:    31.2K
 * Total Submissions: 156.5K
 * Testcase Example:  '[4,2,3]'
 *
 *
 * Given an array with n integers, your task is to check if it could become
 * non-decreasing by modifying at most 1 element.
 *
 *
 *
 * We define an array is non-decreasing if array[i]  holds for every i (1
 *
 * Example 1:
 *
 * Input: [4,2,3]
 * Output: True
 * Explanation: You could modify the first 4 to 1 to get a non-decreasing
 * array.
 *
 *
 *
 * Example 2:
 *
 * Input: [4,2,1]
 * Output: False
 * Explanation: You can't get a non-decreasing array by modify at most one
 * element.
 *
 *
 *
 * Note:
 * The n belongs to [1, 10,000].
 *
 */

/*

* 判断一个数组是否是`non-decreasing`，即最大只改一个数据，就可以达到`array[i] <= array[i + 1]`（665）
  * 首先判断逆序数，如果超过1，肯定false，如果是0，肯定是true
  * 对于逆序数有1个的地方
    * 如果这个逆序发生在数组的首、尾，肯定是true的
    * 因为是要求递增的，所以逆序发生的两个元素的
      * 左边那个小于这两个元素
      * 或者右边那个大于这两个元素
      那么是true

*/

func checkPossibility(nums []int) bool {
	n := 0
	i := -1

	for k := 1; k < len(nums); k++ {
		if nums[k] < nums[k-1] {
			n++
			i = k
		}
	}

	if n == 0 {
		return true
	} else if n > 1 {
		return false
	}

	if i == 1 || i == len(nums)-1 {
		return true
	}
	if nums[i+1] > nums[i-1] {
		return true
	}
	if nums[i-2] < nums[i] {
		return true
	}

	return false
}

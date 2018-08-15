package leetcode

import "fmt"

/*
 * [748] Largest Number At Least Twice of Others
 *
 * https://leetcode.com/problems/largest-number-at-least-twice-of-others/description/
 *
 * algorithms
 * Easy (40.44%)
 * Total Accepted:    26.1K
 * Total Submissions: 64.8K
 * Testcase Example:  '[0,0,0,1]'
 *
 * In a given integer array nums, there is always exactly one largest element.
 *
 * Find whether the largest element in the array is at least twice as much as
 * every other number in the array.
 *
 * If it is, return the index of the largest element, otherwise return -1.
 *
 * Example 1:
 *
 *
 * Input: nums = [3, 6, 1, 0]
 * Output: 1
 * Explanation: 6 is the largest integer, and for every other number in the
 * array x,
 * 6 is more than twice as big as x.  The index of value 6 is 1, so we return
 * 1.
 *
 *
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [1, 2, 3, 4]
 * Output: -1
 * Explanation: 4 isn't at least as big as twice the value of 3, so we return
 * -1.
 *
 * Note:
 *
 * nums will have a length in the range [1, 50].
 * Every nums[i] will be an integer in the range [0, 99].
 *
 */

/*

 给一个数字的数组，判断是否有一个数，是其他所有数字的2倍或者以上

 直接找出最大的数字和第二大的数字，然后比较即可

*/

func dominantIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	} else if len(nums) == 1 {
		return 0
	}
	max1index := -1
	max2index := -1
	for k, v := range nums {
		if max1index == -1 {
			max1index = k
		} else if max2index == -1 {
			max2index = k
			if nums[max1index] < nums[max2index] {
				max1index, max2index = max2index, max1index
			}
		} else {
			// 1比2大，先判断1
			if v > nums[max1index] {
				max2index = max1index
				max1index = k
			} else if v > nums[max2index] {
				max2index = k
			}
		}

	}

	fmt.Printf("%v %v\n", max1index, max2index)
	if nums[max1index] >= nums[max2index]*2 {
		return max1index
	}
	return -1
}

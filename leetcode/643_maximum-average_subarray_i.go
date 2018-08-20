package leetcode

/*
 * [643] Maximum Average Subarray I
 *
 * https://leetcode.com/problems/maximum-average-subarray-i/description/
 *
 * algorithms
 * Easy (37.81%)
 * Total Accepted:    34.7K
 * Total Submissions: 91.4K
 * Testcase Example:  '[1,12,-5,-6,50,3]\n4'
 *
 *
 * Given an array consisting of n integers, find the contiguous subarray of
 * given length k that has the maximum average value. And you need to output
 * the maximum average value.
 *
 *
 * Example 1:
 *
 * Input: [1,12,-5,-6,50,3], k = 4
 * Output: 12.75
 * Explanation: Maximum average is (12-5-6+50)/4 = 51/4 = 12.75
 *
 *
 *
 * Note:
 *
 * 1 k n
 * Elements of the given array will be in the range [-10,000, 10,000].
 *
 *
 */

/*

 给一个数组和一个个数k，求将数组中连续k个数组相加求平均数，求最大的平均数

 因为k固定，所以是求最大的和

 为了优化，可以用一个数组sums的i项存储nums的前i项和
 这样i-k到i的和就是sums[i] - sums[i-k]


*/

func findMaxAverage(nums []int, k int) float64 {
	var sums = make([]int, len(nums))

	sums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sums[i] = sums[i-1] + nums[i]
	}

	max_sum := sums[k-1]
	for i := k; i < len(nums); i++ {
		if he := sums[i] - sums[i-k]; he > max_sum {
			max_sum = he
		}
	}

	return float64(max_sum) / float64(k)
}

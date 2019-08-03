package leetcode

/*
 * @lc app=leetcode id=416 lang=golang
 *
 * [416] Partition Equal Subset Sum
 *
 * https://leetcode.com/problems/partition-equal-subset-sum/description/
 *
 * algorithms
 * Medium (40.86%)
 * Total Accepted:    96.7K
 * Total Submissions: 235K
 * Testcase Example:  '[1,5,11,5]'
 *
 * Given a non-empty array containing only positive integers, find if the array
 * can be partitioned into two subsets such that the sum of elements in both
 * subsets is equal.
 *
 * Note:
 *
 *
 * Each of the array element will not exceed 100.
 * The array size will not exceed 200.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [1, 5, 11, 5]
 *
 * Output: true
 *
 * Explanation: The array can be partitioned as [1, 5, 5] and [11].
 *
 *
 *
 *
 * Example 2:
 *
 *
 * Input: [1, 2, 3, 5]
 *
 * Output: false
 *
 * Explanation: The array cannot be partitioned into equal sum subsets.
 *
 *
 *
 *
 */
func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum&1 == 1 {
		return false
	}

	sum = sum >> 1
	n := len(nums)

	// states := make([n][w+1]bool)
	states := make([][]bool, n+1)
	for i := range states {
		states[i] = make([]bool, sum+1)
	}

	//for i := 0; i <= n+1; i++ {
	for i := 0; i < n+1; i++ {
		states[i][0] = true
	}

	for j := 1; j < sum+1; j++ {
		states[0][j] = false
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < sum+1; j++ {
			states[i][j] = states[i-1][j]
			if j >= nums[i-1] {
				states[i][j] = states[i][j] || states[i-1][j-nums[i-1]]
			}
		}
	}

	return states[n][sum]
}

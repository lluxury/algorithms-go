package leetcode

/*
 * [216] Combination Sum III
 *
 * https://leetcode.com/problems/combination-sum-iii/description/
 *
 * algorithms
 * Medium (48.36%)
 * Total Accepted:    97.2K
 * Total Submissions: 200.8K
 * Testcase Example:  '3\n7'
 *
 *
 * Find all possible combinations of k numbers that add up to a number n, given
 * that only numbers from 1 to 9 can be used and each combination should be a
 * unique set of numbers.
 *
 * Note:
 *
 *
 * All numbers will be positive integers.
 * The solution set must not contain duplicate combinations.
 *
 *
 * Example 1:
 *
 *
 * Input: k = 3, n = 7
 * Output: [[1,2,4]]
 *
 *
 * Example 2:
 *
 *
 * Input: k = 3, n = 9
 * Output: [[1,2,6], [1,3,5], [2,3,4]]
 *
 *
 *
 */

func combinationSum3_rec(k, shuru, n int, shuzi []int, tmp []int, result *[][]int) {
	if n == 0 && len(tmp) == k {
		*result = append(*result, tmp)
		return
	}

	if n < 0 || len(shuzi) == 0 {
		return
	}

	for kk, v := range shuzi {
		if v > shuru {
			continue
		}
		combinationSum3_rec(k, v, n-v, delete_int_slice_index(copy_int_slice(shuzi), kk), append([]int{v}, tmp...), result)
	}
}

func combinationSum3(k int, n int) [][]int {
	var result [][]int
	var shuzi = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for kk, v := range shuzi {
		combinationSum3_rec(k, v, n-v, delete_int_slice_index(copy_int_slice(shuzi), kk), []int{v}, &result)
	}

	return result
}

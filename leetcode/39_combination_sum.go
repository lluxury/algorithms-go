package leetcode

import "github.com/lluxury/algorithms-go/lib"

/*
 * [39] Combination Sum
 *
 * https://leetcode.com/problems/combination-sum/description/
 *
 * algorithms
 * Medium (43.00%)
 * Total Accepted:    243.3K
 * Total Submissions: 565K
 * Testcase Example:  '[2,3,6,7]\n7'
 *
 * Given a set of candidate numbers (candidates) (without duplicates) and a
 * target number (target), find all unique combinations in candidates where the
 * candidate numbers sums to target.
 *
 * The same repeated number may be chosen from candidates unlimited number of
 * times.
 *
 * Note:
 *
 *
 * All numbers (including target) will be positive integers.
 * The solution set must not contain duplicate combinations.
 *
 *
 * Example 1:
 *
 *
 * Input: candidates = [2,3,6,7], target = 7,
 * A solution set is:
 * [
 * ⁠ [7],
 * ⁠ [2,2,3]
 * ]
 *
 *
 * Example 2:
 *
 *
 * Input: candidates = [2,3,5], target = 8,
 * A solution set is:
 * [
 * [2,2,2,2],
 * [2,3,3],
 * [3,5]
 * ]
 *
 *
 */

/*

* 无序的
* 对于[a, b]来说，先选a再选b和先选b再选a，效果是一样的，所以需要一个游标start记录遍历开始的位置
* 如果是有序的，就不要游标

 */

func backtrackCombinationSum(list *[][]int, nums []int, tempList []int, target int, start int) {
	if target == 0 {
		*list = append(*list, lib.IntArrayDeepCopy(tempList))
		return
	}

	for i := start; i < len(nums); i++ {
		c := target - nums[i]
		if c < 0 {
			continue
		} else if c == 0 {

		}
		backtrackCombinationSum(list, nums, append(tempList, nums[i]), c, i)
	}
}

func combinationSum(candidates []int, target int) [][]int {
	var list = make([][]int, 0)
	backtrackCombinationSum(&list, candidates, []int{}, target, 0)
	return list
}

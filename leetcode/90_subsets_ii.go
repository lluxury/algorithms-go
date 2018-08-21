package leetcode

import (
	"github.com/Chyroc/algorithms-go/lib"
	"sort"
)

/*
 * [90] Subsets II
 *
 * https://leetcode.com/problems/subsets-ii/description/
 *
 * algorithms
 * Medium (39.36%)
 * Total Accepted:    160.6K
 * Total Submissions: 407.4K
 * Testcase Example:  '[1,2,2]'
 *
 * Given a collection of integers that might contain duplicates, nums, return
 * all possible subsets (the power set).
 * 
 * Note: The solution set must not contain duplicate subsets.
 * 
 * Example:
 * 
 * 
 * Input: [1,2,2]
 * Output:
 * [
 * ⁠ [2],
 * ⁠ [1],
 * ⁠ [1,2,2],
 * ⁠ [2,2],
 * ⁠ [1,2],
 * ⁠ []
 * ]
 * 
 */


/*

* 给一个重复元素数据集set，求这个set的所有子集（不能重复），如[1,2,1,3]返回[[] [3] [2] [3 2] [1] [3 1] [2 1] [3 2 1]]
  * 遍历递归
    * 基本数和78是一样的
    * 不过跳过：前面数据是相同的，但是后面在不同位置取了相同的数据
*/

func backtrackWithDup(list *[][]int, tempList []int, nums []int, start int) {
	*list = append(*list, lib.IntArrayDeepCopy(tempList))
	for i := start; i < len(nums); i++ {
		// skip duplicates
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		backtrackWithDup(list, append(tempList, nums[i]), nums, i+1)
	}
}

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)

	var list [][]int
	backtrackWithDup(&list, []int{}, nums, 0)
	return list
}

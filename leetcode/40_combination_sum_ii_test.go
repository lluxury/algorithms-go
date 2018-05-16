package leetcode

import (
	"github.com/Chyroc/algorithms-go/lib"
	"sort"
)

/*
> https://leetcode.com/problems/combination-sum/description/
> Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in candidates where the candidate numbers sums to target.

Each number in candidates may only be used once in the combination.

Note:

All numbers (including target) will be positive integers.
The solution set must not contain duplicate combinations.
Example 1:

Input: candidates = [10,1,2,7,6,1,5], target = 8,
A solution set is:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]
Example 2:

Input: candidates = [2,5,2,1,2], target = 5,
A solution set is:
[
  [1,2,2],
  [5]
]

* 和39一样
* 有重复数据
* 但是：不能使用重复元素（大小可以一直，但是不能是同一个位置的）
* 所以和47一样，用一个map计算一下是否使用过，或者传递一个used数据
*/

func backtrackCombinationSum2(list *[][]int, nums []int, tempList []int, target int, start int) {
	if target == 0 {
		*list = append(*list, lib.IntArrayDeepCopy(tempList))
		return
	}

	var numsMap = make(map[int]int)
	for _, v := range nums {
		if _, ok := numsMap[v]; ok {
			numsMap[v]++
		} else {
			numsMap[v] = 1
		}
	}
	for _, v2 := range tempList {
		numsMap[v2]--
	}

	for i := start; i < len(nums); i++ {
		if numsMap[nums[i]] <= 0 {
			continue // element already exists, skip
		}
		if i > start && nums[i] == nums[i-1] {
			continue
		}

		c := target - nums[i]
		if c < 0 {
			continue
		} else if c == 0 {

		}
		backtrackCombinationSum2(list, nums, append(tempList, nums[i]), c, i+1)
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	var list = make([][]int, 0)
	backtrackCombinationSum2(&list, candidates, []int{}, target, 0)
	return list
}

func Example_40_combinationSum2() {
	lib.ExamplePrint(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	lib.ExamplePrint(combinationSum2([]int{2, 5, 2, 1, 2}, 5))

	// Output:
	// [[1 1 6] [1 2 5] [1 7] [2 6]]
	// [[1 2 2] [5]]
}

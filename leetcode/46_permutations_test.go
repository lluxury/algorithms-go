package leetcode

import (
	"github.com/Chyroc/algorithms-go/lib"
	"github.com/Chyroc/algorithms-go/test"
	"testing"
)

/*
> https://leetcode.com/problems/permutations/description/
> Given a collection of distinct integers, return all possible permutations.

Input: [1,2,3]
Output:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]

* 给一个非重复元素数据集set，求这个set的排列的集合，如[1, 2, 3]返回[[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2] [3 2 1]]（46）
  * 和78是一样，遍历数组
  * 本集合是有顺序的
  * for v in nums: 只要不在临时数据的就加到临时数据进行下一轮递归
  * tempList长度等于nums长度返回

*/

func backtrackPermute(list *[][]int, tempList []int, nums []int) {
	if len(tempList) == len(nums) {
		*list = append(*list, lib.IntArrayDeepCopy(tempList))
		return
	}

	tempListMap := make(map[int]bool)
	for _, v := range tempList {
		tempListMap[v] = true
	}

	for _, v := range nums {
		if tempListMap[v] {
			continue // element already exists, skip
		}
		backtrackPermute(list, append(tempList, v), nums)
	}
}

func permute(nums []int) (list [][]int) {
	backtrackPermute(&list, []int{}, nums)
	return list
}

func Test_46(t *testing.T) {
	test.Runs(t, permuteUnique, []*test.Case{
		{Input: `[]`, Output: `[[]]`},
		{Input: `[1,2,3]`, Output: `[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]`},
		{Input: `[0,-1,1]`, Output: `[[-1, 0, 1], [-1, 1, 0], [0, -1, 1], [0, 1, -1], [1, -1, 0], [1, 0, -1]]`},
	})
}

package leetcode

import (
	"github.com/Chyroc/algorithms-go/lib"
	"github.com/Chyroc/algorithms-go/test"
	"sort"
	"testing"
)

/*
> https://leetcode.com/problems/permutations-ii/description/
> Given a collection of numbers that might contain duplicates, return all possible unique permutations.


Input: [1,1,2]
Output:
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]

* 给一个重复元素数据集set，求这个set的排列的集合，如[3 3 0 3]返回[[0 3 3 3] [3 0 3 3] [3 3 0 3] [3 3 3 0]]（47）
  * 和46是一样，遍历数组
  * 本集合是有有序的
  * 本集合是有重复的，所以46是map<int,bool>记录访问过的数据，这里需要先记录每个数据的个数map<int,int>，然后访问过的-1，这样大于0的就是46里面的map[v] == true的
  * for v in nums: 只要「不在临时数据的」 / 「不是重复数据」就加到临时数据进行下一轮递归
  * tempList长度等于nums长度返回
*/

func backtrackPermuteUnique(list *[][]int, tempList []int, nums []int) {
	if len(tempList) == len(nums) {
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

	for k, v := range nums {
		if numsMap[v] <= 0 {
			continue // element already exists, skip
		}
		if k > 0 && nums[k] == nums[k-1] {
			continue // skip duplicates
		}
		backtrackPermuteUnique(list, append(tempList, v), nums)
	}
}

func permuteUnique(nums []int) (list [][]int) {
	sort.Ints(nums)

	backtrackPermuteUnique(&list, []int{}, nums)
	return list
}

func Test_47(t *testing.T) {
	test.Runs(t, permuteUnique, []*test.Case{
		{Input: `[]`, Output: `[[]]`},
		{Input: `[1,2,3]`, Output: `[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]`},
		{Input: `[1, 1, 2]`, Output: `[[1,1,2],[1,2,1],[2,1,1]]`},
		{Input: `[3, 3, 0, 3]`, Output: `[[0,3,3,3],[3,0,3,3],[3,3,0,3],[3,3,3,0]]`},
	})
}

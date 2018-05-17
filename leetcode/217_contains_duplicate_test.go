package leetcode

/*
> https://leetcode.com/problems/contains-duplicate/description/

Given an array of integers, find if the array contains any duplicates. Your function should return true if any value appears at least twice in the array, and it should return false if every element is distinct.

- 问题
  - 判断一个数组是否有重复数据
- 思考
  - 哈希

*/
func containsDuplicate(nums []int) bool {
	var m = make(map[int]bool)

	for _, v := range nums {
		if _, ok := m[v]; ok {
			return true
		}

		m[v] = true
	}

	return false
}

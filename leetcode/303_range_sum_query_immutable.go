package leetcode

/*

> https://leetcode.com/problems/range-sum-query-immutable/description/


Given an integer array nums, find the sum of the elements between indices i and j (i ≤ j), inclusive.

Example:
Given nums = [-2, 0, 3, -5, 2, -1]

sumRange(0, 2) -> 1
sumRange(2, 5) -> -1
sumRange(0, 5) -> -3
Note:
You may assume that the array does not change.
There are many calls to sumRange function.


* 求一个数组的中，下标i和j之间的数的和，包含i和j
  * 题目本身不难，直接加起来就行
  * 但是题目是要求多次运行，能够比较节省时间，所以需要缓存结果或者换一个好一点的办法
  * i和j之间的数的和 = 前j数之和 - 前i数之和
*/
type NumArray struct {
	nums []int
}

func Constructor_303(nums []int) NumArray {
	var result = make([]int, len(nums))
	for k, v := range nums {
		if k == 0 {
			result[0] = v
		} else {
			result[k] = result[k-1] + v
		}
	}
	return NumArray{
		result,
	}
}

func (this *NumArray) SumRange(i int, j int) int {
	if i == 0 {
		return this.nums[j]
	}
	return this.nums[j] - this.nums[i-1]
}

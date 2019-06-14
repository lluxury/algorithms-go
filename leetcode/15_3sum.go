package leetcode

// import sort
import "sort"


/*
Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note:

The solution set must not contain duplicate triplets.

Example:

Given array nums = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]*/


// 给出3数之和为0, 可能不止一组解

func threeSum(nums []int) [][]int {
    var ret [][]int
    if len(nums) < 3 {
        return ret
    }
    sort.Ints(nums)
    for i := 0; i < len(nums); i++ {
        if i != 0 && nums[i] == nums[i-1] {
            continue
        }
        target := 0 - nums[i]
        for j := i + 1; j < len(nums); j++ {
            if j > i+1 && nums[j] == nums[j-1] {
                continue
            }
            hit := target - nums[j]
            if hit < nums[j] {
                continue
            }
            for k := j + 1; k < len(nums) && hit >= nums[k]; k++ {
                if hit == nums[k] {
                    tmp := []int{nums[i], nums[j], nums[k]}
                    ret = append(ret, tmp)
                    break
                }
            }
        }
    }
    return ret
}

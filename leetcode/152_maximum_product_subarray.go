package leetcode

import "math"
/*
Given an integer array nums, find the contiguous subarray within an array (containing at least one number) which has the largest product.

Example 1:

Input: [2,3,-2,4]
Output: 6
Explanation: [2,3] has the largest product 6.
Example 2:

Input: [-2,0,-1]
Output: 0
Explanation: The result cannot be 2, because [-2,-1] is not a subarray.*/

// 动态规划会欠很多债,慢慢还

func maxProduct(nums []int) int  {
    max1 := leftMaxProduct(nums)
    max2 := leftMaxProduct(reverse_152(nums))
    if max2 > max1 {
        return max2
    }
    return max1
}

func leftMaxProduct(nums []int) int  {
    max := math.MinInt32
    product := 1
    for i := 0; i < len(nums); i++ {
        product *= nums[i]
        if product > max {
            max = product
        }
        if product == 0 {
            product = 1
        }
    }
    return max
}

func reverse_152(nums []int) []int {
    res := make([]int, len(nums))
    for i := 0; i < len(nums); i++ {
        res[len(nums) -1 -i] = nums[i]
    }
    return res
}

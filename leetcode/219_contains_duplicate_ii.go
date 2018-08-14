package leetcode

import "fmt"

/*
 * [219] Contains Duplicate II
 *
 * https://leetcode.com/problems/contains-duplicate-ii/description/
 *
 * algorithms
 * Easy (33.25%)
 * Total Accepted:    156.5K
 * Total Submissions: 468.8K
 * Testcase Example:  '[1,2,3,1]\n3'
 *
 * Given an array of integers and an integer k, find out whether there are two
 * distinct indices i and j in the array such that nums[i] = nums[j] and the
 * absolute difference between i and j is at most k.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: nums = [1,2,3,1], k = 3
 * Output: true
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: nums = [1,0,1,1], k = 1
 * Output: true
 * 
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: nums = [1,2,3,1,2,3], k = 2
 * Output: false
 * 
 * 
 * 
 * 
 * 
 */

 /*

 有一个数组，判断连续n个数字有没有重复的

 维护一个长度为n的map即可

 */

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]bool)
	fmt.Printf("%#v\n", nums)

	for i, v := range nums {
		fmt.Printf("%#v, %v, %v\n", m, i, k)
		if len(m) > k {
			// 超过容量，删除最前面的
			delete(m, nums[i-k-1])
			fmt.Printf("delete %v\n",nums[i-k-1])
		}

		if m[v] {
			fmt.Printf("data: %#v, current: %v\n", m, v)
			return true
		}

		m[v] = true
	}

	return false
}

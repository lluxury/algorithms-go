package leetcode
import (
	"fmt"
	"math"
	"strings"
)

/*
 * [198] House Robber
 *
 * https://leetcode.com/problems/house-robber/description/
 *
 * algorithms
 * Easy (40.32%)
 * Total Accepted:    229K
 * Total Submissions: 567.8K
 * Testcase Example:  '[1,2,3,1]'
 *
 * You are a professional robber planning to rob houses along a street. Each
 * house has a certain amount of money stashed, the only constraint stopping
 * you from robbing each of them is that adjacent houses have security system
 * connected and it will automatically contact the police if two adjacent
 * houses were broken into on the same night.
 * 
 * Given a list of non-negative integers representing the amount of money of
 * each house, determine the maximum amount of money you can rob tonight
 * without alerting the police.
 * 
 * Example 1:
 * 
 * 
 * Input: [1,2,3,1]
 * Output: 4
 * Explanation: Rob house 1 (money = 1) and then rob house 3 (money =
 * 3).
 * Total amount you can rob = 1 + 3 = 4.
 * 
 * Example 2:
 * 
 * 
 * Input: [2,7,9,3,1]
 * Output: 12
 * Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house
 * 5 (money = 1).
 * Total amount you can rob = 2 + 9 + 1 = 12.
 * 
 * 
 */


/*

- 问题
  - 从一个数组中，取出若干元素，求其和最大值，要求不能有相邻的元素
- 思考
  - 这个问题和那个上楼梯需要多少步的问题一致
  - 先取第一个元素，那么结果就是这个元素加上第三个元素开始的后面的数组的最大值
  - 先取第二个，第四个开始
  - 先取第三个，是第一个分类的子问题

*/



var rob_d = make(map[string]int)

func rob(nums []int) int {

	numss := []string{}
	for _, v := range nums {
		numss = append(numss, fmt.Sprintf("%d", v))
	}
	numsS := strings.Join(numss, "")
	if v, ok := rob_d[numsS]; ok {
		return v
	}

	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		return int(math.Max(float64(nums[0]), float64(nums[1])))
	} else if len(nums) == 3 {
		return int(math.Max(float64(nums[0]+nums[2]), float64(nums[1])))
	}

	f1 := rob(nums[2:]) + nums[0]
	f2 := rob(nums[3:]) + nums[1]

	data := int(math.Max(float64(f1), float64(f2)))

	rob_d[numsS] = data

	return data
}

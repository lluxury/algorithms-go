package leetcode

/*
 * [645] Set Mismatch
 *
 * https://leetcode.com/problems/set-mismatch/description/
 *
 * algorithms
 * Easy (39.73%)
 * Total Accepted:    32.5K
 * Total Submissions: 81.7K
 * Testcase Example:  '[1,2,2,4]'
 *
 *
 * The set S originally contains numbers from 1 to n. But unfortunately, due to
 * the data error, one of the numbers in the set got duplicated to another
 * number in the set, which results in repetition of one number and loss of
 * another number.
 *
 *
 *
 * Given an array nums representing the data status of this set after the
 * error. Your task is to firstly find the number occurs twice and then find
 * the number that is missing. Return them in the form of an array.
 *
 *
 *
 * Example 1:
 *
 * Input: nums = [1,2,2,4]
 * Output: [2,3]
 *
 *
 *
 * Note:
 *
 * The given array size will in the range [2, 10000].
 * The given array's numbers won't have any order.
 *
 *
 */

/*

 给一个数组，里面是从1到n，其中有一个数字变成了另外那个数字，求重复的数字和缺失的数字

 * 遍历给定的数组，可以找到重复的数字
 * 然后遍历1-n，看看他出现的0次的数字是谁

*/

func findErrorNums(nums []int) []int {

	err := make([]int, 2)

	m := make(map[int]int)
	for _, v := range nums {
		m[v]++

		if m[v] == 2 {
			err[0] = v
		}
	}

	for k := 1; k < len(nums)+1; k++ {
		if m[k] == 0 {
			err[1] = k
		}
	}

	return err
}

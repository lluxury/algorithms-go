package leetcode

/*
 * [367] Valid Perfect Square
 *
 * https://leetcode.com/problems/valid-perfect-square/description/
 *
 * algorithms
 * Easy (38.81%)
 * Total Accepted:    80.9K
 * Total Submissions: 208.4K
 * Testcase Example:  '16'
 *
 * Given a positive integer num, write a function which returns True if num is
 * a perfect square else False.
 *
 *
 * Note: Do not use any built-in library function such as sqrt.
 *
 *
 * Example 1:
 *
 * Input: 16
 * Returns: True
 *
 *
 *
 * Example 2:
 *
 * Input: 14
 * Returns: False
 *
 *
 *
 * Credits:Special thanks to @elmirap for adding this problem and creating all
 * test cases.
 */

/*

 判断一个数是不是一个整数的平方

*/

func isPerfectSquare(num int) bool {
	var x float64 = 1
	var d int
	for {
		x = float64(x+float64(num)/x) / 2
		diff := int(x)*int(x) - num
		if diff < 0 {
			diff = -diff
		}

		if diff == 0 {
			return true
		}
		if diff == d {
			return false
		}
		d = diff
	}
}

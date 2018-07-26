package leetcode

/*
 * [342] Power of Four
 *
 * https://leetcode.com/problems/power-of-four/description/
 *
 * algorithms
 * Easy (39.42%)
 * Total Accepted:    91K
 * Total Submissions: 230.8K
 * Testcase Example:  '16'
 *
 *
 * Given an integer (signed 32 bits), write a function to check whether it is a
 * power of 4.
 *
 * Example:
 * Given num = 16, return true.
 * Given num = 5, return false.
 *
 *
 * Follow up: Could you solve it without loops/recursion?
 *
 * Credits:Special thanks to @yukuairoy  for adding this problem and creating
 * all test cases.
 */

func isPowerOfFour(n int) bool {
	return n > 0 && n&(n-1) == 0 && n&0x55555555 != 0
}

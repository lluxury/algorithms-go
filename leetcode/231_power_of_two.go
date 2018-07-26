package leetcode

/*
 * [231] Power of Two
 *
 * https://leetcode.com/problems/power-of-two/description/
 *
 * algorithms
 * Easy (40.99%)
 * Total Accepted:    182.2K
 * Total Submissions: 444.5K
 * Testcase Example:  '1'
 *
 * Given an integer, write a function to determine if it is a power of two.
 *
 * Example 1:
 *
 *
 * Input: 1
 * Output: true
 * Explanation: 20 = 1
 *
 *
 * Example 2:
 *
 *
 * Input: 16
 * Output: true
 * Explanation: 24 = 16
 *
 * Example 3:
 *
 *
 * Input: 218
 * Output: false
 *
 */

/*

 判断一个数是不是2的幂
 * 2的0次幂是1  ...     1   n-1是 0  ...    0
      1     2  ...    10         1  ...    1
      2     4  ...   100         3  ...   11
      3     8  ...  1000         7  ...  111
      4     16 ... 10000         15 ... 1111
 * n&(n-1) == 0

*/

func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}

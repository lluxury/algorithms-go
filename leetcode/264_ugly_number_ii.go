package leetcode

/*
 * [264] Ugly Number II
 *
 * https://leetcode.com/problems/ugly-number-ii/description/
 *
 * algorithms
 * Medium (33.95%)
 * Total Accepted:    82.2K
 * Total Submissions: 241.9K
 * Testcase Example:  '10'
 *
 * Write a program to find the n-th ugly number.
 * 
 * Ugly numbers are positive numbers whose prime factors only include 2, 3,
 * 5. 
 * 
 * Example:
 * 
 * 
 * Input: n = 10
 * Output: 12
 * Explanation: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 is the sequence of the first 10
 * ugly numbers.
 * 
 * Note:  
 * 
 * 
 * 1 is typically treated as an ugly number.
 * n does not exceed 1690.
 * 
 * 
 */

/*


* 求第n个263中的数字是多少
 * 数组的第一个元素是1
 * 分别乘以2，3，5，取最小值加入数组
 * 如果2乘了一次，那么就不能再乘以数组中的这个相同的数字了，所以需要三变量存储2，3，5分别需要与数组中的哪个数字相乘
 * 那么，第n个加入的就是要求的数字
*/

func min3(a, b, c int) int {
	if a > b {
		return min(b, c)
	} else {
		return min(a, c)
	}
}

func nthUglyNumber(n int) int {
	if n == 1 {
		return 1
	} else if n < 1 {
		return -1
	}

	var s = []int{1}

	var t2, t3, t5 int

	for i := 1; i < n; i++ {
		var (
			y2 = s[t2] * 2
			y3 = s[t3] * 3
			y5 = s[t5] * 5
		)

		m := min3(y2, y3, y5)
		s = append(s, m)

		if m == y2 {
			t2++
		}
		if m == y3 {
			t3++
		}
		if m == y5 {
			t5++
		}
	}

	return s[n-1]
}

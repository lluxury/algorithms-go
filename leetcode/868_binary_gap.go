package leetcode

/*
 * [899] Binary Gap
 *
 * https://leetcode.com/problems/binary-gap/description/
 *
 * algorithms
 * Easy (61.07%)
 * Total Accepted:    7.9K
 * Total Submissions: 13.1K
 * Testcase Example:  '22'
 *
 * Given a positive integer N, find and return the longest distance between two
 * consecutive 1's in the binary representation of N.
 * 
 * If there aren't two consecutive 1's, return 0.
 *
 * Example 1:
 * 
 * Input: 22
 * Output: 2
 * Explanation: 
 * 22 in binary is 0b10110.
 * In the binary representation of 22, there are three ones, and two
 * consecutive pairs of 1's.
 * The first consecutive pair of 1's have distance 2.
 * The second consecutive pair of 1's have distance 1.
 * The answer is the largest of these two distances, which is 2.
 *
 * Example 2:
 * 
 * 
 * Input: 5
 * Output: 2
 * Explanation: 
 * 5 in binary is 0b101.
 *
 * Example 3:
 * 
 * Input: 6
 * Output: 1
 * Explanation: 
 * 6 in binary is 0b110.
 * 
 * Example 4:
 *
 * Input: 8
 * Output: 0
 * Explanation: 
 * 8 in binary is 0b1000.
 * There aren't any consecutive pairs of 1's in the binary representation of 8,
 * so we return 0.
 * 
 * Note:
 * 
 * 1 <= N <= 10^9
 *
 */

 /*

 计算一个数字的二进制表示中，1和1之间距离的最大值（这两个1之间没有别的1）

 先转化为二进制，然后求最大值

 */

func i10_i2(n int) []int {
	s := []int{}
	for {
		a := n % 2
		n = (n - a) / 2

		s = append(s, a)
		if n == 0 {
			break
		}
	}

	for i, j := 0, len(s)-1; i < j; {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}

	return s
}

func binaryGap(N int) int {
	s := i10_i2(N)

	preindex := -1
	max := -1
	for k, v := range s {
		if v == 1 {
			if preindex != -1 {
				if k-preindex > max {
					max = k - preindex
				}
			}
			preindex = k
		}
	}

	if max == -1 {
		return 0
	}
	return max
}

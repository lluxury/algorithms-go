package leetcode

/*
 * [841] Shortest Distance to a Character
 *
 * https://leetcode.com/problems/shortest-distance-to-a-character/description/
 *
 * algorithms
 * Easy (60.89%)
 * Total Accepted:    14.1K
 * Total Submissions: 23.1K
 * Testcase Example:  '"loveleetcode"\n"e"'
 *
 * Given a string S and a character C, return an array of integers representing
 * the shortest distance from the character C in the string.
 *
 * Example 1:
 *
 *
 * Input: S = "loveleetcode", C = 'e'
 * Output: [3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0]
 *
 *
 *
 *
 * Note:
 *
 *
 * S string length is in [1, 10000].
 * C is a single character, and guaranteed to be in string S.
 * All letters in S and C are lowercase.
 *
 *
 */

/*

给一个字符串，和一个字符，求字符串各个位置到距离他最近的这个字符的距离
* 遍历字符串
* 记录相同的字符的位置
* 两个字符中间的字符到他们的位置是：1 2 3 4 5和5 4 3 2 1的最小值，即 1 2 3 2 1

*/

func shortestToChar(S string, C byte) []int {
	var sum = make([]int, len(S))

	lastCIndex := -1
	for k, v := range S {
		if byte(v) == C {
			if lastCIndex == -1 {
				for j := 0; j < k; j++ {
					sum[j] = k - j
				}
			} else {
				for j := lastCIndex + 1; j < k; j++ {
					x := j - lastCIndex
					y := k - j
					if x < y {
						sum[j] = x
					} else {
						sum[j] = y
					}
				}
			}

			lastCIndex = k
			continue
		}
	}

	for j := lastCIndex + 1; j <= len(sum)-1; j++ {
		sum[j] = j - lastCIndex
	}

	return sum
}

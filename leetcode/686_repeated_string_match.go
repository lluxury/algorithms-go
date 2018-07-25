package leetcode

import (
	"bytes"
	"strings"
)

/*
 * [686] Repeated String Match
 *
 * https://leetcode.com/problems/repeated-string-match/description/
 *
 * algorithms
 * Easy (31.48%)
 * Total Accepted:    35.5K
 * Total Submissions: 112.7K
 * Testcase Example:  '"abcd"\n"cdabcdab"'
 *
 * Given two strings A and B, find the minimum number of times A has to be
 * repeated such that B is a substring of it. If no such solution, return -1.
 *
 *
 * For example, with A = "abcd" and B = "cdabcdab".
 *
 *
 * Return 3, because by repeating A three times (“abcdabcdabcd”), B is a
 * substring of it; and B is not a substring of A repeated two times
 * ("abcdabcd").
 *
 *
 * Note:
 * The length of A and B will be between 1 and 10000.
 *
 */

/*
 给两个字符串a和b，问求最少重复a几次，b会是其子串，没有返回-1
   * 首先求b比a的倍数
   * 重复的次数 最少 要能够让长度覆盖b，然后判断能否覆盖
   * 如果不能重复一次，再判断一次
   * 再多重复就没有意义了，返回-1
*/

func repeatedStringMatch(A string, B string) int {
	l := len(B) / len(A)

	buf := new(bytes.Buffer)

	for k := 0; k <= l+1; k++ {
		buf.WriteString(A)

		if strings.Contains(buf.String(), B) {
			return k + 1
		}
	}
	return -1
}

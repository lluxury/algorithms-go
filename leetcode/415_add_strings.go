package leetcode

import (
	"strconv"
)

/*
 * [415] Add Strings
 *
 * https://leetcode.com/problems/add-strings/description/
 *
 * algorithms
 * Easy (41.78%)
 * Total Accepted:    62.5K
 * Total Submissions: 149.6K
 * Testcase Example:  '"0"\n"0"'
 *
 * Given two non-negative integers num1 and num2 represented as string, return
 * the sum of num1 and num2.
 *
 * Note:
 *
 * The length of both num1 and num2 is < 5100.
 * Both num1 and num2 contains only digits 0-9.
 * Both num1 and num2 does not contain any leading zero.
 * You must not use any built-in BigInteger library or convert the inputs to
 * integer directly.
 *
 *
 */

/*

 给两个数字的字符串，将他们加起来，返回结果的字符串
   * 从后往前遍历
   * 有可能两个都有值，有可能只有第一个有值，有可能只有第二个有值
   * -'a'之后就是数字的值
   * 用一个变量存储进位

*/

func addStrings(num1 string, num2 string) string {
	s := ""
	var jinwei = 0
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0; {
		data := 0
		if i >= 0 && j >= 0 {
			data = int(num1[i]) - '0' + int(num2[j]) - '0' + jinwei
		} else if i >= 0 {
			data = int(num1[i]) - '0' + jinwei
		} else if j >= 0 {
			data = int(num2[j]) - '0' + jinwei
		} else {
			break
		}

		y := data % 10
		s = strconv.Itoa(y) + s
		jinwei = (data - y) / 10

		if i >= 0 {
			i--
		}
		if j >= 0 {
			j--
		}
	}
	if jinwei > 0 {
		s = strconv.Itoa(jinwei) + s
	}

	return s
}

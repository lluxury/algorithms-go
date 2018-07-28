package leetcode

/*
 * [8] String to Integer (atoi)
 *
 * https://leetcode.com/problems/string-to-integer-atoi/description/
 *
 * algorithms
 * Medium (14.11%)
 * Total Accepted:    250.1K
 * Total Submissions: 1.8M
 * Testcase Example:  '"42"'
 *
 * Implement atoi which converts a string to an integer.
 *
 * The function first discards as many whitespace characters as necessary until
 * the first non-whitespace character is found. Then, starting from this
 * character, takes an optional initial plus or minus sign followed by as many
 * numerical digits as possible, and interprets them as a numerical value.
 *
 * The string can contain additional characters after those that form the
 * integral number, which are ignored and have no effect on the behavior of
 * this function.
 *
 * If the first sequence of non-whitespace characters in str is not a valid
 * integral number, or if no such sequence exists because either str is empty
 * or it contains only whitespace characters, no conversion is performed.
 *
 * If no valid conversion could be performed, a zero value is returned.
 *
 * Note:
 *
 *
 * Only the space character ' ' is considered as whitespace character.
 * Assume we are dealing with an environment which could only store integers
 * within the 32-bit signed integer range: [−231,  231 − 1]. If the numerical
 * value is out of the range of representable values, INT_MAX (231 − 1) or
 * INT_MIN (−231) is returned.
 *
 *
 * Example 1:
 *
 *
 * Input: "42"
 * Output: 42
 *
 *
 * Example 2:
 *
 *
 * Input: "   -42"
 * Output: -42
 * Explanation: The first non-whitespace character is '-', which is the minus
 * sign.
 * Then take as many numerical digits as possible, which gets 42.
 *
 *
 * Example 3:
 *
 *
 * Input: "4193 with words"
 * Output: 4193
 * Explanation: Conversion stops at digit '3' as the next character is not a
 * numerical digit.
 *
 *
 * Example 4:
 *
 *
 * Input: "words and 987"
 * Output: 0
 * Explanation: The first non-whitespace character is 'w', which is not a
 * numerical
 * digit or a +/- sign. Therefore no valid conversion could be performed.
 *
 * Example 5:
 *
 *
 * Input: "-91283472332"
 * Output: -2147483648
 * Explanation: The number "-91283472332" is out of the range of a 32-bit
 * signed integer.
 * Thefore INT_MIN (−231) is returned.
 *
 */

/*

 将一个字符串转成数字，不能以非法字符开头，结尾的非法字符忽略，超过int32范围会截断
 * 遍历，先取数字开始flag，在那之前，允许非法字符，在那之后，不允许
 * 判断正数还是负数
 * 将数字加到结果
 * 判断int32

*/

func overint32(zhen bool, s int) (bool, int) {
	if zhen {
		if s > 2147483647 {
			return true, 2147483647
		}
		return false, s
	} else {
		if -s < -2147483648 {
			return true, -2147483648
		}
		return false, -s
	}
}

func myAtoi(str string) int {
	start := false
	zhen := true
	s := 0
	for k := 0; k < len(str); k++ {

		if !start {
			if str[k] == ' ' {
				continue
			}
			if str[k] == '-' && k+1 < len(str) {
				zhen = false
			}
			if str[k] == '-' || str[k] == '+' || isDigit(str[k]) {
				start = true
				if isDigit(str[k]) {
					s = s*10 + toDigit(str[k])
				}
				continue
			}
			break
		} else {
			over, news := overint32(zhen, s)
			if over {
				return news
			}

			if isDigit(str[k]) {
				s = s*10 + toDigit(str[k])
			} else {
				break
			}
		}
	}

	_, news := overint32(zhen, s)
	return news
}

package leetcode

/*
 * [742] To Lower Case
 *
 * https://leetcode.com/problems/to-lower-case/description/
 *
 * algorithms
 * Easy (75.33%)
 * Total Accepted:    9.9K
 * Total Submissions: 13.1K
 * Testcase Example:  '"Hello"'
 *
 * Implement function ToLowerCase() that has a string parameter str, and
 * returns the same string in lowercase.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: "Hello"
 * Output: "hello"
 *
 *
 *
 * Example 2:
 *
 *
 * Input: "here"
 * Output: "here"
 *
 *
 *
 * Example 3:
 *
 *
 * Input: "LOVELY"
 * Output: "lovely"
 *
 *
 *
 *
 *
 */

func toLowerCase(str string) string {
	var x []uint8
	for k := range str {
		if str[k] >= 'A' && str[k] <= 'Z' {
			x = append(x, str[k]+uint8(32))
		} else {
			x = append(x, str[k])
		}
	}
	return string(x)
}

package leetcode

/*
 * [680] Valid Palindrome II
 *
 * https://leetcode.com/problems/valid-palindrome-ii/description/
 *
 * algorithms
 * Easy (32.32%)
 * Total Accepted:    38.2K
 * Total Submissions: 118.1K
 * Testcase Example:  '"aba"'
 *
 *
 * Given a non-empty string s, you may delete at most one character.  Judge
 * whether you can make it a palindrome.
 *
 *
 * Example 1:
 *
 * Input: "aba"
 * Output: True
 *
 *
 *
 * Example 2:
 *
 * Input: "abca"
 * Output: True
 * Explanation: You could delete the character 'c'.
 *
 *
 *
 * Note:
 *
 * The string will only contain lowercase characters a-z.
 * The maximum length of the string is 50000.
 *
 *
 */

/*

 对一个字符串，最多删除一个字符，判断是不是回文字符串

 如果是判断回文的话：从头和尾向中间遍历即可

 如果还需要可以删除一个字符的话，可以视为头进，尾不进；或者尾进，头不进

 只能删除一个，所以用rec的bool值控制

 如果可以删除多个，rec换成数字

*/

func validPalindrome_rec(s string, l, r int, rec bool) bool {
	for i, j := l, r; i < j; {
		if s[i] == s[j] {
			i++
			j--
			continue
		}

		if rec {
			return validPalindrome_rec(s, i+1, j, false) || validPalindrome_rec(s, i, j-1, false)
		}

		return false
	}

	return true
}

func validPalindrome(s string) bool {
	return validPalindrome_rec(s, 0, len(s)-1, true)
}

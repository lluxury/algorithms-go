package leetcode

/*
 * [409] Longest Palindrome
 *
 * https://leetcode.com/problems/longest-palindrome/description/
 *
 * algorithms
 * Easy (46.19%)
 * Total Accepted:    70.3K
 * Total Submissions: 152.2K
 * Testcase Example:  '"abccccdd"'
 *
 * Given a string which consists of lowercase or uppercase letters, find the
 * length of the longest palindromes that can be built with those letters.
 *
 * This is case sensitive, for example "Aa" is not considered a palindrome
 * here.
 *
 * Note:
 * Assume the length of given string will not exceed 1,010.
 *
 *
 * Example:
 *
 * Input:
 * "abccccdd"
 *
 * Output:
 * 7
 *
 * Explanation:
 * One longest palindrome that can be built is "dccaccd", whose length is 7.
 *
 *
 */

/*

 给一个字符串，问这些字符能够组成的最长的回文串多长
   * 回文字符串就是对称的字符串
   * 所以先统计字符串各个字符的个数
   * 如果是偶数，左右分布，加x
   * 如果是奇数，减去1个，然后左右分布，加x
   * 如果有奇数，还应该加一个放正中间

*/

func longestPalindrome(s string) int {
	var m = make(map[uint8]int)
	for k := range s {
		m[s[k]]++
	}

	x := 0
	jishu := true
	for _, v := range m {
		if v%2 == 0 {
			x += v
		} else {
			if jishu {
				x += 1
				jishu = false
			}

			x += v - 1

		}
	}
	return x
}

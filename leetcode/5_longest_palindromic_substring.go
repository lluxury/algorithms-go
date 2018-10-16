package leetcode

/*
 * [5] Longest Palindromic Substring
 *
 * https://leetcode.com/problems/longest-palindromic-substring/description/
 *
 * algorithms
 * Medium (25.64%)
 * Total Accepted:    357.5K
 * Total Submissions: 1.4M
 * Testcase Example:  '"babad"'
 *
 * Given a string s, find the longest palindromic substring in s. You may
 * assume that the maximum length of s is 1000.
 *
 * Example 1:
 *
 *
 * Input: "babad"
 * Output: "bab"
 * Note: "aba" is also a valid answer.
 *
 *
 * Example 2:
 *
 *
 * Input: "cbbd"
 * Output: "bb"
 *
 *
 */

/*

 * 求一个字符串的最大子回文串
  * 指定字符串的任意一个位置为回文字符串的中点（遍历）
    * 长度可能是奇数也有可能是偶数，所以中点可能是一个元素（奇数aba），也有可能是重合的一个元素（偶数abba）
  * 从中点向两边出发，如果字符相等，那么长度加1
  * 保存最大长度的起点index和长度

*/

func longestPalindrome_ref(s string, start, maxLen *int, left, right int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}

	if right-left-1 > *maxLen {
		// 如果首尾之间的间隔大于最大长度，更新这个值
		*maxLen = right - (left + 1)
		*start = left + 1
	}
}

func longestPalindrome_5(s string) string {
	if len(s) < 2 {
		return s
	}
	var start, maxLen int
	for i := 0; i < len(s)-1; i++ {
		// 以x为中心，向外扩展，检查每一次是不是字符都相等，x从 0 到 n-1
		longestPalindrome_ref(s, &start, &maxLen, i, i)
		// 以x, x+1 为中心，向外扩展，检查每一次是不是字符都相等，x从 0 到 n-1
		longestPalindrome_ref(s, &start, &maxLen, i, i+1)
	}

	return s[start : start+maxLen]
}

package leetcode

import (
	"strings"
)

/*
 * [459] Repeated Substring Pattern
 *
 * https://leetcode.com/problems/repeated-substring-pattern/description/
 *
 * algorithms
 * Easy (38.43%)
 * Total Accepted:    58.2K
 * Total Submissions: 151.5K
 * Testcase Example:  '"abab"'
 *
 * Given a non-empty string check if it can be constructed by taking a
 * substring of it and appending multiple copies of the substring together.
 * You may assume the given string consists of lowercase English letters only
 * and its length  will not exceed 10000. 
 * 
 * Example 1:
 * 
 * Input: "abab"
 * 
 * Output: True
 * 
 * Explanation: It's the substring "ab" twice.
 * 
 * 
 * 
 * Example 2:
 * 
 * Input: "aba"
 * 
 * Output: False
 * 
 * 
 * 
 * Example 3:
 * 
 * Input: "abcabcabcabc"
 * 
 * Output: True
 * 
 * Explanation: It's the substring "abc" four times. (And the substring
 * "abcabc" twice.)
 * 
 * 
 */


 /*

 给定一个字符串，判定其是否可以由子串重复而组成
   * 首先，肯定从第一个字符串开始重复
   * 所以从第一个位置取若干字符串，然后和下一个相同长度的字符串进行比较
   * 如果相等，则递归
   * 如果不相等，则将子串长度加一
   * 因为开始位置肯定是0，所以只需要记录子串的结尾位置
 */

func repeatedSubstringPattern_ref(s, p string) bool {
	for {
		if p == "" {
			return true
		}
		if !strings.HasPrefix(p, s) {
			return false
		}
		p = strings.TrimPrefix(p, s)
	}
}

func repeatedSubstringPattern(s string) bool {
	var length = 1
	for {
		if length+length > len(s) {
			return false
		}

		r := s[0:length]
		next := s[length : length+length]
		left := s[length:]

		if r == next {
			if repeatedSubstringPattern_ref(r, left) {
				return true
			}
		}
		length++
	}
	return false
}

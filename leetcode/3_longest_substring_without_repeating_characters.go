package leetcode

/*
 * [3] Longest Substring Without Repeating Characters
 *
 * https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
 *
 * algorithms
 * Medium (24.92%)
 * Total Accepted:    551.6K
 * Total Submissions: 2.2M
 * Testcase Example:  '"abcabcbb"'
 *
 * Given a string, find the length of the longest substring without repeating
 * characters.
 *
 *
 * Example 1:
 *
 *
 * Input: "abcabcbb"
 * Output: 3
 * Explanation: The answer is "abc", which the length is 3.
 *
 *
 *
 * Example 2:
 *
 *
 * Input: "bbbbb"
 * Output: 1
 * Explanation: The answer is "b", with the length of 1.
 *
 *
 *
 * Example 3:
 *
 *
 * Input: "pwwkew"
 * Output: 3
 * Explanation: The answer is "wke", with the length of 3.
 * ⁠            Note that the answer must be a substring, "pwke" is a
 * subsequence and not a substring.
 *
 *
 */

/*

 非重复字符最长子串查找，输出子串长度

 思路

- 搞一个i,j分别指向最长字符串的首尾，然后将i和j之间的字符存在一个hashTable里面
- 开始的时候字符串长度为0，所以i和j都是0
- 然后将字符串依次放进map（如果这个字符串以前没有碰到的话），j因为指向字符串的尾部，自增
- 如果出现字符串重复，将头部指向的字符串移除map，然后位置后移，i自增
- j和i的差值的最大值就是结果


*/

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]bool)
	l := 0

	for i, j := 0, 0; i < len(s) && j < len(s); {
		if exist, ok := m[s[j]]; ok && exist {
			m[s[i]] = false
			i++
		} else {
			m[s[j]] = true
			j++

			if j-i > l {
				l = j - i
			}
		}
	}

	return l
}

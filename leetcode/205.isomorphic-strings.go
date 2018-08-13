package leetcode

/*
 * [205] Isomorphic Strings
 *
 * https://leetcode.com/problems/isomorphic-strings/description/
 *
 * algorithms
 * Easy (35.18%)
 * Total Accepted:    150.3K
 * Total Submissions: 426.2K
 * Testcase Example:  '"egg"\n"add"'
 *
 * Given two strings s and t, determine if they are isomorphic.
 *
 * Two strings are isomorphic if the characters in s can be replaced to get t.
 *
 * All occurrences of a character must be replaced with another character while
 * preserving the order of characters. No two characters may map to the same
 * character but a character may map to itself.
 *
 * Example 1:
 *
 *
 * Input: s = "egg", t = "add"
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: s = "foo", t = "bar"
 * Output: false
 *
 * Example 3:
 *
 *
 * Input: s = "paper", t = "title"
 * Output: true
 *
 * Note:
 * You may assume both s and t have the same length.
 *
 */

/*

给两个长度一样的字符串，判断他们是不是isomorphic的，也就是将相同的字母换成其他相同的字母，两个字符串能否变成一样

1. 对于a字符串中出现过的字符x，那么对于a中其他x字符的位置，第二个字符串的相同位置的字符应该是一样的
2. 需要做两个遍历

*/

func isIsomorphic_rec(s string, t string, rec bool) bool {
	diff := make(map[uint8]uint8, len(s))

	for k := 0; k < len(s); k++ {
		if d, ok := diff[s[k]]; ok {
			// 已经有了，那么这次的diff值需要一致
			if d != s[k]-t[k] {
				return false
			}
		} else {
			// 还没有，保存该位置的字母的差值
			diff[s[k]] = s[k] - t[k]
		}
	}
	return rec || isIsomorphic_rec(t, s, true)
}

func isIsomorphic(s string, t string) bool {
	return isIsomorphic_rec(s, t, false)
}

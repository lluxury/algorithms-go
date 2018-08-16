package leetcode

import (
	"strings"
)

/*
 * [290] Word Pattern
 *
 * https://leetcode.com/problems/word-pattern/description/
 *
 * algorithms
 * Easy (33.76%)
 * Total Accepted:    112.4K
 * Total Submissions: 332.6K
 * Testcase Example:  '"abba"\n"dog cat cat dog"'
 *
 * Given a pattern and a string str, find if str follows the same pattern.
 *
 * Here follow means a full match, such that there is a bijection between a
 * letter in pattern and a non-empty word in str.
 *
 * Example 1:
 *
 *
 * Input: pattern = "abba", str = "dog cat cat dog"
 * Output: true
 *
 * Example 2:
 *
 *
 * Input:pattern = "abba", str = "dog cat cat fish"
 * Output: false
 *
 * Example 3:
 *
 *
 * Input: pattern = "aaaa", str = "dog cat cat dog"
 * Output: false
 *
 * Example 4:
 *
 *
 * Input: pattern = "abba", str = "dog dog dog dog"
 * Output: false
 *
 * Notes:
 * You may assume pattern contains only lowercase letters, and str contains
 * lowercase letters separated by a single space.
 *
 */

/*

 和 205 题差不多，只是第二个字符串换成空格分割的多字符相连的字符串

*/

func wordPattern_rec(pattern []string, str []string, rec bool) bool {
	d := make(map[string][]int)
	for k := range pattern {
		if _, ok := d[pattern[k]]; ok {
			d[pattern[k]] = append(d[pattern[k]], k)
		} else {
			d[pattern[k]] = []int{k}
		}
	}

	for _, v := range d {
		value := str[v[0]]
		for k := 1; k < len(v); k++ {
			if str[v[k]] != value {
				return false
			}
		}
	}

	return rec || wordPattern_rec(str, pattern, true)
}

func wordPattern(pattern string, str string) bool {
	a, b := strings.Split(pattern, ""), strings.Split(str, " ")
	if len(a) != len(b) {
		return false
	}
	return wordPattern_rec(a, b, false)
}

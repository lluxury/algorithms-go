package leetcode

import (
	"sort"
)

/*
 * [720] Longest Word in Dictionary
 *
 * https://leetcode.com/problems/longest-word-in-dictionary/description/
 *
 * algorithms
 * Easy (41.61%)
 * Total Accepted:    18.3K
 * Total Submissions: 43.8K
 * Testcase Example:  '["w","wo","wor","worl","world"]'
 *
 * Given a list of strings words representing an English Dictionary, find the
 * longest word in words that can be built one character at a time by other
 * words in words.  If there is more than one possible answer, return the
 * longest word with the smallest lexicographical order.  If there is no
 * answer, return the empty string.
 *
 * Example 1:
 *
 * Input:
 * words = ["w","wo","wor","worl", "world"]
 * Output: "world"
 * Explanation:
 * The word "world" can be built one character at a time by "w", "wo", "wor",
 * and "worl".
 *
 *
 *
 * Example 2:
 *
 * Input:
 * words = ["a", "banana", "app", "appl", "ap", "apply", "apple"]
 * Output: "apple"
 * Explanation:
 * Both "apply" and "apple" can be built from other words in the dictionary.
 * However, "apple" is lexicographically smaller than "apply".
 *
 *
 *
 * Note:
 * All the strings in the input will only contain lowercase letters.
 * The length of words will be in the range [1, 1000].
 * The length of words[i] will be in the range [1, 30].
 *
 */

/*

 给一个字符串数组，返回最长的一个字符串，这个字符串所有的前缀字符串都需要在输入数组中（前缀树）

 * 给数组排序，让长度小的，字母顺序小的靠前
 * 遍历数组
   * 如果除了最后一个字符外，其他的已经出现过，加入map
   * 如果长度大于最大长度，记录下来
   * 返回最大的

*/

type sortstring720 struct {
	words []string
}

func (r *sortstring720) Len() int {
	return len(r.words)
}

func wordLess(a, b string) bool {
	if len(a) != len(b) {
		return len(a) < len(b)
	}
	for k := range a {
		if a[k] != b[k] {
			return a[k] < b[k]
		}
	}
	return true
}

func (r *sortstring720) Less(i, j int) bool {
	return wordLess(r.words[i], r.words[j])
}

func (r *sortstring720) Swap(i, j int) {
	r.words[i], r.words[j] = r.words[j], r.words[i]
}

func longestWord(words []string) string {
	x := &sortstring720{words: words}
	sort.Sort(x)

	exist := make(map[string]bool)

	s := ""
	for _, word := range words {
		if len(word) == 1 {
			if len(word) > len(s) {
				s = word
			}
			exist[word] = true
		} else {
			if exist[word[:len(word)-1]] {
				if len(word) > len(s) {
					s = word
				}
				exist[word] = true
			}
		}
	}

	return s
}

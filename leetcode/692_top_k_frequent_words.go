package leetcode

import "sort"

/*
 * [692] Top K Frequent Words
 *
 * https://leetcode.com/problems/top-k-frequent-words/description/
 *
 * algorithms
 * Medium (41.96%)
 * Total Accepted:    27.9K
 * Total Submissions: 66.3K
 * Testcase Example:  '["i", "love", "leetcode", "i", "love", "coding"]\n2'
 *
 * Given a non-empty list of words, return the k most frequent elements.
 * Your answer should be sorted by frequency from highest to lowest. If two
 * words have the same frequency, then the word with the lower alphabetical
 * order comes first.
 *
 * Example 1:
 *
 * Input: ["i", "love", "leetcode", "i", "love", "coding"], k = 2
 * Output: ["i", "love"]
 * Explanation: "i" and "love" are the two most frequent words.
 * ⁠   Note that "i" comes before "love" due to a lower alphabetical order.
 *
 *
 *
 * Example 2:
 *
 * Input: ["the", "day", "is", "sunny", "the", "the", "the", "sunny", "is",
 * "is"], k = 4
 * Output: ["the", "is", "sunny", "day"]
 * Explanation: "the", "is", "sunny" and "day" are the four most frequent
 * words,
 * ⁠   with the number of occurrence being 4, 3, 2 and 1 respectively.
 *
 *
 *
 * Note:
 *
 * You may assume k is always valid, 1 ≤ k ≤ number of unique elements.
 * Input words contain only lowercase letters.
 *
 *
 *
 * Follow up:
 *
 * Try to solve it in O(n log k) time and O(n) extra space.
 *
 *
 */

/*


给一个字符串数组，求n个出现次数最多的字符串
  * 统计每个字符串出现的次数
  * 按照次数和字符串顺序排序
  * 返回前n个字符串
*/

type indexSort struct {
	Index int
	Str   string
}
type indexSorts []indexSort

func (r indexSorts) Len() int { return len(r) }

func (r indexSorts) Less(i, j int) bool {
	if r[i].Index != r[j].Index {
		return r[i].Index > r[j].Index
	}

	return r[i].Str < r[j].Str
}

func (r indexSorts) Swap(i, j int) { r[i], r[j] = r[j], r[i] }

func topKFrequent(words []string, k int) []string {
	var m = make(map[string]int)
	for _, v := range words {
		m[v]++
	}

	var ss indexSorts
	for k, v := range m {
		ss = append(ss, indexSort{Index: v, Str: k})
	}
	sort.Sort(ss)

	var re []string
	for i := 0; i < k; i++ {
		re = append(re, ss[i].Str)
	}

	return re
}

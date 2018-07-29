package leetcode

import "sort"

/*
 * [451] Sort Characters By Frequency
 *
 * https://leetcode.com/problems/sort-characters-by-frequency/description/
 *
 * algorithms
 * Medium (52.39%)
 * Total Accepted:    60.3K
 * Total Submissions: 115K
 * Testcase Example:  '"tree"'
 *
 * Given a string, sort it in decreasing order based on the frequency of
 * characters.
 *
 * Example 1:
 *
 * Input:
 * "tree"
 *
 * Output:
 * "eert"
 *
 * Explanation:
 * 'e' appears twice while 'r' and 't' both appear once.
 * So 'e' must appear before both 'r' and 't'. Therefore "eetr" is also a valid
 * answer.
 *
 *
 *
 * Example 2:
 *
 * Input:
 * "cccaaa"
 *
 * Output:
 * "cccaaa"
 *
 * Explanation:
 * Both 'c' and 'a' appear three times, so "aaaccc" is also a valid answer.
 * Note that "cacaca" is incorrect, as the same characters must be together.
 *
 *
 *
 * Example 3:
 *
 * Input:
 * "Aabb"
 *
 * Output:
 * "bbAa"
 *
 * Explanation:
 * "bbaA" is also a valid answer, but "Aabb" is incorrect.
 * Note that 'A' and 'a' are treated as two different characters.
 *
 *
 */

/*

 给一个字符串，按照字符出现的顺序排序，次数多的排在前面
   * 先统计个数
   * 排序
   * 返回结果

*/

type frequencySort_sort struct {
	f []frequencySort_sort_single
}

func (r *frequencySort_sort) Len() int {
	return len(r.f)
}

func (r *frequencySort_sort) Less(i, j int) bool {
	return r.f[i].count > r.f[j].count
}

func (r *frequencySort_sort) Swap(i, j int) {
	r.f[i], r.f[j] = r.f[j], r.f[i]
}

type frequencySort_sort_single struct {
	k     uint8
	count int
}

func frequencySort(s string) string {
	m := make(map[uint8]int)
	for k := range s {
		m[s[k]]++
	}

	so := &frequencySort_sort{}
	for k, v := range m {
		so.f = append(so.f, frequencySort_sort_single{k: k, count: v})
	}

	sort.Sort(so)

	var x []uint8
	for _, f := range so.f {
		for i := 0; i < f.count; i++ {
			x = append(x, f.k)
		}
	}
	return string(x)
}

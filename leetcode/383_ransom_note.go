package leetcode

/*
 * [383] Ransom Note
 *
 * https://leetcode.com/problems/ransom-note/description/
 *
 * algorithms
 * Easy (48.07%)
 * Total Accepted:    87.3K
 * Total Submissions: 181.2K
 * Testcase Example:  '"a"\n"b"'
 *
 *
 * Given an arbitrary ransom note string and another string containing letters
 * from all the magazines, write a function that will return true if the
 * ransom
 * note can be constructed from the magazines ; otherwise, it will return
 * false.
 *
 *
 * Each letter in the magazine string can only be used once in your ransom
 * note.
 *
 *
 * Note:
 * You may assume that both strings contain only lowercase letters.
 *
 *
 *
 * canConstruct("a", "b") -> false
 * canConstruct("aa", "ab") -> false
 * canConstruct("aa", "aab") -> true
 *
 *
 */

/*

 将第二个字符串的char存到map

 遍历第一个字符串

*/

func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[uint8]int)
	for k := range magazine {
		m[magazine[k]]++
	}

	for k := range ransomNote {
		if m[ransomNote[k]] <= 0 {
			return false
		}
		m[ransomNote[k]]--
	}
	return true
}

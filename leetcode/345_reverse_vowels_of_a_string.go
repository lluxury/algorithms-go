package leetcode

/*
 * [345] Reverse Vowels of a String
 *
 * https://leetcode.com/problems/reverse-vowels-of-a-string/description/
 *
 * algorithms
 * Easy (39.64%)
 * Total Accepted:    119.4K
 * Total Submissions: 300.5K
 * Testcase Example:  '"hello"'
 *
 * Write a function that takes a string as input and reverse only the vowels of
 * a string.
 * 
 * Example 1:
 * 
 * 
 * Input: "hello"
 * Output: "holle"
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: "leetcode"
 * Output: "leotcede"
 * 
 * 
 * Note:
 * The vowels does not include the letter "y".
 * 
 * 
 * 
 */

 /*

 交换一个字符串中的元音

 用i和j指向开头和结尾，如果是元音就交换

 */

func is(s uint8) bool {
	switch s {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}

func reverseVowels(s string) string {
	b := []byte(s)

	for i, j := 0, len(b)-1; i < j; {
		if !is(b[i]) {
			i++
			continue
		}

		if !is(b[j]) {
			j--
			continue
		}

		b[i], b[j] = b[j], b[i]
		i++
		j--
	}

	return string(b)
}

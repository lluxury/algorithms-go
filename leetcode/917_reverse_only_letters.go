package leetcode

import "github.com/Chyroc/algorithms-go/lib"

/*
 * [953] Reverse Only Letters
 *
 * https://leetcode.com/problems/reverse-only-letters/description/
 *
 * algorithms
 * Easy (57.83%)
 * Total Accepted:    6.6K
 * Total Submissions: 11.4K
 * Testcase Example:  '"ab-cd"'
 *
 * Given a string S, return the "reversed" string where all characters that are
 * not a letter stay in the same place, and all letters reverse their
 * positions.
 *
 * Example 1:
 *
 * Input: "ab-cd"
 * Output: "dc-ba"
 *
 *
 * Example 2:
 *
 * Input: "a-bC-dEf-ghIj"
 * Output: "j-Ih-gfE-dCba"
 *
 *
 * Example 3:
 *
 * Input: "Test1ng-Leet=code-Q!"
 * Output: "Qedo1ct-eeLg=ntse-T!"
 *
 *
 * Note:
 *
 *
 * S.length <= 100
 * 33 <= S[i].ASCIIcode <= 122
 * S doesn't contain \ or "
 *
 */

/*

 反转字符串，非字母的数据位置不动

*/

func reverseOnlyLetters(S string) string {
	var s = make([]byte, len([]byte(S)))

	for i, j := 0, len(S)-1; i <= j; {
		if !lib.IsLetter(S[i]) {
			s[i] = S[i]
			i++
			continue
		}

		if !lib.IsLetter(S[j]) {
			s[j] = S[j]
			j--
			continue
		}

		s[i] = S[j]
		s[j] = S[i]

		i++
		j--
	}

	return string(s)
}

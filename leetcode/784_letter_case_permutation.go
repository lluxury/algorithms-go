package leetcode

/*
 * [800] Letter Case Permutation
 *
 * https://leetcode.com/problems/letter-case-permutation/description/
 *
 * algorithms
 * Easy (53.27%)
 * Total Accepted:    22.1K
 * Total Submissions: 41.3K
 * Testcase Example:  '"a1b2"'
 *
 * Given a string S, we can transform every letter individually to be lowercase
 * or uppercase to create another string.  Return a list of all possible
 * strings we could create.
 *
 *
 * Examples:
 * Input: S = "a1b2"
 * Output: ["a1b2", "a1B2", "A1b2", "A1B2"]
 *
 * Input: S = "3z4"
 * Output: ["3z4", "3Z4"]
 *
 * Input: S = "12345"
 * Output: ["12345"]
 *
 *
 * Note:
 *
 *
 * S will be a string with length at most 12.
 * S will consist only of letters or digits.
 *
 */

/*

* 给定一个字符串，可以将其中的大小字母转小写字母，也可以把其中的小写字母转大小字母，返回所有可能的操作结果
  * 和 backtrack 递归题目差不多
  * 递归函数，参数是已经处理的前面的字符串tmp，和即将处理的todoIndex，以及结果list，原始字符串s
  * 在结果tmp长度和s长度相等的时候，假如list
  * 在todoIndex等于s的长度的时候，说明所有的字符都处理完了，返回
  * 将未处理的字符加到tmp，todoIndex+1，进行递归
  * 然后将todoIndex字符串反转一下，加到tmp，todoIndex+1，进行递归

*/

func letterCasePermutation(S string) (list []string) {
	backtrack_784(S, &list, "", 0)
	return
}

func isLetterDigit(x uint8) bool {
	return (x >= 'a' && x <= 'z') || (x >= 'A' && x <= 'Z') || (x >= '0' && x <= '9')
}

func backtrack_784(s string, list *[]string, tmp string, todoIndex int) {
	if len(tmp) == len(s) {
		*list = append(*list, tmp)
	}
	if todoIndex == len(s) {
		return
	}

	u := s[todoIndex]

	backtrack_784(s, list, tmp+string(rune(u)), todoIndex+1)

	var uTo uint8
	if u >= 'a' && u <= 'z' {
		uTo = u + 'A' - 'a'
	}
	if u >= 'A' && u <= 'Z' {
		uTo = u + 'a' - 'A'
	}

	if isLetterDigit(uTo) {
		backtrack_784(s, list, tmp+string(rune(uTo)), todoIndex+1)
	}
}

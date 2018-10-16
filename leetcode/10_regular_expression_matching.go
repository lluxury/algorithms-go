package leetcode

/*
 * [10] Regular Expression Matching
 *
 * https://leetcode.com/problems/regular-expression-matching/description/
 *
 * algorithms
 * Hard (24.40%)
 * Total Accepted:    236.2K
 * Total Submissions: 967.4K
 * Testcase Example:  '"aa"\n"a"'
 *
 * Given an input string (s) and a pattern (p), implement regular expression
 * matching with support for '.' and '*'.
 * 
 * 
 * '.' Matches any single character.
 * '*' Matches zero or more of the preceding element.
 * 
 * 
 * The matching should cover the entire input string (not partial).
 * 
 * Note:
 * 
 * 
 * s could be empty and contains only lowercase letters a-z.
 * p could be empty and contains only lowercase letters a-z, and characters
 * like . or *.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input:
 * s = "aa"
 * p = "a"
 * Output: false
 * Explanation: "a" does not match the entire string "aa".
 * 
 * 
 * Example 2:
 * 
 * 
 * Input:
 * s = "aa"
 * p = "a*"
 * Output: true
 * Explanation: '*' means zero or more of the precedeng element, 'a'.
 * Therefore, by repeating 'a' once, it becomes "aa".
 * 
 * 
 * Example 3:
 * 
 * 
 * Input:
 * s = "ab"
 * p = ".*"
 * Output: true
 * Explanation: ".*" means "zero or more (*) of any character (.)".
 * 
 * 
 * Example 4:
 * 
 * 
 * Input:
 * s = "aab"
 * p = "c*a*b"
 * Output: true
 * Explanation: c can be repeated 0 times, a can be repeated 1 time. Therefore
 * it matches "aab".
 * 
 * 
 * Example 5:
 * 
 * 
 * Input:
 * s = "mississippi"
 * p = "mis*is*p*."
 * Output: false
 * 
 * 
 */

//func isMatch_rec(s string, p string) bool {
//	var (
//		s_len = len(s)
//		p_len = len(p)
//		i := s_len - 1
//		j := p_len - 1
//	)
//
//	for ; j >= 0; i-- {
//		switch p[j] {
//		case '.':
//			// . 匹配单个任意字符，将s游标前移一位
//			i--
//		case '*':
//			if j == 0 {
//				// p只有一个*，没有前置数据，错误
//				return false
//			}
//
//			// 如果是 * ，需要看前一个位置的数据
//
//			switch p[j-1] {
//			case '*':
//				// ** 错误
//				return false
//			case '.':
//				// .*
//				for {
//
//				}
//			default:
//				// a*
//				for k := i; k >= 0; k-- {
//					if s[k] == p[j-1] {
//						// a == a*
//						if isMatch_rec(s[:k],p[:j-1]){
//							return true
//						}
//					}
//				}
//			}
//		default:
//			// a 正常匹配
//			if i < 0 || s[i] != p[j] {
//				return false
//			}
//		}
//	}
//	return true
//}

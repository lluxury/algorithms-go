package leetcode

/*
 * [32] Longest Valid Parentheses
 *
 * https://leetcode.com/problems/longest-valid-parentheses/description/
 *
 * algorithms
 * Hard (23.51%)
 * Total Accepted:    135.6K
 * Total Submissions: 576.5K
 * Testcase Example:  '"(()"'
 *
 * Given a string containing just the characters '(' and ')', find the length
 * of the longest valid (well-formed) parentheses substring.
 *
 * Example 1:
 *
 *
 * Input: "(()"
 * Output: 2
 * Explanation: The longest valid parentheses substring is "()"
 *
 *
 * Example 2:
 *
 *
 * Input: ")()())"
 * Output: 4
 * Explanation: The longest valid parentheses substring is "()()"
 *
 *
 */

/*

最长有效匹配长度

动态规划算法

	首先需要定义 状态 和 状态转移方程

	状态
	Fk = 第k项到整个字符串最后一个字符之间的最长有效匹配长度，包括第k项

	状态转移
	最后一项是 dp[k-1] = 0
	如果第k-1项是（，那么：
		越过第k项所代表的字符串个数后，那个字符是），比如 "( (()) )"，结果dp[k] = dp[k+1]+2
		如果这个也成功了，那么还有一个可能，那就是 dp[k+1]+1+1那个也是一个有效匹配长度，所以长度还需要加上他 dp[k] = dp[2+dp[k+1]] +dp[k]

*/

func longestValidParentheses(s string) int {
	if len(s) == 0 {
		return 0
	}
	m := make([]int, len(s))

	for i := len(s) - 2; i >= 0; i-- {
		if s[i] == '(' {
			j := i + 1 + m[i+1]
			if j < len(s) && s[j] == ')' {
				m[i] = m[i+1] + 2

				j++
				if j < len(s) {
					m[i] += m[j]
				}
			}
		}
	}
	return max(m...)
}

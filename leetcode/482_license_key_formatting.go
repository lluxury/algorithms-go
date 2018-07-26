package leetcode

import "fmt"

/*
 * [482] License Key Formatting
 *
 * https://leetcode.com/problems/license-key-formatting/description/
 *
 * algorithms
 * Easy (39.16%)
 * Total Accepted:    38.4K
 * Total Submissions: 98K
 * Testcase Example:  '"5F3Z-2e-9-w"\n4'
 *
 * You are given a license key represented as a string S which consists only
 * alphanumeric character and dashes. The string is separated into N+1 groups
 * by N dashes.
 *
 * Given a number K, we would want to reformat the strings such that each group
 * contains exactly K characters, except for the first group which could be
 * shorter than K, but still must contain at least one character. Furthermore,
 * there must be a dash inserted between two groups and all lowercase letters
 * should be converted to uppercase.
 *
 * Given a non-empty string S and a number K, format the string according to
 * the rules described above.
 *
 * Example 1:
 *
 * Input: S = "5F3Z-2e-9-w", K = 4
 *
 * Output: "5F3Z-2E9W"
 *
 * Explanation: The string S has been split into two parts, each part has 4
 * characters.
 * Note that the two extra dashes are not needed and can be removed.
 *
 *
 *
 *
 * Example 2:
 *
 * Input: S = "2-5g-3-J", K = 2
 *
 * Output: "2-5G-3J"
 *
 * Explanation: The string S has been split into three parts, each part has 2
 * characters except the first part as it could be shorter as mentioned
 * above.
 *
 *
 *
 * Note:
 *
 * The length of string S will not exceed 12,000, and K is a positive integer.
 * String S consists only of alphanumerical characters (a-z and/or A-Z and/or
 * 0-9) and dashes(-).
 * String S is non-empty.
 *
 *
 */

/*

 格式化license，变成每x个字符用-连接一下，不足x个的段是在最前面
 * 从后往前遍历
 * 本身的-不需要，跳过
 * 临时遍历记录段的长度，等于x的时候，加入-
 * 转化大小写
 * 最后反转顺序转成字符串

*/

func tostring_482(bs []uint8) string {
	b := make([]uint8, len(bs))
	for i := len(bs) - 1; i >= 0; i-- {
		b[len(bs)-i-1] = bs[i]
	}
	return string(b)
}

func licenseKeyFormatting(S string, K int) string {
	var buf []uint8
	var count = 0

	for i := len(S) - 1; i >= 0; i-- {
		if S[i] == '-' {
			continue
		}
		if count >= K {
			buf = append(buf, '-')
			count = 1
		} else {
			fmt.Printf("count %d k %d\n", count, K)
			count++
		}

		if S[i] >= 'a' && S[i] <= 'z' {
			buf = append(buf, S[i]-32)
		} else {
			buf = append(buf, S[i])
		}
	}
	return tostring_482(buf)
}

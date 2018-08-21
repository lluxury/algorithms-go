package leetcode

/*
 * [168] Excel Sheet Column Title
 *
 * https://leetcode.com/problems/excel-sheet-column-title/description/
 *
 * algorithms
 * Easy (27.81%)
 * Total Accepted:    146.1K
 * Total Submissions: 524.9K
 * Testcase Example:  '1'
 *
 * Given a positive integer, return its corresponding column title as appear in
 * an Excel sheet.
 * 
 * For example:
 * 
 * 
 * ⁠   1 -> A
 * ⁠   2 -> B
 * ⁠   3 -> C
 * ⁠   ...
 * ⁠   26 -> Z
 * ⁠   27 -> AA
 * ⁠   28 -> AB 
 * ⁠   ...
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: 1
 * Output: "A"
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: 28
 * Output: "AB"
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: 701
 * Output: "ZY"
 * 
 * 
 */


/*


如图所示，有一个进制转换，

思考
* 就是一个十进制和26进制的转换
* 二进制的转换是：辗转相除，余数依次加到结果的前面（即第一个余数是结果的最后一位）
* 所以本题就是除以26



*/

func convertToTitle(n int) string {
	s := ""

	for n > 0 {
		m := n % 26
		if m == 0 {
			m += 26
		}
		s = string(m+64) + s
		n = (n - m) / 26
	}

	return s
}

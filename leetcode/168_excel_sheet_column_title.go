package leetcode

import (
	// "fmt"
)
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

func toChar(n int) string {
	if n == 0 {
		n = 26
	}
	return string('A'+n-1)
}

func convertToTitle(n int) string {
	res :=""
	for n> 26{
		d := n % 26
		res = toChar(d) + res
		n/=26
		if d==0{
			n-=1
		}
	}
	d := n % 26
	// // fmt.Fprintln(n,d,res)
    // fmt.Println(n, d, res)
	res = toChar(d) + res
	return res
}

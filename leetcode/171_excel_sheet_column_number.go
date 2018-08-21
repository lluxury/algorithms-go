package leetcode

/*
 * [171] Excel Sheet Column Number
 *
 * https://leetcode.com/problems/excel-sheet-column-number/description/
 *
 * algorithms
 * Easy (49.47%)
 * Total Accepted:    181.7K
 * Total Submissions: 367K
 * Testcase Example:  '"A"'
 *
 * Given a column title as appear in an Excel sheet, return its corresponding
 * column number.
 * 
 * For example:
 * 
 * 
 * ⁠   A -> 1
 * ⁠   B -> 2
 * ⁠   C -> 3
 * ⁠   ...
 * ⁠   Z -> 26
 * ⁠   AA -> 27
 * ⁠   AB -> 28 
 * ⁠   ...
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: "A"
 * Output: 1
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: "AB"
 * Output: 28
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: "ZY"
 * Output: 701
 * 
 * 
 */


/*

同 [Excel Sheet Column Title](./168-excel-sheet-column-title.md)

只不过上一次是十进制转26进制，这一次是26进制转10进制

*/

import (
	"math"
)

func titleToNumber(s string) int {
	length := len(s)
	n := 0

	for k, v := range s {
		n += int(math.Pow(26, float64(length-k-1))) * (int(v - 64))
	}

	return n
}


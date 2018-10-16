package leetcode

/*
 * [6] ZigZag Conversion
 *
 * https://leetcode.com/problems/zigzag-conversion/description/
 *
 * algorithms
 * Medium (28.08%)
 * Total Accepted:    224.9K
 * Total Submissions: 801K
 * Testcase Example:  '"PAYPALISHIRING"\n3'
 *
 * The string "PAYPALISHIRING" is written in a zigzag pattern on a given number
 * of rows like this: (you may want to display this pattern in a fixed font for
 * better legibility)
 *
 *
 * P   A   H   N
 * A P L S I I G
 * Y   I   R
 *
 *
 * And then read line by line: "PAHNAPLSIIGYIR"
 *
 * Write the code that will take a string and make this conversion given a
 * number of rows:
 *
 *
 * string convert(string s, int numRows);
 *
 * Example 1:
 *
 *
 * Input: s = "PAYPALISHIRING", numRows = 3
 * Output: "PAHNAPLSIIGYIR"
 *
 *
 * Example 2:
 *
 *
 * Input: s = "PAYPALISHIRING", numRows = 4
 * Output: "PINALSIGYAHRPI"
 * Explanation:
 *
 * P     I    N
 * A   L S  I G
 * Y A   H R
 * P     I
 *
 */
 

 /*

 给定一个字符串，将这么字符串按照Z的格式展开

 */

func convert(s string, numRows int) string {
	if s == "" || numRows == 1 {
		return s
	}

	var result = make([][]byte, numRows)

	index := 0
	add := -1
	for k := 0; k < len(s); k++ {
		if index == numRows-1 || index == 0 {
			add = -add
		}

		result[index] = append(result[index], byte(s[k]))
		index += add
	}

	var data []byte
	for _, v := range result {
		data = append(data, v...)
	}
	return string(data)
}

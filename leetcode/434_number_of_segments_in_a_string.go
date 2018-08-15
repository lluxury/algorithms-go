package leetcode

/*
 * [434] Number of Segments in a String
 *
 * https://leetcode.com/problems/number-of-segments-in-a-string/description/
 *
 * algorithms
 * Easy (36.55%)
 * Total Accepted:    43K
 * Total Submissions: 117.6K
 * Testcase Example:  '"Hello, my name is John"'
 *
 * Count the number of segments in a string, where a segment is defined to be a
 * contiguous sequence of non-space characters.
 *
 * Please note that the string does not contain any non-printable characters.
 *
 * Example:
 *
 * Input: "Hello, my name is John"
 * Output: 5
 *
 *
 */

func countSegments(s string) int {
	count := 0
	isspace := true
	for _, v := range s {
		if v == ' ' {
			if !isspace {
				isspace = true
			}
		} else {
			if isspace {
				isspace = false
				count++
			}
		}
	}
	return count
}

package leetcode

/*
 * [69] Sqrt(x)
 *
 * https://leetcode.com/problems/sqrtx/description/
 *
 * algorithms
 * Easy (29.30%)
 * Total Accepted:    258.5K
 * Total Submissions: 880K
 * Testcase Example:  '4'
 *
 * Implement int sqrt(int x).
 *
 * Compute and return the square root of x, where x is guaranteed to be a
 * non-negative integer.
 *
 * Since the return type is an integer, the decimal digits are truncated and
 * only the integer part of the result is returned.
 *
 * Example 1:
 *
 *
 * Input: 4
 * Output: 2
 *
 *
 * Example 2:
 *
 *
 * Input: 8
 * Output: 2
 * Explanation: The square root of 8 is 2.82842..., and since
 * the decimal part is truncated, 2 is returned.
 *
 *
 */

/*

 求平方根，转成整数


*/
/*
func mySqrt(x int) int {
	var a float64 = 1
	for {
		a = (a + float64(x)/a) / 2

		diff := a*a - float64(x)
		if diff < 0 {
			diff = -diff
		}
		if diff < 1 {
			break
		}
	}

	return int(a)
}*/


func mySqrt(x int) int {
	if (x == 0 || x == 1){
		return x
	}

	l := 1
	r := x
	var res int

	for (l <= r){
		var m int 
		m = (l + r) /2
		// mm = ll + (rr - ll >> 1);
		if (m == x / m) {
			return m
		} else if (m > x / m){
			r = m -1
		} else {
			l = m + 1
			res = m 
		}
	}
    return res
}

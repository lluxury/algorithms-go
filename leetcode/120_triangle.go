package leetcode

/*
 * @lc app=leetcode id=120 lang=golang
 *
 * [120] Triangle
 *
 * https://leetcode.com/problems/triangle/description/
 *
 * algorithms
 * Medium (39.71%)
 * Total Accepted:    187.8K
 * Total Submissions: 471.2K
 * Testcase Example:  '[[2],[3,4],[6,5,7],[4,1,8,3]]'
 *
 * Given a triangle, find the minimum path sum from top to bottom. Each step
 * you may move to adjacent numbers on the row below.
 *
 * For example, given the following triangle
 *
 *
 * [
 * ⁠    [2],
 * ⁠   [3,4],
 * ⁠  [6,5,7],
 * ⁠ [4,1,8,3]
 * ]
 *
 *
 * The minimum path sum from top to bottom is 11 (i.e., 2 + 3 + 5 + 1 = 11).
 *
 * Note:
 *
 * Bonus point if you are able to do this using only O(n) extra space, where n
 * is the total number of rows in the triangle.
 *
 */

 
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	if n == 0 {
		return 0
	}
	i := 0
	for i = 1; i < n; i++ {
		prev := i - 1
		length := len(triangle[prev])
		for j := range triangle[i] {
			min := 2147483647
			if j == 0 {
				min = triangle[prev][j]
			} else if j == length {
				min = triangle[prev][j-1]
			} else {
				if j-1 >= 0 && triangle[prev][j-1] < triangle[prev][j] {
					min = triangle[prev][j-1]
				} else {
					min = triangle[prev][j]
				}
			}
			triangle[i][j] += min
		}
	}
	min := triangle[n-1][0]
	//for i := 1; i<0; i++ {
	for i := 1; i < n; i++ {
		if triangle[n-1][i] < min {
			min = triangle[n-1][i]
		}
	}
	return min
}

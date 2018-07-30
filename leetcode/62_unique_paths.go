package leetcode

/*
 * [62] Unique Paths
 *
 * https://leetcode.com/problems/unique-paths/description/
 *
 * algorithms
 * Medium (43.75%)
 * Total Accepted:    207K
 * Total Submissions: 472.8K
 * Testcase Example:  '3\n2'
 *
 * A robot is located at the top-left corner of a m x n grid (marked 'Start' in
 * the diagram below).
 *
 * The robot can only move either down or right at any point in time. The robot
 * is trying to reach the bottom-right corner of the grid (marked 'Finish' in
 * the diagram below).
 *
 * How many possible unique paths are there?
 *
 *
 * Above is a 7 x 3 grid. How many possible unique paths are there?
 *
 * Note: m and n will be at most 100.
 *
 * Example 1:
 *
 *
 * Input: m = 3, n = 2
 * Output: 3
 * Explanation:
 * From the top-left corner, there are a total of 3 ways to reach the
 * bottom-right corner:
 * 1. Right -> Right -> Down
 * 2. Right -> Down -> Right
 * 3. Down -> Right -> Right
 *
 *
 * Example 2:
 *
 *
 * Input: m = 7, n = 3
 * Output: 28
 *
 */

/*

最大增序子序列的长度

动态规划算法

	状态
	Fi,j = 以(i,j)为终点，(0,0)为起点的路线图的个数

	状态转移
	因为只能向下或者向右移动，所以，Fi,j = 他上面的Fi-1,j + 他左面的Fi,j-1
	Fi,j 初始值为1

*/

func uniquePaths(m int, n int) int {
	data := make([][]int, m)
	for i := 0; i < m; i++ {
		data[i] = make([]int, n)
		for j := 0; j < n; j++ {
			data[i][j] = 1
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			data[i][j] = data[i][j-1] + data[i-1][j]
		}
	}

	return data[m-1][n-1]
}

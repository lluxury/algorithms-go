package leetcode

/*
 * [598] Range Addition II
 *
 * https://leetcode.com/problems/range-addition-ii/description/
 *
 * algorithms
 * Easy (48.05%)
 * Total Accepted:    22.8K
 * Total Submissions: 47.6K
 * Testcase Example:  '3\n3\n[[2,2],[3,3]]'
 *
 * Given an m * n matrix M initialized with all 0's and several update
 * operations.
 * Operations are represented by a 2D array, and each operation is represented
 * by an array with two positive integers a and b, which means M[i][j] should
 * be added by one for all 0  and 0 .
 * You need to count and return the number of maximum integers in the matrix
 * after performing all the operations.
 *
 * Example 1:
 *
 * Input:
 * m = 3, n = 3
 * operations = [[2,2],[3,3]]
 * Output: 4
 * Explanation:
 * Initially, M =
 * [[0, 0, 0],
 * ⁠[0, 0, 0],
 * ⁠[0, 0, 0]]
 *
 * After performing [2,2], M =
 * [[1, 1, 0],
 * ⁠[1, 1, 0],
 * ⁠[0, 0, 0]]
 *
 * After performing [3,3], M =
 * [[2, 2, 1],
 * ⁠[2, 2, 1],
 * ⁠[1, 1, 1]]
 *
 * So the maximum integer in M is 2, and there are four of it in M. So return
 * 4.
 *
 *
 *
 * Note:
 *
 * The range of m and n is [1,40000].
 * The range of a is [1,m], and the range of b is [1,n].
 * The range of operations size won't exceed 10,000.
 *
 *
 */

/*

 给一个二维数组，然后给定一个坐标，(0,0)到该坐标的数据加1，求最后最大值的数据的个数
   * 显然，每一次相当于涂了一次颜料，求最浓墨重彩的地方的次数
   * 只要记住涂颜料的正方形的起点（0,0）和终点（要求的）即可
   * 每一次都可能在缩小面积，记录最小值即可
   * 最后返回面积

*/

func maxCount(m int, n int, ops [][]int) int {
	if len(ops) == 0 {
		return m * n
	}

	for _, v := range ops {
		if v[0] < m {
			m = v[0]
		}
		if v[1] < n {
			n = v[1]
		}
	}
	return m * n
}

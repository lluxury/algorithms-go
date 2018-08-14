package leetcode

/*
 * [898] Transpose Matrix
 *
 * https://leetcode.com/problems/transpose-matrix/description/
 *
 * algorithms
 * Easy (66.89%)
 * Total Accepted:    13.1K
 * Total Submissions: 19.8K
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * Given a matrix A, return the transpose of A.
 *
 * The transpose of a matrix is the matrix flipped over it's main diagonal,
 * switching the row and column indices of the matrix.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [[1,2,3],[4,5,6],[7,8,9]]
 * Output: [[1,4,7],[2,5,8],[3,6,9]]
 *
 *
 *
 * Example 2:
 *
 *
 * Input: [[1,2,3],[4,5,6]]
 * Output: [[1,4],[2,5],[3,6]]
 *
 *
 *
 *
 * Note:
 *
 *
 * 1 <= A.length <= 1000
 * 1 <= A[0].length <= 1000
 *
 *
 *
 */

/*

 反转二维数组

*/

func transpose(A [][]int) [][]int {
	if len(A) == 0 {
		return nil
	}

	a := make([][]int, len(A[0]))
	for k := range A[0] {
		a[k] = make([]int, len(A))
	}

	for i := range A {
		for j := range A[0] {
			//if i < j {
			//	A[i][j], A[j][i] = A[j][i], A[i][j]
			//}
			a[j][i] = A[i][j]
		}
	}
	return a
}

package leetcode

/*
 * [777] Toeplitz Matrix
 *
 * https://leetcode.com/problems/toeplitz-matrix/description/
 *
 * algorithms
 * Easy (58.60%)
 * Total Accepted:    30.8K
 * Total Submissions: 52.5K
 * Testcase Example:  '[[1,2,3,4],[5,1,2,3],[9,5,1,2]]'
 *
 * A matrix is Toeplitz if every diagonal from top-left to bottom-right has the
 * same element.
 * 
 * Now given an M x N matrix, return True if and only if the matrix is
 * Toeplitz.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input:
 * matrix = [
 * [1,2,3,4],
 * [5,1,2,3],
 * [9,5,1,2]
 * ]
 * Output: True
 * Explanation:
 * In the above grid, the diagonals are:
 * "[9]", "[5, 5]", "[1, 1, 1]", "[2, 2, 2]", "[3, 3]", "[4]".
 * In each diagonal all elements are the same, so the answer is True.
 * 
 * 
 * Example 2:
 * 
 * 
 * Input:
 * matrix = [
 * [1,2],
 * [2,2]
 * ]
 * Output: False
 * Explanation:
 * The diagonal "[1, 2]" has different elements.
 * 
 * 
 * 
 * Note:
 * 
 * 
 * matrix will be a 2D array of integers.
 * matrix will have a number of rows and columns in range [1, 20].
 * matrix[i][j] will be integers in range [0, 99].
 * 
 * 
 * 
 * Follow up:
 * 
 * 
 * What if the matrix is stored on disk, and the memory is limited such that
 * you can only load at most one row of the matrix into the memory at once?
 * What if the matrix is so large that you can only load up a partial row into
 * the memory at once?
 * 
 * 
 */

 /*

 给一个二维数组，判断从右上角到左下角，是不是数字都相等

 从最左下角 遍历 到  最坐上角 再遍历到 右上角
 判断从这些位置开始的「从右上角到左下角」的数字是不是都相等

 */

func isToeplitzMatrix_rec(matrix [][]int, i, j int) bool {
	d := matrix[i][j]
	for {
		if i >= len(matrix)-1 || j >= len(matrix[0])-1 {
			break
		}
		i++
		j++
		if d != matrix[i][j] {
			return false
		}
	}

	return true
}

func isToeplitzMatrix(matrix [][]int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return true
	}

	for i, j := len(matrix)-1, 0; i >= 0 && j < len(matrix[0]); {
		if !isToeplitzMatrix_rec(matrix, i, j) {
			return false
		}
		if i > 0 {
			i--
		} else if j < len(matrix[0])-1 {
			j++
		} else {
			break
		}
	}

	return true
}

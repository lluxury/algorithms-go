package leetcode

/*
 * [861] Flipping an Image
 *
 * https://leetcode.com/problems/flipping-an-image/description/
 *
 * algorithms
 * Easy (69.42%)
 * Total Accepted:    23.6K
 * Total Submissions: 33.9K
 * Testcase Example:  '[[1,1,0],[1,0,1],[0,0,0]]'
 *
 * Given a binary matrix A, we want to flip the image horizontally, then invert
 * it, and return the resulting image.
 *
 * To flip an image horizontally means that each row of the image is reversed.
 * For example, flipping [1, 1, 0] horizontally results in [0, 1, 1].
 *
 * To invert an image means that each 0 is replaced by 1, and each 1 is
 * replaced by 0. For example, inverting [0, 1, 1] results in [1, 0, 0].
 *
 * Example 1:
 *
 *
 * Input: [[1,1,0],[1,0,1],[0,0,0]]
 * Output: [[1,0,0],[0,1,0],[1,1,1]]
 * Explanation: First reverse each row: [[0,1,1],[1,0,1],[0,0,0]].
 * Then, invert the image: [[1,0,0],[0,1,0],[1,1,1]]
 *
 *
 * Example 2:
 *
 *
 * Input: [[1,1,0,0],[1,0,0,1],[0,1,1,1],[1,0,1,0]]
 * Output: [[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]
 * Explanation: First reverse each row:
 * [[0,0,1,1],[1,0,0,1],[1,1,1,0],[0,1,0,1]].
 * Then invert the image: [[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]
 *
 *
 * Notes:
 *
 *
 * 1 <= A.length = A[0].length <= 20
 * 0 <= A[i][j] <= 1
 *
 */

/*
 一个图片，二维数组，将每一个左右反转，再黑白反转
 * 经过观察：对于首尾对称的元素，相等的haul，经过两次操作，则只需要经常黑白转；如果不等，则无需操作
 * 0->1，1->0的操作可以使用操作符：x^1
*/

func flipAndInvertImage(A [][]int) [][]int {
	for k := range A {
		for i, j := 0, len(A[k])-1; i <= j; {
			if i == j {
				A[k][i] ^= 1
			}
			if A[k][i] == A[k][j] {
				A[k][i] ^= 1
				A[k][j] ^= 1
			}
			i++
			j--

		}
	}
	return A
}

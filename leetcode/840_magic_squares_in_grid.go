package leetcode

/*
 * [870] Magic Squares In Grid
 *
 * https://leetcode.com/problems/magic-squares-in-grid/description/
 *
 * algorithms
 * Easy (34.33%)
 * Total Accepted:    5.7K
 * Total Submissions: 16.7K
 * Testcase Example:  '[[4,3,8,4],[9,5,1,9],[2,7,6,2]]'
 *
 * A 3 x 3 magic square is a 3 x 3 grid filled with distinct numbers from 1 to
 * 9 such that each row, column, and both diagonals all have the same sum.
 *
 * Given an grid of integers, how many 3 x 3 "magic square" subgrids are
 * there?  (Each subgrid is contiguous).
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [[4,3,8,4],
 * ⁠       [9,5,1,9],
 * ⁠       [2,7,6,2]]
 * Output: 1
 * Explanation:
 * The following subgrid is a 3 x 3 magic square:
 * 438
 * 951
 * 276
 *
 * while this one is not:
 * 384
 * 519
 * 762
 *
 * In total, there is only one magic square inside the given grid.
 *
 *
 * Note:
 *
 *
 * 1 <= grid.length <= 10
 * 1 <= grid[0].length <= 10
 * 0 <= grid[i][j] <= 15
 *
 *
 */

/*

如果一个3*3的矩阵，横，竖，对角线，如果和相等，就认为是一个合法的矩阵，问给定的矩阵有多少个这样的3*3矩阵？

*/

func numMagicSquaresInside_rec(grid [][]int, i, j, sum int) bool {
	// 2横
	for k := 1; k <= 2; k++ {
		if grid[i+k][j]+grid[i+k][j+1]+grid[i+k][j+2] != sum {
			return false
		}
	}

	// 3竖
	for k := 0; k <= 2; k++ {
		if grid[i][j+k]+grid[i+1][j+k]+grid[i+2][j+k] != sum {
			return false
		}
	}

	// 2 斜
	if grid[i][j]+grid[i+1][j+1]+grid[i+2][j+2] != sum {
		return false
	}
	if grid[i+2][j]+grid[i+1][j+1]+grid[i][j+2] != sum {
		return false
	}

	return true
}

func check19(grid [][]int, i, j int) bool {
	m := make([]int, 10) // 1-9
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			data := grid[i+x][j+y]
			if data < 1 || data > 9 {
				return false
			}
			if m[data] == 1 {
				return false
			}
			m[data] = 1
		}
	}
	return true
}

func numMagicSquaresInside(grid [][]int) int {
	if len(grid) < 3 || len(grid[0]) < 3 {
		return 0
	}

	count := 0
	for i := 0; i < len(grid)-2; i++ {
		for j := 0; j < len(grid[0])-2; j++ {
			if check19(grid, i, j) && numMagicSquaresInside_rec(grid, i, j, grid[i][j]+grid[i][j+1]+grid[i][j+2]) {
				count++
			}
		}
	}
	return count
}

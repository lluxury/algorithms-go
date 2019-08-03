package leetcode

/*
 * @lc app=leetcode id=994 lang=golang
 *
 * [994] Rotting Oranges
 *
 * https://leetcode.com/problems/rotting-oranges/description/
 *
 * algorithms
 * Easy (46.48%)
 * Total Accepted:    15.2K
 * Total Submissions: 32.6K
 * Testcase Example:  '[[2,1,1],[1,1,0],[0,1,1]]'
 *
 * In a given grid, each cell can have one of three values:
 * 
 * 
 * the value 0 representing an empty cell;
 * the value 1 representing a fresh orange;
 * the value 2 representing a rotten orange.
 * 
 * 
 * Every minute, any fresh orange that is adjacent (4-directionally) to a
 * rotten orange becomes rotten.
 * 
 * Return the minimum number of minutes that must elapse until no cell has a
 * fresh orange.  If this is impossible, return -1 instead.
 * 
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * 
 * 
 * Input: [[2,1,1],[1,1,0],[0,1,1]]
 * Output: 4
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: [[2,1,1],[0,1,1],[1,0,1]]
 * Output: -1
 * Explanation:  The orange in the bottom left corner (row 2, column 0) is
 * never rotten, because rotting only happens 4-directionally.
 * 
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: [[0,2]]
 * Output: 0
 * Explanation:  Since there are already no fresh oranges at minute 0, the
 * answer is just 0.
 * 
 * 
 * 
 * 
 * Note:
 * 
 * 
 * 1 <= grid.length <= 10
 * 1 <= grid[0].length <= 10
 * grid[i][j] is only 0, 1, or 2.
 * 
 * 
 * 
 * 
 */
var dx = [4]int{0, 0, -1, 1}
var dy = [4]int{-1, 1, 0, 0}
func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	fresh := 0
	rottens := make([][2]int, 0, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			switch grid[i][j] {
			case 1:
				fresh++
			case 2:
				rottens = append(rottens, [2]int{i, j})
			}
		}
	}

	if fresh == 0 {
		return 0
	}

	count := -1
	for len(rottens) > 0 {
		count++
		size := len(rottens)
		for r := 0; r < size; r++ {
			x, y := rottens[r][0], rottens[r][1]
			for k := 0; k < 4; k++ {
				//i, j := x+dx[k], y+dx[k]
				i, j := x+dx[k], y+dy[k]
				if i < 0 || m <= i ||
					j < 0 || n <= j ||
					grid[i][j] != 1 {
					continue
				}
				grid[i][j] = 2
				fresh--
				rottens = append(rottens, [2]int{i, j})
		    }
		}
		rottens = rottens[size:]
	}
	if fresh > 0 {
		return -1
	}
	return count
}



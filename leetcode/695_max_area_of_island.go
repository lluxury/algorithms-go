package leetcode

import "github.com/Chyroc/algorithms-go/lib"

/*
 * [695] Max Area of Island
 *
 * https://leetcode.com/problems/max-area-of-island/description/
 *
 * algorithms
 * Easy (52.22%)
 * Total Accepted:    35.6K
 * Total Submissions: 68.2K
 * Testcase Example:  '[[1,1,0,0,0],[1,1,0,0,0],[0,0,0,1,1],[0,0,0,1,1]]'
 *
 * Given a non-empty 2D array grid of 0's and 1's, an island is a group of 1's
 * (representing land) connected 4-directionally (horizontal or vertical.)  You
 * may assume all four edges of the grid are surrounded by water.
 *
 * Find the maximum area of an island in the given 2D array.
 * (If there is no island, the maximum area is 0.)
 *
 * Example 1:
 *
 * [[0,0,1,0,0,0,0,1,0,0,0,0,0],
 * ⁠[0,0,0,0,0,0,0,1,1,1,0,0,0],
 * ⁠[0,1,1,0,1,0,0,0,0,0,0,0,0],
 * ⁠[0,1,0,0,1,1,0,0,1,0,1,0,0],
 * ⁠[0,1,0,0,1,1,0,0,1,1,1,0,0],
 * ⁠[0,0,0,0,0,0,0,0,0,0,1,0,0],
 * ⁠[0,0,0,0,0,0,0,1,1,1,0,0,0],
 * ⁠[0,0,0,0,0,0,0,1,1,0,0,0,0]]
 *
 * Given the above grid, return 6.
 *
 * Note the answer is not 11, because the island must be connected
 * 4-directionally.
 *
 *
 * Example 2:
 * [[0,0,0,0,0,0,0,0]]
 * Given the above grid, return 0.
 *
 *
 * Note:
 * The length of each dimension in the given grid does not exceed 50.
 *
 */

/*

 给一个二维数组，求联系相连的1的个数（只有上下左右才算，斜的不算）

 * 递归
   * 用深度搜索
   * 用同样大小的二维数组记录哪些数据已经访问过
   * 因为是深度搜索，所以只要将上下左右的数据加起来就是最大值

 * 迭代
    * 可以使用栈模拟深度遍历
    * 用数组记录是否访问过，然后遍历原始数据数组
      * 如果没有访问过，并且是岛屿，将index入栈
      * 如果栈不为空，遍历：
        * 出栈
        * 获取栈的上下左右四个index，入栈
        * x++
      * 那么，当一个栈空的时候，也就是所有相邻的岛屿遍历结束的时候，可以比较个数与max了

*/

func newArray(i, j int) [][]bool {
	var exist = make([][]bool, i)
	for k := 0; k < i; k++ {
		exist[k] = make([]bool, j)
	}
	return exist
}

func maxAreaOfIsland_rec(grid [][]int, exist *[][]bool, i, j int) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || (*exist)[i][j] || grid[i][j] == 0 {
		return 0
	}

	(*exist)[i][j] = true
	return 1 +
		maxAreaOfIsland_rec(grid, exist, i, j+1) +
		maxAreaOfIsland_rec(grid, exist, i, j-1) +
		maxAreaOfIsland_rec(grid, exist, i+1, j) +
		maxAreaOfIsland_rec(grid, exist, i-1, j)
}

// 递归
func maxAreaOfIsland_Recursive(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	var exist = newArray(len(grid), len(grid[0]))

	var max = 0
	for i := range grid {
		for j := range grid[0] {
			m := maxAreaOfIsland_rec(grid, &exist, i, j)
			if m > max {
				max = m
			}
		}
	}

	return max
}

// 迭代
func maxAreaOfIsland_Iterative(grid [][]int) int {
	var exist = newArray(len(grid), len(grid[0]))
	var dr = []int{1, -1, 0, 0}
	var dc = []int{0, 0, 1, -1}

	max := 0
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 1 && !exist[i][j] {
				t := 0
				exist[i][j] = true
				stack := lib.NewStack()
				stack.Push([]int{i, j})
				for !stack.IsEmpty() {
					t++
					x := stack.Pop().([]int)
					m, n := x[0], x[1]
					for k := range dr {
						r := m + dr[k]
						c := n + dc[k]
						if r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0]) && grid[r][c] == 1 && !exist[r][c] {
							exist[r][c] = true
							stack.Push([]int{r, c})
						}
					}
				}

				if t > max {
					max = t
				}
			}

		}
	}

	return max
}

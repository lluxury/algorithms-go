package leetcode

/*
 * @lc app=leetcode id=417 lang=golang
 *
 * [417] Pacific Atlantic Water Flow
 *
 * https://leetcode.com/problems/pacific-atlantic-water-flow/description/
 *
 * algorithms
 * Medium (38.11%)
 * Total Accepted:    49K
 * Total Submissions: 128.4K
 * Testcase Example:  '[[1,2,2,3,5],[3,2,3,4,4],[2,4,5,3,1],[6,7,1,4,5],[5,1,1,2,4]]'
 *
 * Given an m x n matrix of non-negative integers representing the height of
 * each unit cell in a continent, the "Pacific ocean" touches the left and top
 * edges of the matrix and the "Atlantic ocean" touches the right and bottom
 * edges.
 * 
 * Water can only flow in four directions (up, down, left, or right) from a
 * cell to another one with height equal or lower.
 * 
 * Find the list of grid coordinates where water can flow to both the Pacific
 * and Atlantic ocean.
 * 
 * Note:
 * 
 * 
 * The order of returned grid coordinates does not matter.
 * Both m and n are less than 150.
 * 
 * 
 * 
 * 
 * Example:
 * 
 * 
 * Given the following 5x5 matrix:
 * 
 * ⁠ Pacific ~   ~   ~   ~   ~ 
 * ⁠      ~  1   2   2   3  (5) *
 * ⁠      ~  3   2   3  (4) (4) *
 * ⁠      ~  2   4  (5)  3   1  *
 * ⁠      ~ (6) (7)  1   4   5  *
 * ⁠      ~ (5)  1   1   2   4  *
 * ⁠         *   *   *   *   * Atlantic
 * 
 * Return:
 * 
 * [[0, 4], [1, 3], [1, 4], [2, 2], [3, 0], [3, 1], [4, 0]] (positions with
 * parentheses in above matrix).
 * 
 * 
 * 
 * 
 */
 func pacificAtlantic(matrix [][]int) [][]int {
	res := [][]int{}
	// if len(matrix) == 0 || len(matrix[0] == 0) {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
		return res
	}

	m,n := len(matrix), len(matrix[0])
	p, a := make([][]bool,m), make([][]bool,m)
	for i := 0; i<m; i++{
		p[i] = make([]bool,n)
		a[i] = make([]bool,n)
	}

	pQueue := [][]int{}
	aQueue := [][]int{}
	for i := 0; i < m; i++ {
		p[i][0] = true
		pQueue = append(pQueue, []int{i,0})
		a[i][n-1] = true
		aQueue = append(aQueue,[]int{i,n-1})
	}
	for j:= 0; j <n; j++{
		p[0][j] = true
		pQueue = append(pQueue, []int{0,j})
		a[m-1][j] = true
		aQueue = append(aQueue,[]int{m-1,j})
	}
	ds := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	bfs := func(queue [][]int,rec [][]bool){
		for len(queue) > 0 {
			c := queue[0]
			queue = queue[1:]
			for _, d:= range ds{
				i,j := c[0]+d[0],c[1]+d[1]
				if 0<=i && i <m && 0 <=j && j <n && !rec[i][j] && matrix[c[0]][c[1]] <= matrix[i][j]{
					// 少了 = 号,结果出错
					rec[i][j] = true
					queue = append(queue, []int{i,j})
				}
			}
		}
	}
	bfs(pQueue,p)
	bfs(aQueue,a)

	for i :=0; i <m; i++{
		for j:= 0; j < n; j++{
			if p[i][j] && a[i][j] {
				res = append(res, []int{i,j})
			}
		}
	}
	return res
}

package leetcode

/*
 * [130] Surrounded Regions
 *
 * https://leetcode.com/problems/surrounded-regions/description/
 *
 * algorithms
 * Medium (20.41%)
 * Total Accepted:    111K
 * Total Submissions: 543.6K
 * Testcase Example:  '[["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]'
 *
 * Given a 2D board containing 'X' and 'O' (the letter O), capture all regions
 * surrounded by 'X'.
 *
 * A region is captured by flipping all 'O's into 'X's in that surrounded
 * region.
 *
 * Example:
 *
 *
 * X X X X
 * X O O X
 * X X O X
 * X O X X
 *
 *
 * After running your function, the board should be:
 *
 *
 * X X X X
 * X X X X
 * X X X X
 * X O X X
 *
 *
 * Explanation:
 *
 * Surrounded regions shouldn’t be on the border, which means that any 'O' on
 * the border of the board are not flipped to 'X'. Any 'O' that is not on the
 * border and it is not connected to an 'O' on the border will be flipped to
 * 'X'. Two cells are connected if they are adjacent cells connected
 * horizontally or vertically.
 *
 */
func solve_rec(board [][]byte, fangwen [][]bool, isO [][]bool, i, j int) {
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) {
		return
	}
	if fangwen[i][j] {
		return
	}

	fangwen[i][j] = true

	if board[i][j] == 'O' {
		isO[i][j] = true
		solve_rec(board, fangwen, isO, i+1, j)
		solve_rec(board, fangwen, isO, i-1, j)
		solve_rec(board, fangwen, isO, i, j+1)
		solve_rec(board, fangwen, isO, i, j-1)
	}
}

func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}

	board_len := len(board)

	var iso = make([][]bool, board_len)
	var fangwen = make([][]bool, board_len)
	for i, v := range board {
		iso[i] = make([]bool, len(v))
		fangwen[i] = make([]bool, len(v))
	}

	for i := 0; i < board_len; i++ {
		solve_rec(board, fangwen, iso, i, 0)
		solve_rec(board, fangwen, iso, i, len(board[0])-1)
	}
	for j := 1; j < len(board[0])-1; j++ {
		solve_rec(board, fangwen, iso, 0, j)
		solve_rec(board, fangwen, iso, board_len-1, j)
	}

	// 这个时候iso就是o所在的位置
	for i, v := range iso {
		for j, isooooo := range v {
			if !isooooo {
				board[i][j] = 'X'
			}
		}
	}
}

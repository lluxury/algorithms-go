package leetcode

/*
The n-queens puzzle is the problem of placing n queens on an n×n chessboard such that no two queens attack each other.



Given an integer n, return the number of distinct solutions to the n-queens puzzle.

Example:

Input: 4
Output: 2
Explanation: There are two distinct solutions to the 4-queens puzzle as shown below.
[
 [".Q..",  // Solution 1
  "...Q",
  "Q...",
  "..Q."],

 ["..Q.",  // Solution 2
  "Q...",
  "...Q",
  ".Q.."]
]*/

// 有了前一道,这道很简单, 改下返回值就好

// func reject_52(c [][]rune, row, col int) bool {
//     for i := 0; i < col; i++ {
//         if c[row][i] == 'Q' {
//             return true
//         }
//     }

//     for i, j := row, col; i >= 0 && j >= 0; {
//         if c[i][j] == 'Q' {
//             return true
//         }
//         i--
//         j--
//     }

//     for i, j := row, col; i < len(c) && j >= 0; {
//         if c[i][j] == 'Q' {
//             return true
//         }
//         i++
//         j--
//     }
//     return false
// }

// func totalNQueens(n int) int {
//     var result [][]string

//     t := make([][]rune, n)
//     for i := range t {
//         t[i] = make([]rune, n)
//         for j := range t[i] {
//             t[i][j] = '.'
//         }
//     }

//     backtracking_52(&result, t, n, 0)
//     return len(result)
// }

// func backtracking_52(res *[][]string, c [][]rune, r, column int) {
//     //fmt.Printf("r:%d c:%d \n", r, column)
//     // accept
//     if r == 0 {
//         var arr []string
//         for _, v := range c {
//             arr = append(arr, string(v))
//         }
//         *res = append(*res, arr)
//         //fmt.Printf("Accept %d %v\n", r, c)
//         return
//     }

//     // next
//     for i := 0; i < len(c); i++ {
//         // reject
//         if !reject_52(c, i, column) {
//             c[i][column] = 'Q'
//             backtracking_52(res, c, r-1, column+1)
//             c[i][column] = '.'
//         }
//     }
// }

// var count int =0
// func totalNQueens(n int) int {
//     if n < 1 {
//         return 0
//     }

//     dfs_52(n,0,0,0,0)
//     return count
// }

// func dfs_52(n int, row int, cols int, slash int, backslash int) {
//     if row >= n {
//         count++
//         return
//     }
//     bits := (^(cols | slash | backslash)) & ((1 << uint(n)) - 1)
//     for bits > 0 {
//         p := bits & -bits
//         dfs_52(n, row+1, cols|p, (slash|p)<<1, (backslash|p)>>1)
//         bits = bits & (bits - 1)
//     }
// }

// 花了4个小时,甚至抄袭了别人的解法,还是不正确,先放放吧

func totalNQueens(n int) int {
	// var board Board=Board{size:n}
	// 函数比实例化,省了一点内存

	count := Solve(New(n))
	// count := Solve(board)
	return count
}

type Board struct {
	size      int
	row       int
	cols      int // 64-bits
	slash     int
	backslash int
}

func New(size int) Board {
	return Board{size: size}
}

func Ok(board Board, col int) bool {
	return board.cols&(1<<uint(col)) == 0 &&
		board.slash&(1<<uint(board.row+col)) == 0 &&
		board.backslash&(1<<uint(board.row-col+board.size-1)) == 0
}

func Place(board Board, col int) Board {
	return Board{
		size:      board.size,
		row:       board.row + 1,
		cols:      board.cols | (1 << uint(col)),
		slash:     board.slash | (1 << uint(board.row+col)),
		backslash: board.backslash | (1 << uint(board.row-col+board.size-1))}
}

func Solve(board Board) int {
	if board.row == board.size {
		return 1
	} else {
		result := 0
		for col := 0; col < board.size; col++ {
			if Ok(board, col) {
				result += Solve(Place(board, col))
			}
		}
		return result
	}
}


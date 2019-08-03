package leetcode

/*
Write a program to solve a Sudoku puzzle by filling the empty cells.

A sudoku solution must satisfy all of the following rules:

Each of the digits 1-9 must occur exactly once in each row.
Each of the digits 1-9 must occur exactly once in each column.
Each of the the digits 1-9 must occur exactly once in each of the 9 3x3 sub-boxes of the grid.
Empty cells are indicated by the character '.'.

[
["5","3",".",".","7",".",".",".","."],
["6",".",".","1","9","5",".",".","."],
[".","9","8",".",".",".",".","6","."],
["8",".",".",".","6",".",".",".","3"],
["4",".",".","8",".","3",".",".","1"],
["7",".",".",".","2",".",".",".","6"],
[".","6",".",".",".",".","2","8","."],
[".",".",".","4","1","9",".",".","5"],
[".",".",".",".","8",".",".","7","9"]]

A sudoku puzzle...

[
["5","3","4","6","7","8","9","1","2"],
["6","7","2","1","9","5","3","4","8"],
["1","9","8","3","4","2","5","6","7"],
["8","5","9","7","6","1","4","2","3"],
["4","2","6","8","5","3","7","9","1"],
["7","1","3","9","2","4","8","5","6"],
["9","6","1","5","3","7","2","8","4"],
["2","8","7","4","1","9","6","3","5"],
["3","4","5","2","8","6","1","7","9"]]
...and its solution numbers marked in red.

Note:

The given board contain only digits 1-9 and the character '.'.
You may assume that the given Sudoku puzzle will have a single unique solution.
The given board size is always 9x9.*/

// 数独,需要给出解法

func solveSudoku(board [][]byte) {
	partialSolver(board, 0, 0)
}

func next(i, j int) (ni, nj int) {
	nj = j + 1
	if nj == 9 {
		nj = 0
		ni = i + 1
	} else {
		ni = i
	}
	return
}

var entries = []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}

func inCell(board [][]byte, i, j int, e byte) bool {
	ci, cj := i/3, j/3
	for x := range entries {
		if board[3*ci+x/3][3*cj+x%3] == e {
			return true
		}
	}
	return false
}

func inRow(board [][]byte, i int, e byte) bool {
	for x := range entries {
		if board[i][x] == e {
			return true
		}
	}
	return false
}

func inCol(board [][]byte, j int, e byte) bool {
	for x := range entries {
		if board[x][j] == e {
			return true
		}
	}
	return false
}

func partialSolver(board [][]byte, i, j int) bool {
	if i == 9 {
		return true
	}
	ni, nj := next(i, j)
	if board[i][j] != '.' {
		return partialSolver(board, ni, nj)
	}
	for _, e := range entries {
		if !inCell(board, i, j, e) && !inRow(board, i, e) && !inCol(board, j, e) {
			board[i][j] = e
			if partialSolver(board, ni, nj) {
				return true
			}
		}
	}
	board[i][j] = '.'
	return false
}

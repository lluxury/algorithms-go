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

func reject_52(c [][]rune, row, col int) bool {
    for i := 0; i < col; i++ {
        if c[row][i] == 'Q' {
            return true
        }
    }

    for i, j := row, col; i >= 0 && j >= 0; {
        if c[i][j] == 'Q' {
            return true
        }
        i--
        j--
    }

    for i, j := row, col; i < len(c) && j >= 0; {
        if c[i][j] == 'Q' {
            return true
        }
        i++
        j--
    }
    return false
}

func totalNQueens(n int) int {
    var result [][]string

    t := make([][]rune, n)
    for i := range t {
        t[i] = make([]rune, n)
        for j := range t[i] {
            t[i][j] = '.'
        }
    }

    backtracking_52(&result, t, n, 0)
    return len(result)
}

func backtracking_52(res *[][]string, c [][]rune, r, column int) {
    //fmt.Printf("r:%d c:%d \n", r, column)
    // accept
    if r == 0 {
        var arr []string
        for _, v := range c {
            arr = append(arr, string(v))
        }
        *res = append(*res, arr)
        //fmt.Printf("Accept %d %v\n", r, c)
        return
    }

    // next
    for i := 0; i < len(c); i++ {
        // reject
        if !reject_52(c, i, column) {
            c[i][column] = 'Q'
            backtracking_52(res, c, r-1, column+1)
            c[i][column] = '.'
        }
    }
}

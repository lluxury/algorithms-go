package leetcode

/*
The n-queens puzzle is the problem of placing n queens on an n×n chessboard such that no two queens attack each other.



Given an integer n, return all distinct solutions to the n-queens puzzle.

Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.' both indicate a queen and an empty space respectively.

Example:

Input: 4
Output: [
 [".Q..",  // Solution 1
  "...Q",
  "Q...",
  "..Q."],

 ["..Q.",  // Solution 2
  "Q...",
  "...Q",
  ".Q.."]
]
Explanation: There exist two distinct solutions to the 4-queens puzzle as shown above.*/

// 比较复杂,需要时间分解,测试也没写

func reject(c [][]rune, row, col int) bool {
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

func solveNQueens(n int) [][]string {
    var result [][]string

    t := make([][]rune, n)
    for i := range t {
        t[i] = make([]rune, n)
        for j := range t[i] {
            t[i][j] = '.'
        }
    }

    backtracking(&result, t, n, 0)
    return result
}
func backtracking(res *[][]string, c [][]rune, r, column int) {
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
        if !reject(c, i, column) {
            c[i][column] = 'Q'
            backtracking(res, c, r-1, column+1)
            c[i][column] = '.'
        }
    }
}

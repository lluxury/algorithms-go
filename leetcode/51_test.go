package leetcode

import (
    "testing"
    "github.com/Chyroc/algorithms-go/test"
)

func Test_51(t *testing.T) {
    test.Runs(t, solveNQueens, []*test.Case{
        {Input: `4`, Output: `[[..Q.,Q...,...Q,.Q..],[.Q..,...Q,Q...,..Q.]]`},
        } )
}

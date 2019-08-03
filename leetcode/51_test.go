package leetcode

import (
	"github.com/lluxury/algorithms-go/test"
	"testing"
)

func Test_51(t *testing.T) {
	test.Runs(t, solveNQueens, []*test.Case{
		{Input: `4`, Output: `[[..Q.,Q...,...Q,.Q..],[.Q..,...Q,Q...,..Q.]]`},
	})
}

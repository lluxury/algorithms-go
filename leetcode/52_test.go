package leetcode

import (
    "testing"
    "github.com/Chyroc/algorithms-go/test"
)

func Test_52(t *testing.T) {
    test.Runs(t, totalNQueens, []*test.Case{
        {Input: `1`, Output: "1"},
        {Input: `2`, Output: "0"},
        {Input: `4`, Output: "2"},
        {Input: `5`, Output: "10"},
        })
}

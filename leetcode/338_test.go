package leetcode

import (
    "testing"
    "github.com/Chyroc/algorithms-go/test"
)

func Test_338(t *testing.T) {
    test.Runs(t, countBits, []*test.Case{
    {Input: `2`, Output: `[0,1,1]`},
    {Input: `5`, Output: `[0,1,1,2,1,2]`},
    })
}

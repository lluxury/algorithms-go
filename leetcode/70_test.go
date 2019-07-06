package leetcode

import (
    "testing"
    "github.com/Chyroc/algorithms-go/test"
)

func Test_70(t *testing.T) {
    test.Runs(t, climbStairs, []*test.Case{
        {Input: `2`, Output: "2"},
        {Input: `3`, Output: "3"},
        })
}

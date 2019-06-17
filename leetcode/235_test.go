package leetcode

import (
    "github.com/Chyroc/algorithms-go/test"
    "testing"
)

func Test_235(t *testing.T) {
    test.Runs(t, lowestCommonAncestor, []*test.Case{
        {Input: `(6, (0,2,(3,4,5)), (7,8,9)) \n 2 \n 8`, Output: `(6, (0,2,(3,4,5)), (7,8,9))`},
        {Input: `(6, (0,2,(3,4,5)), (7,8,9)) \n 2 \n 4`, Output: `(3,4,5)`},
        // {Input: `(6, (0,2,(3,4,5)), (7,8,9)) \n 2 \n 4`, Output: `2`},
        })
}


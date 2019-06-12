package leetcode

import (
    "testing"
    "github.com/Chyroc/algorithms-go/lib"
    "github.com/Chyroc/algorithms-go/test"
    // "github.com/stretchr/testify/assert"
)

func Test_25(t *testing.T) {
    for _, f := range []func(head *lib.ListNode, k int) *lib.ListNode{reverseKGroup} {
        test.Runs(t, f, []*test.Case{
            {Input: (`1->2->3->4->5 \n 2`), Output: `2->1->4->3->5`},
            {Input: (`1->2->3->4->5 \n 3`), Output: `3->2->1->4->5`},
            })
    }

}

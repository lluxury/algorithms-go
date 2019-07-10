package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lluxury/algorithms-go/lib"
	"github.com/lluxury/algorithms-go/test"
)

func Test_23(t *testing.T) {
	test.Runs(t, mergeKLists, []*test.Case{
		{Input: `[]`, Output: ``},
		{Input: `[1->4->5, 1->3->4, 2->6]`, Output: `1->1->2->3->4->4->5->6`},
		{Input: `[0->2->5]`, Output: `0->2->5`},
		{Input: `[1->2->3, 4->5->6->7]`, Output: `1->2->3->4->5->6->7`},
	})

	assert.Nil(t, mergeKLists([]*lib.ListNode{nil}))
}

package leetcode

import (
	"github.com/Chyroc/algorithms-go/test"
	"testing"
)

func Test_203(t *testing.T) {
	test.Runs(t, removeElements, []*test.Case{
		{Input: `1->2->6->3->4->5->6 \n 6`, Output: `1->2->3->4->5`},
	})
}

package leetcode

import (
	"github.com/Chyroc/algorithms-go/test"
	"testing"
)

func Test_92(t *testing.T) {
	test.Runs(t, reverseBetween, []*test.Case{
		{Input: `1->2->3->4->5 \n 2 \n 4`, Output: `1->4->3->2->5`},
	})
}

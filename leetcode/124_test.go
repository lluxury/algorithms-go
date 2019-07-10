package leetcode

import (
	"github.com/lluxury/algorithms-go/test"
	"testing"
)

func Test_124(t *testing.T) {
	test.Runs(t, maxPathSum, []*test.Case{
		{Input: `(-3,,)`, Output: `-3`},
		{Input: `(1,2,3)`, Output: `6`},
		{Input: `(-10,9,(20,15,7))`, Output: `42`},
	})
}

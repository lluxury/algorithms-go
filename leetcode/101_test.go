package leetcode

import (
	"github.com/lluxury/algorithms-go/test"
	"testing"
)

func Test_101(t *testing.T) {
	test.Runs(t, isSymmetric, []*test.Case{
		{Input: `(1,(2,3,4),(2,4,3))`, Output: `true`},
		{Input: `(1,(2,,3),(2,,3))`, Output: `false`},
	})
}

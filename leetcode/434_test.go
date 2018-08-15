package leetcode

import (
	"github.com/Chyroc/algorithms-go/test"
	"testing"
)

func Test_434(t *testing.T) {
	test.Runs(t, countSegments, []*test.Case{
		{Input: `"Hello, my name is John"`, Output: `5`},
		//{Input: `a b c`, Output: `3`},
		//{Input: `a b   c`, Output: `3`},
		//{Input: `a b   c  `, Output: `3`},
		//{Input: `  a b   c  `, Output: `3`},
	})
}

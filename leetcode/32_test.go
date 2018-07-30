package leetcode

import (
	"testing"

	"github.com/Chyroc/algorithms-go/test"
)

func Test_32(t *testing.T) {
	test.Runs(t, longestValidParentheses, []*test.Case{
		{Input: `(()`, Output: `2`},
		{Input: `()`, Output: `2`},
		{Input: `)()())`, Output: `4`},
	})
}

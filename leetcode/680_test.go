package leetcode

import (
	"github.com/Chyroc/algorithms-go/test"
	"testing"
)

func Test_680(t *testing.T) {
	test.Runs(t, validPalindrome, []*test.Case{
		{Input: `abca`, Output: `true`},
		{Input: `aba`, Output: `true`},
		{Input: `abaa`, Output: `true`},
		{Input: `acbaaaa`, Output: `false`},
	})
}

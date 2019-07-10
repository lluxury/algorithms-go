package leetcode

import (
	"testing"

	"github.com/lluxury/algorithms-go/test"
)

func Test_459(t *testing.T) {
	test.Runs(t, repeatedSubstringPattern, []*test.Case{
		{Input: `abab`, Output: `true`},
		{Input: `aba`, Output: `false`},
		{Input: `abcabcabcabc`, Output: `true`},
	})
}

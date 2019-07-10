package leetcode

import (
	"github.com/lluxury/algorithms-go/test"
	"testing"
)

func Test_234(t *testing.T) {
	test.Runs(t, isPalindrome_234, []*test.Case{
		{Input: `1->2->3`, Output: `false`},
		{Input: `1->2->2->1`, Output: `true`},
	})
}

package leetcode

import (
	"testing"

	"github.com/lluxury/algorithms-go/test"
)

func Test_409(t *testing.T) {
	test.Runs(t, longestPalindrome, []*test.Case{
		{Input: `abccccdd`, Output: `7`},
	})
}

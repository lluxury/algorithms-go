package leetcode

import (
	"github.com/lluxury/algorithms-go/test"
	"testing"
)

func Test_784(t *testing.T) {
	test.Runs(t, letterCasePermutation, []*test.Case{
		{Input: `a1`, Output: `[a1, A1]`},
		{Input: `a11`, Output: `[a11, A11]`},
		{Input: `a1b2`, Output: `[a1b2, a1B2, A1b2, A1B2]`},
		{Input: "3z4", Output: "[3z4, 3Z4]"},
		{Input: "12345", Output: "[12345]"},
	})
}

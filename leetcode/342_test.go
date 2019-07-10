package leetcode

import (
	"testing"

	"github.com/lluxury/algorithms-go/test"
)

func Test_342(t *testing.T) {
	test.Runs(t, isPowerOfFour, []*test.Case{
		{Input: `16`, Output: `true`},
		{Input: `5`, Output: `false`},
		{Input: `-2147483648`, Output: `false`},
	})
}

package leetcode

import (
	"github.com/lluxury/algorithms-go/test"
	"testing"
)

func Test_367(t *testing.T) {
	test.Runs(t, isPerfectSquare, []*test.Case{
		{Input: `16`, Output: `true`},
		{Input: `14`, Output: `false`},
	})
}

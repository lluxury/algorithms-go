package leetcode

import (
	"testing"

	"github.com/lluxury/algorithms-go/test"
)

func Test_456(t *testing.T) {
	test.Runs(t, find132pattern, []*test.Case{
		{Input: `[1, 2, 3, 4]`, Output: `False`},
		{Input: `[3, 1, 4, 2]`, Output: `True`},
		{Input: `[-1, 3, 2, 0]`, Output: `True`},
	})
}

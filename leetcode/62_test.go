package leetcode

import (
	"testing"

	"github.com/lluxury/algorithms-go/test"
)

func Test_62(t *testing.T) {
	test.Runs(t, uniquePaths, []*test.Case{
		{Input: `3 \n 2`, Output: `3`},
		{Input: `7 \n 3`, Output: `28`},
	})
}

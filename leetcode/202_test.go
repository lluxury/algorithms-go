package leetcode

import (
	"testing"

	"github.com/Chyroc/algorithms-go/test"
)

func Test_202(t *testing.T) {
	test.Runs(t, isHappy, []*test.Case{
		//{Input: `19`, Output: `true`},
		//{Input: `2`, Output: `false`},
		{Input: `23`, Output: `true`},
	})
}


package leetcode

import (
	"github.com/Chyroc/algorithms-go/test"
	"testing"
)

func Test_541(t *testing.T) {
	test.Runs(t, reverseStr, []*test.Case{
		{Input: `abcdefg \n 2`, Output: "bacdfeg"},
	})
}

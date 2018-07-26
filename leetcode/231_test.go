package leetcode

import (
	"github.com/Chyroc/algorithms-go/test"
	"testing"
)

func Test_231(t *testing.T) {
	test.Runs(t, isPowerOfTwo, []*test.Case{
		{Input: `2`, Output: `true`},
	})
}

package leetcode

import (
	"testing"

	"github.com/lluxury/algorithms-go/test"
)

func Test_606(t *testing.T) {
	test.Runs(t, tree2str, []*test.Case{
		{Input: `(1, (2,4,), 3)`, Output: `1(2(4))(3)`},
		{Input: `(1, (2,,4), 3)`, Output: `1(2()(4))(3)`},
	})
}

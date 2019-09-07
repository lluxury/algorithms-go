package leetcode

import (
	"testing"

	"github.com/lluxury/algorithms-go/test"
)

func Test_629(t *testing.T) {
	test.Runs(t, kInversePairs, []*test.Case{
	    {Input: `3\n1`, Output: `2`},
	})
}


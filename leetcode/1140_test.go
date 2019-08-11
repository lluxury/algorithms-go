package leetcode

import (
	"testing"

	"github.com/lluxury/algorithms-go/test"
)

func Test_1140(t *testing.T) {
	test.Runs(t, stoneGameII, []*test.Case{
	    {Input: `[2,7,9,4,4]`, Output: `10`},
	})
}


package leetcode

import (
	"testing"

	"github.com/lluxury/algorithms-go/test"
)

func Test_805(t *testing.T) {
	test.Runs(t, splitArraySameAverage, []*test.Case{
	    {Input: `[1,2,3,4,5,6,7,8]`, Output: `true`},
	})
}


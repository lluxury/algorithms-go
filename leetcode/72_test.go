package leetcode

import (
	"testing"

	"github.com/lluxury/algorithms-go/test"
)

func Test_72(t *testing.T) {
	test.Runs(t, minDistance, []*test.Case{
        {Input: `"horse"\n"ros"`, Output: `3`},
        {Input: `"b"\n""`, Output: `1`},
        {Input: `"intention"\n"execution"`, Output: `5`},
	})
}


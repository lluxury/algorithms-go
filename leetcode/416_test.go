package leetcode

import (
	"testing"

	"github.com/lluxury/algorithms-go/test"
)

func Test_416(t *testing.T) {
	test.Runs(t, canPartition, []*test.Case{
        {Input: `[1, 5, 11, 5]`, Output: `true`},
	    {Input: `[1, 2, 3, 5]`, Output: `false`},
	})
}


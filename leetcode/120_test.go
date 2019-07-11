package leetcode

import (
	"testing"

	"github.com/lluxury/algorithms-go/test"
)

func Test_120(t *testing.T) {
	test.Runs(t, minimumTotal, []*test.Case{
	    {Input: `[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]`, Output: `11`},
	})
}


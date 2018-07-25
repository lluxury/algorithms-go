package leetcode

import (
	"github.com/Chyroc/algorithms-go/test"
	"testing"
)

func Test_111(t *testing.T) {
	test.Runs(t, minDepth, []*test.Case{
		//{Input: `(3, 9, (20,15,7))`, Output: `2`},
		//{Input: `(1, (2,4,), (3,,5))`, Output: `3`},
		//{Input: `(1,2,)`, Output: `2`},
		{Input: `(1,(2,(3,(4,5,))))`, Output: `5`},
	})
}

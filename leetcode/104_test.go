package leetcode

import (
	"github.com/Chyroc/algorithms-go/test"
	"testing"
)

func Test_104(t *testing.T) {
	test.Runs(t, maxDepth, []*test.Case{
		{Input: `(1,(2,3,4),(2,4,3))`, Output: `3`},
		{Input: `(1,(2,,3),(2,(1,2,3),3))`, Output: `4`},
	})
}

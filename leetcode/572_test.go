package leetcode

import (
	"testing"

	"github.com/Chyroc/algorithms-go/test"
)

func Test_572(t *testing.T) {
	test.Runs(t, isSubtree, []*test.Case{
		{Input: `(3, (4,1,2), 5) \n (4,1,2)`, Output: `true`},
		{Input: `(3, (4,1,(2,0,)), 5) \n (4,1,2)`, Output: `false`},
		{Input: `(3,(4,1,),(5,2,)) \n (3,1,2)`, Output: `false`},
	})
}

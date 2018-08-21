package leetcode

import (
	"github.com/Chyroc/algorithms-go/test"
	"testing"
)

func Test_653(t *testing.T) {
	test.Runs(t, findTarget, []*test.Case{
		{Input: `(2,1,3) \n 4`, Output: `true`},
		{Input: `(5,(3,2,4),(6,,7)) \n 9`, Output: `true`},
	})
}

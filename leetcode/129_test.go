package leetcode

import (
	"testing"

	"github.com/lluxury/algorithms-go/test"
)

func Test_129(t *testing.T) {
	test.Runs(t, sumNumbers, []*test.Case{
		{Input: `(1,2,3)`, Output: `25`},
		{Input: `(4,(9,5,1),0)`, Output: `1026`},
	})
}

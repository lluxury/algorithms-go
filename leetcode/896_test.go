package leetcode

import (
	"testing"

	"github.com/lluxury/algorithms-go/test"
)

func Test_896(t *testing.T) {
	test.Runs(t, isMonotonic, []*test.Case{
		{Input: `[1,2,2,3]`, Output: `true`},
		{Input: `[6,5,4,4]`, Output: `true`},
		{Input: `[1,3,2]`, Output: `false`},
		{Input: `[1,2,4,5]`, Output: `true`},
		{Input: `[1,1,1]`, Output: `true`},
		{Input: `[1,2,5,3,3]`, Output: `false`},
	})
}

package leetcode

import (
	"testing"

	"github.com/Chyroc/algorithms-go/test"
)

func Test_69(t *testing.T) {
	test.Runs(t, mySqrt, []*test.Case{
		{Input: `4`, Output: `2`},
		{Input: `8`, Output: `2`},
	})
}

package leetcode

import (
	"testing"

	"github.com/Chyroc/algorithms-go/test"
)

func Test_515(t *testing.T) {
	test.Runs(t, largestValues, []*test.Case{
		{Input: `(1,(3,5,3),(2,,9))`, Output: `[1,3,9]`},
	})
}

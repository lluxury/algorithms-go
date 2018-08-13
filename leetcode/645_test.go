package leetcode

import (
	"github.com/Chyroc/algorithms-go/test"
	"testing"
)

func Test_645(t *testing.T) {
	test.Runs(t, findErrorNums, []*test.Case{
		{Input: `[1,2,2,4]`, Output: `[2, 3]`},
		{Input: `[1,1]`, Output: `[1, 2]`},
		{Input: `[2,2]`, Output: `[2, 1]`},
	})
}

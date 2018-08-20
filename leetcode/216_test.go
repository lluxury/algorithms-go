package leetcode

import (
	"github.com/Chyroc/algorithms-go/test"
	"testing"
)

func Test_216(t *testing.T) {
	test.Runs(t, combinationSum3, []*test.Case{
		{Input: `3 \n 7`, Output: `[[1,2,4]]`},
		{Input: `3 \n 9`, Output: `[[2,3,4], [1,3,5], [1,2,6]]`},
	})
}

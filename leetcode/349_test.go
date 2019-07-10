package leetcode

import (
	"github.com/lluxury/algorithms-go/test"
	"testing"
)

func Test_349(t *testing.T) {
	test.Runs(t, intersection, []*test.Case{
		{Input: `[1,2,2,1] \n [2,2]`, Output: `[2]`},
		{Input: `[4,9,5] \n [9,4,9,8,4]`, Output: `[9,4]`},
	})
}

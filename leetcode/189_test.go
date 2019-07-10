package leetcode

import (
	"testing"

	"github.com/lluxury/algorithms-go/test"
)

func Test_189(t *testing.T) {
	test.Runs(t, rotate, []*test.Case{
		{Input: `[1,2,3,4,5,6,7] \n 3`, Output: `[5,6,7,1,2,3,4]`},
		{Input: `[-1] \n 2`, Output: `[-1]`},
	})
}

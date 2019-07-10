package leetcode

import (
	"github.com/lluxury/algorithms-go/test"
	"testing"
)

func Test_113(t *testing.T) {
	test.Runs(t, pathSum, []*test.Case{
		{Input: `(5, (4,(11,7,2)), (8,14,(4,5,1))) \n 22`, Output: `[ [5,4,11,2], [5,8,4,5] ]`},
	})
}

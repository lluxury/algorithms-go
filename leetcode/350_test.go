package leetcode

import (
	"github.com/lluxury/algorithms-go/test"
	"testing"
)

func Test_350(t *testing.T) {
	test.Runs(t, intersect, []*test.Case{
		{Input: `[1, 2, 2, 1] \n [2, 2]`, Output: "[2, 2]"},
	})
}

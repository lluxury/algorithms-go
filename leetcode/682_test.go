package leetcode

import (
	"testing"
	"github.com/Chyroc/algorithms-go/test"
)

func Test_682(t *testing.T) {
	test.Runs(t, calPoints, []*test.Case{
		{Input: `[5, 2, C, D, +]`, Output: "30"},
		{Input: "[5, -2, 4, C, D, 9, +, +]", Output: "27"},
	})
}

package leetcode

import (
	"github.com/Chyroc/algorithms-go/test"
	"testing"
)

func Test_198(t *testing.T) {
	test.Runs(t, rob, []*test.Case{
		{Input: `[4, 2, 3]`, Output: "7"},
		{Input: `[1, 3, 1]`, Output: "3"},
		{Input: `[1, 1, 1, 1]`, Output: "2"},
	})
}

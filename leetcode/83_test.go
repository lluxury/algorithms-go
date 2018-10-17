package leetcode

import (
	"github.com/Chyroc/algorithms-go/test"
	"testing"
)

func Test_83(t *testing.T) {
	test.Runs(t, deleteDuplicates, []*test.Case{
		{Input: `1 -> 1 -> 2`, Output: "1 -> 2"},
		{Input: ``, Output: ""},
	})
}

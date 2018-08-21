package leetcode

import (
	"testing"
	"github.com/Chyroc/algorithms-go/test"
)

func Test_476(t *testing.T) {
	test.Runs(t, findComplement, []*test.Case{
		{Input: `5`, Output: "2"},
		{Input: `1`, Output: "0"},
	})
}

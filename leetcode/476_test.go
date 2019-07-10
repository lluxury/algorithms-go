package leetcode

import (
	"github.com/lluxury/algorithms-go/test"
	"testing"
)

func Test_476(t *testing.T) {
	test.Runs(t, findComplement, []*test.Case{
		{Input: `5`, Output: "2"},
		{Input: `1`, Output: "0"},
	})
}

package leetcode

import (
	"github.com/lluxury/algorithms-go/test"
	"testing"
)

func Test_168(t *testing.T) {
	test.Runs(t, convertToTitle, []*test.Case{
		{Input: `26`, Output: "Z"},
		{Input: `28`, Output: "AB"},
	})
}

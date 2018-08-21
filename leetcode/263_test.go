package leetcode

import (
	"github.com/Chyroc/algorithms-go/test"
	"testing"
)

func Test_263(t *testing.T) {
	test.Runs(t, isUgly, []*test.Case{
		{Input: "6", Output: "true"},
		{Input: "8", Output: "true"},
		{Input: "14", Output: "false"},
	})
}

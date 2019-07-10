package leetcode

import (
	"github.com/lluxury/algorithms-go/test"
	"testing"
)

func Test_242(t *testing.T) {
	test.Runs(t, isAnagram, []*test.Case{
		{Input: `abc \n bac`, Output: "true"},
		{Input: `aacc \n aaac`, Output: "false"},
	})
}

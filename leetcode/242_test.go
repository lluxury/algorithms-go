package leetcode

import (
	"testing"
	"github.com/Chyroc/algorithms-go/test"
)

func Test_242(t *testing.T) {
	test.Runs(t, isAnagram, []*test.Case{
		{Input: `abc \n bac`, Output: "true"},
		{Input: `aacc \n aaac`, Output: "false"},
	})
}

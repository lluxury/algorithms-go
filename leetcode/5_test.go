package leetcode

import (
	"github.com/Chyroc/algorithms-go/test"
	"testing"
)

func Test_5(t *testing.T) {
	test.Runs(t, longestPalindrome_5, []*test.Case{
		{Input: "", Output: ""},
		{Input: "a", Output: "a"},
		{Input: "abba", Output: "abba"},
		{Input: "cbbd", Output: "bb"},
	})
}

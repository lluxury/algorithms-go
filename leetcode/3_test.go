package leetcode

import (
	"testing"

	"github.com/Chyroc/algorithms-go/test"
)

func Test_3(t *testing.T) {
	test.Runs(t, lengthOfLongestSubstring, []*test.Case{
		{Input: `abcabcbb`, Output: "3"},
		{Input: `bbbbb`, Output: "1"},
		{Input: `pwwkew`, Output: "3"},
	})
}

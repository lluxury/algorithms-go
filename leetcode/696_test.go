package leetcode

import (
	"testing"
	"github.com/Chyroc/algorithms-go/test"
)

func Test_696(t *testing.T) {
	test.Runs(t, countBinarySubstrings, []*test.Case{
		{Input: `00110011`, Output: "6"},
		{Input: `00110`, Output: "3"},
		{Input: `10101`, Output: "4"},
	})
}

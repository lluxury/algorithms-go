package leetcode

import (
	"github.com/Chyroc/algorithms-go/test"
	"testing"
)

func Test_717(t *testing.T) {
	test.Runs(t, isOneBitCharacter, []*test.Case{
		{Input: `[1, 0, 0]`, Output: "true"},
		{Input: `[0, 0]`, Output: "true"},
		{Input: `[1, 1, 1, 0]`, Output: "false"},
	})
}

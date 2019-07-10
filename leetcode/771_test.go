package leetcode

import (
	"github.com/lluxury/algorithms-go/test"
	"testing"
)

func Test_771(t *testing.T) {
	test.Runs(t, numJewelsInStones, []*test.Case{
		{Input: `aA\naAAbbbb`, Output: "3"},
		{Input: `z\nZZ`, Output: "0"},
	})
}

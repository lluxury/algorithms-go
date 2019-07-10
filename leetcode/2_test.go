package leetcode

import (
	"testing"

	"github.com/lluxury/algorithms-go/test"
)

func Test_2(t *testing.T) {
	test.Runs(t, addTwoNumbers, []*test.Case{
		{Input: `5 -> 5 \n 5`, Output: "0 -> 6"},
		{Input: `2 -> 4 -> 3 \n 5 -> 6 -> 4`, Output: "7 -> 0 -> 8"},
		{Input: `5 \n 5`, Output: "0 -> 1"},
	})
}

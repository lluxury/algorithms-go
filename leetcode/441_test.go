package leetcode

import (
	"github.com/lluxury/algorithms-go/test"
	"testing"
)

func Test_441(t *testing.T) {
	test.Runs(t, arrangeCoins, []*test.Case{
		{Input: `5`, Output: "2"},
		{Input: `8`, Output: "3"},
	})
}

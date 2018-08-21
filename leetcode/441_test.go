package leetcode

import (
	"testing"
	"github.com/Chyroc/algorithms-go/test"
)

func Test_441(t *testing.T) {
	test.Runs(t, arrangeCoins, []*test.Case{
		{Input: `5`, Output: "2"},
		{Input: `8`, Output: "3"},
	})
}

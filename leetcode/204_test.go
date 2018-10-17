package leetcode

import (
	"github.com/Chyroc/algorithms-go/test"
	"testing"
)

func Test_204(t *testing.T) {
	test.Runs(t, countPrimes, []*test.Case{
		{Input: `10`, Output: "4"},
		{Input: `0`, Output: "0"},
	})
}

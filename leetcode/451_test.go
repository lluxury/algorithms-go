package leetcode

import (
	"testing"

	"github.com/lluxury/algorithms-go/test"
)

func Test_451(t *testing.T) {
	test.Runs(t, frequencySort, []*test.Case{
		{Input: `tree`, Output: `eetr`},
	})
}

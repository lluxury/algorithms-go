package leetcode

import (
	"testing"

	"github.com/lluxury/algorithms-go/test"
)

func Test_868(t *testing.T) {
	test.Runs(t, binaryGap, []*test.Case{
		{Input: `22`, Output: `2`},
		{Input: `5`, Output: `2`},
		{Input: `6`, Output: `1`},
		{Input: `8`, Output: `0`},
	})
}

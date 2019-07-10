package leetcode

import (
	"testing"

	"github.com/lluxury/algorithms-go/test"
)

func Test_781(t *testing.T) {
	test.Runs(t, numRabbits, []*test.Case{
		{Input: `[1,1,2]`, Output: `5`},
		{Input: `[]`, Output: `0`},
	})
}

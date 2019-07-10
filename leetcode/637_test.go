package leetcode

import (
	"testing"

	"github.com/lluxury/algorithms-go/test"
)

func Test_637(t *testing.T) {
	test.Runs(t, averageOfLevels, []*test.Case{
		{Input: `(3,9,(20,15,7))`, Output: `[3,14.5,11]`},
		{Input: `(5,2,-3)`, Output: `[5,-0.5]`},
	})
}

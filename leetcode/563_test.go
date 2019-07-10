package leetcode

import (
	"testing"

	"github.com/lluxury/algorithms-go/test"
)

func Test_563(t *testing.T) {
	test.Runs(t, findTilt, []*test.Case{
		{Input: `(1, 2, 3)`, Output: `1`},
	})
}

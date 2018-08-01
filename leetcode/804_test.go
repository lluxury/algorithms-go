package leetcode

import (
	"testing"

	"github.com/Chyroc/algorithms-go/test"
)

func Test_804(t *testing.T) {
	test.Runs(t, uniqueMorseRepresentations, []*test.Case{
		{Input: `[gin, zen, gig, msg]`, Output: `2`},
		{Input: `[a]`, Output: `1`},
	})
}

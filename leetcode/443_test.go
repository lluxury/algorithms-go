package leetcode

import (
	"testing"

	"github.com/Chyroc/algorithms-go/test"
)

func Test_443(t *testing.T) {
	test.Runs(t, compress, []*test.Case{
		{Input: `[a,a,b,b,c,c,c]`, Output: `6`},
		{Input: `[a]`, Output: `1`},
		{Input: `[a,b,b,b,b,b,b,b,b,b,b,b,b]`, Output: `4`},
	})
}

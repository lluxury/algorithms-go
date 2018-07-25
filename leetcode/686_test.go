package leetcode

import (
	"github.com/Chyroc/algorithms-go/test"
	"testing"
)

func Test_686(t *testing.T) {
	test.Runs(t, repeatedStringMatch, []*test.Case{
		{Input: `abcd \n cdabcdab`, Output: `3`},
		{Input: `a \n aa`, Output: `2`},
		{Input: `abababaaba \n aabaaba`, Output: `2`},
	})
}

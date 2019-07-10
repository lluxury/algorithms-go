package leetcode

import (
	"testing"

	"github.com/lluxury/algorithms-go/test"
)

func Test_187(t *testing.T) {
	test.Runs(t, findRepeatedDnaSequences, []*test.Case{
		{Input: `AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT`, Output: `[AAAAACCCCC,CCCCCAAAAA]`},
		{Input: `AAAAAAAAAA`, Output: ``},
	})
}

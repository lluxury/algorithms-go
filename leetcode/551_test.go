package leetcode

import (
	"testing"
	"github.com/Chyroc/algorithms-go/test"
)

func Test_551(t *testing.T) {
	test.Runs(t, checkRecord, []*test.Case{
		{Input: `PPALLP`, Output: "true"},
		{Input: `PPALLL`, Output: "false"},
	})
}
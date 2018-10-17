package leetcode

import (
	"github.com/Chyroc/algorithms-go/test"
	"testing"
)

func Test_665(t *testing.T) {
	test.Runs(t, checkPossibility, []*test.Case{
		{Input: `[4, 2, 3]`, Output: "true"},
		{Input: `[4, 2, 1]`, Output: "false"},
		{Input: `[3, 4, 2, 3]`, Output: "false"},
		{Input: `[-1, 4, 2, 3]`, Output: "true"},
	})
}

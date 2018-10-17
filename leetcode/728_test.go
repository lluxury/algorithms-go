package leetcode

import (
	"github.com/Chyroc/algorithms-go/test"
	"testing"
)

func Test_728(t *testing.T) {
	test.Runs(t, selfDividingNumbers, []*test.Case{
		{Input: `1\n22`, Output: "[1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22]"},
	})
}

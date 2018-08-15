package leetcode

import (
	"github.com/Chyroc/algorithms-go/test"
	"testing"
)

func Test_747(t *testing.T) {
	test.Runs(t, dominantIndex, []*test.Case{
		//{Input: `[1,0,0,0]`, Output: `0`},
		//{Input: `[1,0]`, Output: `0`},
		{Input: `[0,1]`, Output: `1`},
		//{Input: `[3, 6, 1, 0]`, Output: `1`},
		//{Input: `[1, 2, 3, 4]`, Output: `-1`},
	})
}

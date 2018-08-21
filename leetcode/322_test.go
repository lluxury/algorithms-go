package leetcode

import (
	"testing"

	"github.com/Chyroc/algorithms-go/test"
)

func Test_322(t *testing.T) {
	test.Runs(t, coinChange, []*test.Case{
		{Input: `[1] \n 0`, Output: `0`},
		{Input: `[1, 2, 5] \n 11`, Output: `3`},
		{Input: `[2] \n 3 `, Output: `-1`},
		{Input: `[186,419,83,408] \n 6249`, Output: `20`},
	})
}

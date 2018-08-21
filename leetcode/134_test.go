package leetcode

import (
	"testing"

	"github.com/Chyroc/algorithms-go/test"
)

func Test_134(t *testing.T) {
	test.Runs(t, canCompleteCircuit, []*test.Case{
		{Input: `[1,2,3,4,5] \n [3,4,5,1,2]`, Output: `3`},
		{Input: `[2,3,4] \n [3,4,3]`, Output: `-1`},
		{Input: `[2] \n [2]`, Output: `0`},
	})
}

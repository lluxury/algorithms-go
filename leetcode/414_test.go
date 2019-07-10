package leetcode

import (
	"github.com/lluxury/algorithms-go/test"
	"testing"
)

func Test_414(t *testing.T) {
	test.Runs(t, thirdMax, []*test.Case{
		{Input: `[5,2,4,1,3,6,0]`, Output: `4`},
		{Input: `[1, 1, 2]`, Output: `2`},
		{Input: `[1,2,2,5,3,5]`, Output: `2`},
		{Input: `[3, 2, 1]`, Output: `1`},
		{Input: `[1, 2]`, Output: `2`},
		{Input: `[2, 2, 3, 1]`, Output: `1`},
	})
}

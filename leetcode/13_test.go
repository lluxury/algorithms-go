package leetcode

import (
	"testing"

	"github.com/lluxury/algorithms-go/test"
)

func Test_13(t *testing.T) {
	test.Runs(t, romanToInt, []*test.Case{
		{Input: `III`, Output: `3`},
		{Input: `IV`, Output: `4`},
		{Input: `IX`, Output: `9`},
		{Input: `LVIII`, Output: `58`},
		{Input: `MCMXCIV`, Output: `1994`},
		{Input: `DCXXI`, Output: `621`},
	})
}

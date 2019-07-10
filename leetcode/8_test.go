package leetcode

import (
	"testing"

	"github.com/lluxury/algorithms-go/test"
)

func Test_8(t *testing.T) {
	test.Runs(t, myAtoi, []*test.Case{
		{Input: `42`, Output: `42`},
		{Input: `   -42`, Output: `-42`},
		{Input: `4193 with words`, Output: `4193`},
		{Input: `words and 987`, Output: `0`},
		{Input: `-91283472332`, Output: `-2147483648`},
		{Input: `9223372036854775808`, Output: `2147483647`},
	})
}

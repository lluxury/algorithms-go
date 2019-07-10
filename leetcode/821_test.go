package leetcode

import (
	"testing"

	"github.com/lluxury/algorithms-go/test"
)

func Test_821(t *testing.T) {
	test.Runs(t, shortestToChar, []*test.Case{
		{Input: `loveleetcode \n e`, Output: `[3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0]`},
		{Input: `abaa \n b`, Output: `[1,0,1,2]`},
	})
}

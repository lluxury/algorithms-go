package leetcode

import (
	"github.com/Chyroc/algorithms-go/test"
	"testing"
)

func Test_752(t *testing.T) {
	test.Runs(t, openLock, []*test.Case{
		{Input: `[0201,0101,0102,1212,2002] \n 0202`, Output: `6`},
		{Input: `[8888] \n 0009`, Output: `1`},
		{Input: `[8887,8889,8878,8898,8788,8988,7888,9888] \n 8888`, Output: `-1`},
		{Input: `[0000] \n 8888`, Output: `-1`},
	})
}

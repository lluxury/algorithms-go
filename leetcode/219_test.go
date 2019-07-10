package leetcode

import (
	"github.com/lluxury/algorithms-go/test"
	"testing"
)

func Test_219(t *testing.T) {
	test.Runs(t, containsNearbyDuplicate, []*test.Case{
		{Input: `[1,2,3,1] \n 3`, Output: `true`},
		{Input: `[1,0,1,1] \n 1`, Output: `true`},
		{Input: `[1,2,3,1,2,3] \n 2`, Output: `false`},
	})
}

package leetcode

import (
	"github.com/Chyroc/algorithms-go/test"
	"testing"
)

func Test_100(t *testing.T) {
	test.Runs(t, isSameTree, []*test.Case{
		{Input: `() \n ()`, Output: `true`},
		{Input: `() \n (1)`, Output: `false`},
		{Input: `(1,2,3) \n (1,2,3)`, Output: `true`},
		{Input: `(1,2,(2)) \n (1,2,(2))`, Output: `true`},
		{Input: `(1,2,(2,,2)) \n (1,2,(2,,2))`, Output: `true`},
		{Input: `(1,(1,,1),(22,22,2)) \n (1,(1,,1),(22,22,2))`, Output: `true`},
	})
}

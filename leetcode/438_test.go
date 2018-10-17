package leetcode

import (
	"github.com/Chyroc/algorithms-go/test"
	"testing"
)

func Test_438(t *testing.T) {
	test.Runs(t, findAnagrams, []*test.Case{
		{Input: `cbaebabacd \n abc`, Output: `[0,6]`},
		{Input: `abab \n ab`, Output: `[0,1,2]`},
	})
}

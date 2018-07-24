package leetcode

import (
	"testing"
	"github.com/Chyroc/algorithms-go/test"
)

func Test_700(t *testing.T) {
	test.Runs(t, searchBST, []*test.Case{
		{Input: `(4, (2,1,3), 7) \n 2`, Output: `(2,1,3)`},
	})
}
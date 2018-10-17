package leetcode

import (
	"testing"

	"github.com/Chyroc/algorithms-go/test"
)

func Test_257(t *testing.T) {
	test.Runs(t, binaryTreePaths, []*test.Case{
		{Input: `(1, (2,,5),3)`, Output: `[1->2->5, 1->3]`},
	})
}

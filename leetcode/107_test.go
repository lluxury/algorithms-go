package leetcode

import (
	"github.com/lluxury/algorithms-go/test"
	"testing"
)

func Test_107(t *testing.T) {
	test.Runs(t, levelOrderBottom, []*test.Case{
		{Input: `(3, 9, (20,15,7))`, Output: `[[15,7],[9,20],[3]]`},
	})
}

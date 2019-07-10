package leetcode

import (
	"testing"

	"github.com/lluxury/algorithms-go/test"
)

func Test_112(t *testing.T) {
	test.Runs(t, hasPathSum, []*test.Case{
		{Input: `(5,(4,(11,7,2)),(8,13,(4,,1))) \n 22`, Output: `true`},
	})
}

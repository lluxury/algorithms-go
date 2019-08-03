package leetcode

import (
	"github.com/lluxury/algorithms-go/test"
	"testing"
)

func Test_22(t *testing.T) {
	test.Runs(t, generateParenthesis, []*test.Case{
		{Input: `3`, Output: `[((())),(()()),(())(),()(()),()()()]`},
	})
}

// 有机会还是看一下test的代码吧

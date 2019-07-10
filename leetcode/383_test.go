package leetcode

import (
	"github.com/lluxury/algorithms-go/test"
	"testing"
)

func Test_383(t *testing.T) {
	test.Runs(t, canConstruct, []*test.Case{
		{Input: `a \n b`, Output: `false`},
		{Input: `aa \n ab`, Output: `false`},
		{Input: `aa \n aab`, Output: `true`},
	})
}

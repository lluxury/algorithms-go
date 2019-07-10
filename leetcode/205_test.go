package leetcode

import (
	"github.com/lluxury/algorithms-go/test"
	"testing"
)

func Test_205(t *testing.T) {
	test.Runs(t, isIsomorphic, []*test.Case{
		{Input: `ab \n aa`, Output: `false`},
		//{Input: `egg \n add`, Output: `true`},
		//{Input: `foo \n bar`, Output: `false`},
		//{Input: `paper \n title`, Output: `true`},
	})
}

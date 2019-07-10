package leetcode

import (
	"github.com/lluxury/algorithms-go/test"
	"testing"
)

func Test_290(t *testing.T) {
	test.Runs(t, wordPattern, []*test.Case{
		{Input: `aaa \n aa aa aa aa`, Output: `false`},
		{Input: `abba \n dog cat cat dog`, Output: `true`},
		{Input: `abba \n dog cat cat fish`, Output: `false`},
		{Input: `aaaa \n dog cat cat dog`, Output: `false`},
		{Input: `abba \n dog dog dog dog`, Output: `false`},
	})
}

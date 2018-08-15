package leetcode

import (
	"testing"

	"github.com/Chyroc/algorithms-go/test"
)

func Test_345(t *testing.T) {
	test.Runs(t, reverseVowels, []*test.Case{
		{Input: `hello`, Output: `holle`},
		{Input: `leetcode`, Output: `leotcede`},
	})
}

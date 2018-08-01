package leetcode

import (
	"testing"

	"github.com/Chyroc/algorithms-go/test"
)

func Test_709(t *testing.T) {
	test.Runs(t, toLowerCase, []*test.Case{
		{Input: `Hello`, Output: `hello`},
		{Input: `here`, Output: `here`},
		{Input: `LOVELY`, Output: `lovely`},
	})
}

package leetcode

import (
	"testing"

	"github.com/Chyroc/algorithms-go/test"
)

func Test_922(t *testing.T) {
	test.Runs(t, sortArrayByParityII, []*test.Case{
		{Input: `[4,2,5,7]`, Output: `[4,5,2,7]`},
	})
}

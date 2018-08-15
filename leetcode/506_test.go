package leetcode

import (
	"testing"

	"github.com/Chyroc/algorithms-go/test"
)

func Test_506(t *testing.T) {
	test.Runs(t, findRelativeRanks, []*test.Case{
		{Input: `[5,4,3,2,1]`, Output: `[Gold, Silver, Bronze, 4, 5]`},
	})
}

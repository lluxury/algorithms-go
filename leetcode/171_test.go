package leetcode

import (
	"github.com/Chyroc/algorithms-go/test"
	"testing"
)

func Test_171(t *testing.T) {
	test.Runs(t, titleToNumber, []*test.Case{
		{Input: "Z", Output: `26`},
		{Input: "AB", Output: `28`},
	})
}

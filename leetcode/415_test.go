package leetcode

import (
	"testing"

	"github.com/lluxury/algorithms-go/test"
)

func Test_415(t *testing.T) {
	test.Runs(t, addStrings, []*test.Case{
		{Input: `0 \n 0`, Output: `0`},
		{Input: `3876620623801494171 \n 6529364523802684779`, Output: `10405985147604178950`},
	})
}

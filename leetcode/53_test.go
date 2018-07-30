package leetcode

import (
	"testing"

	"github.com/Chyroc/algorithms-go/test"
)

func Test_53(t *testing.T) {
	test.Runs(t, maxSubArray, []*test.Case{
		{Input: `[-2,1,-3,4,-1,2,1,-5,4]`, Output: `6`},
	})
}

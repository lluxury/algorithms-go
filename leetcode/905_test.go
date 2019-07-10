package leetcode

import (
	"testing"

	"github.com/lluxury/algorithms-go/test"
)

func Test_905(t *testing.T) {
	test.Runs(t, sortArrayByParity, []*test.Case{
		{Input: `[3,1,2,4]`, Output: `[2,4,1,3]`},
	})
}

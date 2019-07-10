package leetcode

import (
	"testing"

	"github.com/lluxury/algorithms-go/test"
)

func Test_313(t *testing.T) {
	test.Runs(t, nthSuperUglyNumber, []*test.Case{
		{Input: `12 \n [2,7,13,19]`, Output: `32`},
	})
}

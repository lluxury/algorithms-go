package leetcode

import (
	"testing"

	"github.com/lluxury/algorithms-go/test"
)

func Test_598(t *testing.T) {
	test.Runs(t, maxCount, []*test.Case{
		{Input: `3 \n 3 \n [[2,2],[3,3]]`, Output: `4`},
		{Input: `3 \n 3 \n []`, Output: `9`},
		{Input: `39999 \n 39999 \n [[19999,19999]]`, Output: `399960001`},
	})
}

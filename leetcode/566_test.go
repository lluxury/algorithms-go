package leetcode

import (
	"testing"

	"github.com/lluxury/algorithms-go/test"
)

func Test_566(t *testing.T) {
	test.Runs(t, matrixReshape, []*test.Case{
		{Input: `[[1, 2], [3, 4]] \n 1 \n 4`, Output: `[[1, 2, 3, 4]]`},
		{Input: `[[1, 2], [3, 4]] \n 2 \n 4`, Output: `[[1, 2], [3, 4]]`},
	})
}

package leetcode

import (
	"testing"

	"github.com/Chyroc/algorithms-go/test"
)

func Test_404(t *testing.T) {
	test.Runs(t, sumOfLeftLeaves, []*test.Case{
		{Input: `(3, 9, (20,15,7))`, Output: `24`},
	})
}

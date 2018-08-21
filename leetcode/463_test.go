package leetcode

import (
	"testing"

	"github.com/Chyroc/algorithms-go/test"
)

func Test_463(t *testing.T) {
	test.Runs(t, islandPerimeter, []*test.Case{
		{Input: `[[0, 1, 0, 0], [1, 1, 1, 0], [0, 1, 0, 0], [1, 1, 0, 0]]`, Output: `16`},
		{Input: `[[1, 0]]`, Output: `4`},
	})
}

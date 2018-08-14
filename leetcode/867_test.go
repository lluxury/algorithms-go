package leetcode

import (
	"github.com/Chyroc/algorithms-go/test"
	"testing"
)

func Test_867(t *testing.T) {
	test.Runs(t, transpose, []*test.Case{
		{Input: `[[1,2,3],[4,5,6],[7,8,9]]`, Output: `[[1,4,7],[2,5,8],[3,6,9]]`},
		{Input: `[[1,2,3],[4,5,6]]`, Output: `[[1,4],[2,5],[3,6]]`},
	})
}

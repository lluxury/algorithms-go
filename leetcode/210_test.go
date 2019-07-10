package leetcode

import (
	"github.com/lluxury/algorithms-go/test"
	"testing"
)

func Test_210(t *testing.T) {
	test.Runs(t, findOrder, []*test.Case{
		{Input: `1 \n []`, Output: `[0]`},
		{Input: `2 \n [[1,0]]`, Output: `[0, 1]`},
		{Input: `4 \n [[1,0],[2,0],[3,1],[3,2]]`, Output: `[0,1,2,3]`},
		{Input: `2 \n [[1,0],[0,1]]`, Output: `[]`},
		{Input: `3 \n [[1,0],[0,2],[2,1]]`, Output: `[]`},
	})
}

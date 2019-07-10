package leetcode

import (
	"github.com/lluxury/algorithms-go/test"
	"testing"
)

func Test_840(t *testing.T) {
	test.Runs(t, numMagicSquaresInside, []*test.Case{
		{Input: `[[4,3,8,4],[9,5,1,9],[2,7,6,2]]`, Output: `1`},
		{Input: `[[10,3,5],[1,6,11],[7,9,2]]`, Output: `0`},
		{Input: `[[7,0,5],[2,4,6],[3,8,1]]`, Output: `0`},
		{Input: `[[7,6,2,2,4],[4,4,9,2,10],[9,7,8,3,10],[8,1,9,7,5],[7,10,4,11,6]]`, Output: `0`},
	})
}

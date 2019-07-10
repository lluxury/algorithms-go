package leetcode

import (
	"testing"

	"github.com/lluxury/algorithms-go/test"
)

func Test_643(t *testing.T) {
	test.Runs(t, findMaxAverage, []*test.Case{
		{Input: `[1,12,-5,-6,50,3] \n 4`, Output: `12.75`},
		{Input: `[5] \n 1`, Output: `5`},
		{Input: `[4,0,4,3,3] \n 5`, Output: `2.8`},
	})
}

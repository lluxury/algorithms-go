package leetcode

import (
	"github.com/lluxury/algorithms-go/test"
	"testing"
)

func Test_207(t *testing.T) {
	test.Runs(t, canFinish, []*test.Case{
		{Input: `2 \n [[1,0]]`, Output: `true`},
		{Input: `2 \n [[1,0],[0,1]]`, Output: `false`},
		{Input: `3 \n [[1,0],[0,2],[2,1]]`, Output: `false`},
	})
}

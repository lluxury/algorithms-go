package leetcode

import (
	"testing"

	"github.com/lluxury/algorithms-go/test"
)

func Test_695(t *testing.T) {
	c := `
[[0,0,1,0,0,0,0,1,0,0,0,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,1,1,0,1,0,0,0,0,0,0,0,0],
 [0,1,0,0,1,1,0,0,1,0,1,0,0],
 [0,1,0,0,1,1,0,0,1,1,1,0,0],
 [0,0,0,0,0,0,0,0,0,0,1,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,0,0,0,0,0,0,1,1,0,0,0,0]]`

	test.Runs(t, maxAreaOfIsland_Recursive, []*test.Case{
		{Input: c, Output: `6`},
	})
	test.Runs(t, maxAreaOfIsland_Iterative, []*test.Case{
		{Input: c, Output: `6`},
	})
}

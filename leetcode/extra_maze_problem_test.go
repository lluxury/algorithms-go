package leetcode

import (
	"testing"

	"github.com/lluxury/algorithms-go/test"
)

func Test_extra_maze_problem(t *testing.T) {
	test.Runs(t, maze_problem, []*test.Case{
		{Input: `[[0,0,0,0,0],[0,1,0,1,0],[0,1,1,0,0],[0,1,1,0,1],[0,0,0,0,0]]`, Output: `[[0,0], [1,0], [2,0], [3,0], [4,0], [4,1], [4,2], [4,3], [4,4]]`},
	})
}

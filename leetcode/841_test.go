package leetcode

import (
	"testing"

	"github.com/Chyroc/algorithms-go/test"
)

func Test_841(t *testing.T) {
	test.Runs(t, canVisitAllRooms, []*test.Case{
		{Input: `[[1],[2],[3],[]]`, Output: `true`},
		{Input: `[[1,3],[3,0,1],[2],[0]]`, Output: `false`},
	})
}

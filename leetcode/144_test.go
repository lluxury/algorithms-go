package leetcode

import (
	"testing"

	"github.com/lluxury/algorithms-go/test"
)

func Test_144(t *testing.T) {
	t.Run("", func(t *testing.T) {
		test.Runs(t, preorderTraversal_Recursive_1, []*test.Case{
			{Input: `(1,,(2,3,))`, Output: `[1,2,3]`},
		})
	})
	t.Run("", func(t *testing.T) {
		test.Runs(t, preOrder_Iterative, []*test.Case{
			{Input: `(1,(2,3,4),(5,6,7))`, Output: `[1, 2, 3, 4, 5, 6, 7]`},
		})
		test.Runs(t, inOrder_Iterative, []*test.Case{
			{Input: `(1,(2,3,4),(5,6,7))`, Output: `[3, 2, 4, 1, 6, 5, 7]`},
		})
		test.Runs(t, postOrder_Iterative, []*test.Case{
			{Input: `(1,(2,3,4),(5,6,7))`, Output: `[3, 4, 2, 6, 7, 5, 1]`},
			{Input: `(1,,(2,3,))`, Output: `[3,2,1]`},
			{Input: ``, Output: ``},
		})
	})
}

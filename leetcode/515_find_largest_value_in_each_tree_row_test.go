package leetcode

import (
	"testing"

	"github.com/Chyroc/algorithms-go/lib"
	"github.com/Chyroc/algorithms-go/test"
)

/*
> https://leetcode.com/problems/find-largest-value-in-each-tree-row/description/


You need to find the largest value in each row of a binary tree.

Example:
Input:

          1
         / \
        3   2
       / \   \
      5   3   9

Output: [1, 3, 9]

* 求一个二叉树里面每一层的最大值，返回一个数组
  * 递归
  * 维护一个数组，将每一层的最大值加到数组中

*/

func largestValues(root *lib.TreeNode) []int {
	var values []int
	back_largestValues(root, &values, 0)
	return values
}

func back_largestValues(root *lib.TreeNode, values *[]int, dep int) {
	if root == nil {
		return
	}
	if len(*values) <= dep {
		*values = append(*values, root.Val)
	} else {
		if (*values)[dep] < root.Val {
			(*values)[dep] = root.Val
		}
	}
	back_largestValues(root.Left, values, dep+1)
	back_largestValues(root.Right, values, dep+1)
}

func Test_515(t *testing.T) {
	test.Runs(t, largestValues, []*test.Case{
		{Input: `(1,(3,5,3),(2,,9))`, Output: `[1,3,9]`},
	})
}

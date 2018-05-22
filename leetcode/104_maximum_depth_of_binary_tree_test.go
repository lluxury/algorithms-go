package leetcode

import (
	"github.com/Chyroc/algorithms-go/lib"
	"github.com/Chyroc/algorithms-go/test"
	"testing"
)

/*
> https://leetcode.com/problems/maximum-depth-of-binary-tree/description/


Given a binary tree, find its maximum depth.

The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

找出二叉树的最大深度

- 递归
- 树的问题可以分解为左右子树的问题

*/

func maxDepth(root *lib.TreeNode) int {
	if root == nil {
		return 0
	}
	leftDep := maxDepth(root.Left)
	rightDep := maxDepth(root.Right)

	if leftDep > rightDep {
		return 1 + leftDep
	}

	return 1 + rightDep
}

func Test_104(t *testing.T) {
	test.Runs(t, maxDepth, []*test.Case{
		{Input: `(1,(2,3,4),(2,4,3))`, Output: `3`},
		{Input: `(1,(2,,3),(2,(1,2,3),3))`, Output: `4`},
	})
}

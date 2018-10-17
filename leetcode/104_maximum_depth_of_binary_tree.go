package leetcode

import "github.com/Chyroc/algorithms-go/lib"

/*
 * [104] Maximum Depth of Binary Tree
 *
 * https://leetcode.com/problems/maximum-depth-of-binary-tree/description/
 *
 * algorithms
 * Easy (56.40%)
 * Total Accepted:    368.4K
 * Total Submissions: 652.3K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, find its maximum depth.
 *
 * The maximum depth is the number of nodes along the longest path from the
 * root node down to the farthest leaf node.
 *
 * Note: A leaf is a node with no children.
 *
 * Example:
 *
 * Given binary tree [3,9,20,null,null,15,7],
 *
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 * return its depth = 3.
 *
 */

/*

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

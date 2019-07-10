package leetcode

import (
	"github.com/lluxury/algorithms-go/lib"
)

/*
 * [111] Minimum Depth of Binary Tree
 *
 * https://leetcode.com/problems/minimum-depth-of-binary-tree/description/
 *
 * algorithms
 * Easy (33.84%)
 * Total Accepted:    227.9K
 * Total Submissions: 673.4K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, find its minimum depth.
 *
 * The minimum depth is the number of nodes along the shortest path from the
 * root node down to the nearest leaf node.
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
 * return its minimum depth = 2.
 *
 */

/*
 求根节点到叶子节点的记录的最小值
   * 节点为nil，返回0
   * 是叶子节点，返回1
   * 本身不为nil，左右子节点为nil的才是叶子节点
   * 如果本身是nil，说明根本不是合法节点，应该返回他的兄弟的递归的结果
*/
func isLeft(root *lib.TreeNode) bool {
	return root.Left == nil && root.Right == nil
}

func minDepth_rec(root *lib.TreeNode, level int) int {
	if root == nil {
		return 0 + level
	}
	if isLeft(root) {
		return 1 + level
	}

	left := root.Left
	right := root.Right

	if left == nil {
		return minDepth_rec(right, 1+level)
	}
	if right == nil {
		return minDepth_rec(left, 1+level)
	}

	if (left != nil && isLeft(left)) || (right != nil && isLeft(right)) {
		return 2 + level
	}

	a1 := minDepth_rec(left, level+1)
	a2 := minDepth_rec(right, level+1)

	if a1 < a2 {
		return a1
	}
	return a2
}

func minDepth(root *lib.TreeNode) int {
	if root == nil {
		return 0
	}
	return minDepth_rec(root, 0)
}

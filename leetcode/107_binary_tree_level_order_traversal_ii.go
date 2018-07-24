package leetcode

import "github.com/Chyroc/algorithms-go/lib"

/*
 * [107] Binary Tree Level Order Traversal II
 *
 * https://leetcode.com/problems/binary-tree-level-order-traversal-ii/description/
 *
 * algorithms
 * Easy (43.12%)
 * Total Accepted:    174.9K
 * Total Submissions: 405.6K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, return the bottom-up level order traversal of its
 * nodes' values. (ie, from left to right, level by level from leaf to root).
 *
 *
 * For example:
 * Given binary tree [3,9,20,null,null,15,7],
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 *
 *
 * return its bottom-up level order traversal as:
 *
 * [
 * ⁠ [15,7],
 * ⁠ [9,20],
 * ⁠ [3]
 * ]
 *
 *
 */

/*
 将二叉树的各级加入同一个数组，返回各级的数组
 * 简单，二叉树的遍历
*/

func levelOrderBottom_rec(root *lib.TreeNode, level int, data *[][]int) {
	if root == nil {
		return
	}

	if len(*data) <= level {
		(*data) = append((*data), []int{root.Val})
	} else {
		(*data)[level] = append((*data)[level], root.Val)
	}

	if root.Left != nil {
		levelOrderBottom_rec(root.Left, level+1, data)
	}
	if root.Right != nil {
		levelOrderBottom_rec(root.Right, level+1, data)
	}
}

func levelOrderBottom(root *lib.TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var data [][]int

	levelOrderBottom_rec(root, 0, &data)
	r := make([][]int, len(data))
	for k := len(data) - 1; k >= 0; k-- {
		r[len(data)-k-1] = data[k]
	}

	return r
}

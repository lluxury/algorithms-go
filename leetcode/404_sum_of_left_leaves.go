package leetcode

import (
	"github.com/Chyroc/algorithms-go/lib"
)

/*
 * [404] Sum of Left Leaves
 *
 * https://leetcode.com/problems/sum-of-left-leaves/description/
 *
 * algorithms
 * Easy (47.83%)
 * Total Accepted:    95.1K
 * Total Submissions: 198.8K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Find the sum of all left leaves in a given binary tree.
 *
 * Example:
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 * There are two left leaves in the binary tree, with values 9 and 15
 * respectively. Return 24.
 *
 *
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*

 求二叉树的所有左叶子节点的和

*/

func sumOfLeftLeaves_rec(root *lib.TreeNode, isLeft bool, data *int) {
	if root == nil {
		return
	}

	if isLeft && root.Left == nil && root.Right == nil {
		(*data) += root.Val
		return
	}

	sumOfLeftLeaves_rec(root.Left, true, data)
	sumOfLeftLeaves_rec(root.Right, false, data)
}

func sumOfLeftLeaves(root *lib.TreeNode) int {
	var data = 0
	sumOfLeftLeaves_rec(root, false, &data)
	return data
}

package leetcode

import "github.com/Chyroc/algorithms-go/lib"

/*
 * [563] Binary Tree Tilt
 *
 * https://leetcode.com/problems/binary-tree-tilt/description/
 *
 * algorithms
 * Easy (46.88%)
 * Total Accepted:    35.2K
 * Total Submissions: 75.2K
 * Testcase Example:  '[1,2,3]'
 *
 * Given a binary tree, return the tilt of the whole tree.
 * 
 * The tilt of a tree node is defined as the absolute difference between the
 * sum of all left subtree node values and the sum of all right subtree node
 * values. Null node has tilt 0.
 * 
 * The tilt of the whole tree is defined as the sum of all nodes' tilt.
 * 
 * Example:
 * 
 * Input: 
 * ⁠        1
 * ⁠      /   \
 * ⁠     2     3
 * Output: 1
 * Explanation: 
 * Tilt of node 2 : 0
 * Tilt of node 3 : 0
 * Tilt of node 1 : |2-3| = 1
 * Tilt of binary tree : 0 + 0 + 1 = 1
 * 
 * 
 * 
 * Note:
 * 
 * The sum of node values in any subtree won't exceed the range of 32-bit
 * integer. 
 * All the tilt values won't exceed the range of 32-bit integer.
 * 
 * 
 */

 /*

 定义diff：一个树的左子树的所有节点的和与右子树的所有节点的和的差
 定义树的diff：所有节点的diff的和

 递归：计算左子树和 和 右子树和，然后求diff，然后递归求左子树 和 右子树

 */

func sum_tree_rec(root *lib.TreeNode, sum *int) {
	if root == nil {
		return
	}
	*sum += root.Val
	sum_tree_rec(root.Left, sum)
	sum_tree_rec(root.Right, sum)
}

func sum_tree(root *lib.TreeNode) int {
	sum := 0
	sum_tree_rec(root, &sum)
	return sum
}

func findTilt_rec(root *lib.TreeNode, diff *int) {
	if root == nil {
		return
	}

	l := sum_tree(root.Left)
	r := sum_tree(root.Right)
	if l > r {
		*diff += l - r
	} else {
		*diff += r - l
	}

	findTilt_rec(root.Left, diff)
	findTilt_rec(root.Right, diff)
}

func findTilt(root *lib.TreeNode) int {
	diff := 0
	findTilt_rec(root, &diff)
	return diff
}

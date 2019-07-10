package leetcode

import "github.com/lluxury/algorithms-go/lib"

/*
 * [112] Path Sum
 *
 * https://leetcode.com/problems/path-sum/description/
 *
 * algorithms
 * Easy (35.58%)
 * Total Accepted:    237.7K
 * Total Submissions: 667.2K
 * Testcase Example:  '[5,4,8,11,null,13,4,7,2,null,null,null,1]\n22'
 *
 * Given a binary tree and a sum, determine if the tree has a root-to-leaf path
 * such that adding up all the values along the path equals the given sum.
 *
 * Note: A leaf is a node with no children.
 *
 * Example:
 *
 * Given the below binary tree and sum = 22,
 *
 *
 * ⁠     5
 * ⁠    / \
 * ⁠   4   8
 * ⁠  /   / \
 * ⁠ 11  13  4
 * ⁠/  \      \
 * 7    2      1
 *
 *
 * return true, as there exist a root-to-leaf path 5->4->11->2 which sum is 22.
 *
 */

/*

- 求二叉树从根节点到叶子节点，存在和为sum的路径
- 递归，将 root.Val - sum作为新的sum传递给子树
- 因为要求是到叶子节点而不是任意路径，所以需要判断当前节点是叶子节点的时候（左右子树为空），才可以判断当前sum是否为当前节点的值
- 中序遍历，先中，然后左右子树


类似： https://leetcode.com/problems/path-sum-ii/description/

*/

func hasPathSum(root *lib.TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return sum == root.Val
	}

	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}

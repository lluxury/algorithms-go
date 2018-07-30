package leetcode

import (
	"github.com/Chyroc/algorithms-go/lib"
	"math"
	"sort"
)

/*
 * [799] Minimum Distance Between BST Nodes
 *
 * https://leetcode.com/problems/minimum-distance-between-bst-nodes/description/
 *
 * algorithms
 * Easy (48.05%)
 * Total Accepted:    13.7K
 * Total Submissions: 28.6K
 * Testcase Example:  '[4,2,6,1,3,null,null]'
 *
 * Given a Binary Search Tree (BST) with the root node root, return the minimum
 * difference between the values of any two different nodes in the tree.
 *
 * Example :
 *
 *
 * Input: root = [4,2,6,1,3,null,null]
 * Output: 1
 * Explanation:
 * Note that root is a TreeNode object, not an array.
 *
 * The given tree [4,2,6,1,3,null,null] is represented by the following
 * diagram:
 *
 * ⁠         4
 * ⁠       /   \
 * ⁠     2      6
 * ⁠    / \
 * ⁠   1   3
 *
 * while the minimum difference in this tree is 1, it occurs between node 1 and
 * node 2, also between node 3 and node 2.
 *
 *
 * Note:
 *
 *
 * The size of the BST will be between 2 and 100.
 * The BST is always valid, each node's value is an integer, and each node's
 * value is different.
 *
 */

/*

 求一个BST搜索二叉树的任意两个节点的差值的最小值
   * 先获取所有节点
   * 排序
   * 获取相邻节点的差值
   * 获取差值最小值

*/

func diff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
func minDiffInBST_rec(root *lib.TreeNode, data *[]int) {
	if root == nil {
		return
	}

	(*data) = append(*data, root.Val)

	minDiffInBST_rec(root.Left, data)
	minDiffInBST_rec(root.Right, data)
}

func minDiffInBST(root *lib.TreeNode) int {
	if root == nil {
		return 0
	}

	var data []int
	minDiffInBST_rec(root, &data)

	sort.Ints(data)

	min := math.MaxInt32
	for i := 0; i < len(data)-1; i++ {
		x := diff(data[i], data[i+1])
		if x < min {
			min = x
		}
	}
	return min
}

package leetcode

import (
	"github.com/Chyroc/algorithms-go/lib"
)

/*
 * [904] Leaf-Similar Trees
 *
 * https://leetcode.com/problems/leaf-similar-trees/description/
 *
 * algorithms
 * Easy (65.04%)
 * Total Accepted:    9.6K
 * Total Submissions: 15.3K
 * Testcase Example:  '[3,5,1,6,2,9,8,null,null,7,4]\n[3,5,1,6,7,4,2,null,null,null,null,null,null,9,8]'
 *
 * Consider all the leaves of a binary tree.  From left to right order, the
 * values of those leaves form a leaf value sequence.
 *
 *
 *
 * For example, in the given tree above, the leaf value sequence is (6, 7, 4,
 * 9, 8).
 *
 * Two binary trees are considered leaf-similar if their leaf value sequence is
 * the same.
 *
 * Return true if and only if the two given trees with head nodes root1 and
 * root2 are leaf-similar.
 *
 *
 *
 * Note:
 *
 *
 * Both of the given trees will have between 1 and 100 nodes.
 *
 */

/*

 按照从左往右的顺序，列出所有的叶子节点，组成一个数组；如果两个树的这个数组是一样的，返回true

 遍历获取这个数组
 然后比较这个数组

*/

func getLeaf_rec(root *lib.TreeNode, l *[]int) {
	if root == nil {
		return
	}

	getLeaf_rec(root.Left, l)

	if root.Left == nil && root.Right == nil {
		*l = append(*l, root.Val)
	}

	getLeaf_rec(root.Right, l)
}

func getLeaf(root *lib.TreeNode) []int {
	l := []int{}
	getLeaf_rec(root, &l)
	return l
}

func leafSimilar(root1 *lib.TreeNode, root2 *lib.TreeNode) bool {
	l1 := getLeaf(root1)
	l2 := getLeaf(root2)

	if len(l1) != len(l2) {
		return false
	}

	for k := range l1 {
		if l1[k] != l2[k] {
			return false
		}
	}

	return true
}

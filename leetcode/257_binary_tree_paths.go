package leetcode

import (
	"github.com/Chyroc/algorithms-go/lib"
	"strconv"
)

/*
 * [257] Binary Tree Paths
 *
 * https://leetcode.com/problems/binary-tree-paths/description/
 *
 * algorithms
 * Easy (43.33%)
 * Total Accepted:    186.8K
 * Total Submissions: 430.9K
 * Testcase Example:  '[1,2,3,null,5]'
 *
 * Given a binary tree, return all root-to-leaf paths.
 * 
 * Note: A leaf is a node with no children.
 * 
 * Example:
 * 
 * 
 * Input:
 * 
 * ⁠  1
 * ⁠/   \
 * 2     3
 * ⁠\
 * ⁠ 5
 * 
 * Output: ["1->2->5", "1->3"]
 * 
 * Explanation: All root-to-leaf paths are: 1->2->5, 1->3
 * 
 */

func binaryTreePaths(root *lib.TreeNode) []string {
	var result []string
	binaryTreePaths_rec(root, "", &result)
	return result
}

func binaryTreePaths_rec(root *lib.TreeNode, s string, result *[]string) {
	if root == nil {
		return
	}

	if s == "" {
		s = s + strconv.Itoa(root.Val)
	} else {
		s = s + "->" + strconv.Itoa(root.Val)
	}

	if root.Left == nil && root.Right == nil {
		*result = append(*result, s)
		return
	}

	binaryTreePaths_rec(root.Left, s, result)
	binaryTreePaths_rec(root.Right, s, result)
}

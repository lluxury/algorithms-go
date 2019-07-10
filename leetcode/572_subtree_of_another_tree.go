package leetcode

import (
	"github.com/lluxury/algorithms-go/lib"
)

/*
 * [572] Subtree of Another Tree
 *
 * https://leetcode.com/problems/subtree-of-another-tree/description/
 *
 * algorithms
 * Easy (40.29%)
 * Total Accepted:    56.5K
 * Total Submissions: 140.2K
 * Testcase Example:  '[3,4,5,1,2]\n[4,1,2]'
 *
 *
 * Given two non-empty binary trees s and t, check whether tree t has exactly
 * the same structure and node values with a subtree of s. A subtree of s is a
 * tree consists of a node in s and all of this node's descendants. The tree s
 * could also be considered as a subtree of itself.
 *
 *
 * Example 1:
 *
 * Given tree s:
 *
 * ⁠    3
 * ⁠   / \
 * ⁠  4   5
 * ⁠ / \
 * ⁠1   2
 *
 * Given tree t:
 *
 * ⁠  4
 * ⁠ / \
 * ⁠1   2
 *
 * Return true, because t has the same structure and node values with a subtree
 * of s.
 *
 *
 * Example 2:
 *
 * Given tree s:
 *
 * ⁠    3
 * ⁠   / \
 * ⁠  4   5
 * ⁠ / \
 * ⁠1   2
 * ⁠   /
 * ⁠  0
 *
 * Given tree t:
 *
 * ⁠  4
 * ⁠ / \
 * ⁠1   2
 *
 * Return false.
 *
 */

func isSubtree_strict(s *lib.TreeNode, t *lib.TreeNode) bool {
	if s == nil {
		return t == nil
	}
	if t == nil {
		return s == nil
	}
	if s.Val != t.Val {
		return false
	}
	return isSubtree_strict(s.Left, t.Left) && isSubtree_strict(s.Right, t.Right)
}

func isSubtree(s *lib.TreeNode, t *lib.TreeNode) bool {
	if s == nil {
		return t == nil
	}
	if t == nil {
		return s == nil
	}

	if s.Val == t.Val {
		if isSubtree_strict(s.Left, t.Left) && isSubtree_strict(s.Right, t.Right) {
			return true
		}
	}

	return isSubtree(s.Left, t) || isSubtree(s.Right, t)

}

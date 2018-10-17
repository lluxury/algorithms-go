package leetcode

import "github.com/Chyroc/algorithms-go/lib"

/*
 * [515] Find Largest Value in Each Tree Row
 *
 * https://leetcode.com/problems/find-largest-value-in-each-tree-row/description/
 *
 * algorithms
 * Medium (56.04%)
 * Total Accepted:    46.1K
 * Total Submissions: 82.3K
 * Testcase Example:  '[1,3,2,5,3,null,9]'
 *
 * You need to find the largest value in each row of a binary tree.
 *
 * Example:
 *
 * Input:
 *
 * ⁠         1
 * ⁠        / \
 * ⁠       3   2
 * ⁠      / \   \
 * ⁠     5   3   9
 *
 * Output: [1, 3, 9]
 *
 *
 *
 */

/*

* 求一个二叉树里面每一层的最大值，返回一个数组
  * 递归
  * 维护一个数组，将每一层的最大值加到数组中

*/

func largestValues(root *lib.TreeNode) []int {
	var values []int
	back_largestValues(root, &values, 0)
	return values
}

func back_largestValues(root *lib.TreeNode, values *[]int, dep int) {
	if root == nil {
		return
	}
	if len(*values) <= dep {
		*values = append(*values, root.Val)
	} else {
		if (*values)[dep] < root.Val {
			(*values)[dep] = root.Val
		}
	}
	back_largestValues(root.Left, values, dep+1)
	back_largestValues(root.Right, values, dep+1)
}

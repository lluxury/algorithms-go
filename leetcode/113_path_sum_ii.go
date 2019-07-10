package leetcode

import "github.com/lluxury/algorithms-go/lib"

/*
 * [113] Path Sum II
 *
 * https://leetcode.com/problems/path-sum-ii/description/
 *
 * algorithms
 * Medium (36.96%)
 * Total Accepted:    179.9K
 * Total Submissions: 485.7K
 * Testcase Example:  '[5,4,8,11,null,13,4,7,2,null,null,5,1]\n22'
 *
 * Given a binary tree and a sum, find all root-to-leaf paths where each path's
 * sum equals the given sum.
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
 * ⁠/  \    / \
 * 7    2  5   1
 *
 *
 * Return:
 *
 *
 * [
 * ⁠  [5,4,11,2],
 * ⁠  [5,8,4,5]
 * ]
 *
 *
 */

/*

- 求二叉树从根节点到叶子节点，存在和为sum的路径，返回所有可能的数组
同112

类似： https://leetcode.com/problems/path-sum/description/

*/

func pathSum(root *lib.TreeNode, sum int) [][]int {
	var result [][]int
	back_pathSum(root, sum, []int{}, &result)
	return result
}

func CloneAppendSlice(a []int, i ...int) []int {
	var x []int
	for _, v := range a {
		x = append(x, v)
	}
	x = append(x, i...)
	return x
}

func back_pathSum(root *lib.TreeNode, sum int, tmp []int, result *[][]int) {
	if root == nil {
		return
	}

	newtmp := CloneAppendSlice(tmp, root.Val)

	if root.Left == nil && root.Right == nil {
		// 叶子节点
		if root.Val == sum {
			*result = append(*result, newtmp)
		}
		return
	}

	back_pathSum(root.Left, sum-root.Val, newtmp, result)
	back_pathSum(root.Right, sum-root.Val, newtmp, result)
}

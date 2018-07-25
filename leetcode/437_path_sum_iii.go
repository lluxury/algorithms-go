package leetcode

import (
	"github.com/Chyroc/algorithms-go/lib"
)

/*
 * [437] Path Sum III
 *
 * https://leetcode.com/problems/path-sum-iii/description/
 *
 * algorithms
 * Easy (40.28%)
 * Total Accepted:    64.9K
 * Total Submissions: 161.1K
 * Testcase Example:  '[10,5,-3,3,2,null,11,3,-2,null,1]\n8'
 *
 * You are given a binary tree in which each node contains an integer value.
 *
 * Find the number of paths that sum to a given value.
 *
 * The path does not need to start or end at the root or a leaf, but it must go
 * downwards
 * (traveling only from parent nodes to child nodes).
 *
 * The tree has no more than 1,000 nodes and the values are in the range
 * -1,000,000 to 1,000,000.
 *
 * Example:
 *
 * root = [10,5,-3,3,2,null,11,3,-2,null,1], sum = 8
 *
 * ⁠     10
 * ⁠    /  \
 * ⁠   5   -3
 * ⁠  / \    \
 * ⁠ 3   2   11
 * ⁠/ \   \
 * 3  -2   1
 *
 * Return 3. The paths that sum to 8 are:
 *
 * 1.  5 -> 3
 * 2.  5 -> 2 -> 1
 * 3. -3 -> 11
 *
 *
 */

/*
 给一个二叉树，从上往下找个路径，使得和为sum
   * 路径是从上往下的
   * 将所有的这种路径的和求出来
   * 求等于sum的个数
*/

func pathSum_rec(root *lib.TreeNode, stack []int, data *map[int]int) {
	if root == nil {
		return
	}

	var n = make([]int, len(stack))
	copy(n, stack)
	n = append(n, root.Val)

	var s int
	for k := len(n) - 1; k >= 0; k-- {
		s += n[k]
		(*data)[s]++
	}

	pathSum_rec(root.Left, n, data)
	pathSum_rec(root.Right, n, data)
}

func pathSum_437(root *lib.TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	var stack []int              // 存储根节点到当前节点的数据
	var data = make(map[int]int) // 存储每个节点到叶子节点的数据之和，是个数组

	pathSum_rec(root, stack, &data)

	return data[sum]
}

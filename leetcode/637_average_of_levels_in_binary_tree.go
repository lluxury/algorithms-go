package leetcode

import "github.com/Chyroc/algorithms-go/lib"

/*
 * [637] Average of Levels in Binary Tree
 *
 * https://leetcode.com/problems/average-of-levels-in-binary-tree/description/
 *
 * algorithms
 * Easy (56.45%)
 * Total Accepted:    54.3K
 * Total Submissions: 96.2K
 * Testcase Example:  '[3,9,20,15,7]'
 *
 * Given a non-empty binary tree, return the average value of the nodes on each
 * level in the form of an array.
 *
 * Example 1:
 *
 * Input:
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 * Output: [3, 14.5, 11]
 * Explanation:
 * The average value of nodes on level 0 is 3,  on level 1 is 14.5, and on
 * level 2 is 11. Hence return [3, 14.5, 11].
 *
 *
 *
 * Note:
 *
 * The range of node's value is in the range of 32-bit signed integer.
 *
 *
 */

/*

* 求一个二叉树的每一级的平均数
  * 递归
  * 用一个[][]int表示每一层的元素的数组，最后求每一层的平均数
  * 用dep表示深度，也是上面那个数组的下标
*/

func pinjunshu(l ...int) float64 {
	var s int
	for _, v := range l {
		s += v
	}

	return float64(s) / float64(len(l))
}

func averageOfLevels(root *lib.TreeNode) []float64 {
	var sum1 = make([][]int, 0)
	back_637(root, 0, &sum1)

	var sum2 []float64
	for _, v := range sum1 {
		sum2 = append(sum2, pinjunshu(v...))
	}
	return sum2
}

func back_637(root *lib.TreeNode, dep int, sum *[][]int) {
	if root == nil {
		return
	}

	if len(*sum) > dep {
		(*sum)[dep] = append((*sum)[dep], root.Val)
	} else {
		(*sum) = append((*sum), []int{root.Val})
	}

	back_637(root.Left, dep+1, sum)
	back_637(root.Right, dep+1, sum)
}

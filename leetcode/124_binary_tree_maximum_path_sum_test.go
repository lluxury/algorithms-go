package leetcode

import (
	"github.com/Chyroc/algorithms-go/lib"
	"github.com/Chyroc/algorithms-go/test"
	"math"
	"testing"
)

/*

> https://leetcode.com/problems/binary-tree-maximum-path-sum/description/

Given a non-empty binary tree, find the maximum path sum.

For this problem, a path is defined as any sequence of nodes from some starting node to any node in the tree along the parent-child connections. The path must contain at least one node and does not need to go through the root.

Example 1:

Input: [1,2,3]

       1
      / \
     2   3

Output: 6
Example 2:

Input: [-10,9,20,null,null,15,7]

   -10
   / \
  9  20
    /  \
   15   7

Output: 42


* 在一个二叉树中，求一个路径，这个路径上的数字的和是所有可能的结果的最大值；不需要经过根节点

思路
* 从下到上进行计算
* 计算的过程中，.Val会存储那个节点为输入参数的所有path的最大值: root.Val += max(l, r)
* 计算的过程中，最大值会存储到max_: if l + r + root.Val > *max_ { *max_ = l + r + root.Val }

类似：https://leetcode.com/problems/longest-univalue-path/description/

*/

func maxPathSum(root *lib.TreeNode) int {
	var max = math.MinInt32
	back_maxPathSum(root, &max)
	return max
}

func back_maxPathSum(root *lib.TreeNode, max_ *int) int {
	if root == nil {
		return 0
	}
	l := back_maxPathSum(root.Left, max_)
	r := back_maxPathSum(root.Right, max_)

	if l < 0 {
		l = 0
	}
	if r < 0 {
		r = 0
	}

	tmp := l + r + root.Val
	if tmp > *max_ {
		*max_ = tmp
	}
	root.Val += max_all(l, r)
	return root.Val
}

func Test_124(t *testing.T) {
	test.Runs(t, maxPathSum, []*test.Case{
		{Input: `(-3,,)`, Output: `-3`},
		{Input: `(1,2,3)`, Output: `6`},
		{Input: `(-10,9,(20,15,7))`, Output: `42`},
	})
}

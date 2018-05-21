package leetcode

import (
	"github.com/Chyroc/algorithms-go/lib"
	"testing"
	"github.com/stretchr/testify/assert"
)

/*

> https://leetcode.com/problems/two-sum-iv-input-is-a-bst/description/


Given a Binary Search Tree and a target number, return true if there exist two elements in the BST such that their sum is equal to the given target.

Example 1:
Input:
    5
   / \
  3   6
 / \   \
2   4   7

Target = 9

Output: True
Example 2:
Input:
    5
   / \
  3   6
 / \   \
2   4   7

Target = 28

Output: False

* 判断一个二叉树的的两个数，加起来能不能等于k
  * 这一题和第一题是一样的：给定一个数组和一个数，从数组里面找出两个数，他们的和是这个数，返回这两个数字的index（答案唯一，同一个数只能使用一次）
  * 遍历二叉树
    * 用m保存已经访问过的数据
    * 判断已经访问过的数据m里面有没有 k-root.Val ，有的话，返回true
    * 没有的话，设置当前值为已经访问过，递归
*/

func findTarget(root *lib.TreeNode, k int) bool {
	var m = make(map[int]bool)
	return back_653(root, k, &m)
}

func back_653(root *lib.TreeNode, k int, m *map[int]bool) bool {
	if root == nil {
		return false
	}
	if (*m)[k-root.Val] {
		return true
	}
	(*m)[root.Val] = true

	return back_653(root.Left, k, m) || back_653(root.Right, k, m)
}

func Test_653(t *testing.T) {
	as := assert.New(t)

	{
		x := findTarget(&lib.TreeNode{
			Val: 2,
			Left: &lib.TreeNode{
				Val: 1,
			},
			Right: &lib.TreeNode{
				Val: 3,
			},
		}, 4)

		as.True(x)
	}

	t.SkipNow()

	{
		x := findTarget(&lib.TreeNode{
			Val: 5,
			Left: &lib.TreeNode{
				Val:   3,
				Left:  &lib.TreeNode{Val: 2},
				Right: &lib.TreeNode{Val: 4},
			},
			Right: &lib.TreeNode{
				Val:   6,
				Right: &lib.TreeNode{Val: 7},
			},
		}, 9)

		as.True(x)
	}

}

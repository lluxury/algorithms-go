package leetcode

import (
	"testing"

	"github.com/Chyroc/algorithms-go/lib"
	"github.com/Chyroc/algorithms-go/test"
)

/*

> https://leetcode.com/problems/sum-root-to-leaf-numbers/description/

Given a binary tree containing digits from 0-9 only, each root-to-leaf path could represent a number.

An example is the root-to-leaf path 1->2->3 which represents the number 123.

Find the total sum of all root-to-leaf numbers.

Note: A leaf is a node with no children.

Example:

Input: [1,2,3]
    1
   / \
  2   3
Output: 25
Explanation:
The root-to-leaf path 1->2 represents the number 12.
The root-to-leaf path 1->3 represents the number 13.
Therefore, sum = 12 + 13 = 25.
Example 2:

Input: [4,9,0,5,1]
    4
   / \
  9   0
 / \
5   1
Output: 1026
Explanation:
The root-to-leaf path 4->9->5 represents the number 495.
The root-to-leaf path 4->9->1 represents the number 491.
The root-to-leaf path 4->0 represents the number 40.
Therefore, sum = 495 + 491 + 40 = 1026.

* 在一个二叉树中，只有0到9，然后从根节点到叶子节点直接的path认为是一个数字，求所有可能的path的值的和
  * 遍历二叉树
  * 用一个值记录从上到下形成的数
  * 用一个指针形式的sum存储和

*/

func sumNumbers(root *lib.TreeNode) int {
	if root == nil {
		return 0
	}
	var sum int
	back_sumNumbers(root, 0, &sum)
	return sum
}

func back_sumNumbers(root *lib.TreeNode, tmp int, sum *int) {
	tmp = tmp*10 + root.Val
	if root.Left == nil && root.Right == nil {
		*sum += tmp
		return
	}

	if root.Left != nil {
		back_sumNumbers(root.Left, tmp, sum)
	}
	if root.Right != nil {
		back_sumNumbers(root.Right, tmp, sum)
	}
}

func Test_129(t *testing.T) {
	test.Runs(t, sumNumbers, []*test.Case{
		{Input: `(1,2,3)`, Output: `25`},
		{Input: `(4,(9,5,1),0)`, Output: `1026`},
	})
}

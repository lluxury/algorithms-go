package leetcode

import "github.com/lluxury/algorithms-go/lib"

/*
> https://leetcode.com/problems/invert-binary-tree/description/

Invert a binary tree.
```
     4
   /   \
  2     7
 / \   / \
1   3 6   9
```
to
```
     4
   /   \
  7     2
 / \   / \
9   6 3   1
```
Trivia:
This problem was inspired by [this original tweet](https://twitter.com/mxcl/status/608682016205344768) by [Max Howell](https://twitter.com/mxcl):
> Google: 90% of our engineers use the software you wrote (Homebrew), but you can’t invert a binary tree on a whiteboard so fuck off.

- 反转二叉树
- 将右边反转给左边就行

*/

func invertTree(root *lib.TreeNode) *lib.TreeNode {
	if root == nil {
		return nil
	}

	left := root.Left
	root.Left = invertTree(root.Right)
	root.Right = invertTree(left)

	return root
}

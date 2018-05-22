package leetcode

import (
	"github.com/Chyroc/algorithms-go/lib"
	"github.com/Chyroc/algorithms-go/test"
	"testing"
)

/*
> https://leetcode.com/problems/symmetric-tree/description/


Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).

For example, this binary tree `[1,2,2,3,4,4,3]` is symmetric:

检测一个二叉树是否左右对称

```
    1
   / \
  2   2
 / \ / \
3  4 4  3
```
But the following [1,2,2,null,3,null,3] is not:
```
    1
   / \
  2   2
   \   \
   3    3
```


Note:
Bonus points if you could solve it both recursively and iteratively.

- 和100题配合，先把树反转，然后判断是不是和原来的树相等

*/

func reverse_101(root *lib.TreeNode) *lib.TreeNode {
	if root == nil {
		return root
	}

	var root2 = &lib.TreeNode{Val: root.Val}

	root2.Right = reverse_101(root.Left)
	root2.Left = reverse_101(root.Right)

	return root2
}

func isSymmetric(root *lib.TreeNode) bool {
	return isSameTree(root, reverse_101(root))
}

func Test_101(t *testing.T) {
	test.Runs(t, isSymmetric, []*test.Case{
		{Input: `(1,(2,3,4),(2,4,3))`, Output: `true`},
		{Input: `(1,(2,,3),(2,,3))`, Output: `false`},
	})
}

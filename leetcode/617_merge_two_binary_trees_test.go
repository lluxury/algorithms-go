package leetcode

import "github.com/lluxury/algorithms-go/lib"

/*
> https://leetcode.com/problems/merge-two-binary-trees/description/

Given two binary trees and imagine that when you put one of them to cover the other, some nodes of the two trees are overlapped while the others are not.

You need to merge them into a new binary tree. The merge rule is that if two nodes overlap, then sum node values up as the new value of the merged node. Otherwise, the NOT null node will be used as the node of new tree.

Example 1:
```
Input:
	Tree 1                     Tree 2
          1                         2
         / \                       / \
        3   2                     1   3
       /                           \   \
      5                             4   7
Output:
Merged tree:
	     3
	    / \
	   4   5
	  / \   \
	 5   4   7
```

Note: The merging process must start from the root nodes of both trees.

- 先处理空树，然后处理根节点，然后处理左右子树就行


*/

func mergeTrees(t1 *lib.TreeNode, t2 *lib.TreeNode) *lib.TreeNode {
	if t2 == nil {
		return t1
	}
	if t1 == nil {
		return t2
	}

	t1.Val = t1.Val + t2.Val
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)

	return t1
}

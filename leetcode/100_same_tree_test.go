package leetcode

import (
	"github.com/Chyroc/algorithms-go/lib"
	"github.com/Chyroc/algorithms-go/test"
	"testing"
)

/*
> https://leetcode.com/problems/same-tree/description/

Given two binary trees, write a function to check if they are equal or not.

Two binary trees are considered equal if they are structurally identical and the nodes have the same value.

比较两个二叉树

- 递归

*/

func isSameTree(p *lib.TreeNode, q *lib.TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if q == nil || p == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func Test_100(t *testing.T) {
	test.Runs(t, isSameTree, []*test.Case{
		{Input: `() \n ()`, Output: `true`},
		{Input: `() \n (1)`, Output: `false`},
		{Input: `(1,2,3) \n (1,2,3)`, Output: `true`},
		{Input: `(1,2,(2)) \n (1,2,(2))`, Output: `true`},
		{Input: `(1,2,(2,,2)) \n (1,2,(2,,2))`, Output: `true`},
		{Input: `(1,(1,,1),(22,22,2)) \n (1,(1,,1),(22,22,2))`, Output: `true`},
	})
}

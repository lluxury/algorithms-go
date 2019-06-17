package leetcode

import (
   "github.com/Chyroc/algorithms-go/lib"
)

/*
Given a binary search tree (BST), find the lowest common ancestor (LCA) of two given nodes in the BST.

According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”

Given binary search tree:  root = [6,2,8,0,4,7,9,null,null,3,5]


 

Example 1:

Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
Output: 6
Explanation: The LCA of nodes 2 and 8 is 6.
Example 2:

Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
Output: 2
Explanation: The LCA of nodes 2 and 4 is 2, since a node can be a descendant of itself according to the LCA definition.
 

Note:

All of the nodes' values will be unique.
p and q are different and both values will exist in the BST.*/


// 找最小公共祖先,需要考虑几种可能性,然后重新传参,递归

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
// func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode  {
func lowestCommonAncestor(root, p, q *lib.TreeNode) *lib.TreeNode  {
    if p.Val > q.Val {
        p,q = q,p
    }

    if root == nil || p.Val<= root.Val && root.Val <= q.Val {
        return root
    } else if q.Val <= root.Val{
        return lowestCommonAncestor(root.Left, p, q)
    } else {
        return lowestCommonAncestor(root.Right, p, q)
    }
}

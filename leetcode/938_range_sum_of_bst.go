package leetcode

import "github.com/lluxury/algorithms-go/lib"
/*
 * @lc app=leetcode id=938 lang=golang
 *
 * [938] Range Sum of BST
 *
 * https://leetcode.com/problems/range-sum-of-bst/description/
 *
 * algorithms
 * Easy (78.11%)
 * Total Accepted:    92.6K
 * Total Submissions: 118.5K
 * Testcase Example:  '[10,5,15,3,7,null,18]\n7\n15'
 *
 * Given the root node of a binary search tree, return the sum of values of all
 * nodes with value between L and R (inclusive).
 * 
 * The binary search tree is guaranteed to have unique values.
 * 
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: root = [10,5,15,3,7,null,18], L = 7, R = 15
 * Output: 32
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: root = [10,5,15,3,7,13,18,1,null,6], L = 6, R = 10
 * Output: 23
 * 
 * 
 * 
 * 
 * Note:
 * 
 * 
 * The number of nodes in the tree is at most 10000.
 * The final answer is guaranteed to be less than 2^31.
 * 
 * 
 * 
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *lib.TreeNode, L int, R int) int {
// func rangeSumBST(root *TreeNOde,L R int) int{
    if root == nil {
        return 0
    }

    sum := 0

    switch  {
    case root.Val < L:
        sum = rangeSumBST(root.Right, L, R)
    case R < root.Val:
        sum = rangeSumBST(root.Left, L, R)
    default:
        sum = root.Val
        sum += rangeSumBST(root.Left, L, R)
        sum += rangeSumBST(root.Right,L, R)
    }
    return sum
}

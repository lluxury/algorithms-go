package leetcode

import "github.com/Chyroc/algorithms-go/lib"

/*
> https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/description/

Given an array where elements are sorted in ascending order, convert it to a height balanced BST.

- bst 二叉搜索树
- balanced bst 二叉平衡搜索树
- 所以已经排序的数组中位数就是root节点

*/

func tobst(nums []int, left, right int) *lib.TreeNode {
	if left > right {
		return nil
	}

	mid := (left + right) / 2
	root := &lib.TreeNode{Val: nums[mid]}
	root.Left = tobst(nums, left, mid-1)
	root.Right = tobst(nums, mid+1, right)

	return root
}

func sortedArrayToBST(nums []int) *lib.TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return tobst(nums, 0, len(nums)-1)
}

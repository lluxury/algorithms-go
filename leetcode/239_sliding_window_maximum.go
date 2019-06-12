package leetcode

import (
    "fmt"
    "math"
    "strconv"
    "strings"
    "unicode"
)
// leetcod时要去掉
/*
Given an array nums, there is a sliding window of size k which is moving from the very left of the array to the very right. You can only see the k numbers in the window. Each time the sliding window moves right by one position. Return the max sliding window.

Example:

Input: nums = [1,3,-1,-3,5,3,6,7], and k = 3
Output: [3,3,5,5,6,7] 
Explanation: 

Window position                Max
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
Note:
You may assume k is always valid, 1 ≤ k ≤ input array's size for non-empty array.*/

// 滑动窗口,持续给出窗口中最大的数,反回一个切片

// 这个解法是用树做出来的, 
// 我一直想用切片和数组的关系来做,不知道能不能做出来

type TreeNode struct {
        Val int
        Left *TreeNode
        // right *TreeNode
        Right *TreeNode
    }

    type ListNode struct {
        Val int
        Next *ListNode
    }

    type Point struct {
        X int
        Y int
    }

    func useLib()  {
        fmt.Println(strconv.Itoa(1))
        fmt.Println(strings.Compare("1", "2"))
        // fmt.Println(match.Abs(1.0))
        fmt.Println(math.Abs(1.0))
        fmt.Println(unicode.IsDigit('1'))
    }

    func buildTree(nums []int) *TreeNode  {
        if len(nums) == 0 {
            return nil
        }
        root := new(TreeNode)
        root.Val = nums[0]
        ch := make(chan *TreeNode, len(nums))
        ch <- root
        nums = nums[1:]
        for i := 0; i < len(nums); i++ {
            tree := <-ch
            if nums[i] == -1{
                tree.Left = nil
            } else {
                // tree.Left = &TreeNode{Val:nums[i]}
                // tree.Left = &TreeNode{Val:nums[i]
                tree.Left = &TreeNode{Val:nums[i],
                            }
                            ch <- tree.Left
            }
        i++
        if i == len(nums) || nums[i] == -1{
            tree.Right = nil
        } else {
            tree.Right = &TreeNode{
                Val: nums[i],
            }
            ch <- tree.Right
        }

    }
    return root
}

func buildList(nums []int) *ListNode  {
    if len(nums) == 0 {
        return nil
    }
    root := &ListNode{
        // Val := nums[0],
        Val : nums[0],
    }
    tmp := root
    // for i := 0; i < len(nums); i++ {
    for i := 1; i < len(nums); i++ {
        tmp.Next = &ListNode{
                Val: nums[i],
        }
        tmp = tmp.Next
    }
    return root
}

// func min(a, b int) int  {
//     if a > b {
//         return b
//     }
//     return a
// }

// func max(a, b int) int {
//     if a < b {
//         return b
//     } 
//     return a
// }
// 貌似与0号代码重复了,在leetcode里要加上

// func maxSlidingWindow(nmus []int, k int) []int  {
func maxSlidingWindow(nums []int, k int) []int  {
    if len(nums) == 0 || k == 0{
        return nums
}
// left,ritht := make([]int,len(nums)), make(make([]int, len(nums)))
left,right := make([]int,len(nums)),make([]int,len(nums))    
left[0],right[len(nums)-1] = nums[0], nums[len(nums)-1]
for i,j:= 1,len(nums)-2; i<len(nums);i,j = i+1,j-1{
    if i %k == 0{
        // left = nums[i]
        left[i] = nums[i]
    } else{
        left[i] = max(left[i-1],nums[i])
    }
    if j%k == 0 {
        // right[i] = nums[j]
        right[j] = nums[j]
    } else {
        right[j] = max(right[j+1],nums[j])
    }
}
ret := make([]int, len(nums)-k+1)
// for i:=0; i < i+k<=len(nums); i++ {
for i:=0;i+k<=len(nums);i++{
    ret[i] = max(right[i],left[i+k-1])
}
return ret
}

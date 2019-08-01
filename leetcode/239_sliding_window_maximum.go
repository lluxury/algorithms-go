package leetcode

import (
    // "fmt"
    // "math"
    // "strconv"
    // "strings"
    // "unicode"
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

type Queue struct {
    stored []int
    maxInt  int
}

func NewQueue() *Queue {
    return &Queue {
        stored: []int{},
        maxInt: math.MinInt32,
    }
}

func (q *Queue) Enqueue(nums int) {
    if nums > q.maxInt{
        q.maxInt = nums
    }
    q.stored = append(q.stored, nums)
}

func (q *Queue) Dequeue() {
    if len(q.stored) == 0 {
    return
}
    if q.maxInt == q.stored[0]{
        q.maxInt = math.MinInt32
        for _, cur := range q.stored[1:]{
            if cur > q.maxInt{
                q.maxInt = cur
            }
        }
    }
    q.stored = q.stored[1:]
}

func maxSlidingWindow(nums []int, k int) []int {
    res := []int{}
    queue := NewQueue()

    for i:=0;i<len(nums);i++{
        if i<k-1{
            queue.Enqueue(nums[i])
            continue
        }

        if i>k-1{
            queue.Dequeue()
        }
        queue.Enqueue(nums[i])
        res = append(res, queue.maxInt)
    }
    return res
}


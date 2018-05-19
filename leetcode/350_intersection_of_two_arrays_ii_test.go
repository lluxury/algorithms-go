package leetcode

/*
> https://leetcode.com/problems/intersection-of-two-arrays-ii/description/

Given two arrays, write a function to compute their intersection.

Example:

Given nums1 = `[1, 2, 2, 1]`, nums2 = `[2, 2]`, return `[2, 2]`.

Note:
- Each element in the result should appear as many times as it shows in both arrays.
- The result can be in any order.

Follow up:
- What if the given array is already sorted? How would you optimize your algorithm?
- What if nums1's size is small compared to nums2's size? Which algorithm is better?
- What if elements of nums2 are stored on disk, and the memory is limited such that you cannot load all elements into the memory at once?

- 题目
  - 取两个数据的公共元素
  - 每个数据可能有好几个，不要求顺序
- 思路
  - 哈希表

*/

import (
	"testing"

	"github.com/Chyroc/algorithms-go/test"
)

func intersect(nums1 []int, nums2 []int) []int {
	var m = make(map[int]int)
	var d = []int{}

	for _, v := range nums1 {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}

	for _, v := range nums2 {
		if data, ok := m[v]; ok {
			if data > 0 {
				d = append(d, v)
			}
			m[v]--
		}
	}

	return d
}

func Test_350(t *testing.T) {
	test.Runs(t, intersect, []*test.Case{
		{Input: `[1, 2, 2, 1] \n [2, 2]`, Output: "[2, 2]"},
	})
}

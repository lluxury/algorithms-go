package leetcode

import (
	"github.com/lluxury/algorithms-go/lib"
)

/*
 * [628] Maximum Product of Three Numbers
 *
 * https://leetcode.com/problems/maximum-product-of-three-numbers/description/
 *
 * algorithms
 * Easy (44.76%)
 * Total Accepted:    43.6K
 * Total Submissions: 97.4K
 * Testcase Example:  '[1,2,3]'
 *
 * Given an integer array, find three numbers whose product is maximum and
 * output the maximum product.
 *
 * Example 1:
 *
 * Input: [1,2,3]
 * Output: 6
 *
 *
 *
 * Example 2:
 *
 * Input: [1,2,3,4]
 * Output: 24
 *
 *
 *
 * Note:
 *
 * The length of the given array will be in range [3,104] and all elements are
 * in the range [-1000, 1000].
 * Multiplication of any three numbers in the input won't exceed the range of
 * 32-bit signed integer.
 *
 *
 */

/*

 给定一个数组，里面有整数负数，现在求里面的三个数，要求他们的乘积是最大的

 如果有两个负数，那么结果仍然是整数

 所以用两个堆
 一个堆存正数，长度3
 一个堆存负数，长度2

 如果最后负数的堆长度等于2，那么就取出来，去掉负号，再加入正数堆

 最后正数堆就是三个最大的数，求乘积即可

*/

func maximumProduct(nums []int) int {
	maxheap := lib.NewHeap(func(a, b interface{}) bool { return a.(int) < b.(int) })
	minheap := lib.NewHeap(func(a, b interface{}) bool { return a.(int) > b.(int) })
	for _, v := range nums {
		if v > 0 {
			maxheap.Add(v)
			if maxheap.Len() > 3 {
				maxheap.Pop()
			}
		} else if v < 0 {
			minheap.Add(v)
			if minheap.Len() > 2 {
				minheap.Pop()
			}
		}
	}

	if minheap.Len() >= 2 {
		for i := 0; i < 2; i++ {
			v, _ := minheap.Pop()
			maxheap.Add(v)
			if maxheap.Len() > 3 {
				maxheap.Pop()
			}
		}
	}

	sum := 1
	for i := 0; i < 3; i++ {
		if maxheap.Len() == 0 {
			break
		}
		v, _ := maxheap.Pop()
		sum *= v.(int)
	}

	return sum
}

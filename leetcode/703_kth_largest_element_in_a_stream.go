package leetcode

import "github.com/lluxury/algorithms-go/lib"

/*
 * [789] Kth Largest Element in a Stream
 *
 * https://leetcode.com/problems/kth-largest-element-in-a-stream/description/
 *
 * algorithms
 * Easy (20.36%)
 * Total Accepted:    874
 * Total Submissions: 4.3K
 * Testcase Example:  '["KthLargest","add","add","add","add","add"]\n[[3,[4,5,8,2]],[3],[5],[10],[9],[4]]'
 *
 * Design a class to find the kth largest element in a stream. Note that it is
 * the kth largest element in the sorted order, not the kth distinct element.
 *
 * Your KthLargest class will have a constructor which accepts an integer k and
 * an integer array nums, which contains initial elements from the stream. For
 * each call to the method KthLargest.add, return the element representing the
 * kth largest element in the stream.
 *
 * Example:
 *
 *
 * int k = 3;
 * int[] arr = [4,5,8,2];
 * KthLargest kthLargest = new KthLargest(3, arr);
 * kthLargest.add(3);   // returns 4
 * kthLargest.add(5);   // returns 5
 * kthLargest.add(10);  // returns 5
 * kthLargest.add(9);   // returns 8
 * kthLargest.add(4);   // returns 8
 *
 *
 * Note:
 * You may assume that nums' length ≥ k-1 and k ≥ 1.
 *
 */

/*

 可以不停的添加数字，求所有的数字中第k大的数字
 * 用最小堆实现
 * 堆的长度限制为k，所以第k大的数字就是最小的数字

*/

type KthLargest struct {
	k int
	h *lib.Heap
}

func Constructor_703(k int, nums []int) KthLargest {
	var less = func(a, b interface{}) bool {
		return a.(int) < b.(int)
	}
	var input []interface{}
	for _, v := range nums {
		input = append(input, v)
	}

	c := KthLargest{
		k: k,
		h: lib.NewHeap(less, input...),
	}
	for c.h.Len() > k {
		c.h.Pop()
	}

	return c
}

func (this *KthLargest) Add(val int) int {
	this.h.Add(val)
	if this.h.Len() > this.k {
		this.h.Pop()
	}
	return this.h.Peek().(int)
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

// integer min heap
// type IntHeap []int

// func (h IntHeap) Len() int  					{ return len(h)}
// func (h IntHeap) Less (i,j int) bool 	{ return h[i] < h[j]}
// func (h IntHeap) Swap (i,j int) 			{h[i],h[j] = h[j],h[i]}

// func (h *IntHeap) Push (x interface{})  {
// 	*h = append(*h, x.(int))
// }

// func (h *IntHeap) Pop() interface{}  {
// 	old := *h
// 	n := len(old)
// 	x:= old[0]-1
// 	*h = old[0 : n-1]
// 	return x
// }

// type KthLargest struct {
// 	Nums *IntHeap
// 	K 	int
// }

// // func Constructor(k int, nums []int) KthLargest {
// func Constructor_703(k int, nums []int) KthLargest {
// 	h := &IntHeap{}
// 	heap.Init(h)
// 	for i := 0; i < len(nums); i++ {
// 		heap.Push(h, nums[i])
// 	}

// 	for len(*h) > k{
// 		heap.Pop(h)
// 	}
// 	return KthLargest{h, k}
// }

// func (this *KthLargest) Add(val int) int  {
// 	if len(*this.Nums) < this.K {
// 		heap.Push(this.Nums, val)
// 	} else if val >(*this.Nums)[0]{
// 		heap.Push(this.Nums,val)
// 		heap.Pop(this.Nums)
// 	}
// 	return (*this.Nums)[0]
// }

// .\703_kth_largest_element_in_a_stream.go:117:3: undefined: heap
// .\703_kth_largest_element_in_a_stream.go:119:4: undefined: heap
// 问题来了, leetcode里的 heap是哪来的呢?

package leetcode

import (
	"github.com/lluxury/algorithms-go/lib"
	"math"
)

/*
 * [802] K-th Smallest Prime Fraction
 *
 * https://leetcode.com/problems/k-th-smallest-prime-fraction/description/
 *
 * algorithms
 * Hard (33.83%)
 * Total Accepted:    3.7K
 * Total Submissions: 10.9K
 * Testcase Example:  '[1,2,3,5]\n3'
 *
 * A sorted list A contains 1, plus some number of primes.  Then, for every p <
 * q in the list, we consider the fraction p/q.
 *
 * What is the K-th smallest fraction considered?  Return your answer as an
 * array of ints, where answer[0] = p and answer[1] = q.
 *
 *
 * Examples:
 * Input: A = [1, 2, 3, 5], K = 3
 * Output: [2, 5]
 * Explanation:
 * The fractions to be considered in sorted order are:
 * 1/5, 1/3, 2/5, 1/2, 3/5, 2/3.
 * The third fraction is 2/5.
 *
 * Input: A = [1, 7], K = 1
 * Output: [1, 7]
 *
 *
 * Note:
 *
 *
 * A will have length between 2 and 2000.
 * Each A[i] will be between 1 and 30000.
 * K will be between 1 and A.length * (A.length - 1) / 2.
 *
 */

/*


总是超时。。。。


*/

func kthSmallestPrimeFraction_1(A []int, K int) []int {
	var xuhao = make([]int, len(A)) //记录着A中各个位置作为分子的时候，分母所在的index
	for k := range A {
		xuhao[k] = len(A) - 1
	}

	var min_index []int
	for i := 0; i < K; i++ {
		var min = math.MaxFloat64 / 2

		for i := 0; i < len(A); i++ {
			// i          // 分子的index
			j := xuhao[i] //分母的index

			// 因为A是递增的，所以需要i<j
			if i < j {
				d := float64(A[i]) / float64(A[j]) // 分子/分母
				if d <= min {
					min = d
					min_index = []int{i, j}
				}
			}
		}

		// 取最小值的index，让其在xuhao中加1
		fenzi_index := min_index[0]
		xuhao[fenzi_index]--
	}

	return []int{A[min_index[0]], A[min_index[1]]}
}

// 使用最小堆
// 1/47   < 1/29    < 1/23 < 1/7
// 7/47   < 7/29    < 7/23
// 23/47  < 23/29
// 29/47
// 将每个队列的最小值加入最小堆
func kthSmallestPrimeFraction(A []int, K int) []int {
	lastindex := len(A) - 1

	var intpairless = func(a, b interface{}) bool {
		x := a.([]int)
		y := b.([]int)
		return float64(A[x[0]])/float64(A[x[1]]) < float64(A[y[0]])/float64(A[y[1]])
	}
	var input []interface{}
	for k := range A {
		input = append(input, []int{k, lastindex})
	}

	heap := lib.NewHeap(intpairless, input...)

	var t []int
	for i := 0; i < K; i++ {
		top, _ := heap.Pop()
		t = top.([]int)

		if y := t[1] - 1; y > t[0] {
			heap.Add([]int{t[0], y})
		}
	}

	return []int{A[t[0]], A[t[1]]}
}

package leetcode

import "github.com/Chyroc/algorithms-go/lib"

/*
 * [941] Sort Array By Parity
 *
 * https://leetcode.com/problems/sort-array-by-parity/description/
 *
 * algorithms
 * Easy (70.92%)
 * Total Accepted:    22.6K
 * Total Submissions: 31.9K
 * Testcase Example:  '[3,1,2,4]'
 *
 * Given an array A of non-negative integers, return an array consisting of all
 * the even elements of A, followed by all the odd elements of A.
 * 
 * You may return any answer array that satisfies this condition.
 *
 * 
 * Example 1:
 *
 * Input: [3,1,2,4]
 * Output: [2,4,3,1]
 * The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be accepted.
 * 
 * Note:
 * 
 * 
 * 1 <= A.length <= 5000
 * 0 <= A[i] <= 5000
 *
 */

func sortArrayByParity(A []int) []int {
	s := make([]int, len(A))

	a := 0
	b := len(A) - 1

	for i := 0; i < len(A); i++ {
		if !lib.IsOdd(A[i]) {
			// 偶数
			s[a] = A[i]
			a++
		} else {
			// 奇数
			s[b] = A[i]
			b--
		}
	}
	return s
}

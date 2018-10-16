package leetcode

import "github.com/Chyroc/algorithms-go/lib"

/*
 * [958] Sort Array By Parity II
 *
 * https://leetcode.com/problems/sort-array-by-parity-ii/description/
 *
 * algorithms
 * Easy (68.27%)
 * Total Accepted:    5K
 * Total Submissions: 7.3K
 * Testcase Example:  '[4,2,5,7]'
 *
 * Given an array A of non-negative integers, half of the integers in A are
 * odd, and half of the integers are even.
 * 
 * Sort the array so that whenever A[i] is odd, i is odd; and whenever A[i] is
 * even, i is even.
 * 
 * You may return any answer array that satisfies this condition.
 * 
 *
 * Example 1:
 *
 * Input: [4,2,5,7]
 * Output: [4,5,2,7]
 * Explanation: [4,7,2,5], [2,5,4,7], [2,7,4,5] would also have been
 * accepted.
 * 
 *
 * Note:
 * 
 * 
 * 2 <= A.length <= 20000
 * A.length % 2 == 0
 * 0 <= A[i] <= 1000
 *
 */

/*

一个数组，一半奇数一半偶数，将奇数的数据放在index为奇数的位置，偶数的数据放在index为偶数的位置

*/

func sortArrayByParityII(A []int) []int {

	for i := 0; i < len(A); i++{
		if i%2!=0 {
			// index 是奇数

			if lib.IsOdd(A[i]) {
				// 数据是奇数
			} else {
				// 从当前位置向后寻找一个奇数，和当前数据交换
				var oddIndex = i
				for {
					oddIndex++
					if lib.IsOdd(A[oddIndex]) {
						break
					}
				}
				A[i], A[oddIndex] = A[oddIndex], A[i]
			}

		} else {
			// index 是偶数

			if lib.IsOdd(A[i]) {
				// 数据是奇数
				// 从当前位置向后寻找一个偶数，和当前数据交换
				var evenIndex = i
				for {
					evenIndex++
					if evenIndex >= len(A) {
						break
					}
					if !lib.IsOdd(A[evenIndex]) {
						break
					}
				}
				A[i], A[evenIndex] = A[evenIndex], A[i]
			} else {
			}
		}

	}

	return A
}

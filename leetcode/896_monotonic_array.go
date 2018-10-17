package leetcode

/*
 * [932] Monotonic Array
 *
 * https://leetcode.com/problems/monotonic-array/description/
 *
 * algorithms
 * Easy (54.45%)
 * Total Accepted:    14.5K
 * Total Submissions: 26.7K
 * Testcase Example:  '[1,2,2,3]'
 *
 * An array is monotonic if it is either monotone increasing or monotone
 * decreasing.
 *
 * An array A is monotone increasing if for all i <= j, A[i] <= A[j].  An array
 * A is monotone decreasing if for all i <= j, A[i] >= A[j].
 *
 * Return true if and only if the given array A is monotonic.
 *
 *
 * Example 1:
 *
 * Input: [1,2,2,3]
 * Output: true
 *
 *
 * Example 2:
 *
 * Input: [6,5,4,4]
 * Output: true
 *
 *
 * Example 3:
 *
 * Input: [1,3,2]
 * Output: false
 *
 *
 * Example 4:
 *
 * Input: [1,2,4,5]
 * Output: true
 *
 *
 * Example 5:
 *
 * Input: [1,1,1]
 * Output: true
 *
 *
 *
 * Note:
 *
 *
 * 1 <= A.length <= 50000
 * -100000 <= A[i] <= 100000
 *
 */

func isMonotonic(A []int) bool {
	switch len(A) {
	case 0, 1, 2:
		return true
	}

	if A[0] == A[1] {
		return isMonotonic(A[1:])
	}

	for i := 1; i < len(A)-1; i++ {
		d := A[i+1] - A[i]
		if A[0] > A[1] {
			// 递减，d<=0
			if d > 0 {
				return false
			}
		} else if A[0] < A[1] {
			// 递增, d>=0
			if d < 0 {
				return false
			}
		}
	}
	return true
}

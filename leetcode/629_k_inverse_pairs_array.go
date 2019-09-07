package leetcode

/*
 * @lc app=leetcode id=629 lang=golang
 *
 * [629] K Inverse Pairs Array
 *
 * https://leetcode.com/problems/k-inverse-pairs-array/description/
 *
 * algorithms
 * Hard (29.67%)
 * Total Accepted:    8.4K
 * Total Submissions: 28.5K
 * Testcase Example:  '3\n0'
 *
 * Given two integers n and k, find how many different arrays consist of
 * numbers from 1 to n such that there are exactly k inverse pairs.
 * 
 * We define an inverse pair as following: For ith and jth element in the
 * array, if i < j and a[i] > a[j] then it's an inverse pair; Otherwise, it's
 * not.
 * 
 * Since the answer may be very large, the answer should be modulo 109 + 7.
 * 
 * Example 1:
 * 
 * 
 * Input: n = 3, k = 0
 * Output: 1
 * Explanation: 
 * Only the array [1,2,3] which consists of numbers from 1 to 3 has exactly 0
 * inverse pair.
 * 
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: n = 3, k = 1
 * Output: 2
 * Explanation: 
 * The array [1,3,2] and [2,1,3] have exactly 1 inverse pair.
 * 
 * 
 * 
 * 
 * Note:
 * 
 * 
 * The integer n is in the range [1, 1000] and k is in the range [0, 1000].
 * 
 * 
 * 
 * 
 */

const m = 1000000007

func kInversePairs(n int, k int) int {

    dp := make([][]int, n+1)
    for i := range dp {
        dp[i] = make([]int, k+1)
        dp[i][0] = 1
    }

    for i := 1; i <= n; i++ {
        maxJ := min_629(k, i*(i-1)/2)
        for j := 1; j <= maxJ; j++ {
            dp[i][j] = (dp[i][j-1] + dp[i-1][j]) % m
            if j >= i {
                dp[i][j] -= dp[i-1][j-i]
                if dp[i][j] < 0 {
                    dp[i][j] += m
                }
            }
        }
    }

    return dp[n][k]
}

func min_629(a, b int) int {
    if a < b {
        return a
    }
    return b
}

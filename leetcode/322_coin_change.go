package leetcode

/*
 * [322] Coin Change
 *
 * https://leetcode.com/problems/coin-change/description/
 *
 * algorithms
 * Medium (27.07%)
 * Total Accepted:    111.4K
 * Total Submissions: 411K
 * Testcase Example:  '[1,2,5]\n11'
 *
 * You are given coins of different denominations and a total amount of money
 * amount. Write a function to compute the fewest number of coins that you need
 * to make up that amount. If that amount of money cannot be made up by any
 * combination of the coins, return -1.
 * 
 * Example 1:
 * 
 * 
 * Input: coins = [1, 2, 5], amount = 11
 * Output: 3 
 * Explanation: 11 = 5 + 5 + 1
 * 
 * Example 2:
 * 
 * 
 * Input: coins = [2], amount = 3
 * Output: -1
 * 
 * 
 * Note:
 * You may assume that you have an infinite number of each kind of coin.
 * 
 */


/*

* 使用指定面值的硬币，拼成一个指定数值（求和）；要求所用的硬币个数最少，求这个数；如果没有，返回-1
  * 对于这种求极值的问题，可以考虑使用动态规划
  * 递推式是：dp[i] = min(dp[i], dp[i - coins[j]] + 1);
    * 如果和是x，那么最多需要x个硬币（x个1块钱的）
    * 如果x减去一个可获得的硬币之后的y，这个y已经有答案了，那么看看y需要多少个，如果y+1小于现在的x需要的个数，就复制给x
*/
func coinChange(coins []int, amount int) int {
	max := amount + 1

	var dp = make([]int, max)
	for i := 0; i < max; i++ {
		dp[i] = max
	}

	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, v := range coins {
			d := i - v
			if d >= 0 {
				dp[i] = min(dp[i], dp[d]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1
	} else {
		return dp[amount]
	}
}

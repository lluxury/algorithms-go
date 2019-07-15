package leetcode

/*
 * @lc app=leetcode id=714 lang=golang
 *
 * [714] Best Time to Buy and Sell Stock with Transaction Fee
 *
 * https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/description/
 *
 * algorithms
 * Medium (50.61%)
 * Total Accepted:    41.9K
 * Total Submissions: 82.4K
 * Testcase Example:  '[1,3,2,8,4,9]\n2'
 *
 * Your are given an array of integers prices, for which the i-th element is
 * the price of a given stock on day i; and a non-negative integer fee
 * representing a transaction fee.
 * You may complete as many transactions as you like, but you need to pay the
 * transaction fee for each transaction.  You may not buy more than 1 share of
 * a stock at a time (ie. you must sell the stock share before you buy again.)
 * Return the maximum profit you can make.
 *
 * Example 1:
 *
 * Input: prices = [1, 3, 2, 8, 4, 9], fee = 2
 * Output: 8
 * Explanation: The maximum profit can be achieved by:
 * Buying at prices[0] = 1Selling at prices[3] = 8Buying at prices[4] =
 * 4Selling at prices[5] = 9The total profit is ((8 - 1) - 2) + ((9 - 4) - 2) =
 * 8.
 *
 *
 *
 * Note:
 * 0 < prices.length .
 * 0 < prices[i] < 50000.
 * 0 <= fee < 50000.
 *
 */

// 有手续费, 63 没看明白

// func maxProfit(prices []int, fee int) int {
func maxProfit_714(prices []int, fee int) int {
	empty := 0
	hold := -1 << 63

	for _, p := range prices {
		temp := empty
		empty = max_714(empty, hold+p)
		hold = max_714(hold, temp-p-fee)
	}
	return empty
}

func max_714(a, b int) int {
	if a > b {
		return a
	}
	return b
}
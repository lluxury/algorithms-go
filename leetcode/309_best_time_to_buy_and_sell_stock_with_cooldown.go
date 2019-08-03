package leetcode

/*
 * @lc app=leetcode id=309 lang=golang
 *
 * [309] Best Time to Buy and Sell Stock with Cooldown
 *
 * https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/description/
 *
 * algorithms
 * Medium (44.19%)
 * Total Accepted:    95.5K
 * Total Submissions: 215.5K
 * Testcase Example:  '[1,2,3,0,2]'
 *
 * Say you have an array for which the ith element is the price of a given
 * stock on day i.
 *
 * Design an algorithm to find the maximum profit. You may complete as many
 * transactions as you like (ie, buy one and sell one share of the stock
 * multiple times) with the following restrictions:
 *
 *
 * You may not engage in multiple transactions at the same time (ie, you must
 * sell the stock before you buy again).
 * After you sell your stock, you cannot buy stock on next day. (ie, cooldown 1
 * day)
 *
 *
 * Example:
 *
 *
 * Input: [1,2,3,0,2]
 * Output: 3
 * Explanation: transactions = [buy, sell, cooldown, buy, sell]
 *
 */

//  还是用动态规划, 代码很简单, 但思路必须清楚,
//  表格, 递归图, 总结代码, 多做几次熟练

func maxProfit_309(prices []int) int {
	// func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	buy := make([]int, n+1)
	buy[1] = 0 - prices[0]
	sel := make([]int, n+1)

	// for i := 2; i < n; i++ {
	for i := 2; i <= n; i++ {
		buy[i] = max_309(buy[i-1], sel[i-2]-prices[i-1])
		sel[i] = max_309(sel[i-1], buy[i-1]+prices[i-1])
	}
	return sel[n]

}

func max_309(a, b int) int {
	if a > b {
		return a
	}
	return b
}

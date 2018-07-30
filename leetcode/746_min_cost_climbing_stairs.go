package leetcode

import "fmt"

/*
 * [747] Min Cost Climbing Stairs
 *
 * https://leetcode.com/problems/min-cost-climbing-stairs/description/
 *
 * algorithms
 * Easy (43.80%)
 * Total Accepted:    35.4K
 * Total Submissions: 80.7K
 * Testcase Example:  '[0,0,0,0]'
 *
 *
 * On a staircase, the i-th step has some non-negative cost cost[i] assigned (0
 * indexed).
 *
 * Once you pay the cost, you can either climb one or two steps. You need to
 * find minimum cost to reach the top of the floor, and you can either start
 * from the step with index 0, or the step with index 1.
 *
 *
 * Example 1:
 *
 * Input: cost = [10, 15, 20]
 * Output: 15
 * Explanation: Cheapest is start on cost[1], pay that cost and go to the
 * top.
 *
 *
 *
 * Example 2:
 *
 * Input: cost = [1, 100, 1, 1, 1, 100, 1, 1, 100, 1]
 * Output: 6
 * Explanation: Cheapest is start on cost[0], and only step on 1s, skipping
 * cost[3].
 *
 *
 *
 * Note:
 *
 * cost will have a length in the range [2, 1000].
 * Every cost[i] will be an integer in the range [0, 999].
 *
 *
 */

/*

最大增序子序列的长度

动态规划算法

	首先需要定义 状态 和 状态转移方程

	状态
	如果有x个cost，就有x+1级台阶；假
	Fk = 达到第k级台阶所花费的最小费用

	状态转移
	达到第k级台阶，要么是达到F(k-1)加上第cost(k-1)的费用，要么是F(k-2)加上第cost(k-2)的费用，取最小值

	状态转移方程
	F0, F1 = cost[0]
	Fk = min(Fk-1 + cost[k-1], Fk-2 + cost[k-2])
	求Fk+1

*/

func minCostClimbingStairs(cost []int) int {
	switch len(cost) {
	case 0:
		return 0
	case 1, 2:
		return cost[0]
	default:
		m := make([]int, len(cost)+2)

		// index 1 - len

		// l是cost的序号   0 -  len(cost)
		// index是台阶序号  0 - l+1
		// l表示从index到index+1所花费的序号

		m[1] = 0 // 达到台阶1所花费
		m[2] = 0 // 达到台阶2所花费

		// i是台阶序号
		for i := 3; i <= len(cost)+1; i++ {
			m[i] = min(m[i-1]+cost[i-2], m[i-2]+cost[i-3])
		}
		fmt.Printf("m %#v\n", m)
		return m[len(m)-1]
	}
}

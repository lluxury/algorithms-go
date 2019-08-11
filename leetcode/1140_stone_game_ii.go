package leetcode

/*
 * @lc app=leetcode id=1140 lang=golang
 *
 * [1140] Stone Game II
 *
 * https://leetcode.com/problems/stone-game-ii/description/
 *
 * algorithms
 * Medium (60.61%)
 * Total Accepted:    3K
 * Total Submissions: 4.9K
 * Testcase Example:  '[2,7,9,4,4]'
 *
 * Alex and Lee continue their games with piles of stones.  There are a number
 * of piles arranged in a row, and each pile has a positive integer number of
 * stones piles[i].  The objective of the game is to end with the most
 * stones. 
 * 
 * Alex and Lee take turns, with Alex starting first.  Initially, M = 1.
 * 
 * On each player's turn, that player can take all the stones in the first X
 * remaining piles, where 1 <= X <= 2M.  Then, we set M = max(M, X).
 * 
 * The game continues until all the stones have been taken.
 * 
 * Assuming Alex and Lee play optimally, return the maximum number of stones
 * Alex can get.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: piles = [2,7,9,4,4]
 * Output: 10
 * Explanation:  If Alex takes one pile at the beginning, Lee takes two piles,
 * then Alex takes 2 piles again. Alex can get 2 + 4 + 4 = 10 piles in total.
 * If Alex takes two piles at the beginning, then Lee can take all three piles
 * left. In this case, Alex get 2 + 7 = 9 piles in total. So we return 10 since
 * it's larger. 
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * 1 <= piles.length <= 100
 * 1 <= piles[i] <= 10 ^ 4
 * 
 */
func stoneGameII(piles []int) int {
	n := len(piles)
	for i := n - 2; i >= 0; i-- {
		piles[i] += piles[i+1]
	}

	mem := [101][33]int{}
	var dp func(int, int) int
	dp = func(i, m int) int {
		if i+2*m >= n {
			return piles[i]
		}
		if mem[i][m] > 0 {
			return mem[i][m]
		}
		res := 0
		for x := 1; x <= 2*m; x++ {
			res = max_1140(
				res,
				piles[i]-dp(i+x, max_1140(m,x)),
			)
		}
		mem[i][m] = res
		return res
	}

	return dp(0, 1)

}

func max_1140(a,b int) int {
	if a > b {
		return a
	}
	return b
}

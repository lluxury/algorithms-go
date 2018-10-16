package leetcode

import "sort"

/*
 * [950] X of a Kind in a Deck of Cards
 *
 * https://leetcode.com/problems/x-of-a-kind-in-a-deck-of-cards/description/
 *
 * algorithms
 * Easy (32.91%)
 * Total Accepted:    5.9K
 * Total Submissions: 18.1K
 * Testcase Example:  '[1,2,3,4,4,3,2,1]'
 *
 * In a deck of cards, each card has an integer written on it.
 * 
 * Return true if and only if you can choose X >= 2 such that it is possible to
 * split the entire deck into 1 or more groups of cards, where:
 * 
 * 
 * Each group has exactly X cards.
 * All the cards in each group have the same integer.
 *
 * Example 1:
 * 
 * Input: [1,2,3,4,4,3,2,1]
 * Output: true
 * Explanation: Possible partition [1,1],[2,2],[3,3],[4,4]
 * 
 * 
 * 
 * Example 2:
 * 
 * Input: [1,1,1,2,2,2,3,3]
 * Output: false
 * Explanation: No possible partition.
 * 
 * 
 * 
 * Example 3:
 * 
 * Input: [1]
 * Output: false
 * Explanation: No possible partition.
 * 
 * 
 * 
 * Example 4:
 *
 * Input: [1,1]
 * Output: true
 * Explanation: Possible partition [1,1]
 * 
 * 
 * 
 * Example 5:
 * 
 * Input: [1,1,2,2,2,2]
 * Output: true
 * Explanation: Possible partition [1,1],[2,2],[2,2]
 *
 * Note:
 * 
 * 
 * 1 <= deck.length <= 10000
 * 0 <= deck[i] < 10000
 *
 * 
 */

func hasGroupsSizeX(deck []int) bool {
	var m = make(map[int]int)
	for _, v := range deck {
		m[v]++
	}

	var l []int
	for _, v := range m {
		l= append(l, v)
	}

	// l 是相同的数字的个数
	// 求l的最大公约数

	sort.Ints(l)

	for i := 2; i <= l[0]; i++ {
		hasMax := true
		for _, j := range l {
			if j%i != 0 {
				hasMax = false
				break
			}
		}
		if hasMax{
			return true
		}
	}
	return false
}

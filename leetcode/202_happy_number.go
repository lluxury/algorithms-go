package leetcode

import (
	"math"
	"sort"
	"strconv"
)

/*
 * [202] Happy Number
 *
 * https://leetcode.com/problems/happy-number/description/
 *
 * algorithms
 * Easy (42.96%)
 * Total Accepted:    188K
 * Total Submissions: 437.6K
 * Testcase Example:  '19'
 *
 * Write an algorithm to determine if a number is "happy".
 * 
 * A happy number is a number defined by the following process: Starting with
 * any positive integer, replace the number by the sum of the squares of its
 * digits, and repeat the process until the number equals 1 (where it will
 * stay), or it loops endlessly in a cycle which does not include 1. Those
 * numbers for which this process ends in 1 are happy numbers.
 * 
 * Example: 
 * 
 * 
 * Input: 19
 * Output: true
 * Explanation: 
 * 12 + 92 = 82
 * 82 + 22 = 68
 * 62 + 82 = 100
 * 12 + 02 + 02 = 1
 * 
 */

/*

各位平方之和等于1

*/

func area_sum(a []int) int {
	var sum int
	for _, v := range a {
		sum += int(math.Pow(float64(v), 2))
	}
	return sum
}

func join(a []int) string {
	var s string
	sort.Ints(a)
	for _, v := range a {
		s += strconv.Itoa(v)
	}
	return s
}

func isHappy_rec(n int, cache map[string]bool) bool {
	var a []int
	for n > 0 {
		geweishu := n % 10
		a = append(a, geweishu)
		n = (n - geweishu) / 10
	}
	newx := area_sum(a)
	cache_key := join(a)

	switch {
	case newx == 1:
		return true
	case cache[cache_key]:
		return false
	}

	cache[cache_key] = true

	return isHappy_rec(newx, cache)
}

func isHappy(n int) bool {
	return isHappy_rec(n, map[string]bool{})
}

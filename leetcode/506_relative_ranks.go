package leetcode

import (
	"strconv"
)

/*
 * [506] Relative Ranks
 *
 * https://leetcode.com/problems/relative-ranks/description/
 *
 * algorithms
 * Easy (47.07%)
 * Total Accepted:    33.8K
 * Total Submissions: 71.7K
 * Testcase Example:  '[5,4,3,2,1]'
 *
 *
 * Given scores of N athletes, find their relative ranks and the people with
 * the top three highest scores, who will be awarded medals: "Gold Medal",
 * "Silver Medal" and "Bronze Medal".
 *
 * Example 1:
 *
 * Input: [5, 4, 3, 2, 1]
 * Output: ["Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"]
 * Explanation: The first three athletes got the top three highest scores, so
 * they got "Gold Medal", "Silver Medal" and "Bronze Medal". For the left two
 * athletes, you just need to output their relative ranks according to their
 * scores.
 *
 *
 *
 * Note:
 *
 * N is a positive integer and won't exceed 10,000.
 * All the scores of athletes are guaranteed to be unique.
 *
 *
 *
 */

func QuickSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	mid := data[0]
	head, tail := 0, len(data)-1
	for i := 1; i <= tail; {
		if data[i] < mid {
			data[i], data[tail] = data[tail], data[i]
			tail--
		} else {
			data[i], data[head] = data[head], data[i]
			head++
			i++
		}
	}
	QuickSort(data[:head])
	QuickSort(data[head+1:])

	return data
}

func findRelativeRanks(nums []int) []string {
	nnn := make([]int, len(nums))
	copy(nnn, nums)

	QuickSort(nums)
	m := make(map[int]int)
	for k, v := range nums {
		m[v] = k
	}

	var s []string
	for _, v := range nnn {
		switch m[v] {
		case 0:
			s = append(s, "Gold")
		case 1:
			s = append(s, "Silver")
		case 2:
			s = append(s, "Bronze")
		default:
			s = append(s, strconv.Itoa(m[v]+1))
		}
	}
	return s
}

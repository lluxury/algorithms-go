package leetcode

/*
 * [879] Maximize Distance to Closest Person
 *
 * https://leetcode.com/problems/maximize-distance-to-closest-person/description/
 *
 * algorithms
 * Easy (37.97%)
 * Total Accepted:    7.1K
 * Total Submissions: 18.6K
 * Testcase Example:  '[1,0,0,0,1,0,1]'
 *
 * In a row of seats, 1 represents a person sitting in that seat, and 0
 * represents that the seat is empty.
 *
 * There is at least one empty seat, and at least one person sitting.
 *
 * Alex wants to sit in the seat such that the distance between him and the
 * closest person to him is maximized.
 *
 * Return that maximum distance to closest person.
 *
 *
 * Example 1:
 *
 *
 * Input: [1,0,0,0,1,0,1]
 * Output: 2
 * Explanation:
 * If Alex sits in the second open seat (seats[2]), then the closest person has
 * distance 2.
 * If Alex sits in any other open seat, the closest person has distance 1.
 * Thus, the maximum distance to the closest person is 2.
 *
 *
 * Example 2:
 *
 *
 * Input: [1,0,0,0]
 * Output: 3
 * Explanation:
 * If Alex sits in the last seat, the closest person is 3 seats away.
 * This is the maximum distance possible, so the answer is 3.
 *
 *
 * Note:
 *
 *
 * 1 <= seats.length <= 20000
 * seats contains only 0s or 1s, at least one 0, and at least one 1.
 *
 *
 *
 *
 */

/*

统计连续0的个数
  * 靠墙（index等于0的个数，这是靠前面的；index等于len，这是靠后墙），返回个数
  * 否则，人会加1除以2
  * 记录上面的值的最大值

*/

func maxDistToClosest(seats []int) int {
	max, count := 0, 0
	for i := 0; i < len(seats); i++ {
		if seats[i] == 1 {
			count = 0
			continue
		}
		count++
		dist := 0
		if count-(i+1) == 0 || i == len(seats)-1 {
			dist = count
		} else {
			dist = (count + 1) / 2
		}
		if dist > max {
			max = dist
		}
	}
	return max
}

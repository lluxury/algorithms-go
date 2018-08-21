package leetcode

import "fmt"

/*
 * [753] Open the Lock
 *
 * https://leetcode.com/problems/open-the-lock/description/
 *
 * algorithms
 * Medium (41.03%)
 * Total Accepted:    7.2K
 * Total Submissions: 17.6K
 * Testcase Example:  '["0201","0101","0102","1212","2002"]\n"0202"'
 *
 *
 * You have a lock in front of you with 4 circular wheels.  Each wheel has 10
 * slots: '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'.  The wheels can
 * rotate freely and wrap around: for example we can turn '9' to be '0', or '0'
 * to be '9'.  Each move consists of turning one wheel one slot.
 *
 * The lock initially starts at '0000', a string representing the state of the
 * 4 wheels.
 *
 * You are given a list of deadends dead ends, meaning if the lock displays any
 * of these codes, the wheels of the lock will stop turning and you will be
 * unable to open it.
 *
 * Given a target representing the value of the wheels that will unlock the
 * lock, return the minimum total number of turns required to open the lock, or
 * -1 if it is impossible.
 *
 *
 * Example 1:
 *
 * Input: deadends = ["0201","0101","0102","1212","2002"], target = "0202"
 * Output: 6
 * Explanation:
 * A sequence of valid moves would be "0000" -> "1000" -> "1100" -> "1200" ->
 * "1201" -> "1202" -> "0202".
 * Note that a sequence like "0000" -> "0001" -> "0002" -> "0102" -> "0202"
 * would be invalid,
 * because the wheels of the lock become stuck after the display becomes the
 * dead end "0102".
 *
 *
 *
 * Example 2:
 *
 * Input: deadends = ["8888"], target = "0009"
 * Output: 1
 * Explanation:
 * We can turn the last wheel in reverse to move from "0000" -> "0009".
 *
 *
 *
 * Example 3:
 *
 * Input: deadends = ["8887","8889","8878","8898","8788","8988","7888","9888"],
 * target = "8888"
 * Output: -1
 * Explanation:
 * We can't reach the target without getting stuck.
 *
 *
 *
 * Example 4:
 *
 * Input: deadends = ["0000"], target = "8888"
 * Output: -1
 *
 *
 *
 * Note:
 *
 * The length of deadends will be in the range [1, 500].
 * target will not be in the list deadends.
 * Every string in deadends and the string target will be a string of 4 digits
 * from the 10,000 possibilities '0000' to '9999'.
 *
 *
 */

/*

四个旋转锁，每个可以选0-9，求转到给定的四个数字（不能经过死锁的数字）（要求路径最小）

和迷宫问题一致
* bfs

*/

func openLock_add(x int) int {
	if x < 9 {
		return x + 1
	}
	return 0
}

func openLock_sub(x int) int {
	if x > 0 {
		return x - 1
	}
	return 9
}

func openLock_tostring(s []int) string {
	return fmt.Sprintf("%d%d%d%d", s[0], s[1], s[2], s[3])
}

func openLock(deadends []string, target string) int {
	done := make(map[string]bool)   // 访问过的
	pres := make(map[string]string) // 前置节点
	deadends_map := make(map[string]bool)
	for _, v := range deadends {
		deadends_map[v] = true
	}

	if deadends_map["0000"] {
		return -1
	}

	queues := [][]int{{0, 0, 0, 0}}
	done["0000"] = true
	for {
		if len(queues) == 0 {
			break
		}

		queue := queues[0]
		queues = append([][]int{}, queues[1:]...)
		//fmt.Printf("queues %#v\n", queues)
		//fmt.Printf("访问： %v,%v,%v,%v\n", queue[0], queue[1], queue[2], queue[3])

		queue_string := openLock_tostring(queue)
		//fmt.Printf("expect %v, got %v\n", target, queue_string)
		if queue_string == target {
			dep := 0
			for {
				queue_string = pres[queue_string]
				dep++
				if queue_string == "0000" {
					return dep
				}
			}
		}

		for i := 0; i < 4; i++ {
			for j := 0; j < 2; j++ {
				newqueue := copy_int_slice(queue)
				if j == 0 {
					//fmt.Printf("%v, %v add %v\n", openLock_tostring(newqueue), newqueue[i], openLock_add(newqueue[i]))
					newqueue[i] = openLock_add(newqueue[i])
					//fmt.Printf("i %v, j %v, old: %v, next: %v\n", i, j, queue_string, openLock_tostring(newqueue))
				} else {
					//fmt.Printf("%v, %v sub %v\n", openLock_tostring(newqueue), newqueue[i], openLock_sub(newqueue[i]))
					newqueue[i] = openLock_sub(newqueue[i])
					//fmt.Printf("i %v, j %v, old: %v, next: %v\n", i, j, queue_string, openLock_tostring(newqueue))
				}

				newqueue_string := openLock_tostring(newqueue)
				if !deadends_map[newqueue_string] && !done[newqueue_string] {
					pres[newqueue_string] = queue_string
					//fmt.Printf("newqueue %#v\n", newqueue)
					queues = append(queues, newqueue)
				}
				done[newqueue_string] = true
			}
		}
	}

	return -1
}

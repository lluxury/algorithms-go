package leetcode

import (
	"fmt"
	"math"
)

/*
 * [906] Walking Robot Simulation
 *
 * https://leetcode.com/problems/walking-robot-simulation/description/
 *
 * algorithms
 * Easy (25.35%)
 * Total Accepted:    2.1K
 * Total Submissions: 8.4K
 * Testcase Example:  '[4,-1,3]\n[]'
 *
 * A robot on an infinite grid starts at point (0, 0) and faces north.  The
 * robot can receive one of three possible types of commands:
 *
 *
 * -2: turn left 90 degrees
 * -1: turn right 90 degrees
 * 1 <= x <= 9: move forward x units
 *
 *
 * Some of the grid squares are obstacles.
 *
 * The i-th obstacle is at grid point (obstacles[i][0], obstacles[i][1])
 *
 * If the robot would try to move onto them, the robot stays on the previous
 * grid square instead (but still continues following the rest of the route.)
 *
 * Return the square of the maximum Euclidean distance that the robot will be
 * from the origin.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: commands = [4,-1,3], obstacles = []
 * Output: 25
 * Explanation: robot will go to (3, 4)
 *
 *
 *
 * Example 2:
 *
 *
 * Input: commands = [4,-1,4,-2,4], obstacles = [[2,4]]
 * Output: 65
 * Explanation: robot will be stuck at (1, 4) before turning left and going to
 * (1, 8)
 *
 *
 *
 *
 *
 * Note:
 *
 *
 * 0 <= commands.length <= 10000
 * 0 <= obstacles.length <= 10000
 * -30000 <= obstacle[i][0] <= 30000
 * -30000 <= obstacle[i][1] <= 30000
 * The answer is guaranteed to be less than 2 ^ 31.
 *
 */

func degreTo(d int) (int, int) {
	if d < -180 {
		return degreTo(d + 360)
	}
	if d > 180 {
		return degreTo(d - 360)
	}

	switch d {
	case -180, 180:
		return 0, -1
	case -90:
		return -1, 0
	case 0, 360:
		return 0, 1
	case 90:
		return 1, 0
	default:
		panic(fmt.Sprintf("error d %d", d))
	}
}

func robotSim(commands []int, obstacles [][]int) int {
	x, y := 0, 0
	d := 0
	maaa := make(map[[2]int]bool)
	for _, v := range obstacles {
		maaa[[2]int{v[0], v[1]}] = true
	}
	for _, v := range commands {
		switch v {
		case -1: // 右转90
			d += 90
		case -2: // 左转90
			d -= 90
		default:
			m, n := degreTo(d)
			for k := 0; k < v; k++ {
				x2 := x + m
				y2 := y + n

				if maaa[[2]int{x2, y2}] {
					break
				}
				x = x2
				y = y2
			}
		}
	}

	return int(math.Pow(float64(x), 2) + math.Pow(float64(y), 2))
}

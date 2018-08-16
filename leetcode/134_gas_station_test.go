package leetcode

import (
	"testing"

	"github.com/Chyroc/algorithms-go/test"
)

/*

> https://leetcode.com/problems/gas-station/description/


There are N gas stations along a circular route, where the amount of gas at station i is gas[i].

You have a car with an unlimited gas tank and it costs cost[i] of gas to travel from station i to its next station (i+1). You begin the journey with an empty tank at one of the gas stations.

Return the starting gas station's index if you can travel around the circuit once in the clockwise direction, otherwise return -1.

Note:

If there exists a solution, it is guaranteed to be unique.
Both input arrays are non-empty and have the same length.
Each element in the input arrays is a non-negative integer.
Example 1:

Input:
gas  = [1,2,3,4,5]
cost = [3,4,5,1,2]

Output: 3

Explanation:
Start at station 3 (index 3) and fill up with 4 unit of gas. Your tank = 0 + 4 = 4
Travel to station 4. Your tank = 4 - 1 + 5 = 8
Travel to station 0. Your tank = 8 - 2 + 1 = 7
Travel to station 1. Your tank = 7 - 3 + 2 = 6
Travel to station 2. Your tank = 6 - 4 + 3 = 5
Travel to station 3. The cost is 5. Your gas is just enough to travel back to station 3.
Therefore, return 3 as the starting index.
Example 2:

Input:
gas  = [2,3,4]
cost = [3,4,3]

Output: -1

Explanation:
You can't start at station 0 or 1, as there is not enough gas to travel to the next station.
Let's start at station 2 and fill up with 4 unit of gas. Your tank = 0 + 4 = 4
Travel to station 0. Your tank = 4 - 3 + 2 = 3
Travel to station 1. Your tank = 3 - 3 + 3 = 3
You cannot travel back to station 2, as it requires 4 unit of gas but you only have 3.
Therefore, you can't travel around the circuit once no matter where you start.



题目
  * 车子绕着一个圈圈开车，圈圈上有若干加油站，每个加油站可以加若干油，从一个加油站到另外一个加油站，会消耗一些油
  * 输入是两个数组，分别是每个加油站可以加的油，和从当前加油站到下一站所需要的油
  * 思路
    * 首先出发点的油需要比耗费的油多
    * 然后从当前站点加上下一次加的，减去下一次耗费的
    * 只要到终点的时候，一直大于0，就说明可以成功
    * 为了节省计算，「加上下一次加的，减去下一次耗费的」可以先计算出来
*/

func canCompleteCircuit(gas []int, cost []int) int {
	gas_len := len(gas)
	if gas_len != len(cost) || gas_len == 0 {
		return -1
	}
	if gas_len == 1 {
		if gas[0] >= cost[0] {
			return 0
		}
		return -1
	}

	var diff = make([]int, gas_len)
	var posi_index []int
	for i := 0; i < gas_len; i++ {
		if i == 0 {
			diff[i] = gas[i] - cost[gas_len-1]
		} else {
			diff[i] = gas[i] - cost[i-1]
		}

		if gas[i] > cost[i] {
			posi_index = append(posi_index, i)
		}

	}

	for _, index := range posi_index {
		x := 0
		end := false
		for i := index; i < gas_len+index && x >= 0; i++ {
			j := i
			if i >= gas_len {
				j = i - gas_len
			}
			x = x + gas[j] - cost[j]

			if i == gas_len+index-1 && x >= 0 {
				end = true
			}
		}

		if end {
			return index
		}
	}

	return -1
}

func Test_134(t *testing.T) {
	test.Runs(t, canCompleteCircuit, []*test.Case{
		{Input: `[1,2,3,4,5] \n [3,4,5,1,2]`, Output: `3`},
		{Input: `[2,3,4] \n [3,4,3]`, Output: `-1`},
		{Input: `[2] \n [2]`, Output: `0`},
	})
}

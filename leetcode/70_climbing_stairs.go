package leetcode

/*
You are climbing a stair case. It takes n steps to reach to the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Note: Given n will be a positive integer.

Example 1:

Input: 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps
Example 2:

Input: 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step*/

// 动态规划,基础题
// DP 状态定义
// DP 状态方程

// func climbStairs(n int) int {
//     if (n == 0 || n == 1 || n == 2){
//         return n
//     }

//     mem := make([]int,n)
//     mem[0] = 1
//     mem[1] = 2
//     for i := 2; i < n; i++{
//         mem[i]=mem[i-1]+ mem[i-2]
//     }
//     return mem[n-1]
// }

func climbStairs(n int) int {

	// slice1 := make([]type, len) 长度必须
	d := make([]int, n+1)
	d[0] = 1
	d[1] = 1

	for i := 2; i <= n; i++ {
		d[i] = d[i-1] + d[i-2]
	}
	return d[n]
}

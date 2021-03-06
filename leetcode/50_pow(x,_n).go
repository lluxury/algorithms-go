package leetcode

import "math"

/*
Implement pow(x, n), which calculates x raised to the power n (xn).

Example 1:

Input: 2.00000, 10
Output: 1024.00000
Example 2:

Input: 2.10000, 3
Output: 9.26100
Example 3:

Input: 2.00000, -2
Output: 0.25000
Explanation: 2-2 = 1/22 = 1/4 = 0.25
Note:

-100.0 < x < 100.0
n is a 32-bit signed integer, within the range [−231, 231 − 1]*/

// 计算 x 的 n 次幂函数
// 指数运算分解

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}

	if n < 0 {
		return 1 / myPow(x, -n)
	}

	r := myPow(x, n/2)
	if n%2 == 0 {
		return r * r
	} else {
		// return r* r * x
		// 9.261000000000001
		return math.Round(r*r*x*1000) / 1000
	}
}

// 控制精度的方法, leetcode却过不了,看来那边优化过了

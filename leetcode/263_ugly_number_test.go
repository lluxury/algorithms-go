package leetcode

import (
	"github.com/Chyroc/algorithms-go/test"
	"testing"
)

/*
> https://leetcode.com/problems/ugly-number/description/


Write a program to check whether a given number is an ugly number.

Ugly numbers are positive numbers whose prime factors only include 2, 3, 5.

Example 1:

Input: 6
Output: true
Explanation: 6 = 2 × 3
Example 2:

Input: 8
Output: true
Explanation: 8 = 2 × 2 × 2
Example 3:

Input: 14
Output: false
Explanation: 14 is not ugly since it includes another prime factor 7.
Note:

1 is typically treated as an ugly number.
Input is within the 32-bit signed integer range: [−231,  231 − 1].

* 判断一个数，能不能被2，3，5除尽（1认为是可以）
  * 不停的整除2，3，5即可
  * 如果商最后是1，返回true
*/
func isUgly(num int) bool {
	if num > 0 {
		for _, v := range []int{2, 3, 5} {
			for num%v == 0 {
				num /= v
			}
		}
	}

	return num == 1
}

func Test_263(t *testing.T) {
	test.Runs(t, isUgly, []*test.Case{
		{Input: "6", Output: "true"},
		{Input: "8", Output: "true"},
		{Input: "14", Output: "false"},
	})
}
package leetcode

import (
	"testing"

	"github.com/Chyroc/algorithms-go/test"
)

/*
> https://leetcode.com/problems/super-ugly-number/description/

Write a program to find the nth super ugly number.

Super ugly numbers are positive numbers whose all prime factors are in the given prime list primes of size k.

Example:

Input: n = 12, primes = [2,7,13,19]
Output: 32
Explanation: [1, 2, 4, 7, 8, 13, 14, 16, 19, 26, 28, 32]  is the sequence of the first 12
             super ugly numbers given primes = [2,7,13,19] of size 4.
Note:

1 is a super ugly number for any given primes.
The given numbers in primes are in ascending order.
0 < k ≤ 100, 0 < n ≤ 106, 0 < primes[i] < 1000.
The nth super ugly number is guaranteed to fit in a 32-bit signed integer.


* 给定一个素数数组，求一个数组（只能由这个素数数组乘出来，从小到大排列），求这个数组中的第n个数
  * 和264题很像：那个是给定了素数数组235，固定的，解法也是一样的
  * 将素数数组和这个慢慢求出来的数组相乘，最小数加到结果数组中
  * 素数长度a
  * 用一个长度和素数数组长度一致的数组保存每个素数已经乘的结果数组的index，名字是x数组，长度a
    * 肯定是从小开始乘结果最小
    * 已经用过的就不再用
    * 所以保存一个index
  * 然后用x数组取出a个数，分别和素数数组对应的位置相乘，得到a个数的数组y，然后计算出最小值m
  * 将y数组中结果等于m的index对应的数组a的对应位置的数字+1
*/

func min_all(is ...int) int {
	switch len(is) {
	case 0:
		panic("length 0")
	case 1:
		return is[0]
	case 2:
		if is[0] < is[1] {
			return is[0]
		} else {
			return is[1]
		}
	default:
		return min_all(append([]int{min_all(is[0], is[1])}, is[2:]...)...)
	}
}

func nthSuperUglyNumber(n int, primes []int) int {
	var values = []int{1} // 1是第一个
	var chenguodeshu = make([]int, len(primes))

	for len(values) < n {
		var jisuanjieguo []int
		for k, v := range primes {
			jisuanjieguo = append(jisuanjieguo, values[chenguodeshu[k]]*v)
		}

		m := min_all(jisuanjieguo...)
		values = append(values, m)
		for k, v := range jisuanjieguo {
			if v == m {
				chenguodeshu[k]++
			}
		}
	}

	return values[n-1]
}

func Test_313(t *testing.T) {
	test.Runs(t, nthSuperUglyNumber, []*test.Case{
		{Input: `12 \n [2,7,13,19]`, Output: `32`},
	})
}

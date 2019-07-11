package leetcode

import "fmt"

/*
> https://leetcode.com/problems/base-7/description/


Given an integer, return its base 7 string representation.

Example 1:
```
Input: 100
Output: "202"
```
Example 2:
```
Input: -7
Output: "-10"
```
Note: The input will be in range of `[-1e7, 1e7]`.

- 将10进制数转化为7进制字符串
- 注意0和负数的处理
- 循环，取余，「减去取余的再除以7」

*/
func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}

	fu := false
	if num < 0 {
		fu = true
		num = -num
	}

	r := ""
	for num != 0 {
		r = fmt.Sprintf("%d", num%7) + r
		num = (num - num%7) / 7
	}

	if fu {
		r = "-" + r
	}

	return r
}

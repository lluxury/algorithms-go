package leetcode

import (
	"github.com/Chyroc/algorithms-go/test"
	"testing"
)

/*
> https://leetcode.com/problems/2-keys-keyboard/description/


Initially on a notepad only one character 'A' is present. You can perform two operations on this notepad for each step:

Copy All: You can copy all the characters present on the notepad (partial copy is not allowed).
Paste: You can paste the characters which are copied last time.
Given a number n. You have to get exactly n 'A' on the notepad by performing the minimum number of steps permitted. Output the minimum number of steps to get n 'A'.

Example 1:
Input: 3
Output: 3
Explanation:
Intitally, we have one character 'A'.
In step 1, we use Copy All operation.
In step 2, we use Paste operation to get 'AA'.
In step 3, we use Paste operation to get 'AAA'.
Note:
The n will be in the range [1, 1000].

* 现在有一个A，你每次只能进行两个操作："复制前面所有的字符串"，"粘贴字符串"，请问经过至少多少步，可以获得n个A
  * 如果现在有一半也就是n/2个A，那么只需要复制+粘贴，2步即可
  * 如果现在有n/3个A，那么：复制+粘贴+粘贴，3步
  * 也就是还有m个A的时候，就需要m步，并且n/x，x越小，需要的步骤越小
  * 那么d从2开始循环到n（n是n个1，需要n步）
*/
func minSteps(n int) int {
	var s int
	for d := 2; d <= n; d++ {
		for n%d == 0 {
			s += d
			n /= d
		}
	}
	return s
}

func Test_650(t *testing.T) {
	test.Runs(t, minSteps, []*test.Case{
		{Input: "3", Output: "3"},
	})
}

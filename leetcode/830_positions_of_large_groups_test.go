package leetcode

import (
	"testing"

	"github.com/Chyroc/algorithms-go/test"
)

/*

> https://leetcode.com/problems/positions-of-large-groups/description/


In a string S of lowercase letters, these letters form consecutive groups of the same character.

For example, a string like S = "abbxxxxzyy" has the groups "a", "bb", "xxxx", "z" and "yy".

Call a group large if it has 3 or more characters.  We would like the starting and ending positions of every large group.

The final answer should be in lexicographic order.



Example 1:

Input: "abbxxxxzzy"
Output: [[3,6]]
Explanation: "xxxx" is the single large group with starting  3 and ending positions 6.
Example 2:

Input: "abc"
Output: []
Explanation: We have "a","b" and "c" but no large group.
Example 3:

Input: "abcdddeeeeaabbbcd"
Output: [[3,5],[6,9],[12,14]]


给定一个字符串，求连续有三个字符以及以上相等出现的次数
  * 遍历字符串
  * 需要记录前一次的开始index，前一次的字符值，前一次字符出现的次数
  * 遍历的时候
    * 如果值和前一次的值相等，字符数+1
    * 如果不相等，重置前一次的开始index，前一次的字符值，前一次字符出现的次数
  * 注意这种遍历还需要处理最后一个值的情况


*/

func largeGroupPositions(S string) [][]int {
	if len(S) < 3 {
		return nil
	}

	var (
		pre_index = 0
		pre_value = S[0]
		length    = 1
		r         [][]int
	)
	for k := 1; k < len(S); k++ {
		if S[k] == pre_value {
			length++
		} else {
			if length >= 3 {
				r = append(r, []int{pre_index, k - 1})
			}

			pre_value = S[k]
			pre_index = k
			length = 1
		}
	}

	if S[len(S)-1] == S[len(S)-2] && length >= 3 {
		r = append(r, []int{pre_index, len(S) - 1})
	}

	return r
}

func Test_830(t *testing.T) {
	test.Runs(t, largeGroupPositions, []*test.Case{
		{Input: `abbxxxzxzzy`, Output: `[[3,5]]`},
		{Input: `abbxxzxxzzy`, Output: `[]`},
		{Input: `abbxxxxzzy`, Output: `[[3,6]]`},
		{Input: `abc`, Output: `[]`},
		{Input: `bababbabaa`, Output: `[]`},
		{Input: `aaa`, Output: `[[0,2]]`},
		{Input: `abcdddeeeeaabbbcd`, Output: `[[3,5],[6,9],[12,14]]`},
	})
}

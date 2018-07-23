package leetcode

import (
	"testing"
	"github.com/Chyroc/algorithms-go/test"
)

/*

> https://leetcode.com/problems/find-all-anagrams-in-a-string/description/


Given a string s and a non-empty string p, find all the start indices of p's anagrams in s.

Strings consists of lowercase English letters only and the length of both strings s and p will not be larger than 20,100.

The order of output does not matter.

Example 1:

Input:
s: "cbaebabacd" p: "abc"

Output:
[0, 6]

Explanation:
The substring with start index = 0 is "cba", which is an anagram of "abc".
The substring with start index = 6 is "bac", which is an anagram of "abc".
Example 2:

Input:
s: "abab" p: "ab"

Output:
[0, 1, 2]

Explanation:
The substring with start index = 0 is "ab", which is an anagram of "ab".
The substring with start index = 1 is "ba", which is an anagram of "ab".
The substring with start index = 2 is "ab", which is an anagram of "ab".

给定两个字符串，如果第一个字符串的连续子串可以由第二个字符串组成，则添加子串的第一个字符的index
  * 很显然，需要遍历第一个字符串，添加新的字符，删除第二个字符长度之前的字符
  * 然后比较当前map和第二个字符串的各个字符的个数是否相等
  * 因为字符的范围固定是a-z，利用"xxx-'a'"技巧可以将数据存储在一个长度为26的[]uint8数组中，而不是一个map中
  * 数组的index是字符的值，数组的值是该index表示的字符出现的次数

*/
func equalUint8Map(m1, m2 []uint8) bool {
	for k := range m1 {
		if m1[k] != m2[k] {
			return false
		}
	}
	return true
}

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return nil
	}
	var scount = make([]uint8, 26)
	var pcount = make([]uint8, 26)
	for k := range p {
		scount[s[k]-'a']++
		pcount[p[k]-'a']++
	}

	var r []int
	if equalUint8Map(pcount, scount) {
		r = append(r, 0)
	}

	for i := len(p); i < len(s); i++ {
		scount[s[i]-'a']++
		scount[s[i-len(p)]-'a']--

		if equalUint8Map(pcount, scount) {
			r = append(r, i-len(p)+1)
		}
	}

	return r
}

func Test_438(t *testing.T) {
	test.Runs(t, findAnagrams, []*test.Case{
		{Input: `cbaebabacd \n abc`, Output: `[0,6]`},
		{Input: `abab \n ab`, Output: `[0,1,2]`},
	})
}

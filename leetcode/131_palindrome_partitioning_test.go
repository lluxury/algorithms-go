package leetcode

import (
	"github.com/Chyroc/algorithms-go/lib"
	"github.com/Chyroc/algorithms-go/test"
	"testing"
)

/*
> https://leetcode.com/problems/palindrome-partitioning/description/

Given a string s, partition s such that every substring of the partition is a palindrome.

Return all possible palindrome partitioning of s.

Example:

Input: "aab"
Output:
[
  ["aa","b"],
  ["a","a","b"]
]

* 给一个字符串，求子串为回文单词的集合
  * 遍历字符串，

*/
func partition_isPalindrome(s string, low, high int) bool {
	for low < high {
		if s[low] != s[high] {
			return false
		}
		low++
		high--
	}
	return true
}

func backtrack_partition(list *[][]string, tempList []string, s string, start int) {
	if start == len(s) {
		*list = append(*list, lib.StringArrayDeepCopy(tempList))
	}

	for i := start; i < len(s); i++ {
		if partition_isPalindrome(s, start, i) {
			backtrack_partition(list, append(tempList, s[start:i+1]), s, i+1)
		}
	}
}

func partition(s string) (list [][]string) {
	backtrack_partition(&list, []string{}, s, 0)
	return
}

func Test_131(t *testing.T) {
	test.Runs(t, partition, []*test.Case{
		{Input: `aab`, Output: `[[a,a,b],[aa,b]]`},
	})
}

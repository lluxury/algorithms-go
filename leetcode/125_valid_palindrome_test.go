package leetcode

import "github.com/Chyroc/algorithms-go/lib"

/*
> https://leetcode.com/problems/valid-palindrome/description/

Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.

Note: For the purpose of this problem, we define empty string as valid palindrome.

Example 1:

Input: "A man, a plan, a canal: Panama"
Output: true
Example 2:

Input: "race a car"
Output: false

* 回文字符串判断，忽略大小写并假设只有数字和字母
- i从前往后，j从后往前扫描，只要不一样就返回false
- 这个循环里面，注意跳过非字母和数字
*/

// isAlpha 是否是字母
func isAlpha(s rune) bool {
	return (s >= 97 && s <= 122) || (s >= 65 && s <= 90) || (s >= 48 && s <= 57)
}

// isAlphaEqual 是否是字母，并且转为小写
func isAlphaEqual(s1, s2 rune) bool {
	if s1 >= 97 && s1 <= 122 {
		s1 = s1 - 32
	}
	if s2 >= 97 && s2 <= 122 {
		s2 = s2 - 32
	}

	return s1 == s2
}

func isPalindrome(s string) bool {
	if s == "" {
		return true
	}

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		for i < j && !isAlpha(int32(s[i])) {
			i++
		}
		for i < j && !isAlpha(int32(s[j])) {
			j--
		}

		if !isAlphaEqual(int32(s[i]), int32(s[j])) {
			return false
		}
	}

	return true
}


func Example_125_isPalindrome() {
	lib.ExamplePrint(isPalindrome("A man, a plan, a canal: Panama"))
	lib.ExamplePrint(isPalindrome("race a car"))

	// Output:
	// true
	// false
}

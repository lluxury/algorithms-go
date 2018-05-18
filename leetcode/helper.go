package leetcode

func min(s ...int) int {
	if len(s) == 0 {
		panic("no value")
	} else if len(s) == 1 {
		return s[0]
	}

	if s[0] > s[1] {
		return min(append(s[2:], s[1])...)
	} else {
		return min(append(s[2:], s[0])...)
	}
}

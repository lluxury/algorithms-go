package lib

func Max(s ...int) int {
	switch len(s) {
	case 0:
		panic("max")
	case 1:
		return s[0]
	case 2:
		if s[0] > s[1] {
			return s[0]
		}
		return s[1]
	default:
		return Max(append([]int{Max(s[0], s[1])}, s[2:]...)...)
	}
}

func Min(s ...int) int {
	switch len(s) {
	case 0:
		panic("min")
	case 1:
		return s[0]
	case 2:
		if s[0] < s[1] {
			return s[0]
		}
		return s[1]
	default:
		return Min(append([]int{Min(s[0], s[1])}, s[2:]...)...)
	}
}

// 是不是字母
func IsLetter(s byte) bool {
	return (s >= 'a' && s <= 'z') || (s >= 'A' && s <= 'Z')
}

// 是不是奇数
func IsOdd(i int) bool {
	return i%2 != 0
}

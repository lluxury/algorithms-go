package leetcode

func sqrt(a int) float64 {
	x := float64(1)

	for {
		x = float64(float64(a)/x+x) / 2
		check := x*x - float64(a)

		if check < 0 {
			check = -check
		}
		if check <= 0.001 {
			break
		}
	}

	return x
}

func isDigit(s uint8) bool {
	return s >= '0' && s <= '9'
}

func toDigit(s uint8) int {
	return int(s - '0')
}

func max(s ...int) int {
	switch len(s) {
	case 0:
		panic("max all")
	case 1:
		return s[0]
	case 2:
		if s[0] > s[1] {
			return s[0]
		}
		return s[1]
	default:
		return max(append([]int{max(s[0], s[1])}, s[2:]...)...)
	}
}

func min(s ...int) int {
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
		return min(append([]int{min(s[0], s[1])}, s[2:]...)...)
	}
}

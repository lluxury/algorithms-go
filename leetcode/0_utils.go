package leetcode

// 求平方根的方法
//   * 先假设1是a的平方根x
//   * 将a除以x，再加上x，除以2，得到新的假设值x：x = ( x + a/x ) / 2
//   * 当假设值x的平方和a越来越接近的时候，说明答案越来越准确
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
		panic("max")
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

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

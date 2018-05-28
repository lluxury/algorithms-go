package lib

import "math/rand"
import "time"

func sliceOutOfOrder(rr *rand.Rand, a1 [2]int, in []int) []int {
	l := len(in)
	for i := l - 1; i > 0; i-- {
		r := rr.Intn(i)
		in[r], in[i] = in[i], in[r]
	}
	a1[in[0]] += 1
	return in
}

func randslice(rr *rand.Rand, a2 []int,n int) {
	in := rr.Perm(n)
	a2[in[0]] += 1
}

func RandSlice(n int) []int {
	var a1 = [2]int{}
	var a2 = make([]int, n)
	var rr = rand.New(rand.NewSource(time.Now().UnixNano()))

	a := []int{0, 1}
	for i := 0; i < 10000; i++ {
		randslice(rr, a2,n)
		sliceOutOfOrder(rr, a1, a)
	}
	return a2
}

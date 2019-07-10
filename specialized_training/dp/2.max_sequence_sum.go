package dp

import (
	"fmt"
	"github.com/lluxury/algorithms-go/lib"
)

func MaxSequenceSum(a []int) int {
	var sum = make([]int, len(a))
	for i := 0; i < len(a); i++ {
		if i == 0 {
			sum[i] = a[i]
		} else {
			sum[i] = lib.Max(sum[i-1]+a[i], a[i])
		}
	}

	fmt.Println(sum)

	return lib.Max(sum...)
}

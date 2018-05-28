package lib

import (
	"testing"
	"fmt"

	"github.com/stretchr/testify/assert"
)

func Benchmark_Heap(b *testing.B) {
	as := assert.New(b)

	for i := 0; i < b.N; i++ {
		for _, v := range []int{1, 2, 3, 5, 10, 20} {
			input := RandSlice(v)
			output := []int{}
			h := NewMaxHeap(input...)
			for range input {
				x, ok := h.Pop()
				as.True(ok)
				output = append(output, x)
			}

			//fmt.Printf("output %v\n", output)
			for i := 0; i < len(output)-1; i++ {
				as.True(output[i] >= output[i+1], fmt.Sprintf("array(%#v), %d should >= %d", output, output[i], output[i+1]))
			}
		}
	}
}

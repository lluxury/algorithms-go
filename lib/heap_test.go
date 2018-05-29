package lib

import (
	"testing"
	"fmt"

	"github.com/stretchr/testify/assert"
)

func intLess(less bool) func(a, b interface{}) bool {
	if less {
		return func(a, b interface{}) bool {
			return a.(int) < b.(int)
		}
	} else {
		return func(a, b interface{}) bool {
			return a.(int) > b.(int)
		}
	}
}

func Benchmark(b *testing.B) {
	b.Run("MaxHeap int", func(b *testing.B) {
		as := assert.New(b)

		for i := 0; i < b.N; i++ {
			for _, v := range []int{1, 2, 3, 5, 10, 20} {
				input1 := RandSlice(v)
				var input []interface{}
				for _, v := range input1 {
					input = append(input, v)
				}
				output := []int{}
				h := NewHeap(intLess(false), input...)
				for range input {
					x, ok := h.Pop()
					as.True(ok)
					output = append(output, x.(int))
				}

				for i := 0; i < len(output)-1; i++ {
					as.True(output[i] >= output[i+1], fmt.Sprintf("array(%#v), %d should >= %d", output, output[i], output[i+1]))
				}
			}
		}
	})

	b.Run("MinHeap int", func(b *testing.B) {
		as := assert.New(b)

		for i := 0; i < b.N; i++ {
			for _, v := range []int{1, 2, 3, 5, 10, 20} {
				input1 := RandSlice(v)
				var input []interface{}
				for _, v := range input1 {
					input = append(input, v)
				}
				output := []int{}
				h := NewHeap(intLess(true), input...)
				for range input {
					x, ok := h.Pop()
					as.True(ok)
					output = append(output, x.(int))
				}

				for i := 0; i < len(output)-1; i++ {
					as.True(output[i] <= output[i+1], fmt.Sprintf("array(%#v), %d should >= %d", output, output[i], output[i+1]))
				}
			}
		}
	})

	b.Run("MinHeap int pair", func(b *testing.B) {
		as := assert.New(b)

		var intpairless = func(a, b interface{}) bool {
			x := a.([]int)
			y := b.([]int)
			return float64(x[0])/float64(x[1]) < float64(y[0])/float64(y[1])
		}

		for i := 0; i < b.N; i++ {
			input1 := RandSlice(20)
			input2 := RandSlice(20)
			var input []interface{}
			for k := range input1 {
				input = append(input, []int{input1[k], input2[k]})
			}

			output := []float64{}
			h := NewHeap(intpairless, input...)
			for range input {
				x, ok := h.Pop()
				as.True(ok)
				output = append(output, float64(x.([]int)[0])/float64(x.([]int)[1]))
			}

			for i := 0; i < len(output)-1; i++ {
				as.True(output[i] <= output[i+1], fmt.Sprintf("array(%#v), %.3f should >= %.3f", output, output[i], output[i+1]))
			}
		}
	})
}

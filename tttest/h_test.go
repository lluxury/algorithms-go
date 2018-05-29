package tttest

import (
	"testing"
	"fmt"

	"github.com/stretchr/testify/assert"
	"github.com/Chyroc/algorithms-go/lib"
)

/*

参考：http://128kj.iteye.com/blog/1728555

《满二叉树》
* 一个二叉树，如果每一个层的结点数都达到最大值，则这个二叉树就是满二叉树

《完全二叉树》
* 若设二叉树的深度为h，除第 h 层外，其它各层 (1～h-1) 的结点数都达到最大个数，第 h 层所有的结点都连续集中在最左边，这就是完全二叉树。
* 第n个节点的父节点是[n/2]
  * []是向下取整
  * 一个有n个节点的完全二叉树的最后一个非叶节点（最后一个节点的父节点）是节点[n/2]
  * 7节点，(6,7)->3 (4,5)->2
*/

type LessFunc func(a, b interface{}) bool

type heap struct {
	heap []interface{} //第一个位置不存储数据，下标从1开始，因为0index求子节点不好求
	less LessFunc
}

// 使用n个数据初始化最大堆
// 首先将所有数据加到数组，然后排序最大堆
func NewHeap(less LessFunc, xs ...interface{}) *heap {
	r := new(heap)
	r.heap = append(r.heap, 0) //第一个位置没有数据
	for _, v := range xs {
		r.heap = append(r.heap, v)
	}
	r.less = less

	r.adjust()

	return r
}



// 添加节点
// 将x添加到最后一个叶子节点，然后将最后一个节点上浮
func (r *heap) Add(x interface{}) {
	if r.length() == 0 {
		r.heap = append(r.heap, 0)
	}
	r.heap = append(r.heap, x)
	r.up(r.length())
}

// 获取最大的数据
func (r *heap) Pop() (interface{}, bool) {
	//fmt.Printf("pop %s\n", r)
	if r.length() > 0 {
		x := r.heap[1]
		r.swap(1, r.length())
		r.heap = r.heap[:len(r.heap)-1]
		r.adjust()
		return x, true
	}

	return 0, false
}

// 上浮
// 从index获取父节点的p-index，如果大于父节点，则上浮并递归，否则返回
// p-index = index/2
func (r *heap) up(index int) {
	parent := index / 2
	if r.less(r.heap[index], r.heap[parent]) {
		r.swap(index, parent) // 子节点index上浮
		r.up(parent)
	}
}

// 下沉
// 获取他的叶子节点
// 如果叶子节点不为空，且
// 左孩节点l-child=index*2   (3 -> 3*2)
func (r *heap) down(index int) {
	var child int

	// 如果有子节点的话，求左孩与右孩中较大的子节点的index
	if d := index*2 - r.length(); d < 0 {
		// 有左右孩
		child = index * 2 //左孩节点
		//fmt.Printf("%s/%d %d %d\n", r, len(r.heap), index, child)
		if r.less(r.heap[child+1], r.heap[child]) {
			child++
		}
	} else if d == 0 {
		// 只有左孩
		child = index * 2
	} else {
		// 没有子节点
		return
	}

	// 子节点大于当前节点，下沉
	if r.less(r.heap[child], r.heap[index]) {
		r.swap(child, index)
		r.down(child)
	}
}

// 对对进行排序
// 即将所有节点满足：「大于等于子节点的值」即可
// 然后因为叶子节点没有子节点，所以不需要满足：「比所有子节点大」的条件
// 所以从最后一个非叶子节点（也就是lastindex/2）开始调整：保证该节点比左孩和右孩大（也就是下沉该节点）
func (r *heap) adjust() {
	//fmt.Printf("adjust %s\n", r)

	// 从最后一个叶子节点的父节点开始，到第一个节点的顺序做
	for index := r.length() / 2; index > 0; index-- {
		r.down(index)
	}
}

// 长度
func (r *heap) length() int {
	if len(r.heap) == 0 {
		return 0
	} else {
		return len(r.heap) - 1 //第一个位置没有数据
	}
}

// 交换第a，b位置的数
func (r *heap) swap(a, b int) {
	tmp := r.heap[a]
	r.heap[a] = r.heap[b]
	r.heap[b] = tmp
}

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
				input1 := lib.RandSlice(v)
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
				input1 := lib.RandSlice(v)
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
			input1 := lib.RandSlice(20)
			input2 := lib.RandSlice(20)
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

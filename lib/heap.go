package lib

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
	r.heap = append(r.heap, nil) //第一个位置没有数据

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
	if len(r.heap) == 0 {
		r.heap = append(r.heap, nil)
	}
	r.heap = append(r.heap, x)
	r.up(r.Len())
}

// 获取最大的数据
func (r *heap) Pop() (interface{}, bool) {
	if len(r.heap) > 1 {
		x := r.heap[1]
		r.swap(1, r.Len())
		r.heap = r.heap[:len(r.heap)-1]
		r.adjust()
		return x, true
	}

	return nil, false
}

// 上浮
// 从index获取父节点的p-index，如果大于父节点，则上浮并递归，否则返回
// p-index = index/2
func (r *heap) up(index int) {
	if index > 1 { // index 1 是根节点
		parent := index / 2
		if r.less(r.heap[index], r.heap[parent]) {
			r.swap(index, parent) // 子节点index上浮
			r.up(parent)
		}
	}

}

// 下沉
// 获取他的叶子节点
// 如果叶子节点不为空，且
// 左孩节点l-child=index*2   (3 -> 3*2)
func (r *heap) down(index int) {
	var child int

	// 如果有子节点的话，求左孩与右孩中较大的子节点的index
	if d := index*2 - r.Len(); d < 0 {
		// 有左右孩
		child = index * 2 //左孩节点
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
	// 从最后一个叶子节点的父节点开始，到第一个节点的顺序做
	for index := r.Len() / 2; index > 0; index-- {
		r.down(index)
	}
}

// 长度
func (r *heap) Len() int {
	switch x := len(r.heap); x {
	case 0, 1:
		return 0
	default:
		return x - 1 //第一个位置没有数据
	}
}

// 交换第a，b位置的数
func (r *heap) swap(a, b int) {
	tmp := r.heap[a]
	r.heap[a] = r.heap[b]
	r.heap[b] = tmp
}

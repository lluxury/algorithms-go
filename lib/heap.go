package lib

// 最小/最大堆实现

/*

参考：http://128kj.iteye.com/blog/1728555

满二叉树
  * 一个二叉树，如果每一个层的结点数都达到最大值，则这个二叉树就是满二叉树

完全二叉树
  * 解释1：将所有节点按照从上往下，从左往右的顺序，往二叉树上安装，那么结果就是一个完全二叉树
  * 解释2：若设二叉树的深度为h，除第 h 层外，其它各层 (1～h-1) 的结点数都达到最大个数，第 h 层所有的结点都连续集中在最左边，这就是完全二叉树。

可以使用数组模拟完全二叉树
  * 第n个节点的父节点是 n/2向下取整
    * 一个有n个节点的完全二叉树的最后一个非叶节点（最后一个节点的父节点）是节点 n/2向下取整
    * 示例：7节点，(6,7)->3 (4,5)->2
  * 第n个节点的左孩/右孩节点分别是index*2和index*2+1

堆
* 也叫做优先队列
* 队列是先入先出的，堆的数据是按照一定的顺序先入先出的（比如最小堆，永远是最小的先出）
* 可以使用完全二叉树（即可以用数组实现）实现堆

用数组实现堆
* 理解原理
  * 用数组模拟完全二叉树，用完全二叉树保证，所有的的节点都比孩子节点小（假设是最小堆）
* 首先理解3个操作
  * 上浮，因为要保证「所有的的节点都比孩子节点小」，所以如果当前节点比父节点小，就交换当前节点和父节点，然后递归处理父节点
  * 下沉，因为要保证「所有的的节点都比孩子节点小」，所以
    * 如果有两个子节点，取较大的那个；如果只有一个子节点，用这个操作；没有子节点，肯定满足，直接返回
    * 如果当前节点比上面取到的子节点大，交换两个节点，并递归处理子节点
  * 什么时候上浮
    * 添加节点的时候
    * 先加到数组的最后，然后上浮最后一个节点
  * 什么时候下沉
    * 「在调整」中调用
  * 什么是调整
    * 从数组的中部开始，一直到根节点
    * 一直调用下沉操作
    * 结果就是使得堆有序
* 堆的操作
  * 添加：添加节点到最后，然后对最后一个节点递归上浮
  * 获取：返回第一个数据；然后将最后一个数据移动到第一个位置；然后从数组中部到根节点，递归调用下沉

*/

type LessFunc func(a, b interface{}) bool

type Heap struct {
	heap []interface{} //第一个位置不存储数据，下标从1开始，因为0index求子节点不好求
	less LessFunc
}

// 使用n个数据初始化最大堆
// 首先将所有数据加到数组，然后排序最大堆
func NewHeap(less LessFunc, xs ...interface{}) *Heap {
	r := new(Heap)
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
func (r *Heap) Add(x interface{}) {
	if len(r.heap) == 0 {
		r.heap = append(r.heap, nil)
	}
	r.heap = append(r.heap, x)
	r.up(r.Len())
}

// 获取最大的数据
func (r *Heap) Pop() (interface{}, bool) {
	if len(r.heap) > 1 {
		x := r.heap[1]
		r.swap(1, r.Len())
		r.heap = r.heap[:len(r.heap)-1]
		r.adjust()
		return x, true
	}

	return nil, false
}

// 获取最大的数据
func (r *Heap) Peek() interface{} {
	if len(r.heap) > 1 {
		return r.heap[1]
	}

	return nil
}

// 上浮
// 从index获取父节点的p-index，如果大于父节点，则上浮并递归，否则返回
// p-index = index/2
func (r *Heap) up(index int) {
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
func (r *Heap) down(index int) {
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
func (r *Heap) adjust() {
	// 从最后一个叶子节点的父节点开始，到第一个节点的顺序做
	for index := r.Len() / 2; index > 0; index-- {
		r.down(index)
	}
}

// 长度
func (r *Heap) Len() int {
	switch x := len(r.heap); x {
	case 0, 1:
		return 0
	default:
		return x - 1 //第一个位置没有数据
	}
}

// 交换第a，b位置的数
func (r *Heap) swap(a, b int) {
	tmp := r.heap[a]
	r.heap[a] = r.heap[b]
	r.heap[b] = tmp
}

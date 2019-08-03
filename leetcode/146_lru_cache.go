package leetcode

/*

> https://leetcode.com/problems/lru-cache/description/

Design and implement a data structure for Least Recently Used (LRU) cache. It should support the following operations: get and put.

get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
put(key, value) - Set or insert the value if the key is not already present. When the cache reached its capacity, it should invalidate the least recently used item before inserting a new item.

Follow up:
Could you do both operations in O(1) time complexity?

Example:

LRUCache cache = new LRUCache( 2 /* capacity  );

cache.put(1, 1);
cache.put(2, 2);
cache.get(1);       // returns 1
cache.put(3, 3);    // evicts key 2
cache.get(2);       // returns -1 (not found)
cache.put(4, 4);    // evicts key 1
cache.get(1);       // returns -1 (not found)
cache.get(3);       // returns 3
cache.get(4);       // returns 4


实现一个LRU Cache
（leetcode测试通过）
    * 需要使用两种数据结构：双向链表和哈希表
    * 思路：
    * 因为需要是O(1)操作，所以一般是哈希表存储数据
    * 为了能够更新活跃数据的生命，删除老旧数据，需要维护一个有序表，
    	在这里选择双向链表，时间复杂度低
    * 双向链表：
        1. 头节点和尾节点都不存储数据，正式数据存储于头节点和尾节点之间的节点之间
        2. 在链表中，靠前的数据是「最少」用到的，靠后的是「用的多」的
        3. 所以，移除无用数据的时候，请删除头结点的下一个节点；
        添加新节点的时候，添加到尾节点的前一个；
        更新一个存储的节点的生命的时候，先移除他，然后再添加他（即自动的到最后去了）
    * 业务：
        * get数据：更新数据的生命：删除并添加节点
        * set数据：如果存在，则更新数据租期，并修改map的值；
        如果没有，则添加数据（链表+map），
        如果长度超出，删除最旧的数据（头结点的下一个节点）
*/
type Link struct {
	key   int
	value int
	prev  *Link
	next  *Link
}

type LRUCache struct {
	map1     map[int]*Link
	capacity int
	head     *Link
	tail     *Link
}

// 1. 头节点和尾节点都不存储数据，
// 正式数据存储于头节点和尾节点之间的节点之间
// 2. 在链表中，靠前的数据是「最少」用到的，靠后的是「用的多」的
// 3. 所以，移除无用数据的时候，请删除头结点的下一个节点；
// 添加新节点的时候，添加到尾节点的前一个；
// 更新一个存储的节点的生命的时候，先移除他，然后再添加他（即自动的到最后去了）

// func Constructor(capacity int) LRUCache {
func Constructor_146(capacity int) LRUCache {
	cap := LRUCache{
		map1:     make(map[int]*Link),
		capacity: capacity,
		head:     new(Link),
		tail:     new(Link),
	}
	cap.head.next = cap.tail
	cap.tail.prev = cap.head
	return cap
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.map1[key]
	if ok {
		this.delete(v)
		this.add(v)
		return v.value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	v, ok := this.map1[key]
	if ok {
		this.delete(v)
		this.add(v)
		this.map1[key].value = value
	} else {
		link := &Link{
			key:   key,
			value: value,
		}
		this.add(link)
		this.map1[key] = link
		if len(this.map1) > this.capacity {
			data := this.head.next
			this.delete(data)
			delete(this.map1, data.key)
		}
	}
}

// 删除一个节点：将这个节点的pre和next链接起来
func (this *LRUCache) delete(node *Link) {
	prev := node.prev
	next := node.next
	next.prev = prev
	prev.next = next
}

// 添加一个节点到最后（实际上是最后一个节点和倒数第二个节点之间添加一个节点）：
//   * 先拿到最后一个节点
//   * 将最后一个节点的pre设置为node，倒数第二个节点的next设置为nodex  [] -> [] <- []
//   * node的pre和next分别设置为倒数第二个节点和最后一个节点          [] <- [] -> []
func (this *LRUCache) add(node *Link) {
	prev := this.tail.prev
	prev.next = node
	this.tail.prev = node
	node.prev = prev
	node.next = this.tail
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

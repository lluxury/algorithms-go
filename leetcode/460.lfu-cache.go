package leetcode

import "sort"

/*
 * [460] LFU Cache
 *
 * https://leetcode.com/problems/lfu-cache/description/
 *
 * algorithms
 * Hard (25.59%)
 * Total Accepted:    23.7K
 * Total Submissions: 92.6K
 * Testcase Example:  '["LFUCache","put","put","get","put","get","get","put","get","get","get"]\n[[2],[1,1],[2,2],[1],[3,3],[2],[3],[4,4],[1],[3],[4]]'
 *
 * Design and implement a data structure for Least Frequently Used (LFU) cache.
 * It should support the following operations: get and put.
 *
 *
 *
 * get(key) - Get the value (will always be positive) of the key if the key
 * exists in the cache, otherwise return -1.
 * put(key, value) - Set or insert the value if the key is not already present.
 * When the cache reaches its capacity, it should invalidate the least
 * frequently used item before inserting a new item. For the purpose of this
 * problem, when there is a tie (i.e., two or more keys that have the same
 * frequency), the least recently used key would be evicted.
 *
 *
 * Follow up:
 * Could you do both operations in O(1) time complexity?
 *
 * Example:
 *
 * LFUCache cache = new LFUCache( 2 cap );
 *
 * cache.put(1, 1);
 * cache.put(2, 2);
 * cache.get(1);       // returns 1
 * cache.put(3, 3);    // evicts key 2
 * cache.get(2);       // returns -1 (not found)
 * cache.get(3);       // returns 3.
 * cache.put(4, 4);    // evicts key 1.
 * cache.get(1);       // returns -1 (not found)
 * cache.get(3);       // returns 3
 * cache.get(4);       // returns 4
 *
 *
 */

// m 存储key以及使用的次数

// m 存储数据, k-v
// heap 存储使用量，最小堆
// capacity 堆容量，超过的删除堆顶
type LFUCache struct {
	m        map[int]*heap_useage
	capacity int
}

type heap_useage struct {
	Key   int
	Value int
	Count int
}

type heap_useages struct {
	s []*heap_useage
}

func (r heap_useages) Len() int { return len(r.s) }

func (r heap_useages) Less(i, j int) bool { return r.s[i].Count < r.s[j].Count }

func (r heap_useages) Swap(i, j int) { r.s[i], r.s[j] = r.s[j], r.s[i] }

func getMinCount(m map[int]*heap_useage) int {
	s := []*heap_useage{}
	for _, v := range m {
		s = append(s, v)
	}
	h := heap_useages{s}
	sort.Sort(h)
	return h.s[0].Key
}

func Constructor_460(capacity int) LFUCache {
	return LFUCache{
		m:        make(map[int]*heap_useage),
		capacity: capacity,
	}
}

func (this *LFUCache) Get(key int) int {
	if v, ok := this.m[key]; ok {
		v.Count++
		return v.Value
	}
	return -1
}

func (this *LFUCache) Put(key int, value int) {
	v, ok := this.m[key]
	if ok {
		v.Count++
		v.Value = value
		return
	}

	if len(this.m) >= this.capacity {
		minKey := getMinCount(this.m)
		delete(this.m, minKey)
	}

	this.m[key] = &heap_useage{
		Key:   key,
		Value: value,
		Count: 1,
	}
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

package leetcode

// import "sort"
import "time"
import "container/heap"

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
	m 			map[int]*entry
	pq			PQ
	cap	int
}

func Constructor_460(capacity int) LFUCache  {
	return LFUCache{
		m: 		make(map[int]*entry, capacity),
		pq:		make(PQ, 0, capacity),
		cap:	capacity,
	}
}

func (c *LFUCache) Get(key int) int  {
	if ep, ok := c.m[key];ok{
		c.pq.update(ep)
		return ep.value
	}
	return -1
}

func (c *LFUCache) Put(key int, value int)  {
	if c.cap <= 0 {
		return 
	}
	
	ep,ok := c.m[key]
	if ok {
		c.m[key].value = value
		c.pq.update(ep)
		return 
	}

	ep = &entry{key:key, value:value}
	if len(c.pq) == c.cap{
		temp := heap.Pop(&c.pq).(*entry)
		delete(c.m, temp.key)
	}
	
	c.m[key] = ep
	heap.Push(&c.pq, ep)
}


type entry struct {
	key 		int
	value		int
	frequence	int
	index		int
	date	time.Time
}

type PQ []*entry

func (pq PQ) Len() int  { return len(pq) }
func (pq PQ) Less(i,j int) bool {
	if pq[i].frequence == pq[j].frequence{
		return pq[i].date.Before(pq[j].date)
	}
	return pq[i].frequence < pq[j].frequence
}

func (pq PQ) Swap(i,j int) {
	pq[i],pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PQ) Push(x interface{})  {
	n:= len(*pq)
	entry := x.(*entry)
	entry.index = n
	entry.date = time.Now()
	*pq = append(*pq, entry)
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	entry := old[n-1]
	entry.index = -1 // for safety
	*pq = old[0 : n-1]
	return entry
}

// update modifies the priority of an entry in the queue.
func (pq *PQ) update(entry *entry) {
	entry.frequence++
	entry.date = time.Now()
	heap.Fix(pq, entry.index)
}





/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

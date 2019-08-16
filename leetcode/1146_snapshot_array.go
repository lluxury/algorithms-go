package leetcode

/*
 * @lc app=leetcode id=1146 lang=golang
 *
 * [1146] Snapshot Array
 *
 * https://leetcode.com/problems/snapshot-array/description/
 *
 * algorithms
 * Medium (22.13%)
 * Total Accepted:    4.1K
 * Total Submissions: 13.7K
 * Testcase Example:  '["SnapshotArray","set","snap","set","get"]\n[[3],[0,5],[],[0,6],[0,0]]'
 *
 * Implement a SnapshotArray that supports the following interface:
 * 
 * 
 * SnapshotArray(int length) initializes an array-like data structure with the
 * given length.  Initially, each element equals 0.
 * void set(index, val) sets the element at the given index to be equal to
 * val.
 * int snap() takes a snapshot of the array and returns the snap_id: the total
 * number of times we called snap() minus 1.
 * int get(index, snap_id) returns the value at the given index, at the time we
 * took the snapshot with the given snap_id
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: ["SnapshotArray","set","snap","set","get"]
 * [[3],[0,5],[],[0,6],[0,0]]
 * Output: [null,null,0,null,5]
 * Explanation: 
 * SnapshotArray snapshotArr = new SnapshotArray(3); // set the length to be 3
 * snapshotArr.set(0,5);  // Set array[0] = 5
 * snapshotArr.snap();  // Take a snapshot, return snap_id = 0
 * snapshotArr.set(0,6);
 * snapshotArr.get(0,0);  // Get the value of array[0] with snap_id = 0, return
 * 5
 * 
 * 
 * Constraints:
 * 
 * 
 * 1 <= length <= 50000
 * At most 50000 calls will be made to set, snap, and get.
 * 0 <= index < length
 * 0 <= snap_id < (the total number of times we call snap())
 * 0 <= val <= 10^9
 * 
 * 
 */
type SnapshotArray struct {
    snaps   []map[int]int
    current map[int]int
    id      int
}


// func Constructor(length int) SnapshotArray {
func Constructor_1146(length int) SnapshotArray {
    s := make([]map[int]int, 0,128)
    c := make(map[int]int,32)
    return SnapshotArray{
        snaps :     s,
        current:    c,
        id:         0,
    }
}


func (this *SnapshotArray) Set(index int, val int)  {
    this.current[index] = val
}


func (this *SnapshotArray) Snap() int {
    this.snaps = append(this.snaps, this.current)
    this.current = make(map[int]int)
    this.id++
    return this.id -1
}


func (this *SnapshotArray) Get(index int, snap_id int) int {
    for id := snap_id; id >= 0; id --{
        history := this.snaps[id]
        if t, ok := history[index]; ok {
            return t
        }
    }
    return 0
}


/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */

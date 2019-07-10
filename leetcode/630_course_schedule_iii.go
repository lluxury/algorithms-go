package leetcode

import (
	"github.com/lluxury/algorithms-go/lib"
	"sort"
)

/*
 * [630] Course Schedule III
 *
 * https://leetcode.com/problems/course-schedule-iii/description/
 *
 * algorithms
 * Hard (29.52%)
 * Total Accepted:    7.2K
 * Total Submissions: 24.5K
 * Testcase Example:  '[[100,200],[200,1300],[1000,1250],[2000,3200]]'
 *
 *
 * There are n different online courses numbered from 1 to n. Each course has
 * some duration(course length)  t and closed on dth day. A course should be
 * taken continuously for t days and must be finished before or on the dth day.
 * You will start at the 1st day.
 *
 *
 *
 * Given n online courses represented by pairs (t,d), your task is to find the
 * maximal number of courses that can be taken.
 *
 *
 *
 * Example:
 *
 * Input: [[100, 200], [200, 1300], [1000, 1250], [2000, 3200]]
 * Output: 3
 * Explanation:
 * There're totally 4 courses, but you can take 3 courses at most:
 * First, take the 1st course, it costs 100 days so you will finish it on the
 * 100th day, and ready to take the next course on the 101st day.
 * Second, take the 3rd course, it costs 1000 days so you will finish it on the
 * 1100th day, and ready to take the next course on the 1101st day.
 * Third, take the 2nd course, it costs 200 days so you will finish it on the
 * 1300th day.
 * The 4th course cannot be taken now, since you will finish it on the 3300th
 * day, which exceeds the closed date.
 *
 *
 *
 *
 * Note:
 *
 * The integer 1
 * You can't take two courses simultaneously.
 *
 *
 */

type sort_int_pair struct {
	pair [][]int
}

func (r sort_int_pair) Len() int {
	return len(r.pair)
}

func (r sort_int_pair) Less(i, j int) bool {
	return r.pair[i][1] < r.pair[j][1]
}

func (r sort_int_pair) Swap(i, j int) {
	r.pair[i], r.pair[j] = r.pair[j], r.pair[i]
}

func scheduleCourse(courses [][]int) int {
	// 按照结束时间排序
	p := sort_int_pair{courses}
	sort.Sort(p)

	// 最大堆，按照耗费时间排序；
	course := lib.NewHeap(func(a, b interface{}) bool { return a.([]int)[0] > b.([]int)[0] })

	// 遍历，一个个添加课程
	// 如果一个课程的结束时间大于预期的时间，说明不能完成这个课程
	// 这个时候，不是删除这个课程，而是删除过去耗费时间最长的一个课程（由最大堆决定）
	start := 0
	for _, pair := range p.pair {
		course.Add(pair)
		start += pair[0]
		if start > pair[1] {
			x, _ := course.Pop()
			start -= x.([]int)[0]
		}
	}
	return course.Len()
}

package leetcode

/*
 * [207] Course Schedule
 *
 * https://leetcode.com/problems/course-schedule/description/
 *
 * algorithms
 * Medium (34.55%)
 * Total Accepted:    141.1K
 * Total Submissions: 408.1K
 * Testcase Example:  '2\n[[1,0]]'
 *
 * There are a total of n courses you have to take, labeled from 0 to n-1.
 *
 * Some courses may have prerequisites, for example to take course 0 you have
 * to first take course 1, which is expressed as a pair: [0,1]
 *
 * Given the total number of courses and a list of prerequisite pairs, is it
 * possible for you to finish all courses?
 *
 * Example 1:
 *
 *
 * Input: 2, [[1,0]]
 * Output: true
 * Explanation: There are a total of 2 courses to take.
 * To take course 1 you should have finished course 0. So it is possible.
 *
 * Example 2:
 *
 *
 * Input: 2, [[1,0],[0,1]]
 * Output: false
 * Explanation: There are a total of 2 courses to take.
 * To take course 1 you should have finished course 0, and to take course 0 you
 * should
 * also have finished course 1. So it is impossible.
 *
 *
 * Note:
 *
 *
 * The input prerequisites is a graph represented by a list of edges, not
 * adjacency matrices. Read more about how a graph is represented.
 * You may assume that there are no duplicate edges in the input
 * prerequisites.
 *
 *
 */

/*

给一个数组，每个数组有两个数字，表示学习第一个课程的话，必须要先学习第二个课程（从第二个课程指向第一个课程的图）
问：能否完成所有给定的课程

即问一个图里面有没有有环图

入度：指向这个节点的个数

从所有的入度为0的节点出发，将经历过的节点的入度-1

如果没有环，那么最后所有的节点的入度都是0
如果还有节点的入度大于0，说明有环


*/

func canFinish(numCourses int, prerequisites [][]int) bool {
	tu := make(map[int][]int)
	rudu := make(map[int]int)
	for i := 0; i < numCourses; i++ {
		rudu[i] = 0
	}
	for _, vv := range prerequisites {
		rudu[vv[0]]++
		tu[vv[1]] = append(tu[vv[1]], vv[0])
	}

	// 队列是所有入度为0的节点
	queues := []int{}
	for k, v := range rudu {
		if v == 0 {
			queues = append(queues, k)
		}
	}

	// 遍历所有入度为0的节点
	for {
		if len(queues) == 0 {
			break
		}

		queue := queues[0]
		queues = append([]int{}, queues[1:]...)

		for _, v := range tu[queue] {
			rudu[v]--
			if rudu[v] == 0 {
				queues = append(queues, v)
			}
		}
	}

	for _, v := range rudu {
		if v != 0 {
			return false
		}
	}
	return true

}

package leetcode

/*
 * [210] Course Schedule II
 *
 * https://leetcode.com/problems/course-schedule-ii/description/
 *
 * algorithms
 * Medium (31.24%)
 * Total Accepted:    101.9K
 * Total Submissions: 325.6K
 * Testcase Example:  '2\n[[1,0]]'
 *
 * There are a total of n courses you have to take, labeled from 0 to n-1.
 * 
 * Some courses may have prerequisites, for example to take course 0 you have
 * to first take course 1, which is expressed as a pair: [0,1]
 * 
 * Given the total number of courses and a list of prerequisite pairs, return
 * the ordering of courses you should take to finish all courses.
 * 
 * There may be multiple correct orders, you just need to return one of them.
 * If it is impossible to finish all courses, return an empty array.
 * 
 * Example 1:
 * 
 * 
 * Input: 2, [[1,0]] 
 * Output: [0,1]
 * Explanation: There are a total of 2 courses to take. To take course 1 you
 * should have finished   
 * course 0. So the correct course order is [0,1] .
 * 
 * Example 2:
 * 
 * 
 * Input: 4, [[1,0],[2,0],[3,1],[3,2]]
 * Output: [0,1,2,3] or [0,2,1,3]
 * Explanation: There are a total of 4 courses to take. To take course 3 you
 * should have finished both     
 * ⁠            courses 1 and 2. Both courses 1 and 2 should be taken after you
 * finished course 0. 
 * So one correct course order is [0,1,2,3]. Another correct ordering is
 * [0,2,1,3] .
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

func findOrder(numCourses int, prerequisites [][]int) []int {
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
	resp := []int{}
	for {
		if len(queues) == 0 {
			break
		}

		queue := queues[0]
		queues = append([]int{}, queues[1:]...)
		resp = append(resp, queue)

		for _, v := range tu[queue] {
			rudu[v]--
			if rudu[v] == 0 {
				queues = append(queues, v)
			}
		}
	}

	for _, v := range rudu {
		if v != 0 {
			return nil
		}
	}
	return resp

}

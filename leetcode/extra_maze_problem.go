package leetcode

// 迷宫问题
//
// BFS，广度优先遍历
//
// 将起点加入队列
// 遍历：如果队列不为空，取出队列第一个元素a
//   取出该点所能达到的所有点，将其前置节点标记为a
//   将这些点加入队列

type pair struct {
	x int
	y int
}

type prePointerPair struct {
	pre pair
}

func maze_problem(input [][]int) [][]int {
	done := make(map[[2]int]bool, len(input)*len(input[2]))
	prePointerPairs := make([][]prePointerPair, len(input))
	for k := range input {
		prePointerPairs[k] = make([]prePointerPair, len(input[0]))
	}

	queue := []pair{{0, 0}}
	for {
		if len(queue) == 0 {
			break
		}

		// 取出队首元素
		q := queue[0]
		queue = append([]pair{}, queue[1:]...)

		x, y := q.x, q.y
		if done[[2]int{x, y}] {
			continue
		}
		done[[2]int{x, y}] = true

		if x == len(input)-1 && y == len(input[0])-1 {
			var path = [][]int{{x, y}}
			for {
				p := prePointerPairs[x][y].pre
				x, y = p.x, p.y
				path = append([][]int{{p.x, p.y}}, path...)
				if x == 0 && y == 0 {
					return path
				}
			}
		}

		if x < 0 || y < 0 || x >= len(input) || y >= len(input[0]) || input[x][y] == 1 {
			continue
		}

		for _, v := range []pair{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			next_x := x + v.x
			next_y := y + v.y

			if next_x < 0 || next_y < 0 || next_x >= len(input) || next_y >= len(input[0]) || input[next_x][next_y] == 1 || done[[2]int{next_x, next_y}] {
				continue
			}

			prePointerPairs[next_x][next_y].pre = pair{x, y}
			queue = append(queue, pair{next_x, next_y})
		}
	}

	return nil
}

package leetcode

/*
 * @lc app=leetcode id=391 lang=golang
 *
 * [391] Perfect Rectangle
 *
 * https://leetcode.com/problems/perfect-rectangle/description/
 *
 * algorithms
 * Hard (28.55%)
 * Total Accepted:    19.3K
 * Total Submissions: 67.7K
 * Testcase Example:  '[[1,1,3,3],[3,1,4,2],[3,2,4,4],[1,3,2,4],[2,3,3,4]]'
 *
 * Given N axis-aligned rectangles where N > 0, determine if they all together
 * form an exact cover of a rectangular region.
 * 
 * Each rectangle is represented as a bottom-left point and a top-right point.
 * For example, a unit square is represented as [1,1,2,2]. (coordinate of
 * bottom-left point is (1, 1) and top-right point is (2, 2)).
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * rectangles = [
 * ⁠ [1,1,3,3],
 * ⁠ [3,1,4,2],
 * ⁠ [3,2,4,4],
 * ⁠ [1,3,2,4],
 * ⁠ [2,3,3,4]
 * ]
 * 
 * Return true. All 5 rectangles together form an exact cover of a rectangular
 * region.
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * rectangles = [
 * ⁠ [1,1,2,3],
 * ⁠ [1,3,2,4],
 * ⁠ [3,1,4,2],
 * ⁠ [3,2,4,4]
 * ]
 * 
 * Return false. Because there is a gap between the two rectangular
 * regions.
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 * Example 3:
 * 
 * 
 * rectangles = [
 * ⁠ [1,1,3,3],
 * ⁠ [3,1,4,2],
 * ⁠ [1,3,2,4],
 * ⁠ [3,2,4,4]
 * ]
 * 
 * Return false. Because there is a gap in the top center.
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 * Example 4:
 * 
 * 
 * rectangles = [
 * ⁠ [1,1,3,3],
 * ⁠ [3,1,4,2],
 * ⁠ [1,3,2,4],
 * ⁠ [2,2,4,4]
 * ]
 * 
 * Return false. Because two of the rectangles overlap with each other.
 * 
 * 
 * 
 */

 type point struct {
	x,y int
}

func isRectangleCover(rectangles [][]int) bool {
	if len(rectangles) ==0 || len(rectangles[0]) == 0{
		return false
	}

	minX, maxX := math.MaxInt32, math.MinInt32
	minY, maxY := minX, maxX
	area :=0

	isCorner := make(map[point]bool)

	var p, p1,p2,p3,p4 point
	var x, y, x1,x2,y1,y2 int

	for _, r := range rectangles {
		x1,y1,x2,y2 = r[0], r[1], r[2], r[3]

		minX = min(x1, minX)
		minY = min(y1, minY)
		maxX = max(x2, maxX)
		maxY = max(y2, maxY)

		area += (x2 - x1) * (y2 -y1)

		for _, x = range []int{x1,x2} {
			for _, y = range []int{y1,y2}{
				p = point{x,y}
				if isCorner[p] {
					delete(isCorner,p)
				} else {
					isCorner[p] = true
				}
			}
		}
	}
	p1 = point{minX, minY}
	p2 = point{minX, maxY}
	p3 = point{maxX, minY}
	p4 = point{maxX, maxY}

	if !isCorner[p1] ||
		!isCorner[p2] ||
		!isCorner[p3] ||
		!isCorner[p4] ||
		len(isCorner) != 4 {
		return false
	}
	return area == (maxX-minX)*(maxY-minY)
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a,b int) int  {
	if a < b {
		return a
	}
	return b
}

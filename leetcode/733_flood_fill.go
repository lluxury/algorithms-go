package leetcode

/*
 * [733] Flood Fill
 *
 * https://leetcode.com/problems/flood-fill/description/
 *
 * algorithms
 * Easy (47.74%)
 * Total Accepted:    19.8K
 * Total Submissions: 41.3K
 * Testcase Example:  '[[1,1,1],[1,1,0],[1,0,1]]\n1\n1\n2'
 *
 *
 * An image is represented by a 2-D array of integers, each integer
 * representing the pixel value of the image (from 0 to 65535).
 *
 * Given a coordinate (sr, sc) representing the starting pixel (row and column)
 * of the flood fill, and a pixel value newColor, "flood fill" the image.
 *
 * To perform a "flood fill", consider the starting pixel, plus any pixels
 * connected 4-directionally to the starting pixel of the same color as the
 * starting pixel, plus any pixels connected 4-directionally to those pixels
 * (also with the same color as the starting pixel), and so on.  Replace the
 * color of all of the aforementioned pixels with the newColor.
 *
 * At the end, return the modified image.
 *
 * Example 1:
 *
 * Input:
 * image = [[1,1,1],[1,1,0],[1,0,1]]
 * sr = 1, sc = 1, newColor = 2
 * Output: [[2,2,2],[2,2,0],[2,0,1]]
 * Explanation:
 * From the center of the image (with position (sr, sc) = (1, 1)), all pixels
 * connected
 * by a path of the same color as the starting pixel are colored with the new
 * color.
 * Note the bottom corner is not colored 2, because it is not 4-directionally
 * connected
 * to the starting pixel.
 *
 *
 *
 * Note:
 * The length of image and image[0] will be in the range [1, 50].
 * The given starting pixel will satisfy 0  and 0 .
 * The value of each color in image[i][j] and newColor will be an integer in
 * [0, 65535].
 *
 */

/*

 实现一个画图程序的油漆桶：
 * 用一个二维的数组表示一个图片，给一个起始的地址和一个颜色值，然后将这个起始地址的上下左右四个位置和起始位置的颜色相同的数据改成和新的颜色值一样的

 * 图的深度遍历
 * 用一个二维数组记录以及访问过的坐标，访问过，直接返回
 * 坐标不合法，直接返回
 * 值和期望的不一样，直接返回

*/

func floodFill_rec(image *[][]int, done *[][]bool, origin, sr, sc int, newColor int) {
	if sr < 0 || sr >= len(*image) || sc < 0 || sc >= len((*image)[0]) {
		return
	}

	if (*done)[sr][sc] {
		return
	}

	if (*image)[sr][sc] != origin {
		return
	}

	(*image)[sr][sc] = newColor
	(*done)[sr][sc] = true

	for _, v := range [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
		floodFill_rec(image, done, origin, sr+v[0], sc+v[1], newColor)
	}
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	var done = make([][]bool, len(image))
	for k := range image {
		done[k] = make([]bool, len(image[0]))
	}

	floodFill_rec(&image, &done, image[sr][sc], sr, sc, newColor)

	return image
}

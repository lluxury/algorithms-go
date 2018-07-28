package leetcode

import "sort"

/*
 * [780] Max Chunks To Make Sorted
 *
 * https://leetcode.com/problems/max-chunks-to-make-sorted/description/
 *
 * algorithms
 * Medium (48.44%)
 * Total Accepted:    10.8K
 * Total Submissions: 22.4K
 * Testcase Example:  '[4,3,2,1,0]'
 *
 * Given an array arr that is a permutation of [0, 1, ..., arr.length - 1], we
 * split the array into some number of "chunks" (partitions), and individually
 * sort each chunk.  After concatenating them, the result equals the sorted
 * array.
 *
 * What is the most number of chunks we could have made?
 *
 * Example 1:
 *
 *
 * Input: arr = [4,3,2,1,0]
 * Output: 1
 * Explanation:
 * Splitting into two or more chunks will not return the required result.
 * For example, splitting into [4, 3], [2, 1, 0] will result in [3, 4, 0, 1,
 * 2], which isn't sorted.
 *
 *
 * Example 2:
 *
 *
 * Input: arr = [1,0,2,3,4]
 * Output: 4
 * Explanation:
 * We can split into two chunks, such as [1, 0], [2, 3, 4].
 * However, splitting into [1, 0], [2], [3], [4] is the highest number of
 * chunks possible.
 *
 *
 * Note:
 *
 *
 * arr will have length in range [1, 10].
 * arr[i] will be a permutation of [0, 1, ..., arr.length - 1].
 *
 *
 *
 */

/*

 给一个数组，问最大可以分成多少份，然后每一份都升序排序，然后组合起来的结果和整个数组升序排序的结果一致，求这个最大数
   * 很显然，是求最大值
   * 所以如果是单个数字，就直接加1，不用排序
   * 如果不相等，很显然，后面的数字要和这个数组联合起来排序，
   * 真的需要排序吗？不需要，只要相同数字的个数一样就行了
   * 所以这里需要一个标志的flag。和一个记录各个数字的map

*/

func check_map_value_zero(m map[int]int) bool {
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

func map_value_to_zero(m map[int]int) {
	for k := range m {
		m[k] = 0
	}
}

func maxChunksToSorted(arr []int) int {
	arr_sort := make([]int, len(arr))
	copy(arr_sort, arr)

	sort.Ints(arr_sort)

	count := 0
	single := true
	single_map := make(map[int]int)

	for k := range arr_sort {
		if single {
			if arr[k] == arr_sort[k] {
				count++
			} else {
				single = false
				single_map[arr[k]]++
				single_map[arr_sort[k]]--
			}
		} else {
			single_map[arr[k]]++
			single_map[arr_sort[k]]--

			if check_map_value_zero(single_map) {
				count++
				single = true
				map_value_to_zero(single_map)
			}
		}
	}
	return count
}

package leetcode

import (
	"github.com/Chyroc/algorithms-go/test"
	"testing"
)

/*
> https://leetcode.com/problems/median-of-two-sorted-arrays/description/

There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

Example 1:
nums1 = [1, 3]
nums2 = [2]

The median is 2.0
Example 2:
nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5

* 求两个已经排序的数组的中位数，要求时间复杂度是 O(log(m+n)).
  * 参考 http://cs-cjl.com/2016/06_25_leetcode_4_median_of_two_sorted_arrays
*/

// 求两个数最小数
func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

// 找到一个数的中间那个数的序号
// 奇数返回中间那个需要
// 偶数返回中间那两个数的前一个数的序号
// 从1开始计数
// 2 -> 1
// 3 -> 2
// 4 -> 2
// 5 -> 3
// 6 -> 3
func getMedianNumber(length int) int {
	return (length-1)/2 + 1
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	num1_len, num2_len := len(nums1), len(nums2)

	// 假如是一个num1_len + num2_len的有序数组，结果就是第median个数，或者（median，median+1）个数的平均数
	median := getMedianNumber(num1_len + num2_len)

	if (num1_len+num2_len)%2 == 0 {
		// 偶数，需要返回中间两个数的平均数
		return float64(findKth(nums1, nums2, num1_len, num2_len, median, 0, 0)+findKth(nums1, nums2, num1_len, num2_len, median+1, 0, 0)) / float64(2)
	} else {
		// 奇数，返回中间那个数
		return float64(findKth(nums1, nums2, num1_len, num2_len, median, 0, 0))
	}
}

// findKth 返回两个数组中第off个数
// start_1, start_2指的是两个数组中有效数据的起始位置，这样就不用拷贝数组了
//
// 然后将两个数组各分一半，a和b分别是
// 如果：num1[a+start_1-1] <= num2[b+start_2-1]，第一个数组的一半a小于第二个数组的一半b
//      那么：抛弃第一个数组的a前面的部分：num1_len-a，off-a，start_1+a
//      否则：抛弃数组2的前面的部分： num2_len-b，off-b，start_2+b
// num1_len-a 啥意思：数组的长度，不能通过数组计算得出，因为数组这里一直是一个应用
// off-a 啥意思：第off个数，如果抛弃了一个数组的前多少个，那么off就要迁移多少
// start_1+a 啥意思：
//
// 时间复杂度：是一个递归，每次抛弃a或者b个长度，这个长度大概是off的一半，即O(log(m+n))
//
// 还使用到了动态规划：
//   * 短数组长度为0，返回长数组的第off个数
//   * 求第1个数，返回两个数组的第一个元素的最小值
//   * 其余的，从两个数组的前面一共取x个，然后抛弃较小的那一端（x就是k）
func findKth(num1, num2 []int, num1_len, num2_len, off int, start_1, start_2 int) int {
	// 我们规定第一个数组比较小，所以num1_len > num2_len，调换顺序
	if num1_len > num2_len {
		return findKth(num2, num1, num2_len, num1_len, off, start_2, start_1)
	}

	// 如果小的数组，也就是第一个数组长度0，返回第二个数组的的第off个数
	if num1_len == 0 {
		return num2[off-1]
	}

	// 如果要返回的off是1，也就是返回第一个数，那么只要返回两个数组的第一个数的最小值即可
	if off == 1 {
		return min(num1[start_1], num2[start_2])
	}

	a := min(off/2, num1_len)
	b := off - a

	if num1[a+start_1-1] <= num2[b+start_2-1] {
		return findKth(num1, num2, num1_len-a, num2_len, off-a, start_1+a, start_2)
	} else {
		return findKth(num1, num2, num1_len, num2_len-b, off-b, start_1, start_2+b)
	}
}

func Test_4(t *testing.T) {
	t.Run("findMedianSortedArrays", func(t *testing.T) {
		test.Runs(t, findMedianSortedArrays, []*test.Case{
			{Input: `[1,3] \n [2]`, Output: "2"},
			{Input: `[3,4] \n [1,2,3]`, Output: "3"},
			{Input: `[1,2] \n [3,4]`, Output: "2.5"},
			{Input: `[2,3] \n []`, Output: "2.5"},
			{Input: `[2] \n []`, Output: "2"},
			{Input: `[] \n [2,3]`, Output: "2.5"},
			{Input: `[] \n [2]`, Output: "2"},
		})
	})
}

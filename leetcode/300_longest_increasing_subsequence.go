package leetcode

/*
 * [300] Longest Increasing Subsequence
 *
 * https://leetcode.com/problems/longest-increasing-subsequence/description/
 *
 * algorithms
 * Medium (39.07%)
 * Total Accepted:    140.2K
 * Total Submissions: 358.9K
 * Testcase Example:  '[10,9,2,5,3,7,101,18]'
 *
 * Given an unsorted array of integers, find the length of longest increasing
 * subsequence.
 *
 * Example:
 *
 *
 * Input: [10,9,2,5,3,7,101,18]
 * Output: 4
 * Explanation: The longest increasing subsequence is [2,3,7,101], therefore
 * the length is 4.
 *
 * Note:
 *
 *
 * There may be more than one LIS combination, it is only necessary for you to
 * return the length.
 * Your algorithm should run in O(n2) complexity.
 *
 *
 * Follow up: Could you improve it to O(n log n) time complexity?
 *
 */

/*

最大增序子序列的长度

动态规划算法

	https://www.zhihu.com/question/23995189
	https://www.cnblogs.com/steven_oyj/archive/2010/05/22/1741374.html

	首先需要定义 状态 和 状态转移方程

	状态
	Fk = 数组前k项的 最长递增子序列的长度

	状态转移
	在k项之前的所有项中，如果他比第k项小，那么第k项的值要么是max(第i项长度+1)

	状态转移方程
	F1 = 1
	Fk = 在i = 0-k-1的情况下，如果第i项比第k项小，Fk = max(Fi + 1)

*/

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	m := make([]int, len(nums))
	for k := range m {
		m[k] = 1
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				m[i] = max(m[i], m[j]+1)
			}
		}
	}

	return max(m...)
}

// todo
// 优化算法，回头看，现在先集中解决动态规划
//下面我们来看一种优化时间复杂度到O(nlgn)的解法，这里用到了二分查找法，所以才能加快运行时间哇。思路是，我们先建立一个数组ends，把首元素放进去，然后比较之后的元素，如果遍历到的新元素比ends数组中的首元素小的话，替换首元素为此新元素，如果遍历到的新元素比ends数组中的末尾元素还大的话，将此新元素添加到ends数组末尾(注意不覆盖原末尾元素)。如果遍历到的新元素比ends数组首元素大，比尾元素小时，此时用二分查找法找到第一个不小于此新元素的位置，覆盖掉位置的原来的数字，以此类推直至遍历完整个nums数组，此时ends数组的长度就是我们要求的LIS的长度，特别注意的是ends数组的值可能不是一个真实的LIS，比如若输入数组nums为{4, 2， 4， 5， 3， 7}，那么算完后的ends数组为{2， 3， 5， 7}，可以发现它不是一个原数组的LIS，只是长度相等而已，千万要注意这点。参见代码如下：
//
//
//
//解法二：
//
//复制代码
//class Solution {
//public:
//    int lengthOfLIS(vector<int>& nums) {
//        if (nums.empty()) return 0;
//        vector<int> ends{nums[0]};
//        for (auto a : nums) {
//            if (a < ends[0]) ends[0] = a;
//            else if (a > ends.back()) ends.push_back(a);
//            else {
//                int left = 0, right = ends.size();
//                while (left < right) {
//                    int mid = left + (right - left) / 2;
//                    if (ends[mid] < a) left = mid + 1;
//                    else right = mid;
//                }
//                ends[right] = a;
//            }
//        }
//        return ends.size();
//    }
//};

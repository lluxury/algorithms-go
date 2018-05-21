package leetcode

/*
> https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/

答案是和第一题是一样的，但是已经排序，有更优解？

*/
func twoSum(numbers []int, target int) []int {
	var mapInt = make(map[int]int, len(numbers))
	for k, v := range numbers {
		if j, ok := mapInt[target-v]; ok {
			if j > k {
				return []int{k + 1, j + 1}
			} else {
				return []int{j + 1, k + 1}
			}
		}
		mapInt[v] = k
	}

	return []int{}
}

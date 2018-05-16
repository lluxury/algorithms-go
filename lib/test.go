package lib

func ArrayDeepCopy(nums []int) []int {
	temp := make([]int, len(nums))
	copy(temp, nums)
	return temp
}

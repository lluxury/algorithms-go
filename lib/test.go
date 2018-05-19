package lib

import "fmt"

// IntArrayDeepCopy ...
func IntArrayDeepCopy(nums []int) []int {
	temp := make([]int, len(nums))
	copy(temp, nums)
	return temp
}

// StringArrayDeepCopy ...
func StringArrayDeepCopy(nums []string) []string {
	temp := make([]string, len(nums))
	copy(temp, nums)
	return temp
}

// ExamplePrint ...
func ExamplePrint(v interface{}) {
	fmt.Printf("%v\n", v)
}

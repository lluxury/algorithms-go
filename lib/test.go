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

// ArrayEquals ...
func ArrayEquals(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

// ArrayContain ...
func ArrayContain(s [][]int, b []int) bool {
	for _, a := range s {
		if ArrayEquals(a, b) {
			return true
		}
	}
	return false
}

// ExamplePrint ...
func ExamplePrint(v interface{}) {
	fmt.Printf("%v\n", v)
}

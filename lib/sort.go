package lib

// 快速排序，时间复杂度 nlogn
func QuickSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	mid := data[0]
	head, tail := 0, len(data)-1
	for i := 1; i <= tail; {
		if data[i] > mid {
			data[i], data[tail] = data[tail], data[i]
			tail--
		} else {
			data[i], data[head] = data[head], data[i]
			head++
			i++
		}
	}
	QuickSort(data[:head])
	QuickSort(data[head+1:])

	return data
}

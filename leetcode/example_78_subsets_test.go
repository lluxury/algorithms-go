package leetcode

import (
	"fmt"
	"math"

	"github.com/Chyroc/algorithms-go/lib"
)

/*
> https://leetcode.com/problems/subsets/description/
> Given a set of distinct integers, nums, return all possible subsets (the power set)(The solution set must not contain duplicate subsets.)


* 给一个非重复元素数据集set，求这个set的所有子集，如[1, 2]返回[[] [1] [1 2] [2]] (78)
  * 遍历递归
    * 添加空元素（第一个元素）到结果集
    * 然后把第一个元素之后的元素append到第一个元素上（遍历），所得的结果加到返回的结果集
    * 然后把第二个元素之后的结果append到[第一二个元素]上（递归），所得的结果加到返回结果集
    * 递归
    * 所以考虑到维护一个临时数组，作为一个待append的数据集
    * 下面这个函数详解
      * 把临时数据集加到结果集：*list = append(*list, temp)
      * 遍历数组
        * 把遍历的元素和临时数据集合并起来作为参数调用递归函数：想当于1分别于2，3，4合并起来调用递归函数
        * 在递归的函数中，临时结果集会是这么一个变化过程：1,12,123,13,2,23,3
  * bit二进制方法
    * 相当于一个n位的数组，每一位可以取1或者0，求所有的可能性
    * 上面这个也相当于一个n位的2进制，求所有的可能性（一共2^n种，范围0 ~ 2^n-1）
    * 那么就：循环i = 0 ~ 2^n-1
      * 将i转成2进制k
      * 以k的各位为1的index，去set中取元素
      * 返回上面选出的元素组成的数组
*/

func backtrack(list *[][]int, tempList []int, nums []int, start int) {
	*list = append(*list, lib.ArrayDeepCopy(tempList))
	for i := start; i < len(nums); i++ {
		backtrack(list, append(tempList, nums[i]), nums, i+1)
	}
}

func subsets(nums []int) [][]int {
	var list [][]int
	backtrack(&list, []int{}, nums, 0)
	return list
}

func subsets_bit(nums []int) [][]int {
	var list [][]int

	m := int(math.Pow(float64(2), float64(len(nums))))
	for i := 0; i < m; i++ {
		var l []int

		j := i
		numIndex := len(nums) - 1

		for j > 0 {
			m := j % 2
			j = (j - m) / 2
			if m == 1 {
				l = append(l, nums[numIndex])
			}
			numIndex--
		}

		list = append(list, l)
	}
	return list
}

func Example_78_res() {
	fmt.Printf("%v\n", subsets([]int{}))
	fmt.Printf("%v\n", subsets([]int{1}))
	fmt.Printf("%v\n", subsets([]int{1, 2}))
	fmt.Printf("%v\n", subsets([]int{1, 2, 3}))

	// Output:
	// [[]]
	// [[] [1]]
	// [[] [1] [1 2] [2]]
	// [[] [1] [1 2] [1 2 3] [1 3] [2] [2 3] [3]]
}

func Example_78_bit() {
	fmt.Printf("%v\n", subsets_bit([]int{}))
	fmt.Printf("%v\n", subsets_bit([]int{1}))
	fmt.Printf("%v\n", subsets_bit([]int{1, 2}))
	fmt.Printf("%v\n", subsets_bit([]int{1, 2, 3}))

	// Output:
	// [[]]
	// [[] [1]]
	// [[] [2] [1] [2 1]]
	// [[] [3] [2] [3 2] [1] [3 1] [2 1] [3 2 1]]
}

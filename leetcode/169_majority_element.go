package leetcode

func majorityElement(nums []int) int {
    res, cnt := nums[0], 1
    for i := 1; i < len(nums); i++ {
        if nums[i] == res {
            cnt++
        } else {
            cnt--
        }
        if cnt == 0 {
            res = nums[i]
            cnt = 1
        }
    }
    return res
}

package leetcode

import (
    "testing"
    "github.com/lluxury/algorithms-go/test"
)

func Test_50(t *testing.T) {
    test.Runs(t, myPow, []*test.Case{
        {Input: `2.00000 \n 10`, Output: `1024.00000`},
        {Input: `2.10000 \n 3`, Output: `9.26100`},
        {Input: `2.00000 \n -2`, Output: `0.25000`},
        })
}


// 第二个测试有一点误差,应该限定输出位数么?
// 或者是版本的区别?

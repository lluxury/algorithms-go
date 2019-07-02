package leetcode

/*

Write a function that takes an unsigned integer and return the number of '1' bits it has (also known as the Hamming weight).

 
Example 1:

Input: 00000000000000000000000000001011
Output: 3
Explanation: The input binary string 00000000000000000000000000001011 has a total of three '1' bits.
Example 2:

Input: 00000000000000000000000010000000
Output: 1
Explanation: The input binary string 00000000000000000000000010000000 has a total of one '1' bit.
Example 3:

Input: 11111111111111111111111111111101
Output: 31
Explanation: The input binary string 11111111111111111111111111111101 has a total of thirty one '1' bits.
 

Note:

Note that in some languages such as Java, there is no unsigned integer type. In this case, the input will be given as signed integer type and should not affect your implementation, as the internal binary representation of the integer is the same whether it is signed or unsigned.
In Java, the compiler represents the signed integers using 2's complement notation. Therefore, in Example 3 above the input represents the signed integer -3.

*/

// 统计给定的数中有多少个1
// 位运算第一题, n & (n-1) 去末尾 1 操作
// 虽然用除2的方法也可以达成,不过我没有试


func hammingWeight_191(num uint32) int {
    cnt := 0
    for (num != 0){
        num = num & (num - 1)
        cnt=cnt + 1
    }
    return cnt
}

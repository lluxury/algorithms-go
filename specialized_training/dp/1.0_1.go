package dp

/*

* 最优子结构
* 子问题重叠
* 边界
* 子问题独立

*/

/*

01背包

有n个物品，价值v = [v1,v2,v3…]，体积c = [c1,c2,c3…]，放入总容量为 totalCapacity 的背包中，求能获得的最大价值是多少？

找递推式（or 状态转移方程）的思路一般是由最终状态往前回溯，考察解答最终问题需要哪些子问题。

背包问题应用动态规划：

* 子问题

  * 假设f[i][j] = 前i件物品中选取若干件放入空间为j的背包中所能得到的最大价值。

* 递推式

  * 对第i件物品，不选择时，最大价值 = 前i-1件物品获得的最大价值 = f[i-1][j]；
  * 选择时，最大价值 = 前i-1件物品放入j-c[i]获得的最大价值 + i的价值 = f[i-1][j-c[i]] + v[i]；
  * 因此， f[i][j] = max(f[i-1][j],f[i-1][j-c[i]] + v[i])；
  * 当然，若c[i] > 背包总空间j，物品i只能不选，此时和第一种情况一样。

*/

func ZeroOne(v int, c []int, w []int) int {
	return 1
}

/*

最大连续子序列之和

给定K个整数的序列{ N1, N2, ..., NK }，其任意连续子序列可表示为{ Ni, Ni+1, ..., Nj }，其中 1 <= i <= j <= K。
最大连续子序列是所有连续子序中元素和最大的一个， 例如给定序列{ -2, 11, -4, 13, -5, -2 }，其最大连续子序列为{ 11, -4, 13 }，最大和为20。

* 子问题

  * sum[i] = 表示A[0...i]以A[i]结尾的子数组最大和

* 递推式

  * 对于第i个数，将 max(sum[i-1]+a[i], a[i]) 加入 sum[i]

* 状态转移方程

  * sum[i] = max(sum[i-1]+a[i], a[i])

*/


/*

最长递增子序列(LIS)

给定一个序列 An = a1 ,a2 ,  ... , an ，找出最长的子序列使得对所有 i < j ，ai < aj 。

转移方程

b[k] = max(max(b[j] | a[j] < a[k], j<k) + 1, 1)



*/
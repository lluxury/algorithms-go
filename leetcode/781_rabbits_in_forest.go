package leetcode

/*
 * [797] Rabbits in Forest
 *
 * https://leetcode.com/problems/rabbits-in-forest/description/
 *
 * algorithms
 * Medium (49.96%)
 * Total Accepted:    7.2K
 * Total Submissions: 14.4K
 * Testcase Example:  '[1,0,1,0,0]'
 *
 * In a forest, each rabbit has some color. Some subset of rabbits (possibly
 * all of them) tell you how many other rabbits have the same color as them.
 * Those answers are placed in an array.
 * 
 * Return the minimum number of rabbits that could be in the forest.
 * 
 * 
 * Examples:
 * Input: answers = [1, 1, 2]
 * Output: 5
 * Explanation:
 * The two rabbits that answered "1" could both be the same color, say red.
 * The rabbit than answered "2" can't be red or the answers would be
 * inconsistent.
 * Say the rabbit that answered "2" was blue.
 * Then there should be 2 other blue rabbits in the forest that didn't answer
 * into the array.
 * The smallest possible number of rabbits in the forest is therefore 5: 3 that
 * answered plus 2 that didn't.
 * 
 * Input: answers = [10, 10, 10]
 * Output: 11
 * 
 * Input: answers = []
 * Output: 0
 * 
 * 
 * Note:
 * 
 * 
 * answers will have length at most 1000.
 * Each answers[i] will be an integer in the range [0, 999].
 * 
 */


/*


* 题目：
  * 一个森林里面有若干兔子（每只都有颜色），其中有几只会说出森林中其他和他一样颜色兔子的个数，现在给定几个说出的数字，求森林中至少有几只兔子
  * [1,1,2]，前面1和1是两个颜色一样的兔子，2是另一个颜色的兔子（这个颜色的兔子还有两只没有说话），所以答案2+3=5
  * [10,10,10]，有11只颜色一样的兔子，答案11

* 解法
  * 将数组转为map，k是报的个数，v是对应个数出现的次数
  * 如果有一个兔子喊了x，那么就会有x+1只兔子
  * 为了最小化兔子的个数，我们认为其他也喊了x的兔子和他颜色是一样的
  * 所以如果有x+1个兔子喊了x，那么正好；如果有大于x+1个兔子喊了的话，那么剩余的继续上面的过程，也就是求余数和商的过程
  * 因为不一定每只兔子都说话，所以余数不为0的时候，认为有一组颜色的x+1只兔子
*/

func numRabbits(answers []int) int {
	if len(answers) == 0 {
		return 0
	}

	var m = make(map[int]int)
	for _, v := range answers {
		m[v]++
	}

	var r int
	for k, v := range m {
		yu := v % (k + 1)
		shang := v / (k + 1)
		if yu != 0 {
			shang++
		}
		r += shang * (k + 1)
	}

	return r
}


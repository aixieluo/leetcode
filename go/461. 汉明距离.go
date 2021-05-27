// 两个整数之间的汉明距离指的是这两个数字对应二进制位不同的位置的数目。
//
// 给出两个整数 x 和 y，计算它们之间的汉明距离。
//
// 注意：
// 0 ≤ x, y < 231.
//
// 示例:
//
// 输入: x = 1, y = 4
//
// 输出: 2
//
// 解释:
// 1   (0 0 0 1)
// 4   (0 1 0 0)
//       ↑   ↑
//
// 上面的箭头指出了对应二进制位不同的位置。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/hamming-distance

package main

func main()  {

}

func hammingDistance(x int, y int) int {
	sor := x^y
	c := 0
	for sor > 0 {
		if sor & 1 == 1 {
			c++
		}
		sor = sor >> 1
	}
	return c
}

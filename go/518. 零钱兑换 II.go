// 给定不同面额的硬币和一个总金额。写出函数来计算可以凑成总金额的硬币组合数。假设每一种面额的硬币有无限个。 
//
//  
//
// 示例 1:
//
// 输入: amount = 5, coins = [1, 2, 5]
// 输出: 4
// 解释: 有四种方式可以凑成总金额:
// 5=5
// 5=2+2+1
// 5=2+1+1+1
// 5=1+1+1+1+1
// 示例 2:
//
// 输入: amount = 3, coins = [2]
// 输出: 0
// 解释: 只用面额2的硬币不能凑成总金额3。
// 示例 3:
//
// 输入: amount = 10, coins = [10]
// 输出: 1
//  
//
// 注意:
//
// 你可以假设：
//
// 0 <= amount (总金额) <= 5000
// 1 <= coin (硬币面额) <= 5000
// 硬币种类不超过 500 种
// 结果符合 32 位符号整数
//
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/coin-change-2

package main

import "fmt"

func main() {
	fmt.Println(change(5, []int{1,2,5}))
	fmt.Println(change(500, []int{3,5,7,8,9,10,11}))
}

func change(amount int, coins []int) int {
	if amount == 0 {
		return 1
	}
	mp := make([]map[int]int, len(coins))
	for k := range mp {
		mp[k] = make(map[int]int)
	}
	var dfs func(int, int) int
	dfs = func(t int, i int) int {
		if t < 0 {
			return 0
		}
		if t == 0 {
			return 1
		}
		if val, has := mp[i][t]; has {
			return val
		}
		c := 0
		for j:=i;j< len(coins) ;j++  {
			c += dfs(t-coins[j], j)
		}
		mp[i][t] = c
		return c
	}
	return dfs(amount, 0)
}

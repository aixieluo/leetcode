// 大餐 是指 恰好包含两道不同餐品 的一餐，其美味程度之和等于 2 的幂。
//
// 你可以搭配 任意 两道餐品做一顿大餐。
//
// 给你一个整数数组 deliciousness ，其中 deliciousness[i] 是第 i​​​​​​​​​​​​​​ 道餐品的美味程度，返回你可以用数组中的餐品做出的不同 大餐 的数量。结果需要对 109 + 7 取余。
//
// 注意，只要餐品下标不同，就可以认为是不同的餐品，即便它们的美味程度相同。
//
//  
//
// 示例 1：
//
// 输入：deliciousness = [1,3,5,7,9]
// 输出：4
// 解释：大餐的美味程度组合为 (1,3) 、(1,7) 、(3,5) 和 (7,9) 。
// 它们各自的美味程度之和分别为 4 、8 、8 和 16 ，都是 2 的幂。
// 示例 2：
//
// 输入：deliciousness = [1,1,1,3,3,3,7]
// 输出：15
// 解释：大餐的美味程度组合为 3 种 (1,1) ，9 种 (1,3) ，和 3 种 (1,7) 。
//  
//
// 提示：
//
// 1 <= deliciousness.length <= 105
// 0 <= deliciousness[i] <= 220
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/count-good-meals

package main

import (
	"fmt"
	"time"
)

func main() {
	t:= time.Now()
	fmt.Println(time.Since(t))
	fmt.Println(countPairs([]int{2160,1936,3,29,27,5,2503,1593,2,0,16,0,3860,28908,6,2,15,49,6246,1946,23,105,7996,196,0,2,55,457,5,3,924,7268,16,48,4,0,12,116,2628,1468}))
	fmt.Println(countPairs([]int{64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,64}))
	fmt.Println(countPairs([]int{149,107,1,63,0,1,6867,1325,5611,2581,39,89,46,18,12,20,22,234}))
	fmt.Println(countPairs([]int{0, 1,3,5,7,9}))
	fmt.Println(countPairs([]int{1,1,1,3,3,3,7}))
}

func countPairs(deliciousness []int) int {
	const mod  = 1e9+7
	maxVal := deliciousness[0]
	for _, v :=range deliciousness[1:] {
		if v > maxVal {
			maxVal = v
		}
	}
	maxSum := 2* maxVal
	cnt := map[int]int{}
	ans := 0
	for _,v :=range deliciousness {
		for sum:=1; sum<= maxSum; sum<<=1 {
			ans += cnt[sum-v]
			ans %= mod
		}
		cnt[v]++
	}
	return ans
}

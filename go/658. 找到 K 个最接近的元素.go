// 给定一个 排序好 的数组 arr ，两个整数 k 和 x ，从数组中找到最靠近 x（两数之差最小）的 k 个数。返回的结果必须要是按升序排好的。
//
//整数 a 比整数 b 更接近 x 需要满足：
//
//|a - x| < |b - x| 或者
//|a - x| == |b - x| 且 a < b
//
//
//示例 1：
//
//输入：arr = [1,2,3,4,5], k = 4, x = 3
//输出：[1,2,3,4]
//示例 2：
//
//输入：arr = [1,2,3,4,5], k = 4, x = -1
//输出：[1,2,3,4]
//
//
//提示：
//
//1 <= k <= arr.length
//1 <= arr.length <= 104
//arr 按 升序 排列
//-104 <= arr[i], x <= 104
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/find-k-closest-elements
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findClosestElements([]int{1, 2, 3, 4, 5}, 3, 3))
	fmt.Println(findClosestElements([]int{1, 2, 3, 4, 5}, 4, -1))
	fmt.Println(findClosestElements([]int{1, 2, 5, 6}, 3, 10))
}

func findClosestElements(arr []int, k int, x int) []int {
	sc := make([]int, len(arr))
	minV := math.MaxInt
	minK := 0
	for k := range arr {
		sc[k] = abs(arr[k] - x)
		if sc[k] < minV {
			minK = k
			minV = sc[k]
		}
	}
	left, right := minK, minK
	k--
	for k > 0 {
		if left > 0 && right < len(arr)-1 {
			if sc[left-1] <= sc[right+1] {
				left--
			} else {
				right++
			}
		} else if left == 0 {
			right++
		} else {
			left--
		}
		k--
	}
	return arr[left : right+1]
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

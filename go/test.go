// 2、 [算法] 给定一个正整数数组a，以及一个目标数字x，使用你最熟悉的语言，编写一个函数，找出数组中所有可以使数字之和等于x的组合g（组合不能重复）。注：数组中的数字可以无限制被重复选取。
// 例如：a = [2,3,5,6,8]，x = 8，g = [[2,2,2,2], [2,3,3], [2,6], [8], [3,5]]

package main

import "fmt"

func main()  {
	getSum([]int{0,3,4,6,8,1}, 1)
	getSum([]int{2,3,4,6,8}, 8)
}

func getSum(nums []int, target int) {
	var f func(x int, start int)
	var list [][]int
	res := make([]int, 0, 10)
	f= func(x int, start int) {
		if x < 0 {
			return
		}
		if x == 0 && len(res) > 0 {
			list = append(list, append([]int{}, res...))
		}
		for k, v:=range nums {
			if start> k {
				continue
			}
			res = append(res, v)
			f(x-v, k)
			res = res[:len(res)-1]
		}
	}
	f(target, 0)
	fmt.Println(list)
}

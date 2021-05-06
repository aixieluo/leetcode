// n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
//
// 给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
//
// 每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
//
//  
//
// 示例 1：
//
//
// 输入：n = 4
// 输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
// 解释：如上图所示，4 皇后问题存在两个不同的解法。
// 示例 2：
//
// 输入：n = 1
// 输出：[["Q"]]
//  
//
// 提示：
//
// 1 <= n <= 9
// 皇后彼此不能相互攻击，也就是说：任何两个皇后都不能处于同一条横行、纵行或斜线上。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/n-queens

package main

import "fmt"

func main() {
	fmt.Println(solveNQueens(4))
}

func solveNQueens(n int) [][]string {
	Queens := make([]int, n+1)
	N := n
	res := make([][]int, 0)
	var isSafe func(row int) bool
	isSafe = func(row int) bool {
		for j := 1; j < row; j++ {
			isCol := Queens[j] == Queens[row]
			x1 := Queens[j]+j == (Queens[row] + row)
			x2 := Queens[j]-j == (Queens[row] - row)
			if isCol || x1 || x2 {
				return false
			}
		}
		return true
	}
	var backTracking func(row int)
	backTracking = func(row int) {
		if row <= N {
			for col:=1; col<=N; col++ {
				// if row == 1 {
				// 	Queens = make([]int, N+1)
				// }
				Queens[row] = col
				if isSafe(row) {
					backTracking(row+1)
				}
				Queens[row] = 0
			}
		} else {
			res = append(res, append([]int{}, Queens...))
		}
	}
	backTracking(1)
	cnm := make([][]string, 0)
	for _,v :=range res {
		sbs := make([]string, 0)
		for i:=1; i<=N; i++ {
			sb := make([]byte, N)
			row := v[i]
			for j:=0; j<N ;j++  {
				if j == row-1 {
					sb[j] = 'Q'
				} else {
					sb[j] = '.'
				}
			}
			sbs = append(sbs, string(sb))
		}
		cnm = append(cnm, sbs)
	}
	return cnm
}

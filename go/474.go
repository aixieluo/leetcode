package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(findMaxForm2([]string{"10", "0001", "111001", "1", "0"}, 5, 3))
}

func findMaxForm2(strs []string, m int, n int) int {
	res := make([][]int, n+1)
	for k:= range res {
		res[k] = make([]int, m+1)
	}
	for _, v:= range strs {
		one, zero := strings.Count(v, "1"), strings.Count(v, "0")
		for i:=n; i>= one; i-- {
			for j:= m; j>= zero; j-- {
				res[i][j] = max(res[i][j], 1+ res[i-one][j-zero])
			}
		}
	}
	return res[n][m]
}

func findMaxForm(strs []string, m int, n int) int {
	dp:=make([][]int,m+1)
	for index,_:=range dp{
		dp[index]=make([]int,n+1)
	}
	for _,str:=range strs{
		C_zero,C_one:=count(str)
		for zero:=m;zero>=C_zero;zero--{
			for one:=n;one>=C_one;one--{
				dp[zero][one]=max(dp[zero][one],1+dp[zero-C_zero][one-C_one])
			}
		}
	}
	return dp[m][n]
}

func count(str string)(count0,count1 int){
	count0,count1=0,0
	for _,s:=range str{
		if s=='1'{
			count1++
		}else{
			count0++
		}
	}
	return count0,count1
}


func max(x,y int)int{
	if x>y{
		return x
	}
	return y
}

package main

import "fmt"

func main()  {
	fmt.Println(permutation("abc"))
}

func permutation(s string) []string {
	hash:= make(map[string]bool)
	var dfs func(s string)
	ans := make([]string, 0)
	var ss string
	dfs = func(str string) {
		if len(str) < 1 {
			if !hash[ss] {
				ans = append(ans, ss)
				hash[ss] = true
			}
			return
		}
		for k,v :=range str {
			ss += string(v)
			dfs(str[:k] + str[k+1:])
			ss = ss[:len(ss)-1]
		}
	}
	dfs(s)
	return ans
}

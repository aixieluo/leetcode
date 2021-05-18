package main

import "fmt"

func main()  {
	fmt.Println(cuttingRope(10))
}

func cuttingRope(n int) int {
	const mod = 1e9 + 7
	ans := 1
	if n== 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	for n > 4 {
		n -= 3
		ans *= 3
		ans %= mod
	}
	return ans * n % mod
}

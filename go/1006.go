package main

import "fmt"

func main() {
	fmt.Println(clumsy(4))
}

func clumsy2(N int) int {
	var ts func(int, int) int
	ts = func(n int, i int) int {
		if n > 4 {
			sum1 := 0
			if i == 1 {
				sum1 = n*(n-1)/(n-2) + (n - 3)
			} else {
				sum1 = -n*(n-1)/(n-2) + (n - 3)
			}
			s := ts(n-4, i+4)
			return sum1 + s
		} else {
			sum2 := 0
			for ; n >= 1; n, i = n-1, i+1 {
				switch i % 4 {
				case 1:
					if i == 1 {
						sum2 = n
					} else {
						sum2 = -n
					}
				case 2:
					sum2 *= n
				case 3:
					sum2 /= n
				case 0:
					sum2 += n
				default:
					return 0
				}
			}
			return sum2
		}
	}
	return ts(N, 1)
}

func clumsy(N int) int {
	stk := []int{N}
	N--
	index := 0
	for N > 0 {
		switch index % 4 {
		case 0:
			stk[len(stk)-1] *= N
		case 1:
			stk[len(stk)-1] /= N
		case 2:
			stk = append(stk, N)
		default:
			stk = append(stk, -N)
		}
		N--
		index++
	}
	sum := 0
	for _, v := range stk {
		sum += v
	}
	return sum
}

package main

func main() {

}

func findNthDigit(n int) int {
	// 1 10
	// 2 90
	// 3 900
	// 4 9000
	if n <= 10 {
		return n - 1
	}
	jz := 9
	for (n -10 ) / jz != 0 {
		jz*= 10
	}
}

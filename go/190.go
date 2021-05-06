package main

import (
	"fmt"
	"strconv"
)

func main()  {
	fmt.Println(reverseBits(43261596))
}

func reverseBits2(num uint32) uint32 {
	n := reverse(strconv.FormatInt(int64(num), 2))
	l := len(n)
	number := 0
	for i := 0; i < l; i++ {
		number = number*2 + int(n[i] -48)
	}
	number = number << (32 - l)
	return uint32(number)
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes) -1; i < j ; i, j = i +1, j -1  {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return  string(runes)
}

func reverseBits(n uint32) (rev uint32) {
	for i := 0; i < 32 && n > 0; i++ {
		rev |= n & 1 << (31 - i)
		n >>= 1
	}
	return
}

package main

import "fmt"

func main()  {
	fmt.Println(convertToTitle(701))
	fmt.Println(convertToTitle(27))
	fmt.Println(convertToTitle(259))
	fmt.Println(convertToTitle(260))
}

func convertToTitle(columnNumber int) string {
	ans := make([]byte, 0)
	for columnNumber > 0 {
		n := columnNumber % 26
		cn := columnNumber / 26
		columnNumber = cn
		if n > 0 {
			ans = append([]byte{byte(n) + 'A' - 1}, ans...)
		} else {
			var as []byte
			if cn >1 {
				as = []byte{byte(cn) + 'A' - 2, 'Z'}
			} else {
				as = []byte{'Z'}
			}
			ans = append(as, ans...)
			columnNumber = 0
		}
	}
	return string(ans)
}

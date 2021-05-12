package main

import "fmt"

func main()  {
	fmt.Println(replaceSpace("     "))
}

func replaceSpace(s string) string {
	for k:=0; k< len(s); k++ {
		if s[k] == ' ' {
			s = s[:k] + "%20" + s[k+1:]
		}
	}
	return s
}

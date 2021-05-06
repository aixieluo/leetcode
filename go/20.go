package main

import (
	"fmt"
)

func main() {
	fmt.Println(isValid("}{"))
}

func isValid(s string) bool {
	if len(s) %2 == 1 {
		return false
	}
	dic := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := make([]byte, 0)
	for _,v := range s {
		if dv := dic[byte(v)]; dv > 0 {
			if len(stack) > 0 && stack[len(stack)-1] == dv {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, byte(v))
		}
	}
	return len(stack) == 0
}

func isValid2(s string) bool {
	z:= make([]rune, 0)
	for _,v := range s {
		switch string(v) {
		case "{":
			fallthrough
		case "[":
			fallthrough
		case "(":
			z = append(z, v)
		case "}":
			if len(z) <1 {
				return false
			}
			pop:= z[len(z)-1]
			z = z[:len(z)-1]
			if string(pop) != "{" {
				return false
			}
		case "]":
			if len(z) <1 {
				return false
			}
			pop:= z[len(z)-1]
			z = z[:len(z)-1]
			if string(pop) != "[" {
				return false
			}
		case ")":
			if len(z) <1 {
				return false
			}
			pop:= z[len(z)-1]
			z = z[:len(z)-1]
			if string(pop) != "(" {
				return false
			}
		}
	}
	return len(z) == 0
}

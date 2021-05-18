package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isNumber("0"))
	fmt.Println(isNumber("e"))
	fmt.Println(isNumber("."))
	fmt.Println(isNumber("       .1     "))
}

func isNumber(s string) bool {
	i := 0
	for ; i < len(s); i++ {
		if s[i] != ' ' {
			break
		}
	}
	s = s[i:]
	j := len(s) - 1
	for ; j >= 0; j-- {
		if s[j] != ' ' {
			break
		}
	}
	s = s[:j+1]
	ec := strings.Count(s, "e") + strings.Count(s, "E")
	if ec > 1 {
		return false
	} else if ec == 0 {
		return isXs(s) || isZs(s)
	} else {
		eIndex := 0
		for ; eIndex < len(s); eIndex++ {
			if s[eIndex] == 'e' || s[eIndex] == 'E' {
				break
			}
		}
		return (isXs(s[:eIndex]) || isZs(s[:eIndex])) && isZs(s[eIndex+1:])
	}
}

func isXs(s string) bool {
	if len(s) < 2 {
		return false
	}
	if s[0] == '+' || s[0] == '-' {
		s = s[1:]
	}
	for _, v := range s {
		if v == '.' {
			continue
		}
		if v < '0' || v > '9' {
			return false
		}
	}
	pc := strings.Count(s, ".")
	return pc == 1 && len(s) > 1
}

func isZs(s string) bool {
	if len(s) == 0 {
		return false
	}
	if s[0] == '+' || s[0] == '-' {
		s = s[1:]
	}
	for _, v := range s {
		if v < '0' || v > '9' {
			return false
		}
	}
	return len(s) > 0
}

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(deckRevealedIncreasing([]int{17, 13, 11, 2, 3, 5, 7}))
}

func deckRevealedIncreasing2(deck []int) []int {
	if len(deck) < 1 {
		return []int{}
	}
	sort.Ints(deck)
	var res []int
	for right := len(deck) - 1; right >= 0; right-- {
		if len(res) > 1 {
			res = append(res[len(res)-1:], res[:len(res)-1]...)
		}
		res = append([]int{deck[right]}, res...)
	}
	return res
}

func deckRevealedIncreasing(deck []int) []int {
	if len(deck) < 2 {
		return deck
	}
	sort.Ints(deck)
	right := len(deck) - 2
	res := []int{deck[right+1]}
	for right >0  {
		res = append(res, deck[right])
		res = append(res[1:], res[0])
		right--
	}
	res = append(res, deck[0])
	l,r:= 0, len(res) -1
	for l < r {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}
	return res
}

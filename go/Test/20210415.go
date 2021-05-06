package main

import (
	"fmt"
	"sort"
)

func main() {
	c:= Constructor(3)
	c.Push(1)
	c.Push(2)
	fmt.Println(c.Pop())
	c.Push(2)
	c.Push(3)
	c.Push(4)
	c.Increment(5, 100)
	c.Increment(2, 100)
	fmt.Println(c.Pop())
	fmt.Println(c.Pop())
	fmt.Println(c.Pop())
	fmt.Println(c.Pop())

	// o := Constructor2(5)
	// fmt.Println(o.Insert(3, "c"))
	// fmt.Println(o.Insert(1, "a"))
	// fmt.Println(o.Insert(2, "b"))
	// fmt.Println(o.Insert(5, "e"))
	// fmt.Println(o.Insert(4, "d"))
	// fmt.Println(romanToInt("MCMXCIV")) // 1994
}

type CustomStack struct {
	Max int
	List []int
}


func Constructor(maxSize int) CustomStack {
	return CustomStack{maxSize, make([]int, 0, maxSize)}
}


func (this *CustomStack) Push(x int)  {
	if len(this.List) == this.Max {
		return
	}
	this.List = append(this.List, x)
}


func (this *CustomStack) Pop() int {
	l:=len(this.List)
	if l == 0 {
		return  -1
	}
	v:= this.List[l-1]
	this.List = this.List[:l-1]
	return v
}


func (this *CustomStack) Increment(k int, val int)  {
	for i:=0; i<k && i< len(this.List); i++ {
		this.List[i] += val
	}
}

type OrderedStream struct {
	Ptr int
	Hash map[int]string
	List []int
}

func Constructor2(n int) OrderedStream {
	return OrderedStream{0, make(map[int]string), make([]int, 0, n)}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	idKey--
	this.List = append(this.List, idKey)
	this.Hash[idKey] = value
	sort.Ints(this.List)
	if _, ok:= this.Hash[this.Ptr]; ok {
		start := -1
		prev := idKey
		for k,v:= range this.List {
			if v ==idKey {
				start = k
			}
			if prev + 1 == v {
				prev = v
			}
		}
		strs := make([]string, 0)
		for _, v:=range this.List[start:prev+1]  {
			strs = append(strs, this.Hash[v])
		}
		this.Ptr = prev+1
		return strs
	}
	return []string{}
}

func romanToInt(s string) int {
	hash := map[string]int{
		"I":  1,
		"V":  5,
		"IV": 4,
		"X":  10,
		"IX": 9,
		"L":  50,
		"XL": 40,
		"C":  100,
		"XC": 90,
		"D":  500,
		"CD": 400,
		"M":  1000,
		"CM": 900,
	}
	sum := 0
	for len(s) > 0 {
		if len(s) > 1 && hash[string(s[0])] < hash[string(s[1])] {
			sum += hash[string(s[1])] - hash[string(s[0])]
			s = s[2:]
		} else {
			sum += hash[string(s[0])]
			s = s[1:]
		}
	}
	return sum
}

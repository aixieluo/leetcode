package main

type RecentCounter struct {
	Pings []int
}


func Constructor() RecentCounter {
	return RecentCounter{}
}


func (this *RecentCounter) Ping(t int) int {
	this.Pings = append(this.Pings, t)
	res := 1
	for i := len(this.Pings)-2; i >=0; i-- {
		if t - this.Pings[i] <= 3000 {
			res++
		} else {
			break
		}
	}
	return res
}

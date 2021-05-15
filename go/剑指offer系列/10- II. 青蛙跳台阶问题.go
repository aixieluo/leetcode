package main

func main() {

}

func numWays(n int) int {
    const mod = 1e9 + 7
    a,b := 1, 1
    for n > 1 {
        a,b = b, (a+b) % mod
        n--
    }
    return b
}
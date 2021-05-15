package main

func main() {

}

func fib(n int) int {
    const mod = 1e9+7
    if n == 0 {
        return 0
    }
    a,b := 0,1
    for i:=2; i<=n; i++ {
        a,b = b,(a+b) % mod
    }
    return b
}
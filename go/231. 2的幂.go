package main

func main() {

}

func isPowerOfTwo(n int) bool {
    for n > 1 {
        if n & 1 == 1 {
            return false
        }
        n >>= 1
    }
    return n == 1
}
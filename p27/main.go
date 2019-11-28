package main

import "fmt"
import "../common"

func primeRun (a int, b int, primes []int) int {
    for n := 0; n < 1000; n++ {
        v := n*n + a*n + b
        if !common.Contains(primes, v) {
            return n
        }
    }
    return 1000
}

func main() {
    primes := common.PrimesUntil(1000)
    longestRun := 0
    product := 0
    for a := -1000; a < 1001; a++ {
        for b := -1000; b < 1001; b++ {
            run := primeRun(a, b, primes)
            if run > longestRun {
                longestRun = run
                product = a * b
            }
        }
    }
    fmt.Println("product", product)
}

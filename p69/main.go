package main

import "fmt"
import "../common"

func tot (n int, primes []int) int {
    result := n
    for _, p := range primes {
        if p > result { break }

        if n % p == 0 {
            result = (p - 1) * result / p
        }
    }
    return result
}

func main () {
    limit := 1000000
    fmt.Println("pre-generating primes")
    primes := common.PrimesUntil(limit)
    fmt.Println("done")
    highest := 0.0
    highest_n := 0

    for n := 2; n <= limit; n++ {
        phi := tot(n, primes)
        ratio := float64(n) / float64(phi)

        if ratio > highest {
            highest = ratio
            highest_n = n
        }
    }

    fmt.Println("highest", highest_n)
}

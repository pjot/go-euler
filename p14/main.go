package main

import "fmt"
import "../common"

func chainLength(n int, cache map[int]int) int {
    chain := []int{}
    for !common.Contains(chain, n) {
        val, cached := cache[n]
        if cached {
            return len(chain) + val
        }
        chain = append(chain, n)
        if n == 1 {
            return len(chain)
        }
        if n%2 == 0 {
            n = n/2
        } else {
            n = 3 * n + 1
        }
    }
    return len(chain)
}

func main() {
    chains := make(map[int]int)
    longest := 0
    n := 0
    for i := 2; i < 1000000; i++ {
        length := chainLength(i, chains)
        chains[i] = length
        if length > longest {
            longest = length
            n = i
        }
    }
    fmt.Println("longest chain", n)
}

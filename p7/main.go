package main

import "fmt"
import "math"
import "../common"

func lg(f float64) float64 {
	return math.Log(f)
}

func pi(num int) int {
	n := float64(num)
	return int(n * (lg(n) + lg(lg(n))))
}

func main() {
	primes := common.PrimesUntil(pi(10001))
	fmt.Println("prime", primes[10000])
}

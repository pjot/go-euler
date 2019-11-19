package main

import "fmt"
import "../common"

func main() {
	primes := common.PrimesUntil(2000000)
	fmt.Println("sum", common.Sum(primes))
}

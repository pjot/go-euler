package main

import "fmt"

func main() {
	n := 100
	sumOfSquares := 0
	for i := 1; i <= n; i++ {
		sumOfSquares += i * i
	}
	sum := (n + 1) * (n / 2)
	squareOfSum := sum * sum

	fmt.Println("difference", squareOfSum-sumOfSquares)
}
